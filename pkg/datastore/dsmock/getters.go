package dsmock

import (
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

// Getter is the same interface as Getter
type Getter interface {
	Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error)
}

// Stub are data that can be used as a datastore value.
type Stub map[dskey.Key][]byte

// Get implements the Getter interface.
func (s Stub) Get(_ context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	data := map[dskey.Key][]byte(s)
	requested := make(map[dskey.Key][]byte, len(keys))
	for _, k := range keys {
		requested[k] = data[k]
	}
	return requested, nil
}

// Update blocks until the context is done.
//
// It is here to implement the flow.Updater interface.
func (s Stub) Update(ctx context.Context, updateFn func(map[dskey.Key][]byte, error)) {
	<-ctx.Done()
}

// StubWithUpdate is like Stub but the values can be changed via the Send
// method.
//
// It implements the datastore.Source interface.
type StubWithUpdate struct {
	mu sync.RWMutex

	stub Stub
	ch   chan map[dskey.Key][]byte

	getter Getter

	middlewares []Getter
}

// NewStubWithUpdate initializes a the object.
func NewStubWithUpdate(stub Stub, middlewares ...func(Getter) Getter) *StubWithUpdate {
	getter := Getter(stub)
	initialized := make([]Getter, len(middlewares))
	for i, m := range middlewares {
		getter = m(getter)
		initialized[i] = getter
	}

	return &StubWithUpdate{
		stub:        stub,
		ch:          make(chan map[dskey.Key][]byte),
		getter:      getter,
		middlewares: initialized,
	}
}

// Get returns the current value for the given keys.
func (s *StubWithUpdate) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.getter.Get(ctx, keys...)
}

// Update blocks until new data is received via the Send method.
func (s *StubWithUpdate) Update(ctx context.Context) (map[dskey.Key][]byte, error) {
	select {
	case newValues := <-s.ch:
		s.mu.Lock()
		for k, v := range newValues {
			s.stub[k] = v
		}
		s.mu.Unlock()
		return newValues, nil

	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// Send sends keys to the mock that can be received with Update().
func (s *StubWithUpdate) Send(values map[dskey.Key][]byte) {
	s.ch <- values
}

// Middlewares returns a list of Getters that where used in
// NewStubWithUpdates.
//
// For example:
//
//	ds := NewStubWithUpdate(stub, dsmock.NewCounter)
//	counter := ds.Middlewares()[0].(*dsmock.Counter)
func (s *StubWithUpdate) Middlewares() []Getter {
	return s.middlewares
}

// HistoryInformation writes a fake history.
func (s *StubWithUpdate) HistoryInformation(ctx context.Context, fqid string, w io.Writer) error {
	w.Write([]byte(`[{"position":42,"user_id": 5,"information": "motion was created","timestamp: 1234567}]`))
	return nil
}

// Counter counts all keys that where requested.
type Counter struct {
	mu sync.Mutex

	ds       Getter
	requests [][]dskey.Key
}

// NewCounter initializes a Counter.
func NewCounter(ds Getter) Getter {
	return &Counter{ds: ds}
}

// Get sends the keys to the underling getter. Counting the request.
func (ds *Counter) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	ds.requests = append(ds.requests, keys)
	return ds.ds.Get(ctx, keys...)
}

// Reset resets the counter.
func (ds *Counter) Reset() {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	ds.requests = nil
}

// Value returns the number of requests.
func (ds *Counter) Value() int {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	return len(ds.requests)
}

// Requests returns all lists of requested keys.
func (ds *Counter) Requests() [][]dskey.Key {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	return ds.requests
}

// Cache caches all requested keys and only redirects keys, if they where
// not requested before.
type Cache struct {
	mu sync.Mutex

	ds    Getter
	cache map[dskey.Key][]byte
}

// NewCache initializes a Cache.
func NewCache(ds Getter) Getter {
	return &Cache{ds: ds, cache: make(map[dskey.Key][]byte)}
}

// Get redirects the keys to the underling getter. But only, if they where not
// requested before.
func (ds *Cache) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	out := make(map[dskey.Key][]byte, len(keys))
	var needKeys []dskey.Key
	for _, key := range keys {
		v, ok := ds.cache[key]
		if !ok {
			needKeys = append(needKeys, key)
			continue
		}
		out[key] = v
	}

	if len(needKeys) == 0 {
		return out, nil
	}

	upstream, err := ds.ds.Get(ctx, needKeys...)
	if err != nil {
		return nil, fmt.Errorf("upstream: %w", err)
	}

	for k, v := range upstream {
		out[k] = v
		ds.cache[k] = v
	}
	return out, nil
}
