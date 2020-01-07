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
	go s.receiveKeyChanges()
	go s.pruneTopic()
	return s
}

// Close calls the shutdown logic of the service
func (s *Service) Close() {
	close(s.closed)
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

func (s *Service) prepare(ctx context.Context, uid int, kr keysrequest.KeysRequest) (uint64, []string, map[string][]byte, error) {
	b := keysbuilder.Builder{
		User:    uid,
		Restr:   s.restricter,
		Request: kr,
	}
	keys, err := b.Keys()
	if err != nil {
		if errors.Is(err, keysrequest.ErrInvalid{}) {
			err = new400(err)
		}
		return 0, nil, nil, fmt.Errorf("can not build keys: %w", err)
	}
	if len(keys) == 0 {
		return 0, nil, nil, new400(fmt.Errorf("No keys requested"))
	}
	tid := s.topic.LastID()
	data, err := s.restricter.Restrict(ctx, uid, keys)
	if err != nil {
		return 0, nil, nil, fmt.Errorf("can not restrict data: %v", err)
	}
	return tid, keys, data, nil
}

func (s *Service) echo(ctx context.Context, uid int, tid uint64, keys []string) (uint64, map[string][]byte, error) {
	keysSlice := make(map[string]bool, len(keys))
	for _, key := range keys {
		keysSlice[key] = true
	}

	newKeys, tid, err := s.topic.Get(tid)
	if err != nil {
		return 0, nil, fmt.Errorf("can not get new data: %w", err)
	}
	keys = []string{}
	for _, key := range newKeys {
		if !keysSlice[key] {
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
