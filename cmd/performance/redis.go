package main

import (
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

func (p *redisPool) sendKey(key string) {
	conn := p.pool.Get()
	defer p.pool.Close()

	conn.Do("XADD", redisTopic, "*", "updated", key)
}
