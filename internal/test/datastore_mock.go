package test

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

// MockDatastore implements the autoupdate.Datastore interface.
type MockDatastore struct {
	changes chan []string
	done    chan struct{}
	DatastoreValues
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
func (d *MockDatastore) Get(ctx context.Context, keys ...string) ([]json.RawMessage, error) {
	data := make(map[string]json.RawMessage, len(keys))
	for _, key := range keys {
		value, exist, err := d.DatastoreValues.Value(key)
		if err != nil {
			return nil, err
		}

		if !exist {
			continue
		}

		data[key] = value
	}

	values := make([]json.RawMessage, len(keys))
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

// DatastoreValues returns data for the test.MockDatastore and the test.DatastoreServer.
type DatastoreValues struct {
	mu       sync.RWMutex
	Data     map[string]json.RawMessage
	OnlyData bool
}

// Value returns a value for a key. If the value does not exist, the second
// return value is false.
func (d *DatastoreValues) Value(key string) (json.RawMessage, bool, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	v, ok := d.Data[key]
	if ok {
		return v, true, nil
	}

	if d.OnlyData {
		return nil, false, nil
	}

	switch {
	case strings.HasPrefix(key, "error"):
		return nil, true, fmt.Errorf("mock datastore error")
	case strings.HasSuffix(key, "_id"):
		return json.RawMessage(`1`), true, nil
	case strings.HasSuffix(key, "_ids"):
		return json.RawMessage(`[1,2]`), true, nil
	default:
		return json.RawMessage(`"Hello World"`), true, nil
	}
}

// Update updates the values from the Datastore.
//
// This does not send a KeysChanged signal.
func (d *DatastoreValues) Update(data map[string]json.RawMessage) {
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
