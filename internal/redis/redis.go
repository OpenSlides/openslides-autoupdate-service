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

	// blockTimeout is the time in miliseconds, how long the xread command will
	// block.
	blockTimeout = "3600000" // One Hour

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
func (s *Service) Update() (map[string]json.RawMessage, error) {
	id := s.lastAutoupdateID
	if id == "" {
		id = "$"
	}

	id, keys, err := autoupdateStream(s.Conn.XREAD(maxMessages, blockTimeout, fieldChangedTopic, id))
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
	return keys, nil
}

// LogoutEvent is a blocking function that returns, when a session was revoked.
func (s *Service) LogoutEvent() ([]string, error) {
	id := s.lastLogoutID
	if id == "" {
		// Generate an redis ID to get the logout events from the since `lastLogoutDuration`.
		id = strconv.FormatInt(time.Now().Add(-lastLogoutDuration).Unix(), 10)
	}

	id, sessionIDs, err := logoutStream(s.Conn.XREAD(maxMessages, blockTimeout, logoutTopic, id))
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
