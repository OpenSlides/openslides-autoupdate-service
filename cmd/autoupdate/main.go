package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/datastore"
	autoupdateHttp "github.com/openslides/openslides-autoupdate-service/internal/http"
	"github.com/openslides/openslides-autoupdate-service/internal/redis"
	"github.com/openslides/openslides-autoupdate-service/internal/restrict"
)

func main() {
	listenAddr := ":" + getEnv("AUTOUPDATE_PORT", "9012")
	authService := buildAuth()
	datastoreService := buildDatastore()

	service := autoupdate.New(datastoreService, new(restrict.Restricter))

	handler := autoupdateHttp.New(service, authService)
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

// buildDatastore builds the datastore implementation needed by the autoupdate
// service. It uses environment variables to make the decission. Per default, a
// fake server is started and its url is used.
func buildDatastore() autoupdate.Datastore {
	dsService := getEnv("DATASTORE", "service")

	var f *faker
	var url string
	if dsService == "fake" {
		fmt.Println("Fake Datastore")
		f = newFaker(os.Stdin)
		url = f.ts.TS.URL
	} else if dsService == "service" {
		host := getEnv("DATASTORE_READER_HOST", "localhost")
		port := getEnv("DATASTORE_READER_PORT", "9010")
		url = "http://" + host + ":" + port
	} else {
		log.Fatalf("Unknown datastore reader service %s\n", dsService)
	}
	fmt.Println("Datastore URL:", url)
	return datastore.New(url, buildReceiver(f))
}

// buildReceiver builds the receiver needed by the datastore service. It uses
// environment variables to make the decission. Per default, the given faker is
// used.
func buildReceiver(f *faker) datastore.Updater {
	var receiver datastore.Updater
	serviceName := getEnv("MESSAGING_SERVICE", "redis")
	switch serviceName {
	case "redis":
		redisAddress := getEnv("MESSAGE_BUS_HOST", "localhost") + ":" + getEnv("MESSAGE_BUS_PORT", "6379")
		conn := redis.NewConnection(redisAddress)
		if getEnv("REDIS_TEST_CONN", "true") == "true" {
			if err := conn.TestConn(); err != nil {
				log.Fatalf("Can not connect to redis: %v", err)
			}
		}
		receiver = &redis.Service{Conn: conn}
	case "fake":
		receiver = f
		if f == nil {
			serviceName = "none"
		}
	default:
		log.Fatalf("Unknown messagin service %s\n", serviceName)
	}
	fmt.Printf("Messaging Service: %s\n", serviceName)
	return receiver
}

// buildAuth returns the auth service needed by the http server.
//
// Currently, there is only the fakeAuth service.
func buildAuth() autoupdateHttp.Authenticator {
	return fakeAuth(1)
}

// getEnv returns the value of the environment variable env. If it is empty, the
// defaultValue is used.
func getEnv(env, devaultValue string) string {
	value := os.Getenv(env)
	if value == "" {
		return devaultValue
	}
	return value
}
