package dsmock

import (
	"context"
	"fmt"
	"sync"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Stub are data that can be used as a datastore value.
type Stub map[string][]byte

// Get implements the Getter interface.
func (s Stub) Get(_ context.Context, keys ...string) (map[string][]byte, error) {
	if invalid := datastore.InvalidKeys(keys...); len(invalid) > 0 {
		return nil, fmt.Errorf("keys %v are invalid", invalid)
	}

	data := map[string][]byte(s)
	requested := make(map[string][]byte, len(keys))
	for _, k := range keys {
		requested[k] = data[k]
	}
	return requested, nil
}

// Counter counts all keys that where requested.
type Counter struct {
	mu sync.Mutex

	ds       datastore.Getter
	requests [][]string
}

// NewCounter initializes a Counter.
func NewCounter(ds datastore.Getter) *Counter {
	return &Counter{ds: ds}
}

// Get sends the keys to the underling getter. Counting the request.
func (ds *Counter) Get(ctx context.Context, keys ...string) (map[string][]byte, error) {
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
func (ds *Counter) Requests() [][]string {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	return ds.requests
}

// Cache caches all requested keys and only redirects keys, if they where
// not requested before.
type Cache struct {
	mu sync.Mutex

	ds    datastore.Getter
	cache map[string][]byte
}

// NewCache initializes a Cache.
func NewCache(ds datastore.Getter) *Cache {
	return &Cache{ds: ds, cache: make(map[string][]byte)}
}

// Get redirects the keys to the underling getter. But only, if they where not
// requested before.
func (ds *Cache) Get(ctx context.Context, keys ...string) (map[string][]byte, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	out := make(map[string][]byte, len(keys))
	var needKeys []string
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
