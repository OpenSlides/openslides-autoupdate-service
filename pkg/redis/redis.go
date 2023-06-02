// Package redis connects to a redis database to fetch database updates and
// logout events.
package redis

import (
	"context"
	"fmt"
	"math/rand"
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
	envMetricTooOld   = environment.NewVariable("METRIC_TOO_OLD", "5m", "Ignore metric values that are older then this value.")
)

// Redis holds the state of the redis receiver.
type Redis struct {
	pool             *redis.Pool
	lastAutoupdateID string
	lastLogoutID     string

	instanceID  string
	metricToOld time.Duration
	now         func() time.Time
}

// New initializes a Redis instance.
func New(lookup environment.Environmenter, now func() time.Time) (*Redis, error) {
	addr := envMessageBusHost.Value(lookup) + ":" + envMessageBusPort.Value(lookup)

	pool := &redis.Pool{
		MaxActive:   100,
		Wait:        true,
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}

	tooOld, err := environment.ParseDuration(envMetricTooOld.Value(lookup))
	if err != nil {
		return nil, fmt.Errorf("parsing METRIC_TOO_OLD: %w", err)
	}

	if now == nil {
		now = time.Now
	}

	r := Redis{
		pool:        pool,
		instanceID:  buildInstanceID(now),
		metricToOld: tooOld,
		now:         now,
	}

	return &r, nil
}

func buildInstanceID(now func() time.Time) string {
	t := now().UTC().Format("2006-01-02T15:04:05")

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 8

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return fmt.Sprintf("%s-%s", t, b)
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

const metricKeyPrefix = "metric"

// SaveMetric saves a metric value for this instance.
func (r *Redis) SaveMetric(ctx context.Context, name string, v int) error {
	conn := r.pool.Get()
	defer conn.Close()

	valueKey := fmt.Sprintf("%s-%s-values", metricKeyPrefix, name)
	timeStampKey := fmt.Sprintf("%s-%s-timestamp", metricKeyPrefix, name)

	if _, err := redis.DoContext(conn, ctx, "HSET", valueKey, r.instanceID, v); err != nil {
		return fmt.Errorf("redis save value: %w", err)
	}

	now := r.now().UTC().Unix()

	if _, err := redis.DoContext(conn, ctx, "HSET", timeStampKey, r.instanceID, now); err != nil {
		return fmt.Errorf("redis save timestamp: %w", err)
	}

	return nil
}

// Metric returns a metric from redis.
func (r *Redis) Metric(ctx context.Context, name string) (int, error) {
	conn := r.pool.Get()
	defer conn.Close()

	valueKey := fmt.Sprintf("%s-%s-values", metricKeyPrefix, name)
	timeStampKey := fmt.Sprintf("%s-%s-timestamp", metricKeyPrefix, name)

	values, err := redis.IntMap(redis.DoContext(conn, ctx, "HGETALL", valueKey))
	if err != nil {
		return 0, fmt.Errorf("redis HALS: %w", err)
	}

	timeStamps, err := redis.IntMap(redis.DoContext(conn, ctx, "HGETALL", timeStampKey))
	if err != nil {
		return 0, fmt.Errorf("redis HALS: %w", err)
	}

	tooOld := r.now().UTC().Add(-r.metricToOld).Unix()

	fmt.Println(values, timeStamps, tooOld)

	v := 0
	for k, timestamp := range timeStamps {
		if timestamp < int(tooOld) {
			continue
		}

		v += values[k]
	}

	return v, nil
}
