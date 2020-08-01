package test

import (
	"encoding/json"
	"time"
)

// UpdaterMock implements the datastore.Updater interface.
//
// The received keys can be controlled by using the Send-method.
//
// The mock has to be initialized with NewUpdaterMock.
type UpdaterMock struct {
	c chan map[string]json.RawMessage
	t *time.Ticker
}

// NewUpdaterMock creates a new UpdaterMock.
func NewUpdaterMock() *UpdaterMock {
	m := UpdaterMock{}
	m.c = make(chan map[string]json.RawMessage, 1)
	m.t = time.NewTicker(time.Second)
	return &m
}

// Update returnes keys that have changed. Blocks until keys are send with
// the Send-method.
func (m *UpdaterMock) Update() (map[string]json.RawMessage, error) {
	select {
	case v := <-m.c:
		return v, nil
	case <-m.t.C:
		return nil, nil
	}
}

// Send sends keys to the mock that can be received with Update().
func (m *UpdaterMock) Send(values map[string]json.RawMessage) {
	m.c <- values
}

// Close cleans up after the Mock is used.
func (m *UpdaterMock) Close() {
	m.t.Stop()
}
