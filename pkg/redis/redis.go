// Package redis connects to a redis database to fetch database updates and
// logout events.
package redis

import (
	"context"
	"errors"
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

	knownStart int64
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

// Update is a blocking function that returns, when there is new data.
func (r *Redis) Update(ctx context.Context) (map[dskey.Key][]byte, error) {
	id := r.lastAutoupdateID
	if id == "" {
		id = "$"
	}

	conn := r.pool.Get()
	defer conn.Close()

	reply, err := redis.DoContext(conn, ctx, "XREAD", "COUNT", maxMessages, "BLOCK", "0", "STREAMS", fieldChangedTopic, id)
	if err != nil {
		return nil, fmt.Errorf("redis reply: %w", err)
	}

	if reply == nil {
		// This happens, when the redis command times out.
		return nil, nil
	}

	id, data, err := parseMessageBus(reply)
	if err != nil {
		return nil, fmt.Errorf("parsing message bus: %w", err)
	}

	if id != "" {
		// TODO When is id empty????
		r.lastAutoupdateID = id
	}

	return data, nil
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

const (
	connectionKey      = "autoupdate_connection"
	connectionStartKey = "autoupdate_connection_start"
)

// ConnectionAdd adds the connection counter for an user by 1.
func (r *Redis) ConnectionAdd(ctx context.Context, uid int) error {
	conn := r.pool.Get()
	defer conn.Close()

	if err := r.checkRedisStart(ctx, conn); err != nil {
		return err
	}

	if _, err := redis.DoContext(conn, ctx, "HINCRBY", connectionKey, uid, 1); err != nil {
		return fmt.Errorf("count connection: %w", err)
	}

	return nil
}

// ConnectionDone decreases the connection counter for an user by 1.
func (r *Redis) ConnectionDone(ctx context.Context, uid int) error {
	conn := r.pool.Get()
	defer conn.Close()

	if err := r.checkRedisStart(ctx, conn); err != nil {
		return err
	}

	if _, err := redis.DoContext(conn, ctx, "HINCRBY", connectionKey, uid, -1); err != nil {
		return fmt.Errorf("count connection: %w", err)
	}

	return nil
}

// ConnectionShow returns the counter.
//
// Returned is a map from user_id (as string) the the amount of connections of
// this user.
func (r *Redis) ConnectionShow(ctx context.Context) (map[string]int, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := r.checkRedisStart(ctx, conn); err != nil {
		return nil, err
	}

	val, err := redis.IntMap(redis.DoContext(conn, ctx, "HGETALL", connectionKey))
	if err != nil {
		return nil, fmt.Errorf("get hash from redis: %w", err)
	}

	return val, nil
}

// ErrNeedsRestart is returned if redis was restart after the autoupdate service
// was started.
var ErrNeedsRestart = errors.New("Needs service restart")

// checkRedisStart checks, that redis was started before this service.
//
// Autoupdate has to be restarted if:
// * Redis has started after this service has started
func (r *Redis) checkRedisStart(ctx context.Context, conn redis.Conn) error {
	now := time.Now().Unix()

	startTime, err := redis.Int64(redis.DoContext(conn, ctx, "SET", connectionStartKey, now, "NX", "GET"))
	if err != nil {
		return fmt.Errorf("count connection: %w", err)
	}

	if startTime == 0 {
		// Redis SET with the GET option returns the value that was previous in
		// the key. So when the key is set, 0 is returned.
		startTime = now
	}

	if r.knownStart == 0 {
		// First call. Just save the time.
		r.knownStart = startTime
		return nil
	}

	if startTime != r.knownStart {
		return ErrNeedsRestart
	}

	return nil
}
