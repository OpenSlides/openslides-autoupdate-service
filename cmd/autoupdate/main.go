package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/OpenSlides/openslides-permission-service/pkg/permission"
	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	autoupdateHttp "github.com/openslides/openslides-autoupdate-service/internal/http"
	"github.com/openslides/openslides-autoupdate-service/internal/projector"
	"github.com/openslides/openslides-autoupdate-service/internal/projector/slide"
	"github.com/openslides/openslides-autoupdate-service/internal/restrict"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
	"github.com/openslides/openslides-autoupdate-service/pkg/auth"
	"github.com/openslides/openslides-autoupdate-service/pkg/datastore"
	"github.com/openslides/openslides-autoupdate-service/pkg/redis"
)

type messageBus interface {
	datastore.Updater
	auth.LogoutEventer
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("Fatal error: %v", err)
	}
}

const debugKey = "auth-dev-key"

func defaultEnv() map[string]string {
	defaults := map[string]string{
		"AUTOUPDATE_HOST": "",
		"AUTOUPDATE_PORT": "9012",

		"DATASTORE_READER_HOST":     "localhost",
		"DATASTORE_READER_PORT":     "9010",
		"DATASTORE_READER_PROTOCOL": "http",

		"MESSAGING":        "fake",
		"MESSAGE_BUS_HOST": "localhost",
		"MESSAGE_BUS_PORT": "6379",
		"REDIS_TEST_CONN":  "true",

		"AUTH":          "fake",
		"AUTH_PROTOCOL": "http",
		"AUTH_HOST":     "localhost",
		"AUTH_PORT":     "9004",

		"DEACTIVATE_PERMISSION":  "false",
		"OPENSLIDES_DEVELOPMENT": "false",
	}

	for k := range defaults {
		e, ok := os.LookupEnv(k)
		if ok {
			defaults[k] = e
		}
	}
	return defaults
}

func secret(name string, dev bool) (string, error) {
	defaultSecrets := map[string]string{
		"auth_token_key":  debugKey,
		"auth_cookie_key": debugKey,
	}

	d, ok := defaultSecrets[name]
	if !ok {
		return "", fmt.Errorf("unknown secret %s", name)
	}

	s, err := openSecret(name)
	if err != nil {
		if !dev {
			return "", fmt.Errorf("can not read secret %s: %w", s, err)
		}
		s = d
	}
	return s, nil
}

func run() error {
	env := defaultEnv()

	closed := make(chan struct{})
	errHandler := func(err error) {
		// If an error happend, we just close the session.
		var closing interface {
			Closing()
		}
		if !errors.As(err, &closing) {
			log.Printf("Error: %v", err)
		}
	}

	// Receiver for datastore and logout events.
	r, err := buildReceiver(env)
	if err != nil {
		return fmt.Errorf("creating messsaging adapter: %w", err)
	}

	// Datastore Service.
	datastoreService, err := buildDatastore(env, r, closed, errHandler)
	if err != nil {
		return fmt.Errorf("creating datastore adapter: %w", err)
	}

	// Permission Service.
	var perms restrict.Permissioner = &test.MockPermission{Default: true}
	var updater autoupdate.UserUpdater = new(test.UserUpdater)
	permService := "fake"
	if env["DEACTIVATE_PERMISSION"] == "false" {
		permService = "permission"
		p := permission.New(datastoreService)
		perms = p
		updater = p
	}
	fmt.Println("Permission-Service: " + permService)

	// Restricter Service.
	checker := restrict.RelationChecker(restrict.RelationLists, perms)
	restricter := restrict.New(perms, checker)

	// Create http mux to add urls.
	mux := http.NewServeMux()
	autoupdateHttp.Health(mux)

	// Auth Service.
	authService, err := buildAuth(env, r, closed, errHandler)
	if err != nil {
		return fmt.Errorf("creating auth adapter: %w", err)
	}

	// Autoupdate Service.
	service := autoupdate.New(datastoreService, restricter, updater, closed)
	autoupdateHttp.Complex(mux, authService, service, service)
	autoupdateHttp.Simple(mux, authService, service)

	// Projector Service.
	projector.Register(datastoreService, slide.Slides())

	// Create http server.
	listenAddr := ":" + env["AUTOUPDATE_PORT"]
	srv := &http.Server{Addr: listenAddr, Handler: mux}

	// Shutdown logic in separate goroutine.
	wait := make(chan error)
	go func() {
		waitForShutdown()

		close(closed)
		if err := srv.Shutdown(context.Background()); err != nil {
			wait <- fmt.Errorf("HTTP server shutdown: %w", err)
			return
		}
		wait <- nil
	}()

	fmt.Printf("Listen on %s\n", listenAddr)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		return fmt.Errorf("HTTP Server failed: %v", err)
	}

	return <-wait
}

