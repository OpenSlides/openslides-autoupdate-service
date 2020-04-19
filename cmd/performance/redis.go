package main

import (
	"log"

	"github.com/mediocregopher/radix/v3"
)

type redisPool struct {
	client radix.Client
}

// newClient creates a new redis client.
func newClient(addr string) (*redisPool, error) {
	client, err := radix.NewPool("tcp", addr, 10)
	if err != nil {
		return nil, err
	}
	return &redisPool{
		client: client,
	}, nil
}

// sendKey updates the key in redis so an autoupdate is tiggert.
func (p *redisPool) sendKey(key string) {
	if err := p.client.Do(radix.Cmd(nil, "XADD", redisTopic, "*", key, "modified")); err != nil {
		log.Fatalf("Can not send data to redis: %v", err)
	}
}
