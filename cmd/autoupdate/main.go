package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"

	"github.com/openslides/openslides-autoupdate-service/internal/auth"
	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/datastore"
	autoupdateHttp "github.com/openslides/openslides-autoupdate-service/internal/http"
	"github.com/openslides/openslides-autoupdate-service/internal/redis"
	"github.com/openslides/openslides-autoupdate-service/internal/restrict"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

type receiver interface {
	datastore.Updater
	auth.LogoutEventer
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

		"CERT_DIR": "",

		"DATASTORE":                 "fake",
		"DATASTORE_READER_HOST":     "localhost",
		"DATASTORE_READER_PORT":     "9010",
		"DATASTORE_READER_PROTOCOL": "http",

		"MESSAGING":        "fake",
		"MESSAGE_BUS_HOST": "localhost",
		"MESSAGE_BUS_PORT": "6379",
		"REDIS_TEST_CONN":  "true",

		"AUTH":                  "fake",
		"AUTH_KEY_TOKEN":        "auth-dev-key",
		"AUTH_KEY_COOKIE":       "auth-dev-key",
		"AUTH_SERIVCE_PROTOCOL": "http",
		"AUTH_SERIVCE_HOST":     "localhost",
		"AUTH_SERVICE_PORT":     "9004",
	}

	for k := range defaults {
		e, ok := os.LookupEnv(k)
		if ok {
			defaults[k] = e
		}
	}
	return defaults
}

func run() error {
	env := defaultEnv()
	closed := make(chan struct{})
	errHandler := func(err error) {
		// If an error happend, we just close the session.
		var closing interface {
			Closing()
		}
		if !errors.As(err, &closing) {
			log.Printf("Error: %v", err)
		}
	}

	// Receiver for datastore and logout events.
	r, err := buildReceiver(env)
	if err != nil {
		return fmt.Errorf("creating messsaging adapter: %w", err)
	}

	// Datastore Service.
	datastoreService, err := buildDatastore(env, r, closed, errHandler)
	if err != nil {
		return fmt.Errorf("creating datastore adapter: %w", err)
	}

	// Perm Service.
	perms := &test.MockPermission{}
	perms.Default = true

	// Restricter Service.
	restricter := restrict.New(perms, restrict.RelationChecker(restrict.RelationLists, perms))

	// Autoupdate Service.
	service := autoupdate.New(datastoreService, restricter, closed)

	// Auth Service.
	authService, err := buildAuth(env, r, closed, errHandler)
	if err != nil {
		return fmt.Errorf("creating auth adapter: %w", err)
	}

	// Create tls http2 server.
	handler := autoupdateHttp.New(service, authService)
	listenAddr := env["AUTOUPDATE_HOST"] + ":" + env["AUTOUPDATE_PORT"]
	listener, err := buildHTTPListener(env, listenAddr, handler)
	if err != nil {
		return fmt.Errorf("creating http listener: %w", err)
	}
	srv := &http.Server{Addr: listenAddr, Handler: handler}

	// Shutdown logic in separate goroutine.
	shutdownDone := make(chan struct{})
	go func() {
		defer close(shutdownDone)
		waitForShutdown()

		close(closed)
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("Error on HTTP server shutdown: %v", err)
		}
	}()

	if err := srv.Serve(listener); err != http.ErrServerClosed {
		return fmt.Errorf("http server: %w", err)
	}
	<-shutdownDone
	return nil
}

func buildHTTPListener(env map[string]string, addr string, handler http.Handler) (net.Listener, error) {
	cert, err := getCert(env)
	if err != nil {
		return nil, fmt.Errorf("getting http certs: %w", err)
	}

	tlsConf := new(tls.Config)
	tlsConf.NextProtos = []string{"h2"}
	tlsConf.Certificates = []tls.Certificate{cert}

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Can not listen on %s: %v", addr, err)
	}
	defer ln.Close()

	tlsListener := tls.NewListener(ln, tlsConf)
	fmt.Printf("Listen on %s\n", addr)
	return tlsListener, nil
}

