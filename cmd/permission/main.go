package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/OpenSlides/openslides-permission-service/internal/definitions"
	permHTTP "github.com/OpenSlides/openslides-permission-service/internal/http"
	"github.com/OpenSlides/openslides-permission-service/pkg/permission"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Fatal error: %v", err)
	}
}

func defaultEnv() map[string]string {
	defaults := map[string]string{
		"PERMISSION_HOST": "",
		"PERMISSION_PORT": "9005",
	}

	for k := range defaults {
		e, ok := os.LookupEnv(k)
		if ok {
			defaults[k] = e
		}
	}
	return defaults
}

type fakeDataProvider struct{}

func (dp fakeDataProvider) Get(ctx context.Context, fqfields ...definitions.Fqfield) ([]json.RawMessage, error) {
	m := make([]json.RawMessage, len(fqfields))
	for i := range fqfields {
		m[i] = json.RawMessage(strconv.Itoa(i))
	}
	return m, nil
}

func run() error {
	env := defaultEnv()
	myDataProvider := fakeDataProvider{}
	ps := permission.New(myDataProvider)

	// Register handlers
	mux := http.NewServeMux()
	permHTTP.Health(mux)
	permHTTP.IsAllowed(mux, ps)

	listenAddr := env["PERMISSION_HOST"] + ":" + env["PERMISSION_PORT"]
	fmt.Printf("Listen on: %s\n", listenAddr)

	srv := &http.Server{
		Addr:    listenAddr,
		Handler: mux,
	}

	// Shutdown logic in separate goroutine.
	shutdownDone := make(chan struct{})
	go func() {
		defer close(shutdownDone)
		waitForShutdown()

		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("Error on HTTP server shutdown: %v", err)
		}
	}()

	// Start the http service. This blocks until the server is stopped.
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		return fmt.Errorf("http server: %w", err)
	}
	<-shutdownDone
	return nil
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
