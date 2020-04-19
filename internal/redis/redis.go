// Package redis holds the Service type, that implements the KeysChangedReceiver
// interface of the autoupdate package by reading from a redis stream.
package redis

import (
	"fmt"
	"time"

	"github.com/mediocregopher/radix/v3"
)

const (
	// maxMessages desides how many messages are read at once from the stream.
	// The value 0 means all messages at once.
	maxMessages = 10

	// blockTimeout defines how long the XREAD call to redis blocks.
	blockTimeout = time.Hour

	// fieldChangedTopic is the redis key name of the stream.
	fieldChangedTopic = "ModifiedFields"
)

// NewClient returns a new redis client that can be used in the redis.Updater
// struct. It tests, that the connection can be established.
func NewClient(addr string) (*radix.Pool, error) {
	// Create a radix pool without implicit pipelining and with unlimiting read
	// timeout. This is important to call the blocking XREAD redis command.
	return radix.NewPool(
		"tcp",
		addr,
		1,
		radix.PoolPipelineWindow(0, 0), // Disable implicit pipelining
		radix.PoolConnFunc(func(network, addr string) (radix.Conn, error) {
			// Disable read timeout.
			return radix.Dial(network, addr, radix.DialReadTimeout(0))
		}),
	)

}

// Updater holds the state of the redis receiver.
type Updater struct {
	Client radix.Client
	sr     radix.StreamReader
}

// KeysChanged is a blocking function that returns, when there is new data.
func (s *Updater) KeysChanged() ([]string, error) {
	if s.sr == nil {
		s.sr = radix.NewStreamReader(s.Client, radix.StreamReaderOpts{
			Streams: map[string]*radix.StreamEntryID{fieldChangedTopic: nil},
			Block:   blockTimeout,
			Count:   maxMessages,
		})
	}

	stream, entires, ok := s.sr.Next()
	if ok == false {
		return nil, s.sr.Err()
	}

	if stream != fieldChangedTopic {
		return nil, fmt.Errorf("got value from stream %s. Only values from %s are requested", stream, fieldChangedTopic)
	}

	var values []string
	for _, entry := range entires {
		for k, v := range entry.Fields {
			if v != "modified" {
				// Only support modified keys.
				continue
			}
			values = append(values, k)
		}
	}

	return values, nil
}
