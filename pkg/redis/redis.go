// Package redis connects to a redis database to fetch database updates and
// logout events.
package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
	"github.com/gomodule/redigo/redis"
)

const (
	// maxMessages desides how many messages are read at once from the stream.
	maxMessages = "10"

	// fieldChangedTopic is the redis key name of the autoupdate stream.
	fieldChangedTopic = "ModifiedFields"

	// logoutTopic is the redis key name of the logout stream.
	logoutTopic = "logout"

	// lastLogoutDuration decides how many old logout messages are received.
	lastLogoutDuration = 15 * time.Minute
)

var (
	envMessageBusHost = environment.NewVariable("MESSAGE_BUS_HOST", "localhost", "Host of the redis server.")
	envMessageBusPort = environment.NewVariable("MESSAGE_BUS_PORT", "6379", "Port of the redis server.")
)

// Redis holds the state of the redis receiver.
type Redis struct {
	pool             *redis.Pool
	lastAutoupdateID string
	lastLogoutID     string
}

// New initializes a Redis instance.
func New(lookup environment.Environmenter) *Redis {
	addr := envMessageBusHost.Value(lookup) + ":" + envMessageBusPort.Value(lookup)

	pool := &redis.Pool{
		MaxActive:   100,
		Wait:        true,
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}

	return &Redis{
		pool: pool,
	}
}

// Wait blocks until a connection can be established.
func (r *Redis) Wait(ctx context.Context) error {
	var lastErr error
	for {
		conn := r.pool.Get()
		_, err := redis.DoContext(conn, ctx, "PING")
		conn.Close()

		if err == nil {
			return nil
		}
		lastErr = err

		select {
		case <-time.After(200 * time.Millisecond):
		case <-ctx.Done():
			return lastErr
		}
	}
}

// Update implements the Flow interface.
func (r *Redis) Update(ctx context.Context, updateFn func(map[dskey.Key][]byte, error)) {
	id := "$"

	for ctx.Err() == nil {
		newID, data, err := r.singleUpdate(ctx, id)
		if err != nil {
			updateFn(nil, err)
			time.Sleep(5 * time.Second)
			continue
		}
		updateFn(data, nil)
		id = newID
	}
}

func (r *Redis) singleUpdate(ctx context.Context, id string) (string, map[dskey.Key][]byte, error) {
	conn := r.pool.Get()
	defer conn.Close()

	reply, err := redis.DoContext(conn, ctx, "XREAD", "COUNT", maxMessages, "BLOCK", "0", "STREAMS", fieldChangedTopic, id)
	if err != nil {
		return "", nil, fmt.Errorf("redis `XREAD count %s BLOCK 0 STREAMS %s %s: %w", maxMessages, fieldChangedTopic, id, err)
	}

	if reply == nil {
		// This happens, when the redis command times out.
		return id, nil, nil
	}

	id, data, err := parseMessageBus(reply)
	if err != nil {
		return "", nil, fmt.Errorf("parsing message bus: %w", err)
	}

	return id, data, nil
}

// LogoutEvent is a blocking function that returns, when a session was revoked.
func (r *Redis) LogoutEvent(ctx context.Context) ([]string, error) {
	id := r.lastLogoutID
	if id == "" {
		// Generate an redis ID to get the logout events from the since `lastLogoutDuration`.
		id = strconv.FormatInt(time.Now().Add(-lastLogoutDuration).Unix(), 10)
	}

	conn := r.pool.Get()
	defer conn.Close()

	reply, err := redis.DoContext(conn, ctx, "XREAD", "COUNT", maxMessages, "BLOCK", "0", "STREAMS", logoutTopic, id)
	if err != nil {
		return nil, fmt.Errorf("redis reply: %w", err)
	}

	if reply == nil {
		// This happens, when the redis command times out.
		return nil, nil
	}

	id, sessionIDs, err := logoutStream(reply)
	if err != nil {
		// TODO External Error
		return nil, fmt.Errorf("parsing message bus: %w", err)
	}
	if id != "" {
		// TODO When is id empty????
		r.lastLogoutID = id
	}
	return sessionIDs, nil
}
