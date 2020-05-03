package datastore

import (
	"context"
	"sync"
)

// cacheSetFunc is a function to update cache keys.
type cacheSetFunc func(keys []string) (map[string]string, error)

// cache stores the values to the datastore.
//
// Each value of the cache has three states. Either it exists, it does not
// exist, or it is pending. Pending means, that there is a current request to
// the datastore.
//
// A new cache instance has to be created with newCache().
type cache struct {
	mu   sync.RWMutex
	data map[string]*cacheEntry
}

func newCache() *cache {
	return &cache{data: make(map[string]*cacheEntry)}
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
func (c *cache) getOrSet(ctx context.Context, keys []string, set cacheSetFunc) ([]string, error) {
	// Get all requested entries from cache. If entry does not exist, create a
	// new one.
	c.mu.Lock()
	var missingKeys []string
	for _, key := range keys {
		if _, ok := c.data[key]; ok {
			continue
		}

		missingKeys = append(missingKeys, key)
		c.data[key] = &cacheEntry{
			done: make(chan struct{}),
		}
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

	values := make([]string, len(keys))
	for i, key := range keys {
		c.mu.RLock()
		entry := c.data[key]
		c.mu.RUnlock()

		// This blocks until the value is ready.
		value, err := entry.get()

		if err != nil {
			return nil, err
		}
		values[i] = value
	}
	return values, nil
}

func (c *cache) fetchMissing(keys []string, set cacheSetFunc) {
	data, err := set(keys)

	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, key := range keys {
		value, ok := data[key]
		if !ok {
			// TODO: think about this.
			value = "null"
		}
		c.data[key].set(value, err)
	}
}

// setIfExist updates each the cache with the value in the given map. But only
// values that are already in the cache get an update.
func (c *cache) setIfExist(data map[string]string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, value := range data {
		entry, ok := c.data[key]

		if !ok {
			continue
		}

		entry.update(value)
	}
}

type cacheEntry struct {
	mu    sync.RWMutex
	done  chan struct{}
	value string
	err   error
}

// get returns the value and error from the cacheEntry. It block until the value
// is ready.
func (ce *cacheEntry) get() (string, error) {
	<-ce.done

	ce.mu.RLock()
	defer ce.mu.RUnlock()

	return ce.value, ce.err
}

// set sets the cache entry.
func (ce *cacheEntry) set(value string, err error) {
	ce.mu.Lock()
	defer ce.mu.Unlock()

	// Only set entry, if it is not done (by setIfExist).
	select {
	case <-ce.done:
		return
	default:
	}

	defer close(ce.done)

	if err != nil {
		ce.err = err
		return
	}

	ce.value = value
}

// update updates the cache entry. The differece to set is, that is writes the
// value, even when it is done.
func (ce *cacheEntry) update(value string) {
	ce.mu.Lock()
	defer ce.mu.Unlock()

	ce.value = value
	ce.err = nil

	// If it was done before, set it to done.
	select {
	case <-ce.done:
	default:
		close(ce.done)
	}
}
