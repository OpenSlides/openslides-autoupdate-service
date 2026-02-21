package metric

import (
	"context"
	"encoding/json"
	"sync"
	"sync/atomic"
	"time"

	"github.com/OpenSlides/openslides-go/oslog"
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
func Loop(ctx context.Context, d time.Duration) {
	ticker := time.NewTicker(d)
	defer ticker.Stop()

	var lastSize int

	for {
		select {
		case <-ctx.Done():
			return

		case <-ticker.C:
			data := Container{make(map[string]int, lastSize)}

			callbacks.mu.Lock()
			for _, callback := range callbacks.fs {
				callback(data)
			}
			callbacks.mu.Unlock()

			lastSize = len(data.data)

			bs, err := json.Marshal(data)
			if err != nil {
				oslog.Error("Metric failed: converting data to json: %v", err)
				return
			}

			oslog.Metric(bs)
		}
	}
}

// Container is given to the callbacks for them to add the values.
type Container struct {
	data map[string]int
}

// Add adds a metric value.
func (c *Container) Add(key string, value int) {
	c.data[key] = value
}

// MarshalJSON converts the data to json.
func (c Container) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.data)
}

// CurrentCounter is a metric that shows a current value.
type CurrentCounter struct {
	name    string
	current int64
	total   uint64
}

// NewCurrentCounter initializes a current counter with a name that is used as
// prefix.
func NewCurrentCounter(name string) *CurrentCounter {
	return &CurrentCounter{name: name + "_"}
}

// Add increases the counter by one.
func (c *CurrentCounter) Add() {
	atomic.AddInt64(&c.current, 1)
	atomic.AddUint64(&c.total, 1)
}

// Done decreases the counter by one.
func (c *CurrentCounter) Done() {
	atomic.AddInt64(&c.current, -1)
}

// Metric writes the current counter.
func (c *CurrentCounter) Metric(con Container) {
	current := atomic.LoadInt64(&c.current)
	total := atomic.LoadUint64(&c.total)
	con.Add(c.name+"current", int(current))
	con.Add(c.name+"total", int(total))
}
