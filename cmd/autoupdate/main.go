package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate/keysbuilder"
	"github.com/openslides/openslides-autoupdate-service/internal/redis"
	"github.com/openslides/openslides-autoupdate-service/internal/redis/conn"
)

func main() {
	listenAddr := getEnv("LISTEN_HTTP_ADDR", ":8002")

	f := faker{bufio.NewReader(os.Stdin), make(map[string][]byte)}

	// Choose the topic service
	var receiver autoupdate.KeysChangedReceiver
	switch getEnv("MESSAGIN_SERVICE", "fake") {
	case "redis":
		conn := conn.New(getEnv("REDIS_ADDR", "localhost:6379"))
		if getEnv("REDIS_TEST_CONN", "true") == "true" {
			if err := conn.TestConn(); err != nil {
				log.Fatalf("Can not connect to redis: %v", err)
			}
		}
		receiver = &redis.Service{Conn: conn}
	default:
		receiver = f
	}

	// Chose the auth service
	var authService autoupdate.Authenticator
	switch getEnv("AUTH_SERVICE", "fake") {
	default:
		authService = fakeAuth{}
	}

	// Chose the restricter service
	var restricter keysbuilder.Restricter
	switch getEnv("RESTRICTER_SERVICE", "fake") {
	default:
		restricter = f
	}

	aService := autoupdate.New(restricter, receiver)

	handler := autoupdate.NewHandler(aService, authService)
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

func shutdown(srv *http.Server, exit chan struct{}) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint
	go func() {
		<-sigint
		os.Exit(1)
	}()

	// We received an interrupt signal, shut down.
	// Give open connections a timeout of one second to close.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	if err := srv.Shutdown(ctx); err != nil {
		// Error from closing listeners, or context timeout:
		log.Printf("Error on HTTP server shutdown: %v", err)
	}
	cancel()
	close(exit)
}

func getEnv(n, d string) string {
	out := os.Getenv(n)
	if out == "" {
		return d
	}
	return out
}
