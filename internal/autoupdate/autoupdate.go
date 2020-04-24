// Package autoupdate allows clients to request keys and get updates when the
// keys changes.
//
// To register to the autoupdate serive, a client has to receive a Connection
// object by calling the Connect()-method. It is not necessary and therefore not
// possible to close a connection. The client can just stop listening.
package autoupdate

import (
	"context"
	"log"
	"time"

	"github.com/ostcar/topic"
)

// pruneTime defines how long a topic id will be valid. If a client needs more
// time to process the data, it will get an error and has to reconnect. A higher
// value means, that more memory is used.
const pruneTime = time.Minute

// Service holds the state of the autoupdate service. It has to be initialized
// with autoupdate.New().
//
// The service updates its data in the background. To stop this background job,
// the service has to be closed in the end with the Close()-method.
type Service struct {
	restricter Restricter
	keyChanged KeysChangedReceiver
	closed     chan struct{}
	topic      *topic.Topic
}

// New creates a new autoupdate service.
//
// After the service is not needed anymore, it has to be closed with s.Close().
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

// Close calls the shutdown logic of the service. This method is not save for
// concourent use. It not allowed to call it more then once. Only the caller of
// New() should call Close().
func (s *Service) Close() {
	close(s.closed)
}

// Connect has to be called by a client to register to the service. The method
// returns a Connection object, that can be used to receive the data.
//
// There is no need to "close" the Connection object.
func (s *Service) Connect(ctx context.Context, userID int, kb KeysBuilder) *Connection {
	return &Connection{
		autoupdate: s,
		ctx:        ctx,
		uid:        userID,
		kb:         kb,
	}
}

// IDer returns an object, that implements the keysbuilder.IDer interface. It is
// used to return ids for a key. This implementation uses the restricter to get
// the ids.
func (s *Service) IDer(uid int) RestrictedIDs {
	return RestrictedIDs{uid, s.restricter}
}

// pruneTopic removes old data from the topic. Blocks until the service is
// closed.
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

// receiveKeyChanges listens for updates and saves then into the topic. This
// function blocks until the service is closed.
func (s *Service) receiveKeyChanges() {
	for {
		select {
		case <-s.closed:
			return
		default:
		}

		keys, err := s.keyChanged.KeysChanged()
		if err != nil {
			log.Printf("Could not update keys: %v\n", err)
			continue
		}

		s.topic.Publish(keys...)
	}
}
