package dsmock

import (
	"context"
	"fmt"
	"sync"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/flow"
)

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

// Flow is is a mock flow.
type Flow struct {
	mu sync.RWMutex

	stub Stub
	ch   chan map[dskey.Key][]byte

	getter flow.Getter

	middlewares []flow.Getter
}

// NewFlow initializes a stub with Get and Update.
func NewFlow(data map[dskey.Key][]byte, middlewares ...func(flow.Getter) flow.Getter) *Flow {
	stub := Stub(data)
	getter := flow.Getter(stub)
	initialized := make([]flow.Getter, len(middlewares))
	for i, m := range middlewares {
		getter = m(getter)
		initialized[i] = getter
	}

	return &Flow{
		stub:        stub,
		ch:          make(chan map[dskey.Key][]byte),
		getter:      getter,
		middlewares: initialized,
	}
}

// Get returns the current value for the given keys.
func (s *Flow) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.getter.Get(ctx, keys...)
}

// Update blocks until new data is received via the Send method.
func (s *Flow) Update(ctx context.Context, updateFn func(map[dskey.Key][]byte, error)) {
	if updateFn == nil {
		updateFn = func(map[dskey.Key][]byte, error) {}
	}
	for {
		select {
		case newValues := <-s.ch:
			updateFn(newValues, nil)
			continue

		case <-ctx.Done():
			return
		}
	}
}

// Send sends keys to the mock that can be received with Update().
func (s *Flow) Send(values map[dskey.Key][]byte) {
	s.mu.Lock()
	for k, v := range values {
		s.stub[k] = v
	}
	s.mu.Unlock()
	s.ch <- values
}

// Middlewares returns a list of Getters that where used in
// NewStubWithUpdates.
//
// For example:
//
//	ds := NewStubWithUpdate(stub, dsmock.NewCounter)
//	counter := ds.Middlewares()[0].(*dsmock.Counter)
func (s *Flow) Middlewares() []flow.Getter {
	return s.middlewares
}

// Counter counts all keys that where requested.
type Counter struct {
	mu sync.Mutex

	ds       flow.Getter
	requests [][]dskey.Key
}

// NewCounter initializes a Counter.
func NewCounter(ds flow.Getter) flow.Getter {
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

// Count returns the number of requests.
func (ds *Counter) Count() int {
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

	ds    flow.Getter
	cache map[dskey.Key][]byte
}

// NewCache initializes a Cache.
func NewCache(ds flow.Getter) flow.Getter {
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

// Wait is a Getter that blocks until there is a signal on a channel.
//
// If the signal is an error, it is returned
type Wait struct {
	waiter chan error
	getter flow.Getter
}

// NewWait initializes a Wait.
func NewWait(waiter chan error) func(flow.Getter) flow.Getter {
	return func(getter flow.Getter) flow.Getter {
		return &Wait{
			waiter: waiter,
			getter: getter,
		}
	}
}

// Get retuns the values from the getter as soon, as there is a signal on the channel.
func (w *Wait) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()

	case err := <-w.waiter:
		if err != nil {
			return nil, err
		}

		return w.getter.Get(ctx, keys...)
	}
}
