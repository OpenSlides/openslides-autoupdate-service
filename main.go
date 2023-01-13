package main

import (
	"context"
	"fmt"
	"io"
	"log"
	gohttp "net/http"
	"os"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/http"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/metric"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/oserror"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/auth"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/redis"
	"github.com/alecthomas/kong"
)

//go:generate  sh -c "go run main.go build-doc > environment.md"

var (
	envAutoupdatePort = environment.NewVariable("AUTOUPDATE_PORT", "9012", "Port on which the service listen on.")
	envMetricInterval = environment.NewVariable("METRIC_INTERVAL", "5m", "Time in how often the metrics are gathered. Zero disables the metrics.")
)

var cli struct {
	Run      struct{} `cmd:"" help:"Runs the service." default:"withargs"`
	BuildDoc struct{} `cmd:"" help:"Build the environment documentation."`
	Health   struct{} `cmd:"" help:"Runs a health check."`
}

func main() {
	ctx, cancel := environment.InterruptContext()
	defer cancel()

	kongCTX := kong.Parse(&cli, kong.UsageOnError())
	switch kongCTX.Command() {
	case "run":
		if err := run(ctx); err != nil {
			oserror.Handle(err)
			os.Exit(1)
		}

	case "build-doc":
		if err := buildDocu(); err != nil {
			oserror.Handle(err)
			os.Exit(1)
		}

	case "health":
		if err := health(ctx); err != nil {
			oserror.Handle(err)
			os.Exit(1)
		}
	}
}

func run(ctx context.Context) error {
	lookup := new(environment.ForProduction)

	service, err := initService(lookup)
	if err != nil {
		return fmt.Errorf("init services: %w", err)
	}

	return service(ctx)
}

func buildDocu() error {
	lookup := new(environment.ForDocu)

	if _, err := initService(lookup); err != nil {
		return fmt.Errorf("init services: %w", err)
	}

	doc, err := lookup.BuildDoc()
	if err != nil {
		return fmt.Errorf("build doc: %w", err)
	}

	fmt.Println(doc)
	return nil
}

func health(ctx context.Context) error {
	port, found := os.LookupEnv("AUTOUPDATE_PORT")
	if !found {
		port = "9012"
	}

	req, err := gohttp.NewRequestWithContext(ctx, "GET", "http://localhost:"+port+"/system/autoupdate/health", nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	resp, err := gohttp.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("sending request: %w", err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("health returned status %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %w", err)
	}

	expect := `{"healthy": true}`
	got := strings.TrimSpace(string(body))
	if got != expect {
		return fmt.Errorf("got `%s`, expected `%s`", body, expect)
	}

	return nil
}

// initService initializes all packages needed for the autoupdate service.
//
// Returns a the service as callable.
func initService(lookup environment.Environmenter) (func(context.Context) error, error) {
	var backgroundTasks []func(context.Context, func(error))
	listenAddr := ":" + envAutoupdatePort.Value(lookup)

	// Redis as message bus for datastore and logout events.
	messageBus := redis.New(lookup)

	// Datastore Service.
	datastoreService, dsBackground, err := datastore.New(
		lookup,
		messageBus,
		datastore.WithVoteCount(),
		datastore.WithHistory(),
		datastore.WithProjector(),
	)
	if err != nil {
		return nil, fmt.Errorf("init datastore: %w", err)
	}
	backgroundTasks = append(backgroundTasks, dsBackground)

	// Auth Service.
	authService, authBackground := auth.New(lookup, messageBus)
	backgroundTasks = append(backgroundTasks, authBackground)

	// Restricter
	restricter := restrict.New(datastoreService)

	// TODO: Move the history logic out of the datastoreService
	history := datastoreService

	// Autoupdate Service.
	auService, auBackground, err := autoupdate.New(lookup, restricter, history)
	if err != nil {
		return nil, fmt.Errorf("init autoupdate: %w", err)
	}
	backgroundTasks = append(backgroundTasks, auBackground)

	// Start metrics.
	metric.Register(metric.Runtime)
	metricTime, err := environment.ParseDuration(envMetricInterval.Value(lookup))
	if err != nil {
		return nil, fmt.Errorf("invalid value for `METRIC_INTERVAL`, expected duration got %s: %w", envMetricInterval.Value(lookup), err)
	}

	if metricTime > 0 {
		runMetirc := func(ctx context.Context, errorHandler func(error)) {
			metric.Loop(ctx, metricTime, log.Default())
		}
		backgroundTasks = append(backgroundTasks, runMetirc)
	}

	service := func(ctx context.Context) error {
		for _, bg := range backgroundTasks {
			go bg(ctx, oserror.Handle)
		}

		// Start http server.
		fmt.Printf("Listen on %s\n", listenAddr)
		return http.Run(ctx, listenAddr, authService, auService)
	}

	return service, nil
}
