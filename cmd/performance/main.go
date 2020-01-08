package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

const (
	connections = 5000
	url         = "http://localhost:8002/autoupdate/"
	redisAddr   = "localhost:6379"
	redisTopic  = "field_changed"
)

func main() {
	//p := newPool(redisAddr)
	// Create clients
	clients := make([]*client, connections)
	for i := 0; i < connections; i++ {
		clients[i] = &client{}
	}

	// Connect clients
	keys := make(chan string, connections)
	errs := []error{}
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

	// start = time.Now()
	// p.sendKey("user/5/name")
	// readClients(connections, keys)
	// log.Printf("Send and Receive one key took %d milliseconds", time.Since(start)/time.Millisecond)

	for len(errs) == 0 {
		readClients(connections, keys)
		log.Println("All data received")
	}
	fmt.Printf("Errors: %d, first: %v\n", len(errs), errs[0])
}

func readClients(count int, c <-chan string) {
	for i := 0; i < count; i++ {
		<-c
	}
}
