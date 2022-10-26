package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
	"github.com/gomodule/redigo/redis"
)

var (
	envMessageBusHost = environment.NewVariable("MESSAGE_BUS_HOST", "localhost", "Host of the redis server.")
	envMessageBusPort = environment.NewVariable("MESSAGE_BUS_PORT", "6379", "Port of the redis server.")
)

// Pool hold the redis connection.
type Pool struct {
	pool *redis.Pool
}

// NewConn creates a new pool.
func NewConn(lookup environment.Getenver) (*Pool, []environment.Variable) {
	addr := envMessageBusHost.Value(lookup) + ":" + envMessageBusPort.Value(lookup)

	pool := Pool{
		pool: &redis.Pool{
			MaxActive:   100,
			Wait:        true,
			MaxIdle:     10,
			IdleTimeout: 240 * time.Second,
			Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
		},
	}

	usedEnv := []environment.Variable{
		envMessageBusHost,
		envMessageBusPort,
	}

	return &pool, usedEnv
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
func (s *Pool) XREAD(ctx context.Context, count, stream, id string) (interface{}, error) {
	conn := s.pool.Get()
	defer conn.Close()
	return redis.DoContext(conn, ctx, "XREAD", "COUNT", count, "BLOCK", "0", "STREAMS", stream, id)
}
