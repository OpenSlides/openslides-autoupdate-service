// Package autoupdate holds the logik to register the keysrequests of clients and inform them
// about updates.
package autoupdate

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate/keysbuilder"
	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate/keysrequest"
	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate/topic"
)

// pruneTime defines how long a topic id will be valid. If a client needs more time to process
// the data, it will get an error and has to reconnect.
const pruneTime = 10 * time.Second

// Service holds the state of the autoupdate service.
type Service struct {
	restricter Restricter
	keyChanged KeysChangedReceiver
	closed     chan struct{}
	topic      topic.Topic
}

// New creates a new autoupdate service.
func New(restricter Restricter, keyChanges KeysChangedReceiver) *Service {
	s := &Service{
		restricter: restricter,
		keyChanged: keyChanges,
		closed:     make(chan struct{}),
	}
	s.topic = topic.Topic{Closed: s.closed}
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

// Done returns a channel that is closed when the autoupdate service is
// closed.
func (s *Service) Done() <-chan struct{} {
	return s.closed
}

// Connect gives the first data for a list of keysrequests and returns a connection objekt
// to pass to Echo.
func (s *Service) Connect(ctx context.Context, uid int, krs []keysrequest.Body) (*Connection, map[string]string, error) {
	b, err := keysbuilder.New(restrictedIDs{uid, s.restricter}, krs...)
	if err != nil {
		if errors.Is(err, keysrequest.ErrInvalid{}) {
			err = raiseErrInput(err)
		}
		return nil, nil, fmt.Errorf("can not build keys: %w", err)
	}

	lastID := s.topic.LastID()

	data, err := s.restricter.Restrict(ctx, uid, b.Keys())
	if err != nil {
		return nil, nil, fmt.Errorf("can not restrict data: %v", err)
	}

	c := &Connection{
		s:    s,
		ctx:  ctx,
		user: uid,
		tid:  lastID,
		b:    b,
	}
	c.filter(data)
	return c, data, nil
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

		s.topic.Add(keys)
	}
}

func keysDiff(old []string, new []string) []string {
	slice := make(map[string]bool, len(old))
	for _, key := range old {
		slice[key] = true
	}
	added := []string{}
	for _, key := range new {
		if slice[key] {
			continue
		}
		added = append(added, key)
	}
	return added
}
