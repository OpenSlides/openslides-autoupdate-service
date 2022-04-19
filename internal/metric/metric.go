package metric

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"
)

var callbacks struct {
	mu sync.Mutex
	fs []func(Container)
}

// Register registers a function that is called when metrics are gathered.
func Register(f func(Container)) {
	callbacks.mu.Lock()
	callbacks.fs = append(callbacks.fs, f)
	callbacks.mu.Unlock()
}

// Loop gathers the metric data from all registered callbacks.
//
// Blocks until the context is done.
//
// It is not possible to Register new metrics, when the loop is running.
func Loop(ctx context.Context, d time.Duration, logger log.Logger) {
	callbacks.mu.Lock()
	defer callbacks.mu.Unlock()

	ticker := time.NewTicker(d)
	defer ticker.Stop()

	var lastSize int

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			data := Container{make(map[string]any, lastSize)}
			for _, callback := range callbacks.fs {
				callback(data)
			}
			lastSize = len(data.data)

			bs, err := json.Marshal(data)
			if err != nil {
				logger.Printf("Metric failed: converting data to json: %v", err)
				return
			}

			logger.Printf("Metric: %s", bs)
		}
	}
}

// Container is given to the callbacks for them to add the values.
type Container struct {
	data map[string]any
}

// Add adds a metric value.
func (c *Container) Add(key string, value any) {
	c.data[key] = value
}

// MarshalJSON converts the data to json.
func (c Container) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.data)
}

// Sub creates a sub container for adding values under one name.
func (c *Container) Sub(name string) Container {
	nc := Container{make(map[string]any)}
	c.data[name] = nc
	return nc
}

// CurrentCounter is a metric that shows a current value.
type CurrentCounter struct {
	current int
	total   uint64
}

// Add increases the counter by one.
func (c *CurrentCounter) Add() {
	c.current++
	c.total++
}

// Done decreases the counter by one.
func (c *CurrentCounter) Done() {
	c.current--
}

// Metric writes the current counter.
func (c *CurrentCounter) Metric(con Container) {
	s := con.Sub("connection")
	s.Add("current", c.current)
	s.Add("total", c.total)
}
