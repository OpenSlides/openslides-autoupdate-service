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

// connectionCount counts, how many connections a user has.
//
// It holds a local counter and saves it to redis after a connection is created
// or closed.
//
// It also pings redis from time to time to show, that this instance is
// still running.
type connectionCount struct {
	metric RedisMetric

	mu          sync.Mutex
	connections map[int]int
}

func newConnectionCount(r *redis.Redis, tooOld time.Duration) *connectionCount {
	redisMetric := redis.NewMetric[map[int]int](r, "autoupdate_connection_count", mapIntConverter{}, tooOld, time.Now)

	go func() {
		ctx := context.Background()
		for {
			time.Sleep(tooOld / 2)
			if err := redisMetric.KeepAlive(ctx); err != nil {
				oserror.Handle(fmt.Errorf("Send keep alive to redis: %w", err))
			}

		}
	}()

	return &connectionCount{
		metric:      redisMetric,
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
	// Done is callled when the connection is closed. The the context from the
	// request can not be used.
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
		return nil, fmt.Errorf("getting counter from redis: %w", err)
	}

	return data, nil
}

// Metric is a function needed my the openslides metric system to fetch some values.
func (c *connectionCount) Metric(con metric.Container) {
	ctx := context.Background()

	data, err := c.metric.Get(ctx)
	if err != nil {
		oserror.Handle(fmt.Errorf("fetch connection count metric from redis: %w", err))
	}

	localCurrentConnections := 0
	c.mu.Lock()
	totalCurrentConnections := len(c.connections)
	for _, v := range c.connections {
		if v <= 0 {
			continue
		}

		localCurrentConnections++
	}
	c.mu.Unlock()

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
	con.Add(prefix+"current_local", localCurrentConnections)
	con.Add(prefix+"total_local", totalCurrentConnections)
	con.Add(prefix+"average_connections", average)
	con.Add(prefix+"anonymous_connections", data[0])
}

// mapIntConverter tells the redis Metric how to convert the map[int]int to a
// value that can be saved by redis.
type mapIntConverter struct{}

func (mapIntConverter) Convert(v map[int]int) (string, error) {
	converted, err := json.Marshal(v)
	return string(converted), err
}

func (mapIntConverter) Combine(rawValue string, acc map[int]int) (map[int]int, error) {
	if rawValue == "" {
		return acc, nil
	}

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
