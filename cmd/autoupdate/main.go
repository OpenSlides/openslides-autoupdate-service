package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/datastore"
	ahttp "github.com/openslides/openslides-autoupdate-service/internal/http"
	"github.com/openslides/openslides-autoupdate-service/internal/redis"
	"github.com/openslides/openslides-autoupdate-service/internal/restrict"
)

func main() {
	listenAddr := getEnv("LISTEN_HTTP_ADDR", ":8002")

	f := &faker{buf: bufio.NewReader(os.Stdin)}
	authService := buildAuth()
	datastoreService := buildDatastore(f)

	service := autoupdate.New(datastoreService, new(restrict.Restricter))

	handler := ahttp.New(service, authService)
	srv := &http.Server{Addr: listenAddr, Handler: handler}
	defer func() {
		service.Close()
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("Error on HTTP server shutdown: %v", err)
		}
	}()

	go func() {
		fmt.Printf("Listen on %s\n", listenAddr)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server failed: %v", err)
		}
	}()

	waitForShutdown()
}

// waitForShutdown blocks until the service should be waitForShutdown.
//
// It listens on SIGINT and SIGTERM. If the signal is received for a
// second time, the process is killed with statuscode 1.
func waitForShutdown() {
	sigint := make(chan os.Signal, 1)
	// syscall.SIGTERM is not pressent on all plattforms. Since the autoupdate
	// service is only run on linux, this is ok. If other plattforms should be supported,
	// os.Interrupt should be used instead.
	signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
	<-sigint
	go func() {
		<-sigint
		os.Exit(1)
	}()
}

// buildDatastore builds the datastore  needed by the autoupdate service.
// It uses environment variables to make the decission. Per default, the given
// faker is used.
func buildDatastore(f *faker) autoupdate.Datastore {
	var restricter autoupdate.Datastore
	fmt.Print("Datastore Service: ")
	switch getEnv("DATASTORE", "fake") {
	case "service":
		url := getEnv("DATASTORE_URL", "http://localhost:8002")
		fmt.Println(url)
		receiver := buildReceiver(f)
		restricter = datastore.New(url, receiver)
	default:
		restricter = f
		fmt.Println("fake")
	}
	return restricter
}

// buildReceiver builds the receiver needed by the datastore service.
// It uses environment variables to make the decission. Per default, the given
// faker is used.
func buildReceiver(f *faker) datastore.KeysChangedReceiver {
	var receiver datastore.KeysChangedReceiver
	fmt.Print("Messagin Service: ")
	switch getEnv("MESSAGIN_SERVICE", "fake") {
	case "redis":
		conn := redis.NewConnection(getEnv("REDIS_ADDR", "localhost:6379"))
		if getEnv("REDIS_TEST_CONN", "true") == "true" {
			if err := conn.TestConn(); err != nil {
				log.Fatalf("Can not connect to redis: %v", err)
			}
		}
		receiver = &redis.Service{Conn: conn}
		fmt.Println("redis")
	default:
		receiver = f
		fmt.Println("fake")
	}
	return receiver
}

// buildAuth returns the auth service needed by the http server.
//
// Currently, there is only the fakeAuth service.
func buildAuth() ahttp.Authenticator {
	return fakeAuth(1)
}

func getEnv(n, d string) string {
	out := os.Getenv(n)
	if out == "" {
		return d
	}
	return out
}
