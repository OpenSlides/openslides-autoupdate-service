package autoupdate

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/openslides/openslides-autoupdate-service/internal/keysbuilder"
	"github.com/openslides/openslides-autoupdate-service/internal/keysrequest"
	"github.com/openslides/openslides-autoupdate-service/internal/topic"
)

// Service holds the state of the autoupdate service
type Service struct {
	restricter keysbuilder.Restricter
	keyChanged KeysChangedReceiver
	closed     chan struct{}
	topic      topic.Topic
}

// New creates a new autoupdate service
func New(restricter keysbuilder.Restricter, keyChanges KeysChangedReceiver) *Service {
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

// Close calls the shutdown logic of the service
func (s *Service) Close() {
	select {
	case <-s.closed:
	default:
		close(s.closed)
	}
}

func (s *Service) pruneTopic() {
	tick := time.NewTicker(time.Second)
	defer tick.Stop()
	for {
		select {
		case <-s.closed:
			return
		case <-tick.C:
			s.topic.Prune(time.Now().Add(-10 * time.Second))
		}
	}
}

func (s *Service) receiveKeyChanges() {
	for {
		// Test if the service has been closed
		select {
		case <-s.closed:
			return
		default:
		}

		kc, err := s.keyChanged.KeysChanged()
		if err != nil {
			log.Printf("TODO: %v", err)
		}
		if len(kc.Updated) == 0 {
			continue
		}

		s.topic.Save(kc.Updated)
	}
}

func (s *Service) prepare(ctx context.Context, uid int, kr keysrequest.KeysRequest) (uint64, *keysbuilder.Builder, map[string][]byte, error) {
	b, err := keysbuilder.New(uid, s.restricter, kr)
	if err != nil {
		if errors.Is(err, keysrequest.ErrInvalid{}) {
			err = raise400(err)
		}
		return 0, nil, nil, fmt.Errorf("can not build keys: %w", err)
	}

	data, err := s.restricter.Restrict(ctx, uid, b.Keys())
	if err != nil {
		return 0, nil, nil, fmt.Errorf("can not restrict data: %v", err)
	}
	return s.topic.LastID(), b, data, nil
}

func (s *Service) echo(ctx context.Context, uid int, tid uint64, b *keysbuilder.Builder) (uint64, map[string][]byte, error) {
	changedKeys, tid, err := s.topic.Get(ctx, tid)
	if err != nil {
		return 0, nil, fmt.Errorf("can not get new data: %w", err)
	}
	oldKeys := b.Keys()
	if err := b.Update(changedKeys); err != nil {
		return 0, nil, fmt.Errorf("can not update keybuilder: %w", err)
	}
	keys := keysDiff(oldKeys, b.Keys())

	changedSlice := make(map[string]bool, len(changedKeys))
	for _, key := range changedKeys {
		changedSlice[key] = true
	}

	for _, key := range b.Keys() {
		if !changedSlice[key] {
			continue
		}
		keys = append(keys, key)
	}

	data, err := s.restricter.Restrict(ctx, uid, keys)
	if err != nil {
		return 0, nil, fmt.Errorf("can not restrict data: %v", err)
	}
	return tid, data, nil
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
