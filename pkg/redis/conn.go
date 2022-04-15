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

// BlockingConn is a fake implementation of the redis connection. It does not
// create a connection but blocks forever.
type BlockingConn struct{}

// XREAD blocks forever.
func (BlockingConn) XREAD(count, stream, id string) (interface{}, error) {
	select {}
}
