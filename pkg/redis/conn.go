package redis

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

// Pool hold the redis connection.
type Pool struct {
	pool *redis.Pool
}

// NewConnection creates a new pool.
func NewConnection(addr string) *Pool {
	return &Pool{
		pool: &redis.Pool{
			MaxActive:   100,
			Wait:        true,
			MaxIdle:     10,
			IdleTimeout: 240 * time.Second,
			Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
		},
	}
}

// TestConn sends a ping command to redis. Does not return the response, but an
// error if there is no response.
func (s *Pool) TestConn() error {
	conn := s.pool.Get()
	defer conn.Close()

	if _, err := conn.Do("PING"); err != nil {
		return fmt.Errorf("no connection to redis: %w", err)
	}
	return nil
}

// XREAD reads new messages from one stream.
func (s *Pool) XREAD(count, stream, id string) (interface{}, error) {
	conn := s.pool.Get()
	defer conn.Close()
	return conn.Do("XREAD", "COUNT", count, "BLOCK", "0", "STREAMS", stream, id)
}

// ZINCR increments a sorted set by one.
func (s *Pool) ZINCR(key string, value []byte) error {
	conn := s.pool.Get()
	defer conn.Close()
	_, err := conn.Do("ZINCRBY", key, 1, value)
	return err
}

// ZRANGE returns all values from a sorted set with the scores.
func (s *Pool) ZRANGE(key string) (interface{}, error) {
	conn := s.pool.Get()
	defer conn.Close()
	return conn.Do("ZREVRANGE", key, 0, -1, "WITHSCORES")
}

// BlockingConn is a fake implementation of the redis connection. It does not
// create a connection but blocks forever.
type BlockingConn struct{}

// XREAD blocks forever.
func (BlockingConn) XREAD(count, stream, id string) (interface{}, error) {
	select {}
}

// ZINCR does nothing.
func (BlockingConn) ZINCR(key string, value []byte) error {
	return nil
}

// ZRANGE does nothing.
func (BlockingConn) ZRANGE(key string) (interface{}, error) {
	return nil, nil
}
