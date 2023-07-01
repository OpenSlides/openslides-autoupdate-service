package redis

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/gomodule/redigo/redis"
)

const metricKeyPrefix = "metric"

// MetricValueCombiner is a helper for the Metric to combine multible metric
// values into one.
type MetricValueCombiner[T any] interface {
	Combine(value string, acc T) (T, error)
}

// Metric saves a metric value in redis.
type Metric[T any] struct {
	r *Redis

	name      string
	converter MetricValueCombiner[T]

	instanceID string
	tooOld     time.Duration
	now        func() time.Time
}

// NewMetric initializes a redis metric.
func NewMetric[T any](r *Redis, name string, converter MetricValueCombiner[T], tooOld time.Duration, now func() time.Time) Metric[T] {
	if now == nil {
		now = time.Now
	}

	return Metric[T]{
		r: r,

		name:      name,
		converter: converter,

		instanceID: buildInstanceID(now),
		tooOld:     tooOld,
		now:        now,
	}
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

// Save saves a metric value for this instance.
func (m Metric[T]) Save(ctx context.Context, value string) error {
	conn := m.r.pool.Get()
	defer conn.Close()

	valueKey := fmt.Sprintf("%s-%s-values", metricKeyPrefix, m.name)
	timeStampKey := fmt.Sprintf("%s-%s-timestamp", metricKeyPrefix, m.name)

	if _, err := redis.DoContext(conn, ctx, "HSET", valueKey, m.instanceID, value); err != nil {
		return fmt.Errorf("redis save value: %w", err)
	}

	now := m.now().UTC().Unix()

	if _, err := redis.DoContext(conn, ctx, "HSET", timeStampKey, m.instanceID, now); err != nil {
		return fmt.Errorf("redis save timestamp: %w", err)
	}

	return nil
}

// Get returns a metric from redis.
func (m Metric[T]) Get(ctx context.Context) (T, error) {
	var zero T
	conn := m.r.pool.Get()
	defer conn.Close()

	valueKey := fmt.Sprintf("%s-%s-values", metricKeyPrefix, m.name)
	timeStampKey := fmt.Sprintf("%s-%s-timestamp", metricKeyPrefix, m.name)

	values, err := redis.StringMap(redis.DoContext(conn, ctx, "HGETALL", valueKey))
	if err != nil {
		return zero, fmt.Errorf("redis HALS: %w", err)
	}

	timeStamps, err := redis.IntMap(redis.DoContext(conn, ctx, "HGETALL", timeStampKey))
	if err != nil {
		return zero, fmt.Errorf("redis HALS: %w", err)
	}

	tooOld := m.now().UTC().Add(-m.tooOld).Unix()

	var v T
	for k, timestamp := range timeStamps {
		if timestamp < int(tooOld) {
			continue
		}

		v, err = m.converter.Combine(values[k], v)
		if err != nil {
			return zero, fmt.Errorf("combine: %w", err)
		}
	}

	return v, nil
}
