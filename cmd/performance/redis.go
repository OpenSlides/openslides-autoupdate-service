package main

import (
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

type redisPool struct {
	pool *redis.Pool
}

// newPool creates a new redis pool.
func newPool(addr string) *redisPool {
	return &redisPool{
		pool: &redis.Pool{
			MaxActive:   10,
			Wait:        true,
			MaxIdle:     2,
			IdleTimeout: 240 * time.Second,
			Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
		},
	}
}

// sendKey updates the key in redis. This tiggers an autoupdate event.
func (p *redisPool) sendKey(key string) {
	conn := p.pool.Get()
	defer conn.Close()

	if _, err := conn.Do("XADD", redisTopic, "*", key, `"new value"`); err != nil {
		log.Fatalf("Can not send data to redis: %v", err)
	}
}
