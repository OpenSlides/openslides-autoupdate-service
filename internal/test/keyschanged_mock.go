package test

import "time"

// MockKeysChanged implements the datastore.KeysChangedReceiver interface.
//
// The received keys can be controlled by using the Send-method.
//
// The mock has to be initialized with NewMockKeysChanged and be closed at the
// end with the Close method.
type MockKeysChanged struct {
	c chan []string
	t *time.Ticker
}

// NewMockKeysChanged creates a new KeysChanged mock.
func NewMockKeysChanged() *MockKeysChanged {
	m := MockKeysChanged{}
	m.c = make(chan []string, 1)
	m.t = time.NewTicker(time.Second)
	return &m
}

// KeysChanged returnes keys that have changed. Blocks until keys are send with
// the Send-method.
func (m *MockKeysChanged) KeysChanged() ([]string, error) {
	select {
	case v := <-m.c:
		return v, nil
	case <-m.t.C:
		return nil, nil
	}
}

// Send sends keys to the mock that can be received with KeysChanged().
func (m *MockKeysChanged) Send(keys []string) {
	m.c <- keys
}

// Close cleans up after the Mock is used.
func (m *MockKeysChanged) Close() {
	m.t.Stop()
}