// waitForShutdown blocks until the service exists.
//
// It listens on SIGINT and SIGTERM. If the signal is received for a second
// time, the process is killed with statuscode 1.
func waitForShutdown() {
	sigint := make(chan os.Signal, 1)
	// syscall.SIGTERM is not pressent on all plattforms. Since the autoupdate
	// service is only run on linux, this is ok. If other plattforms should be
	// supported, os.Interrupt should be used instead.
	signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
	<-sigint
	go func() {
		<-sigint
		os.Exit(1)
	}()
}

// buildDatastore configures the datastore service.
func buildDatastore(env map[string]string, receiver datastore.Updater, closed <-chan struct{}, errHandler func(error)) (*datastore.Datastore, error) {
	protocol := env["DATASTORE_READER_PROTOCOL"]
	host := env["DATASTORE_READER_HOST"]
	port := env["DATASTORE_READER_PORT"]
	url := protocol + "://" + host + ":" + port
	return datastore.New(url, closed, errHandler, receiver), nil
}

// buildReceiver builds the receiver needed by the datastore service. It uses
// environment variables to make the decission. Per default, the given faker is
// used.
func buildReceiver(env map[string]string) (messageBus, error) {
	serviceName := env["MESSAGING"]
	fmt.Printf("Messaging Service: %s\n", serviceName)

	var conn redis.Connection
	switch serviceName {
	case "redis":
		redisAddress := env["MESSAGE_BUS_HOST"] + ":" + env["MESSAGE_BUS_PORT"]
		c := redis.NewConnection(redisAddress)
		if env["REDIS_TEST_CONN"] == "true" {
			if err := c.TestConn(); err != nil {
				return nil, fmt.Errorf("connect to redis: %w", err)
			}
		}

		conn = c

	case "fake":
		conn = redis.BlockingConn{}
	default:
		return nil, fmt.Errorf("unknown messagin service %s", serviceName)
	}

	return &redis.Redis{Conn: conn}, nil
}

// buildAuth returns the auth service needed by the http server.
func buildAuth(env map[string]string, receiver auth.LogoutEventer, closed <-chan struct{}, errHandler func(error)) (autoupdateHttp.Authenticater, error) {
	method := env["AUTH"]
	switch method {
	case "ticket":
		fmt.Println("Auth Method: ticket")
		tokenKey, err := secret("auth_token_key", env["OPENSLIDES_DEVELOPMENT"] != "false")
		if err != nil {
			return nil, fmt.Errorf("getting token secret: %w", err)
		}

		cookieKey, err := secret("auth_cookie_key", env["OPENSLIDES_DEVELOPMENT"] != "false")
		if err != nil {
			return nil, fmt.Errorf("getting cookie secret: %w", err)
		}

		if tokenKey == debugKey || cookieKey == debugKey {
			fmt.Println("Auth with debug key")
		}

		protocol := env["AUTH_PROTOCOL"]
		host := env["AUTH_HOST"]
		port := env["AUTH_PORT"]
		url := protocol + "://" + host + ":" + port

		fmt.Printf("Auth Service: %s\n", url)
		return auth.New(url, receiver, closed, errHandler, []byte(tokenKey), []byte(cookieKey))
	case "fake":
		fmt.Println("Auth Method: FakeAuth (User ID 1 for all requests)")
		return test.Auth(1), nil
	default:
		return nil, fmt.Errorf("unknown auth method %s", method)
	}
}

func openSecret(name string) (string, error) {
	f, err := os.Open("/run/secrets/" + name)
	if err != nil {
		return "", err
	}

	secret, err := io.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf("reading `/run/secrets/%s`: %w", name, err)
	}

	return string(secret), nil
}
