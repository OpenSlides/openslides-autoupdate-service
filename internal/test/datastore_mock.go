package test

import (
	"context"
	"fmt"
	"strings"
	"sync"
)

// MockDatastore implements the autoupdate.Datastore interface.
type MockDatastore struct {
	mu       sync.Mutex
	Data     map[string]string
	OnlyData bool

	changes chan []string
	done    chan struct{}
}

// NewMockDatastore returns a new MockDatastore.
func NewMockDatastore() *MockDatastore {
	return &MockDatastore{
		changes: make(chan []string),
		done:    make(chan struct{}),
	}
}

// Get returnes the values for the given keys. If the keys exist in the Data
// attribute, the values are returned.
//
// If a key is not present in Data, a default value is returned.
//
// If the key starts with "error", an error it thrown.
//
// If the key ends with "_id", "1" is returned.
//
// If the key ends with "_ids", "[1,2]" is returned.
//
// In any other case, "some value" is returned.
func (d *MockDatastore) Get(ctx context.Context, keys ...string) ([]string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	data := make(map[string]string, len(keys))
	for _, key := range keys {
		v, ok := d.Data[key]
		if ok {
			data[key] = v
			continue
		}

		if d.OnlyData {
			continue
		}
		var value string
		switch {
		case strings.HasPrefix(key, "error"):
			return nil, fmt.Errorf("mock datastore error")
		case strings.HasSuffix(key, "_id"):
			value = `1`
		case strings.HasSuffix(key, "_ids"):
			value = `[1,2]`
		default:
			value = `"Hello World"`
		}
		data[key] = value
	}

	values := make([]string, len(keys))
	for i, key := range keys {
		values[i] = data[key]
	}
	return values, nil
}

// KeysChanged returnes keys that have changed. Blocks until keys are send with
// the Send-method.
func (d *MockDatastore) KeysChanged() ([]string, error) {
	select {
	case v := <-d.changes:
		return v, nil
	case <-d.done:
		return nil, nil
	}
}

// Send sends keys to the mock that can be received with KeysChanged().
func (d *MockDatastore) Send(keys []string) {
	d.changes <- keys
}

// Close cleans up after the Mock is used.
func (d *MockDatastore) Close() {
	close(d.done)
}

// Update updates the values from the Datastore.
//
// This does not send a KeysChanged signal.
func (d *MockDatastore) Update(data map[string]string) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.Data == nil {
		d.Data = data
		return
	}

	for key, value := range data {
		d.Data[key] = value
	}
}
