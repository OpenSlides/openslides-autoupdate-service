// Package redis holds the Service type, that implements the datastore.Updater
// interface of the autoupdate package by reading from a redis stream.
package redis

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

const (
	// maxMessages desides how many messages are read at once from the stream.
	maxMessages = "10"

	// fieldChangedTopic is the redis key name of the autoupdate stream.
	fieldChangedTopic = "ModifiedFields"

	// logoutTopic is the redis key name of the logout stream.
	logoutTopic = "logout"

	// lastLogoutDuration decides how many old logout messages are received.
	lastLogoutDuration = 15 * time.Minute
)

// Service holds the state of the redis receiver.
type Service struct {
	Conn             Connection
	lastAutoupdateID string
	lastLogoutID     string
}

// Update is a blocking function that returns, when there is new data.
func (s *Service) Update(closing <-chan struct{}) (map[string]json.RawMessage, error) {
	id := s.lastAutoupdateID
	if id == "" {
		id = "$"
	}

	var data map[string]json.RawMessage
	err := closingFunc(closing, func() error {
		newID, d, err := autoupdateStream(s.Conn.XREAD(maxMessages, fieldChangedTopic, id))
		if err != nil {
			return err
		}
		id = newID
		data = d
		return nil
	})

	if err != nil {
		if err == errNil {
			// No new data
			return nil, nil
		}
		return nil, fmt.Errorf("get xread data from redis: %w", err)
	}
	if id != "" {
		s.lastAutoupdateID = id
	}
	return data, nil
}

// LogoutEvent is a blocking function that returns, when a session was revoked.
func (s *Service) LogoutEvent(closing <-chan struct{}) ([]string, error) {
	id := s.lastLogoutID
	if id == "" {
		// Generate an redis ID to get the logout events from the since `lastLogoutDuration`.
		id = strconv.FormatInt(time.Now().Add(-lastLogoutDuration).Unix(), 10)
	}

	var sessionIDs []string
	err := closingFunc(closing, func() error {
		newID, sIDs, err := logoutStream(s.Conn.XREAD(maxMessages, logoutTopic, id))
		if err != nil {
			return err
		}
		id = newID
		sessionIDs = sIDs
		return nil
	})

	if err != nil {
		if err == errNil {
			// No new data
			return nil, nil
		}
		return nil, fmt.Errorf("get xread data from redis: %w", err)
	}
	if id != "" {
		s.lastLogoutID = id
	}
	return sessionIDs, nil
}

// closingFunc calls f in a separat goroutine. If closing is closed, the
// function returned, leaving the goroutine behind.
//
// The returned error is either the error returned by f() or an closingError.
func closingFunc(closing <-chan struct{}, f func() error) error {
	received := make(chan struct{})
	var err error
	go func() {
		err = f()
		close(received)
	}()

	select {
	case <-received:
		return err
	case <-closing:
		return closingError{}
	}
}
