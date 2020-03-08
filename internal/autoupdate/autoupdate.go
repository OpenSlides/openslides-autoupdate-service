// Package autoupdate allows clients to request keys and get updates when the keys changes.
package autoupdate

import (
	"context"
	"log"
	"time"

	"github.com/ostcar/topic"
)

// pruneTime defines how long a topic id will be valid. If a client needs more time to process
// the data, it will get an error and has to reconnect.
const pruneTime = 10 * time.Second

// Service holds the state of the autoupdate service. It has to be initialized with autoupdate.New().
type Service struct {
	restricter Restricter
	keyChanged KeysChangedReceiver
	closed     chan struct{}
	topic      *topic.Topic
}

// New creates a new autoupdate service.
//
// After the service is not needed anymore, it has to be closed with
// s.Close().
func New(restricter Restricter, keysChanges KeysChangedReceiver) *Service {
	s := &Service{
		restricter: restricter,
		keyChanged: keysChanges,
		closed:     make(chan struct{}),
	}
	s.topic = topic.New(topic.WithClosed(s.closed))
	go s.receiveKeyChanges()
	go s.pruneTopic()
	return s
}

// Close calls the shutdown logic of the service.
// This method is not save for concourent use. It not allowed to call
// it more then once. Onle the caller of New() should call Close().
func (s *Service) Close() {
	close(s.closed)
}

// Connect returns a Connection object
func (s *Service) Connect(ctx context.Context, uid int, kb KeysBuilder) *Connection {
	return &Connection{
		autoupdate: s,
		ctx:        ctx,
		user:       uid,
		kb:         kb,
	}
}

// IDer returns an object, that can be used to return ids for an related key.
func (s *Service) IDer(uid int) RestrictedIDs {
	return RestrictedIDs{uid, s.restricter}
}

// pruneTopic removes old data from the topic.
// Blocks until the service is closed.
func (s *Service) pruneTopic() {
	tick := time.NewTicker(time.Second)
	defer tick.Stop()
	for {
		select {
		case <-s.closed:
			return
		case <-tick.C:
			s.topic.Prune(time.Now().Add(-pruneTime))
		}
	}
}

// receiveKeyChanges listens for updates and saves then into the topic.
// Blocks until the service is closed.
func (s *Service) receiveKeyChanges() {
	for {
		select {
		case <-s.closed:
			return
		default:
		}

		keys, err := s.keyChanged.KeysChanged()
		if err != nil {
			log.Printf("Could not update keys: %v", err)
			continue
		}

		if len(keys) == 0 {
			continue
		}

		s.topic.Add(keys...)
	}
}
