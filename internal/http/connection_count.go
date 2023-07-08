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

func newConnectionCount(ctx context.Context, r *redis.Redis, saveInterval time.Duration) *connectionCount {
	redisMetric := redis.NewMetric[map[int]int](r, "autoupdate_connection_count", mapIntCombiner{}, saveInterval*2, time.Now)

	c := connectionCount{
		metric:      redisMetric,
		connections: make(map[int]int),
	}

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

	return &c
}

func (c *connectionCount) save(ctx context.Context) error {
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

func (c *connectionCount) increment(uid int, increment int) {
	c.mu.Lock()
	c.connections[uid] += increment
	c.mu.Unlock()
}

func (c *connectionCount) Add(uid int) {
	c.increment(uid, 1)
}

func (c *connectionCount) Done(uid int) {
	c.increment(uid, -1)
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
		return
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

	prefix := "connected_users"
	con.Add(prefix+"_current", currentConnectedUsers)
	con.Add(prefix+"_total", len(data))
	con.Add(prefix+"_current_local", localCurrentUsers)
	con.Add(prefix+"_total_local", totalCurrentConnections)
	con.Add(prefix+"_average_connections", average)
	con.Add(prefix+"_anonymous_connections", data[0])

	prefix = "current_connections"
	con.Add(prefix, currentConnections)
	con.Add(prefix+"_local", localCurrentConnections)
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
