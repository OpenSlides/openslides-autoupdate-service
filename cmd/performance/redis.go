package main

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

type redisPool struct {
	pool *redis.Pool
}

// New creates a new pool
func newPool(addr string) *redisPool {
	return &redisPool{
		pool: &redis.Pool{
			MaxActive:   100,
			Wait:        true,
			MaxIdle:     10,
			IdleTimeout: 240 * time.Second,
			Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
		},
	}
}

// sendKey updates the key in redis so an autoupdate is tiggert.
func (p *redisPool) sendKey(key string) {
	conn := p.pool.Get()
	defer conn.Close()

	if _, err := conn.Do("XADD", redisTopic, "*", "modified", key); err != nil {
		log.Fatalf("Can not send data to redis: %v", err)
	}
}
