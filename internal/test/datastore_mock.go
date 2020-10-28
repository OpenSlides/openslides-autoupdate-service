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
	changeListeners []func(map[string]json.RawMessage) error
	DatastoreValues
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
		value, _, err := d.DatastoreValues.Value(key)
		if err != nil {
			return nil, err
		}

		data[key] = value
	}

	values := make([]json.RawMessage, len(keys))
	for i, key := range keys {
		values[i] = data[key]
	}
	return values, nil
}

// RegisterChangeListener registers a change listener.
func (d *MockDatastore) RegisterChangeListener(f func(map[string]json.RawMessage) error) {
	d.changeListeners = append(d.changeListeners, f)
}

// Send sends keys to the mock that can be received with a listerer registered
// by RegisterChangeListener().
func (d *MockDatastore) Send(keys []string) {
	data := make(map[string]json.RawMessage, len(keys))
	for _, key := range keys {
		data[key] = nil
	}

	for _, f := range d.changeListeners {
		f(data)
	}
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
	case strings.Contains(key, "$_"):
		return json.RawMessage(`"1","2"`), true, nil
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
// This does not send a signal to the listeners.
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
