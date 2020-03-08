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
	ahttp "github.com/openslides/openslides-autoupdate-service/internal/http"
	"github.com/openslides/openslides-autoupdate-service/internal/redis"
	"github.com/openslides/openslides-autoupdate-service/internal/restrict"
)

func main() {
	listenAddr := getEnv("LISTEN_HTTP_ADDR", ":8002")

	f := faker{bufio.NewReader(os.Stdin), make(map[string]string)}
	receiver := buildReceiver(f)
	authService := buildAuth()
	restricter := buildRestricter(f)

	service := autoupdate.New(restricter, receiver)
	defer service.Close()

	handler := ahttp.New(service, authService)
	srv := &http.Server{Addr: listenAddr, Handler: handler}
	defer func() {
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

// buildReceiver builds the receiver needed by the autoupdate service.
// It uses environment variables to make the decission. Per default, the given
// faker is used.
func buildReceiver(f faker) autoupdate.KeysChangedReceiver {
	// Choose the topic service
	var receiver autoupdate.KeysChangedReceiver
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

// buildRestricter builds the restricter needed by the autoupdate service.
// It uses environment variables to make the decission. Per default, the given
// faker is used.
func buildRestricter(f faker) autoupdate.Restricter {
	var restricter autoupdate.Restricter
	switch getEnv("RESTRICTER_SERVICE", "fake") {
	case "backend":
		restricter = &restrict.Service{Addr: getEnv("BACKEND_ADDR", "http://localhost:8000")}
	default:
		restricter = f
	}
	return restricter
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
