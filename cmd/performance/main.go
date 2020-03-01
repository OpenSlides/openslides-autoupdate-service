// Package performance implments a test tool that opens a lot of connections to an
// autoupdate server and measures how long it takes to connect and receive some data
//
// The autoupdate service has to be started with the redis backend.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	connections = 5000
	url         = "http://localhost:8002/system/autoupdate/keys?" + keyName
	redisAddr   = "localhost:6379"
	redisTopic  = "ModifiedFields"
	keyName     = "MyKey"
)

func main() {
	keepOpen := flag.Bool("keep-open", false, "Keeps the connections open after the test succeeded.")
	flag.Parse()

	p := newPool(redisAddr)

	// Create clients
	clients := make([]*client, connections)
	for i := 0; i < connections; i++ {
		clients[i] = &client{}
	}
	keys := make(chan string, connections)
	errs := []error{}

	// Connect test
	start := time.Now()
	for _, c := range clients {
		go func(c *client) {
			if err := c.connect(context.Background(), keys); err != nil {
				errs = append(errs, err)
			}
		}(c)
	}
	readClients(connections, keys)
	log.Printf("Connect %d clients took %d milliseconds", connections, time.Since(start)/time.Millisecond)

	// Update one key
	start = time.Now()
	p.sendKey(keyName)
	readClients(connections, keys)
	log.Printf("Send and Receive one key took %d milliseconds", time.Since(start)/time.Millisecond)

	if *keepOpen {
		fmt.Println("Connections are kept open...")

		for len(errs) == 0 {
			readClients(connections, keys)
			log.Println("Connections received data")
		}
	}
	if len(errs) > 0 {
		fmt.Printf("Errors: %d, first: %v\n", len(errs), errs[0])
		os.Exit(1)
	}
}

func readClients(count int, c <-chan string) {
	for i := 0; i < count; i++ {
		<-c
	}
}
