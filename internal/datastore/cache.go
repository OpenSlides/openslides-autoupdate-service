package datastore

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

// cacheSetFunc is a function to update cache keys.
type cacheSetFunc func(keys []string) (map[string]json.RawMessage, error)

// cache stores the values to the datastore.
//
// Each value of the cache has three states. Either it exists, it does not
// exist, or it is pending. Pending means, that there is a current request to
// the datastore.
//
// A new cache instance has to be created with newCache().
type cache struct {
	mu      sync.RWMutex
	data    map[string]json.RawMessage
	pending map[string]chan struct{}
}

func newCache() *cache {
	return &cache{
		data:    make(map[string]json.RawMessage),
		pending: make(map[string]chan struct{}),
	}
}

// getOrSet returns the values for a list of keys. If one or more keys do not
// exist in the cache, then the missing values are fetched with the given set
// function. If this method is called more then once at the same time, only the
// first calculates the result, the other calles get blocked until it is
// calculated.
//
// All values get returned together. If only one key is missing, this function
// blocks, until all values are retrieved.
//
// The set function is used to create the cache values. It is called only with
// the missing keys.
//
// If the context is done, getOrSet returns. But the set() call is not stopped.
// Other calls to getOrSet may wait for its result.
func (c *cache) getOrSet(ctx context.Context, keys []string, set cacheSetFunc) ([]json.RawMessage, error) {
	// Get all requested entries from cache. If entry does not exist, set it to
	// pending.
	c.mu.Lock()
	var missingKeys []string
	for _, key := range keys {
		if _, ok := c.data[key]; ok {
			continue
		}

		if _, ok := c.pending[key]; ok {
			continue
		}

		missingKeys = append(missingKeys, key)
		c.pending[key] = make(chan struct{})
		c.data[key] = nil
	}
	c.mu.Unlock()

	// Get values that are missing.
	if len(missingKeys) > 0 {
		done := make(chan struct{})

		// Fetch missing keys in the background. Do not stop the fetching. Even
		// when the context is done. Other calls could also request it.
		go func() {
			c.fetchMissing(missingKeys, set)
			close(done)
		}()

		select {
		case <-done:
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}

	values := make([]json.RawMessage, len(keys))
	c.mu.RLock()
	for i, key := range keys {

		v := c.data[key]
		p, pendingOK := c.pending[key]

		if !pendingOK {
			values[i] = v
			continue
		}

		c.mu.RUnlock()
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-p:
		}
		c.mu.RLock()

		v, ok := c.data[key]
		if !ok {
			// The value is not in the cache after pending was done. This
			// happens when the request to the datastore returned with an error.
			return nil, fmt.Errorf("key %s is unknoen", key)
		}
		values[i] = v
	}
	c.mu.RUnlock()
	return values, nil
}

func (c *cache) fetchMissing(keys []string, set cacheSetFunc) {
	data, err := set(keys)

	c.mu.Lock()
	defer c.mu.Unlock()

	if err != nil {
		log.Printf("Can not load keys %v: %v", keys, err)
		for _, k := range keys {
			close(c.pending[k])
			delete(c.pending, k)
		}
		return
	}

	for k, v := range data {
		c.data[k] = v

		p, ok := c.pending[k]
		if !ok {
			continue
		}

		close(p)
		delete(c.pending, k)
	}
}

// setIfExist updates each the cache with the value in the given map. But only
// values that are already in the cache get an update.
//
// Pending keys are not updated
func (c *cache) setIfExist(data map[string]json.RawMessage) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, value := range data {
		_, ok := c.data[key]

		if !ok {
			continue
		}

		p, ok := c.pending[key]

		if ok {
			close(p)
			delete(c.pending, key)
		}

		c.data[key] = value
	}
}