func getCert(env map[string]string) (tls.Certificate, error) {
	const (
		generalCertName = "cert.pem"
		generalKeyName  = "key.pem"
		specialCertName = "autoupdate.pem"
		specialKeyName  = "autoupdate-key.pem"
	)

	certDir := env["CERT_DIR"]
	if certDir == "" {
		cert, err := autoupdateHttp.GenerateCert()
		if err != nil {
			return tls.Certificate{}, fmt.Errorf("creating new certificate: %w", err)
		}
		fmt.Println("Use inmemory self signed certificate")
		return cert, nil
	}
	certFile := path.Join(certDir, specialCertName)
	if _, err := os.Stat(certFile); os.IsNotExist(err) {
		certFile2 := path.Join(certDir, generalCertName)
		if _, err := os.Stat(certFile); os.IsNotExist(err) {
			return tls.Certificate{}, fmt.Errorf("%s or %s has to exist", certFile, certFile2)
		}
		certFile = certFile2
	}

	keyFile := path.Join(certDir, specialKeyName)
	if _, err := os.Stat(keyFile); os.IsNotExist(err) {
		keyFile2 := path.Join(certDir, generalKeyName)
		if _, err := os.Stat(keyFile); os.IsNotExist(err) {
			return tls.Certificate{}, fmt.Errorf("%s or %s has to exist", keyFile, keyFile2)
		}
		keyFile = keyFile2
	}

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("loading certificates from %s and %s: %w", certFile, keyFile, err)
	}
	fmt.Printf("Use certificate %s with key %s\n", certFile, keyFile)

	return cert, nil
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
func buildDatastore(env map[string]string, receiver datastore.Updater, closed <-chan struct{}, errHandler func(error)) (autoupdate.Datastore, error) {
	var url string
	dsService := env["DATASTORE"]
	switch dsService {
	case "fake":
		fmt.Println("Fake Datastore")
		url = test.NewDatastoreServer().TS.URL

	case "service":
		host := env["DATASTORE_READER_HOST"]
		port := env["DATASTORE_READER_PORT"]
		protocol := env["DATASTORE_READER_PROTOCOL"]
		url = protocol + "://" + host + ":" + port

	default:
		return nil, fmt.Errorf("unknown datastore %s", dsService)
	}

	fmt.Println("Datastore URL:", url)

	return datastore.New(url, closed, errHandler, receiver), nil
}

// buildReceiver builds the receiver needed by the datastore service. It uses
// environment variables to make the decission. Per default, the given faker is
// used.
func buildReceiver(env map[string]string) (receiver, error) {
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
		return nil, fmt.Errorf("unknown messagin service %s", serviceName)
	}

	return &redis.Service{Conn: conn}, nil
}

// buildAuth returns the auth service needed by the http server.
func buildAuth(env map[string]string, receiver auth.LogoutEventer, closed <-chan struct{}, errHandler func(error)) (autoupdateHttp.Authenticater, error) {
	method := env["AUTH"]
	switch method {
	case "ticket":
		fmt.Println("Auth Method: token")
		const debugKey = "auth-dev-key"
		tokenKey := env["AUTH_KEY_TOKEN"]
		cookieKey := env["AUTH_KEY_COOKIE"]
		if tokenKey == debugKey || cookieKey == debugKey {
			fmt.Println("Auth with debug key")
		}

		protocol := env["AUTH_SERIVCE_PROTOCOL"]
		host := env["AUTH_SERIVCE_HOST"]
		port := env["AUTH_SERIVCE_PORT"]
		url := protocol + "://" + host + ":" + port

		fmt.Printf("Auth Service: %s\n", url)
		return auth.New(url, receiver, closed, errHandler, []byte(tokenKey), []byte(cookieKey))
	case "fake":
		fmt.Println("Auth Method: FakeAuth (User ID 1 for all requests)")
		return test.Auth(1), nil
	default:
		return nil, fmt.Errorf("unknown auth method %s", method)
	}
}
