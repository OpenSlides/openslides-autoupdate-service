package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/http"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/metric"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/oserror"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/slide"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/auth"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/redis"
	"golang.org/x/sys/unix"
)

func main() {
	if err := run(); err != nil {
		oserror.Handle(err)
		os.Exit(1)
	}
}

func run() error {
	ctx, cancel := interruptContext()
	defer cancel()

	env := defaultEnv()

	// Redis as message bus for datastore and logout events.
	messageBus, err := initRedis(env)
	if err != nil {
		return fmt.Errorf("init redis as message bus: %w", err)
	}

	// Datastore Service.
	datastoreService, err := initDatastore(env, messageBus)
	if err != nil {
		return fmt.Errorf("creating datastore adapter: %w", err)
	}
	go datastoreService.ListenOnUpdates(ctx, oserror.Handle)

	// Register projector in datastore.
	projector.Register(datastoreService, slide.Slides())

	// Auth Service.
	authService, authBackground, err := initAuth(env, messageBus)
	if err != nil {
		return fmt.Errorf("creating auth adapter: %w", err)
	}
	go authBackground(ctx)

	// Autoupdate Service.
	service := autoupdate.New(datastoreService, restrict.Middleware)
	go service.PruneOldData(ctx)
	go service.ResetCache(ctx)

	// Start metrics.
	metric.Register(metric.Runtime)
	metricSeconds := 0
	if got, err := strconv.Atoi(env["METRIC_INTERVAL_SECONDS"]); err == nil {
		metricSeconds = got
	}
	if metricSeconds > 0 {
		go metric.Loop(ctx, time.Duration(metricSeconds)*time.Second, log.Default())
	}

	// Start http server.
	listenAddr := ":" + env["AUTOUPDATE_PORT"]
	fmt.Printf("Listen on %s\n", listenAddr)
	return http.Run(ctx, listenAddr, authService, service)
}

func defaultEnv() map[string]string {
	defaults := map[string]string{
		"AUTOUPDATE_HOST": "",
		"AUTOUPDATE_PORT": "9012",

		"DATASTORE_READER_HOST":     "localhost",
		"DATASTORE_READER_PORT":     "9010",
		"DATASTORE_READER_PROTOCOL": "http",

		"MESSAGE_BUS_HOST": "localhost",
		"MESSAGE_BUS_PORT": "6379",
		"REDIS_TEST_CONN":  "true",

		"VOTE_HOST":     "localhost",
		"VOTE_PORT":     "9013",
		"VOTE_PROTOCOL": "http",

		"AUTH":          "fake",
		"AUTH_PROTOCOL": "http",
		"AUTH_HOST":     "localhost",
		"AUTH_PORT":     "9004",

		"OPENSLIDES_DEVELOPMENT":  "false",
		"METRIC_INTERVAL_SECONDS": "300",
	}

	for k := range defaults {
		e, ok := os.LookupEnv(k)
		if ok {
			defaults[k] = e
		}
	}
	return defaults
}

func secret(name string, dev bool) ([]byte, error) {
	if dev {
		debugSecred := "openslides"
		switch name {
		case "auth_token_key":
			debugSecred = auth.DebugTokenKey
		case "auth_cookie_key":
			debugSecred = auth.DebugCookieKey
		}

		return []byte(debugSecred), nil
	}

	secret, err := os.ReadFile("/run/secrets/" + name)
	if err != nil {
		// TODO EXTERMAL ERROR
		return nil, fmt.Errorf("reading `/run/secrets/%s`: %w", name, err)
	}

	return secret, nil
}

// interruptContext works like signal.NotifyContext. It returns a context that
// is canceled, when a signal is received.
//
// It listens on os.Interrupt and unix.SIGTERM. If the signal is received two
// times, os.Exit(2) is called.
func interruptContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, unix.SIGTERM)
		<-sig
		cancel()
		<-sig
		os.Exit(2)
	}()
	return ctx, cancel
}

func initRedis(env map[string]string) (*redis.Redis, error) {
	redisAddress := env["MESSAGE_BUS_HOST"] + ":" + env["MESSAGE_BUS_PORT"]
	conn := redis.NewConnection(redisAddress)
	if ok, _ := strconv.ParseBool(env["REDIS_TEST_CONN"]); ok {
		if err := conn.TestConn(); err != nil {
			return nil, fmt.Errorf("connect to redis: %w", err)
		}
	}

	return &redis.Redis{Conn: conn}, nil
}

func initDatastore(env map[string]string, mb *redis.Redis) (*datastore.Datastore, error) {
	datastoreSource := datastore.NewSourceDatastore(env["DATASTORE_READER_PROTOCOL"]+"://"+env["DATASTORE_READER_HOST"]+":"+env["DATASTORE_READER_PORT"], mb, 1000)
	voteCountSource := datastore.NewVoteCountSource(env["VOTE_PROTOCOL"] + "://" + env["VOTE_HOST"] + ":" + env["VOTE_PORT"])

	return datastore.New(
		datastoreSource,
		map[string]datastore.Source{
			"poll/vote_count": voteCountSource,
		},
		datastoreSource,
	), nil
}

func initAuth(env map[string]string, messageBus auth.LogoutEventer) (http.Authenticater, func(context.Context), error) {
	method := env["AUTH"]

	switch method {
	case "ticket":
		useDev, _ := strconv.ParseBool(env["OPENSLIDES_DEVELOPMENT"])

		tokenKey, err := secret("auth_token_key", useDev)
		if err != nil {
			return nil, nil, fmt.Errorf("getting token secret: %w", err)
		}

		cookieKey, err := secret("auth_cookie_key", useDev)
		if err != nil {
			return nil, nil, fmt.Errorf("getting cookie secret: %w", err)
		}

		url := fmt.Sprintf("%s://%s:%s", env["AUTH_PROTOCOL"], env["AUTH_HOST"], env["AUTH_PORT"])
		a, err := auth.New(url, tokenKey, cookieKey)
		if err != nil {
			return nil, nil, fmt.Errorf("creating auth service: %w", err)
		}

		backgroundtask := func(ctx context.Context) {
			go a.ListenOnLogouts(ctx, messageBus, oserror.Handle)
			go a.PruneOldData(ctx)
		}

		return a, backgroundtask, nil

	case "fake":
		fmt.Println("Auth Method: FakeAuth (User ID 1 for all requests)")
		return auth.Fake(1), func(context.Context) {}, nil

	default:
		// TODO LAST ERROR
		return nil, nil, fmt.Errorf("unknown auth method: %s", method)
	}
}
