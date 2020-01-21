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
func (s *Service) Close() {
	select {
	case <-s.closed:
	default:
		close(s.closed)
	}
}

// IsClosed returns a channel that is closed when the autoupdate service is
// closed.
func (s *Service) IsClosed() <-chan struct{} {
	return s.closed
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

		s.topic.Save(keys)
	}
}

// Prepare gives the first data for a list of keysrequests and returns a connection objekt
// to pass to Echo.
func (s *Service) Prepare(ctx context.Context, uid int, krs []keysrequest.KeysRequest) (*Connection, map[string][]byte, error) {
	c := &Connection{
		user: uid,
		tid:  s.topic.LastID(),
	}

	b, err := keysbuilder.New(restrictedIDs{uid, s.restricter}, krs...)
	if err != nil {
		if errors.Is(err, keysrequest.ErrInvalid{}) {
			err = raiseErrInput(err)
		}
		return c, nil, fmt.Errorf("can not build keys: %w", err)
	}
	c.b = b

	data, err := s.restricter.Restrict(ctx, uid, b.Keys())
	if err != nil {
		return c, nil, fmt.Errorf("can not restrict data: %v", err)
	}
	return c, data, nil
}

// Echo listens for data changes and blocks until then. When data has changed,
// it returns with the new data.
// When the given context is done, it returns immediately with nil data
func (s *Service) Echo(ctx context.Context, c *Connection) (map[string][]byte, error) {
	changedKeys, tid, err := s.topic.Get(ctx, c.tid)
	if err != nil {
		return nil, fmt.Errorf("can not get new data: %w", err)
	}
	c.tid = tid

	if len(changedKeys) == 0 {
		// Exit early
		return nil, nil
	}

	oldKeys := c.b.Keys()

	// Update keysbuilder get new list of keys
	if err := c.b.Update(changedKeys); err != nil {
		return nil, fmt.Errorf("can not update keysbuilder: %w", err)
	}

	// Start with keys hat are new for the user
	keys := keysDiff(oldKeys, c.b.Keys())

	changedSlice := make(map[string]bool, len(changedKeys))
	for _, key := range changedKeys {
		changedSlice[key] = true
	}

	// Append keys that are old but have been changed.
	for _, key := range c.b.Keys() {
		if !changedSlice[key] {
			continue
		}
		keys = append(keys, key)
	}

	data, err := s.restricter.Restrict(ctx, c.user, keys)
	if err != nil {
		return nil, fmt.Errorf("can not restrict data: %v", err)
	}
	c.filter(data)
	return data, nil
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
