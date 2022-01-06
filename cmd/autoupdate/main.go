package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/autoupdate"
	autoupdateHttp "github.com/OpenSlides/openslides-autoupdate-service/internal/http"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/slide"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/test"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/auth"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/redis"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type messageBus interface {
	datastore.Updater
	auth.LogoutEventer
	autoupdateHttp.RequestMetricer
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("Fatal error: %v", err)
	}
}

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

		"VOTE_HOST":     "localhost",
		"VOTE_PORT":     "9013",
		"VOTE_PROTOCAL": "http",

		"AUTH":          "fake",
		"AUTH_PROTOCOL": "http",
		"AUTH_HOST":     "localhost",
		"AUTH_PORT":     "9004",

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
		"auth_token_key":  auth.DebugTokenKey,
		"auth_cookie_key": auth.DebugCookieKey,
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

// errHandler is called by some background tasts.
func errHandler(err error) {
	// If an error happend, we just close the session.
	var closing interface {
		Closing()
	}
	if errors.As(err, &closing) {
		return
	}

	var errNet *net.OpError
	if errors.As(err, &errNet) {
		if errNet.Op == "dial" {
			log.Printf("Can not connect to redis.")
			return
		}
	}

	log.Printf("Error: %v", err)
}

func run() error {
	env := defaultEnv()

	ctx, cancel := interruptContext()
	defer cancel()

	otelAgentAddr, ok := os.LookupEnv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if !ok {
		otelAgentAddr = "collector:4318"
	}

	traceClient := otlptracehttp.NewClient(
		otlptracehttp.WithInsecure(),
		otlptracehttp.WithEndpoint(otelAgentAddr),
	)

	traceExporter, err := otlptrace.New(ctx, traceClient)
	if err != nil {
		return fmt.Errorf("creating connection to collector: %w", err)
	}

	res, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithProcess(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithAttributes(
			// the service name used to display traces in backends
			semconv.ServiceNameKey.String("autoupdate"),
		),
	)
	if err != nil {
		return fmt.Errorf("creating otlp resouce: %w", err)
	}

	bsp := trace.NewBatchSpanProcessor(traceExporter)
	tp := trace.NewTracerProvider(
		trace.WithSpanProcessor(bsp),
		trace.WithResource(res),
	)
	defer tp.Shutdown(context.Background())
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	// Receiver for datastore and logout events.
	messageBus, err := buildMessagebus(env)
	if err != nil {
		return fmt.Errorf("creating messsaging adapter: %w", err)
	}

	// Datastore Service.
	datastoreService, err := buildDatastore(env)
	if err != nil {
		return fmt.Errorf("creating datastore adapter: %w", err)
	}
	go datastoreService.ListenOnUpdates(ctx, messageBus, errHandler)

	// Create http mux to add urls.
	mux := http.NewServeMux()

	// Auth Service.
	authService, err := buildAuth(ctx, env, messageBus, errHandler)
	if err != nil {
		return fmt.Errorf("creating auth adapter: %w", err)
	}

	voteAddr := fmt.Sprintf("%s://%s:%s", env["VOTE_PROTOCAL"], env["VOTE_HOST"], env["VOTE_PORT"])

	// Autoupdate Service.
	service := autoupdate.New(datastoreService, restrict.Middleware, voteAddr, ctx.Done())
	go service.PruneOldData(ctx)
	go service.ResetCache(ctx)

	autoupdateHttp.Health(mux)
	autoupdateHttp.Autoupdate(mux, authService, service, messageBus)
	autoupdateHttp.MetricRequest(mux, messageBus)

	// Projector Service.
	projector.Register(datastoreService, slide.Slides())

	// Create http server.
	listenAddr := ":" + env["AUTOUPDATE_PORT"]
	srv := &http.Server{Addr: listenAddr, Handler: otelhttp.NewHandler(mux, "server")}

	// Shutdown logic in separate goroutine.
	wait := make(chan error)
	go func() {
		<-ctx.Done()
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

// interruptContext works like signal.NotifyContext
//
// In only listens on os.Interrupt. If the signal is received two times,
// os.Exit(1) is called.
func interruptContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		cancel()

		// If the signal was send for the second time, make a hard cut.
		<-sigint
		os.Exit(1)
	}()
	return ctx, cancel
}

// buildDatastore configures the datastore service.
func buildDatastore(
	env map[string]string,
) (*datastore.Datastore, error) {
	protocol := env["DATASTORE_READER_PROTOCOL"]
	host := env["DATASTORE_READER_HOST"]
	port := env["DATASTORE_READER_PORT"]
	url := protocol + "://" + host + ":" + port
	return datastore.New(url), nil
}

// buildMessagebus builds the receiver needed by the datastore service. It uses
// environment variables to make the decission. Per default, the given faker is
// used.
func buildMessagebus(env map[string]string) (messageBus, error) {
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
		return nil, fmt.Errorf("unknown messagin service %q", serviceName)
	}

	return &redis.Redis{Conn: conn}, nil
}

// buildAuth returns the auth service needed by the http server.
//
// This function is not blocking. The context is used to give it to auth.New
// that uses it to stop background goroutines.
func buildAuth(
	ctx context.Context,
	env map[string]string,
	messageBus auth.LogoutEventer,
	errHandler func(error),
) (autoupdateHttp.Authenticater, error) {
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

		if tokenKey == auth.DebugTokenKey || cookieKey == auth.DebugCookieKey {
			fmt.Println("Auth with debug key")
		}

		protocol := env["AUTH_PROTOCOL"]
		host := env["AUTH_HOST"]
		port := env["AUTH_PORT"]
		url := protocol + "://" + host + ":" + port

		fmt.Printf("Auth Service: %s\n", url)
		a, err := auth.New(url, ctx.Done(), []byte(tokenKey), []byte(cookieKey))
		if err != nil {
			return nil, fmt.Errorf("creating auth service: %w", err)
		}
		go a.ListenOnLogouts(ctx, messageBus, errHandler)
		go a.PruneOldData(ctx)

		return a, nil

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
