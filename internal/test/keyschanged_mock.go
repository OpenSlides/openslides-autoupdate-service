package test

import (
	"encoding/json"
	"fmt"
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
	return &UpdaterMock{
		c: make(chan map[string]json.RawMessage, 1),
		t: time.NewTicker(time.Second),
	}
}

// Update returnes keys that have changed. Blocks until keys are send with
// the Send-method.
func (m *UpdaterMock) Update(closing <-chan struct{}) (map[string]json.RawMessage, error) {
	select {
	case v := <-m.c:
		return v, nil
	case <-m.t.C:
		return nil, fmt.Errorf("Time is up")
	case <-closing:
		return nil, closingError{}
	}
}

// Send sends keys to the mock that can be received with Update().
func (m *UpdaterMock) Send(values map[string]string) {
	conv := make(map[string]json.RawMessage)
	for k, v := range values {
		conv[k] = []byte(v)
	}
	m.c <- conv
}

// Close cleans up after the Mock is used.
func (m *UpdaterMock) Close() {
	m.t.Stop()
}

type closingError struct{}

func (e closingError) Closing()      {}
func (e closingError) Error() string { return "closing" }
