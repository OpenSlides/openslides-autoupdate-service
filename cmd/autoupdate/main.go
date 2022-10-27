package main

import (
	"fmt"
	"log"
	"os"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/http"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/metric"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/oserror"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/slide"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/auth"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/redis"
)

var (
	envAutoupdatePort = environment.NewVariable("AUTOUPDATE_PORT", "9012", "Port on which the service listen on.")
	envMetricInterval = environment.NewVariable("METRIC_INTERVAL", "5m", "Time in how often the metrics are gathered. Zero disables the metrics.")
)

func main() {
	if err := run(); err != nil {
		oserror.Handle(err)
		os.Exit(1)
	}
}

func run() error {
	ctx, cancel := environment.InterruptContext()
	defer cancel()

	lookup := environment.Getenvfunc(os.Getenv)
	var environmentVariables []environment.Variable

	// Redis as message bus for datastore and logout events.
	messageBus, redisEnv := redis.New(lookup)
	environmentVariables = append(environmentVariables, redisEnv...)

	// Datastore Service.
	datastoreService, dsEnv, dsBackground, err := datastore.New(ctx, lookup, messageBus, datastore.WithVoteCount(), datastore.WithHistory())
	if err != nil {
		return fmt.Errorf("init datastore: %w", err)
	}
	environmentVariables = append(environmentVariables, dsEnv...)
	go dsBackground(ctx)

	// Register projector in datastore. (TODO: Should be an option from datastore.New)
	projector.Register(datastoreService, slide.Slides())

	// Auth Service.
	authService, authEnv, authBackground := auth.New(lookup, messageBus)
	environmentVariables = append(environmentVariables, authEnv...)
	go authBackground(ctx)

	// Autoupdate Service.
	service, auBackground := autoupdate.New(datastoreService, restrict.Middleware)
	go auBackground(ctx)

	// Start metrics.
	metric.Register(metric.Runtime)
	metricTime, err := environment.ParseDuration(envMetricInterval.Value(lookup))
	if err != nil {
		return fmt.Errorf("invalid value for `METRIC_INTERVAL`, expected duration got %s: %w", envMetricInterval.Value(lookup), err)
	}

	if metricTime > 0 {
		go metric.Loop(ctx, metricTime, log.Default())
	}

	// Start http server.
	listenAddr := ":" + envAutoupdatePort.Value(lookup)
	fmt.Printf("Listen on %s\n", listenAddr)
	return http.Run(ctx, listenAddr, authService, service)
}
