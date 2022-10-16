package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path"
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
	datastoreService, background, err := initDatastore(ctx, env, messageBus)
	if err != nil {
		return fmt.Errorf("creating datastore adapter: %w", err)
	}
	background(ctx)

	// Register projector in datastore.
	projector.Register(datastoreService, slide.Slides())

	// Auth Service.
	authService, authBackground, err := initAuth(env, messageBus)
	if err != nil {
		return fmt.Errorf("creating auth adapter: %w", err)
	}
	authBackground(ctx)

	// Autoupdate Service.
	service := autoupdate.New(datastoreService, restrict.Middleware)
	go service.PruneOldData(ctx)
	go service.ResetCache(ctx)

	// Start metrics.
	metric.Register(metric.Runtime)
	metricTime, err := parseDuration(env["METRIC_INTERVAL"])
	if err != nil {
		return fmt.Errorf("invalid value for `METRIC_INTERVAL`, expected duration got %s: %w", env["METRIC_INTERVAL"], err)
	}

	if metricTime > 0 {
		go metric.Loop(ctx, metricTime, log.Default())
	}

	// Start http server.
	listenAddr := ":" + env["AUTOUPDATE_PORT"]
	fmt.Printf("Listen on %s\n", listenAddr)
	return http.Run(ctx, listenAddr, authService, service)
}

func defaultEnv() map[string]string {
	defaults := map[string]string{
		"AUTOUPDATE_PORT": "9012",

		"DATASTORE_DATABASE_HOST": "localhost",
		"DATASTORE_DATABASE_PORT": "5432",
		"DATASTORE_DATABASE_USER": "openslides",
		"DATASTORE_DATABASE_NAME": "openslides",

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

		"OPENSLIDES_DEVELOPMENT": "false",
		"SECRETS_PATH":           "/run/secrets",

		"METRIC_INTERVAL":   "5m",
		"MAX_PARALLEL_KEYS": "1000",
		"DATASTORE_TIMEOUT": "3s",
	}

	for k := range defaults {
		e, ok := os.LookupEnv(k)
		if ok {
			defaults[k] = e
		}
	}
	return defaults
}

func secret(env map[string]string, name string) ([]byte, error) {
	useDev, _ := strconv.ParseBool(env["OPENSLIDES_DEVELOPMENT"])

	if useDev {
		debugSecred := "openslides"
		switch name {
		case "auth_token_key":
			debugSecred = auth.DebugTokenKey
		case "auth_cookie_key":
			debugSecred = auth.DebugCookieKey
		}

		return []byte(debugSecred), nil
	}

	path := path.Join(env["SECRETS_PATH"], name)
	secret, err := os.ReadFile(path)
	if err != nil {
		// TODO EXTERMAL ERROR
		return nil, fmt.Errorf("reading `%s`: %w", path, err)
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

func initDatastore(ctx context.Context, env map[string]string, mb *redis.Redis) (*datastore.Datastore, func(context.Context), error) {
	maxParallel, err := strconv.Atoi(env["MAX_PARALLEL_KEYS"])
	if err != nil {
		return nil, nil, fmt.Errorf("environment variable MAX_PARALLEL_KEYS has to be a number, not %s", env["MAX_PARALLEL_KEYS"])
	}

	timeout, err := parseDuration(env["DATASTORE_TIMEOUT"])
	if err != nil {
		return nil, nil, fmt.Errorf("environment variable DATASTORE_TIMEOUT has to be a duration like 3s, not %s: %w", env["DATASTORE_TIMEOUT"], err)
	}

	datastoreSource := datastore.NewSourceDatastore(
		env["DATASTORE_READER_PROTOCOL"]+"://"+env["DATASTORE_READER_HOST"]+":"+env["DATASTORE_READER_PORT"],
		mb,
		maxParallel,
		timeout,
	)
	voteCountSource := datastore.NewVoteCountSource(env["VOTE_PROTOCOL"] + "://" + env["VOTE_HOST"] + ":" + env["VOTE_PORT"])

	password, err := secret(env, "postgres_password")
	if err != nil {
		return nil, nil, fmt.Errorf("getting postgres password: %w", err)
	}

	addr := fmt.Sprintf(
		"postgres://%s@%s:%s/%s",
		env["DATASTORE_DATABASE_USER"],
		env["DATASTORE_DATABASE_HOST"],
		env["DATASTORE_DATABASE_PORT"],
		env["DATASTORE_DATABASE_NAME"],
	)

	postgresSource, err := datastore.NewSourcePostgres(ctx, addr, string(password), datastoreSource)
	if err != nil {
		return nil, nil, fmt.Errorf("creating connection to postgres: %w", err)
	}

	ds := datastore.New(
		postgresSource,
		map[string]datastore.Source{
			"poll/vote_count": voteCountSource,
		},
		datastoreSource,
	)

	eventer := func() (<-chan time.Time, func() bool) {
		timer := time.NewTimer(time.Second)
		return timer.C, timer.Stop
	}

	background := func(ctx context.Context) {
		go voteCountSource.Connect(ctx, eventer, oserror.Handle)
		go ds.ListenOnUpdates(ctx, oserror.Handle)
	}

	return ds, background, nil
}

func initAuth(env map[string]string, messageBus auth.LogoutEventer) (http.Authenticater, func(context.Context), error) {
	method := env["AUTH"]

	switch method {
	case "ticket":
		tokenKey, err := secret(env, "auth_token_key")
		if err != nil {
			return nil, nil, fmt.Errorf("getting token secret: %w", err)
		}

		cookieKey, err := secret(env, "auth_cookie_key")
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

func parseDuration(s string) (time.Duration, error) {
	sec, err := strconv.Atoi(s)
	if err == nil {
		// TODO External error
		return time.Duration(sec) * time.Second, nil
	}

	return time.ParseDuration(s)
}
