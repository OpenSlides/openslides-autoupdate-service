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
	"github.com/openslides/openslides-autoupdate-service/internal/redis/conn"
	"github.com/openslides/openslides-autoupdate-service/internal/restrict"
)

func main() {
	listenAddr := getEnv("LISTEN_HTTP_ADDR", ":8002")

	f := faker{bufio.NewReader(os.Stdin), make(map[string]string)}

	// Choose the topic service
	var receiver autoupdate.KeysChangedReceiver
	fmt.Print("Messagin Service: ")
	switch getEnv("MESSAGIN_SERVICE", "fake") {
	case "redis":
		conn := conn.New(getEnv("REDIS_ADDR", "localhost:6379"))
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

	// Chose the auth service
	var authService ahttp.Authenticator
	switch getEnv("AUTH_SERVICE", "fake") {
	default:
		authService = fakeAuth(1)
	}

	// Chose the restricter service
	var restricter autoupdate.Restricter
	switch getEnv("RESTRICTER_SERVICE", "fake") {
	case "backend":
		restricter = &restrict.Service{Addr: getEnv("BACKEND_ADDR", "http://localhost:8000")}
	default:
		restricter = f
	}

	aService := autoupdate.New(restricter, receiver)

	handler := ahttp.New(aService, authService)
	srv := &http.Server{Addr: listenAddr, Handler: handler}
	srv.RegisterOnShutdown(aService.Close)

	exit := make(chan struct{})
	go shutdown(srv, exit)

	fmt.Printf("Listen on %s\n", listenAddr)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server failed: %v", err)
	}

	<-exit
}

func shutdown(srv *http.Server, exit chan<- struct{}) {
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

	// We received an interrupt signal, shut down.
	if err := srv.Shutdown(context.Background()); err != nil {
		// Error from closing listeners, or context timeout:
		log.Printf("Error on HTTP server shutdown: %v", err)
	}
	close(exit)
}

func getEnv(n, d string) string {
	out := os.Getenv(n)
	if out == "" {
		return d
	}
	return out
}
