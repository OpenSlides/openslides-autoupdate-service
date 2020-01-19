package redis

import (
	"fmt"
)

const (
	maxMessages       = "10"
	blockTimeout      = "1000"
	fieldChangedTopic = "field_changed"
)

// Service holds the state of the redis receiver
type Service struct {
	Conn   Connection
	lastID string
}

// KeysChanged is a blocking function that returns, when there is new data
func (s *Service) KeysChanged() ([]string, error) {
	id := s.lastID
	if id == "" {
		id = "$"
	}
	id, kc, err := stream(s.Conn.XREAD(maxMessages, blockTimeout, fieldChangedTopic, id))
	if err != nil {
		if err == errNil {
			// No new data
			return kc, nil
		}
		return kc, fmt.Errorf("can not get data from redis: %w", err)
	}
	if id != "" {
		s.lastID = id
	}
	return kc, nil
}
