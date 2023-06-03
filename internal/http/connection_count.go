package http

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/metric"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/oserror"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/redis"
)

// RedisMetric is used for the connection counter.
type RedisMetric interface {
	Save(ctx context.Context, v map[int]int) error
	Get(ctx context.Context) (map[int]int, error)
}

// InitConnectionCountMetric initializes a redis metric to use for the
// connection count.
func InitConnectionCountMetric(r *redis.Redis, tooOld time.Duration) RedisMetric {
	// TODO: Save metric regularry.
	return redis.NewMetric[map[int]int](r, "autoupdate_connection_count", mapIntConverter{}, tooOld, time.Now)
}

type mapIntConverter struct{}

func (mapIntConverter) Convert(v map[int]int) (string, error) {
	converted, err := json.Marshal(v)
	return string(converted), err
}

func (mapIntConverter) Combine(rawValue string, acc map[int]int) (map[int]int, error) {
	var value map[int]int
	if err := json.Unmarshal([]byte(rawValue), &value); err != nil {
		return nil, err
	}

	if acc == nil {
		acc = make(map[int]int)
	}

	for k, v := range value {
		acc[k] += v
	}

	return acc, nil
}

type connectionCount struct {
	mu          sync.Mutex
	metric      RedisMetric
	connections map[int]int
}

func newConnectionCount(metric RedisMetric) *connectionCount {
	return &connectionCount{
		metric:      metric,
		connections: make(map[int]int),
	}
}

func (c *connectionCount) Add(ctx context.Context, uid int) {
	c.mu.Lock()
	c.connections[uid]++
	c.mu.Unlock()

	if err := c.metric.Save(ctx, c.connections); err != nil {
		oserror.Handle(fmt.Errorf("save connection count in redis: %w", err))
	}
}

func (c *connectionCount) Done(uid int) {
	ctx := context.Background()

	c.mu.Lock()
	c.connections[uid]--
	c.mu.Unlock()

	if err := c.metric.Save(ctx, c.connections); err != nil {
		oserror.Handle(fmt.Errorf("save connection count in redis: %w", err))
	}
}

func (c *connectionCount) Show(ctx context.Context) (map[int]int, error) {
	data, err := c.metric.Get(ctx)
	if err != nil {
		oserror.Handle(fmt.Errorf("fetch connection count metric from redis: %w", err))
	}

	return data, nil
}

func (c *connectionCount) Metric(con metric.Container) {
	ctx := context.Background()

	data, err := c.metric.Get(ctx)
	if err != nil {
		oserror.Handle(fmt.Errorf("fetch connection count metric from redis: %w", err))
	}

	currentConnections := 0
	averageCount := 0
	averageSum := 0
	for k, v := range data {
		if v <= 0 {
			continue
		}

		currentConnections++

		if k != 0 {
			averageCount++
			averageSum += v
		}
	}

	average := 0
	if averageCount > 0 {
		average = averageSum / averageCount
	}

	prefix := "connected_users_"
	con.Add(prefix+"current", currentConnections)
	con.Add(prefix+"total", len(data))
	con.Add(prefix+"average", average)
	con.Add(prefix+"anonymous", data[0])
}
