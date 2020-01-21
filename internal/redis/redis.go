// Package redis holds the Service type, that implements the KeysChangedReceiver interface
// of the autoupdate package by reading from a redis stream.
package redis

import (
	"fmt"
)

const (
	// maxMessages desides how many messages are read at once from the stream.
	maxMessages = "10"

	// blockTimeout is the time in miliseconds, how long the xread command will block. A longer time means
	// that the service needs longer to shutdown.
	blockTimeout = "1000"

	// fieldChangedTopic is the redis key name of the stream.
	fieldChangedTopic = "field_changed"
)

// Service holds the state of the redis receiver
type Service struct {
	Conn   Connection
	lastID string
}

// KeysChanged is a blocking function that returns, when there is new data.
func (s *Service) KeysChanged() ([]string, error) {
	id := s.lastID
	if id == "" {
		id = "$"
	}
	id, keys, err := stream(s.Conn.XREAD(maxMessages, blockTimeout, fieldChangedTopic, id))
	if err != nil {
		if err == errNil {
			// No new data
			return keys, nil
		}
		return keys, fmt.Errorf("can not get data from redis: %w", err)
	}
	if id != "" {
		s.lastID = id
	}
	return keys, nil
}
