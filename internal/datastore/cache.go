package datastore

import (
	"context"
	"fmt"
	"sync"
)

// cache stores the values to the datastore.
//
// Each value of the cache has three states. Either it exists, it does not
// exist, or it is pending. Pending means, that there is a current request to
// the datastore.
//
// A new cache instance has to be created with newCache().
type cache struct {
	mu   sync.Mutex
	data map[string]*cacheEntry
}

type cacheEntry struct {
	done   chan struct{}
	cancel chan struct{}

	mu    sync.RWMutex
	value string
	err   error
}

func newCache() *cache {
	return &cache{data: make(map[string]*cacheEntry)}
}

// getOrSet returns the values for a list of keys. If one or more keys do not
// exist in the cache, then the missing values are created with the given set
// function. If this method is called more then once at the same time, only the
// first calculates the result, the other calles get blocked until it is
// calculated.
//
// All values get returned at once. If only one key is missing, this function
// blocks, until all values are retrieved.
//
// The set function is used to create the cache values. It is called only with
// the missing keys.
//
// If the context is done, getOrSet returns. But the set() call is not stopped.
// Other calls to getOrSet may wait for its result.
func (c *cache) getOrSet(ctx context.Context, keys []string, set func(keys []string) ([]string, error)) ([]string, error) {
	// entries is a map like cache.data. All values from cache.data are also
	// saved in this map, so cache.data does not have to be locked for long.
	entries := make(map[string]*cacheEntry, len(keys))

	c.mu.Lock()

	// Get all requested entries from cache. If entry does not exist, create a
	// new one.
	var missingKeys []string
	var ok bool
	for _, key := range keys {
		entries[key], ok = c.data[key]
		if ok {
			continue
		}

		missingKeys = append(missingKeys, key)
		entries[key] = &cacheEntry{
			done:   make(chan struct{}),
			cancel: make(chan struct{}),
		}
		c.data[key] = entries[key]
	}
	c.mu.Unlock()

	// Get values that are missing
	if len(missingKeys) > 0 {
		done := make(chan struct{})

		go func() {
			retrievedValues, err := set(missingKeys)
			if err == nil && len(retrievedValues) < len(missingKeys) {
				err = fmt.Errorf("no value returned")
			}
			for i, key := range missingKeys {
				entry := entries[key]
				// Only set enty.value and entry.err if key was not canceled.
				select {
				case <-entry.cancel:
				default:
					// TODO: is this a race condition??? if not:
					// entry.mu has not to be locked. It is guaraneed, that the values
					// can not be written.
					if err != nil {
						entry.err = err

					} else {
						entry.value = retrievedValues[i]
					}
				}
				close(entry.done)
			}
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
		entry := entries[key]
		// Wait until each entry is done.
		<-entry.done

		entry.mu.RLock()
		if err := entry.err; err != nil {
			entry.mu.RUnlock()
			return nil, err
		}
		values[i] = entry.value
		entry.mu.RUnlock()
	}
	return values, nil
}

func (c *cache) setIfExist(data map[string]string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, value := range data {
		entry, ok := c.data[key]

		if !ok {
			return
		}

		entry.mu.Lock()
		entry.value = value
		entry.err = nil
		close(entry.cancel)
		entry.mu.Unlock()
	}
}
