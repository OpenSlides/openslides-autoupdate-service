package test

import (
	"time"
)

// LockoutEventMock implements the datastore.Updater interface.
//
// The received keys can be controlled by using the Send-method.
//
// The mock has to be initialized with NewUpdaterMock.
type LockoutEventMock struct {
	c chan []string
	t *time.Ticker
}

// NewLockoutEventMock creates a new UpdaterMock.
func NewLockoutEventMock() *LockoutEventMock {
	m := LockoutEventMock{}
	m.c = make(chan []string, 1)
	m.t = time.NewTicker(time.Second)
	return &m
}

// LogoutEvent returnes the sessionIDs that have changed. Blocks until
// sessionIDs are send with the Send-method.
func (m *LockoutEventMock) LogoutEvent(closing <-chan struct{}) ([]string, error) {
	select {
	case v := <-m.c:
		return v, nil
	case <-m.t.C:
		return nil, nil
	case <-closing:
		return nil, closingError{}
	}
}

// Send sends sessionIDs to the mock that can be received with LogoutEvent().
func (m *LockoutEventMock) Send(values []string) {
	m.c <- values
}

// Close cleans up after the Mock is used.
func (m *LockoutEventMock) Close() {
	m.t.Stop()
}
