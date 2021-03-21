package datastore

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"sync"
)

const (
	stNotExist = iota
	stExist
	stPending
	stInvalid
)

// cacheSetFunc is a function to update cache keys.
type cacheSetFunc func(keys []string) (map[string]json.RawMessage, error)

// cache stores the values to the datastore.
//
// Each value of the cache has three states. Either it exists, it does not
// exist, or it is pending. Pending means, that there is a current request to
// the datastore. An existing key can have the value `nil` which means, that the
// cache knows, that the key does not exist in the datastore. Each value
// []byte("null") is changed to nil.
//
// cache.keyState() tells, if a key exist or is pending.
//
// A new cache instance has to be created with newCache().
type cache struct {
	mu      sync.RWMutex
	data    map[string]json.RawMessage
	pending map[string]chan struct{}
}

// newCache creates an initialized cache instance.
func newCache() *cache {
	return &cache{
		data:    make(map[string]json.RawMessage),
		pending: make(map[string]chan struct{}),
	}
}

// GetOrSet returns the values for a list of keys. If one or more keys do not
// exist in the cache, then the missing values are fetched with the given set
// function. If this method is called more then once at the same time, only the
// first call fetches the result, the other calles get blocked until it the
// answer was fetched.
//
// A non existing value is returned as nil.
//
// All values get returned together. If only one key is missing, this function
// blocks, until all values are retrieved.
//
// The set function is used to create the cache values. It is called only with
// the missing keys.
//
// If a value is not returned by the set function, it is saved in the cache as
// nil to prevent a second call for the same key.
//
// If the context is done, GetOrSet returns. But the set() call is not stopped.
// Other calls to GetOrSet may wait for its result.
func (c *cache) GetOrSet(ctx context.Context, keys []string, set cacheSetFunc) ([]json.RawMessage, error) {
	c.mu.Lock()
	missingKeys := c.notExistToPending(keys)
	c.mu.Unlock()

	// Fetch missing keys.
	if len(missingKeys) > 0 {
		// Fetch missing keys in the background. Do not stop the fetching. Even
		// when the context is done. Other calls could also request it.
		errChan := make(chan error)
		go func() {
			err := c.fetchMissing(missingKeys, set)
			errChan <- err
		}()

		select {
		case err := <-errChan:
			if err != nil {
				return nil, fmt.Errorf("fetching key: %w", err)
			}
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}

	// Build return values. Blocks until pending keys are fetched.
	values := make([]json.RawMessage, len(keys))
	c.mu.RLock()
	for i, key := range keys {
		switch c.keyState(key) {
		case stExist:
			values[i] = c.data[key]
			continue
		case stInvalid:
			return nil, fmt.Errorf("key `%s` is in invalid state", key)
		case stNotExist:
			return nil, fmt.Errorf("key `%s` does not exist in cache", key)
		}
		p := c.pending[key]

		c.mu.RUnlock()
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-p:
		}
		c.mu.RLock()

		if c.keyState(key) != stExist {
			// The value is not in the cache after pending was done. This
			// happens when the request to the datastore of another
			// GetOrSet-Call returned with an error. Try it once more.
			c.mu.RUnlock()
			_, err := c.GetOrSet(ctx, []string{key}, set)
			if err != nil {
				return nil, fmt.Errorf("fetching keys for a second time: %w", err)
			}
			c.mu.RLock()
		}

		values[i] = c.data[key]
	}
	c.mu.RUnlock()
	return values, nil
}

// fetchMissing loads the given keys with the set method. Does not update keys
// that are already in the cache.
//
// Deletes the keys from the pending map, even when an error happens.
func (c *cache) fetchMissing(keys []string, set cacheSetFunc) error {
	data, err := set(keys)

	c.mu.Lock()
	defer c.mu.Unlock()

	// Make sure all pending keys are closed and deleted. Make also sure, that
	// missing keys are set to nil.
	defer func() {
		for _, k := range keys {
			if c.keyState(k) == stPending {
				p := c.pending[k]
				close(p)
				delete(c.pending, k)
			}
		}
	}()

	if err != nil {
		return fmt.Errorf("fetching missing keys: %w", err)
	}

	for k, v := range data {
		if c.keyState(k) == stPending {
			c.set(k, v)
		}
	}

	// Set all keys, that where not returned to not existing.
	for _, k := range keys {
		if c.keyState(k) == stPending {
			c.set(k, nil)
		}
	}
	return nil
}

// SetIfExist updates the cache with the value in the given map.
//
// Only keys that exist or are pending are updated.
func (c *cache) SetIfExist(data map[string]json.RawMessage) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, value := range data {
		if c.keyState(key) == stNotExist {
			continue
		}
		c.set(key, value)
	}
}

// Returns the state of a key.
//
// The cache has to be in read lock to call this method.
//
// If a key does not exist, data[key] and pending[key] do not exist.
//
// If a key does exist, data[key] exists but pending[key] does not exist.
//
// If a key is pending, data[key] does not exist and panding[key] does exist.
func (c *cache) keyState(key string) int {
	_, dataOK := c.data[key]
	_, pendingOK := c.pending[key]

	if dataOK {
		if pendingOK {
			return stInvalid
		}
		return stExist
	}
	if pendingOK {
		return stPending
	}
	return stNotExist
}

// set sets a key in the cache to a value. Closes the pending state.
func (c *cache) set(key string, value json.RawMessage) {
	// Change "null" values to nil.
	if bytes.Equal(value, []byte("null")) {
		value = nil
	}
	c.data[key] = value
	if p, ok := c.pending[key]; ok {
		close(p)
		delete(c.pending, key)
	}
}

// notExistToPending sets all given keys, that do not exist in the cache, to pending.
// Returns the list of keys that where set to pending.
//
// The cache has to be in write lock to call this method.
func (c *cache) notExistToPending(keys []string) []string {
	var missingKeys []string
	for _, key := range keys {
		if c.keyState(key) == stNotExist {
			missingKeys = append(missingKeys, key)
			c.pending[key] = make(chan struct{})
		}
	}
	return missingKeys
}
