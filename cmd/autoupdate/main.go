package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/datastore"
	autoupdateHttp "github.com/openslides/openslides-autoupdate-service/internal/http"
	"github.com/openslides/openslides-autoupdate-service/internal/redis"
	"github.com/openslides/openslides-autoupdate-service/internal/restrict"
)

func main() {
	listenAddr := getEnv("AUTOUPDATE_HOST", "") + ":" + getEnv("AUTOUPDATE_PORT", "9012")
	keepAliveRaw := getEnv("KEEP_ALIVE_DURATION", "30")
	keepAlive, err := strconv.Atoi(keepAliveRaw)
	if err != nil {
		log.Fatalf("Invalid value for KEEP_ALIVE_DURATION, got %s, expected an int: %v", keepAliveRaw, err)
	}

	authService := buildAuth()
	datastoreService, err := buildDatastore()
	if err != nil {
		log.Fatalf("Can not create datastore service: %v", err)
	}

	service := autoupdate.New(datastoreService, new(restrict.Restricter))

	handler := autoupdateHttp.New(service, authService, keepAlive)
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
func buildDatastore() (autoupdate.Datastore, error) {
	var f *faker
	var url string
	dsService := getEnv("DATASTORE", "fake")
	switch dsService {
	case "fake":
		fmt.Println("Fake Datastore")
		f = newFaker(os.Stdin)
		url = f.ts.TS.URL

	case "service":
		host := getEnv("DATASTORE_READER_HOST", "localhost")
		port := getEnv("DATASTORE_READER_PORT", "9010")
		protocol := getEnv("DATASTORE_READER_PROTOCOL", "http")
		url = protocol + "://" + host + ":" + port

	default:
		return nil, fmt.Errorf("unknown datastore %s", dsService)
	}

	fmt.Println("Datastore URL:", url)
	receiver, err := buildReceiver(f)
	if err != nil {
		return nil, fmt.Errorf("build receiver: %w", err)
	}
	return datastore.New(url, receiver), nil
}

// buildReceiver builds the receiver needed by the datastore service. It uses
// environment variables to make the decission. Per default, the given faker is
// used.
func buildReceiver(f *faker) (datastore.Updater, error) {
	var receiver datastore.Updater
	serviceName := getEnv("MESSAGING", "fake")
	switch serviceName {
	case "redis":
		redisAddress := getEnv("MESSAGE_BUS_HOST", "localhost") + ":" + getEnv("MESSAGE_BUS_PORT", "6379")
		conn := redis.NewConnection(redisAddress)
		if getEnv("REDIS_TEST_CONN", "true") == "true" {
			if err := conn.TestConn(); err != nil {
				return nil, fmt.Errorf("connect to redis: %w", err)
			}
		}
		receiver = &redis.Service{Conn: conn}

	case "fake":
		receiver = f
		if f == nil {
			serviceName = "none"
		}
	default:
		return nil, fmt.Errorf("unknown messagin service %s", serviceName)
	}

	fmt.Printf("Messaging Service: %s\n", serviceName)
	return receiver, nil
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
