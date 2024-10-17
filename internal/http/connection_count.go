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
	Save(ctx context.Context, v string) error
	Get(ctx context.Context) (map[int]int, error)
}

// ConnectionCount counts, how many connections a user has.
//
// It holds a local counter and saves it to redis from time to time. The
// argument `saveInterval` defines, how oftem it is saved.
//
// It also pings redis from time to time to show, that this instance is still
// running.
type ConnectionCount struct {
	metric RedisMetric
	name   string

	mu          sync.Mutex
	connections map[int]int
}

// newConnectionCount creates a new storage for metrics.
//
// If r is set, it saves the metric to redis. If r is nil, then no data is saved to redis.
func newConnectionCount(ctx context.Context, r *redis.Redis, saveInterval time.Duration, name string) *ConnectionCount {
	var redisMetric RedisMetric
	if r != nil {
		redisMetric = redis.NewMetric(r, name, mapIntCombiner{}, saveInterval*2, time.Now)
	}

	c := ConnectionCount{
		metric:      redisMetric,
		name:        name,
		connections: make(map[int]int),
	}

	if redisMetric != nil {
		go func() {
			tick := time.NewTicker(saveInterval)
			defer tick.Stop()

			for {
				select {
				case <-ctx.Done():
					return
				case <-tick.C:
				}

				if err := c.save(ctx); err != nil {
					oserror.Handle(fmt.Errorf("Error: save connection count: %w", err))
				}
			}
		}()
	}

	return &c
}

func (c *ConnectionCount) save(ctx context.Context) error {
	c.mu.Lock()
	converted, err := json.Marshal(c.connections)
	c.mu.Unlock()
	if err != nil {
		return fmt.Errorf("convert connection count to json: %w", err)
	}

	if err := c.metric.Save(ctx, string(converted)); err != nil {
		return fmt.Errorf("save connection count in redis: %w", err)
	}

	return nil
}

func (c *ConnectionCount) increment(uid int, increment int) {
	c.mu.Lock()
	c.connections[uid] += increment
	c.mu.Unlock()
}

// Add adds one connection to the counter.
func (c *ConnectionCount) Add(uid int) {
	c.increment(uid, 1)
}

// Done removes one connection from the counter.
func (c *ConnectionCount) Done(uid int) {
	c.increment(uid, -1)
}

// Show shows the counter.
//
// if a redis connection is set, the data are fetched from redis. In other case,
// only the local data is returned.
func (c *ConnectionCount) Show(ctx context.Context, filter func(ctx context.Context, count map[int]int) error) (map[int]int, error) {
	var data map[int]int
	if c.metric == nil {
		c.mu.Lock()
		data = c.connections
		c.mu.Unlock()
	} else {
		var err error
		data, err = c.metric.Get(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting counter from redis: %w", err)
		}
	}

	if err := filter(ctx, data); err != nil {
		return nil, fmt.Errorf("filtering counter: %w", err)
	}

	return data, nil
}

// Metric is a function needed my the openslides metric system to fetch some values.
func (c *ConnectionCount) Metric(con metric.Container) {
	ctx := context.Background()

	if c.metric != nil {
		data, err := c.metric.Get(ctx)
		if err != nil {
			oserror.Handle(fmt.Errorf("fetch connection count metric from redis: %w", err))
			return
		}

		currentConnectedUsers := 0
		currentConnections := 0
		averageCount := 0
		averageSum := 0
		for k, v := range data {
			if v <= 0 {
				continue
			}

			currentConnectedUsers++
			currentConnections += v

			if k != 0 {
				averageCount++
				averageSum += v
			}
		}

		average := 0
		if averageCount > 0 {
			average = averageSum / averageCount
		}

		con.Add(c.name+"_connected_users_current", currentConnectedUsers)
		con.Add(c.name+"_connected_users_total", len(data))
		con.Add(c.name+"_connected_users_average_connections", average)
		con.Add(c.name+"_connections_public_access", data[0])
		con.Add(c.name+"_current_connections", currentConnections)
	}

	localCurrentUsers := 0
	localCurrentConnections := 0
	c.mu.Lock()
	totalCurrentConnections := len(c.connections)
	for _, v := range c.connections {
		if v <= 0 {
			continue
		}

		localCurrentUsers++
		localCurrentConnections += v
	}
	c.mu.Unlock()

	con.Add(c.name+"_connected_users_current_local", localCurrentUsers)
	con.Add(c.name+"_connected_users_total_local", totalCurrentConnections)
	con.Add(c.name+"_current_connections_local", localCurrentConnections)
}

// mapIntCombiner tells the redis Metric, how to combine the metric values.
type mapIntCombiner struct{}

func (mapIntCombiner) Combine(rawValue string, acc map[int]int) (map[int]int, error) {
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
