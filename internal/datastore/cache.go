package datastore

import (
	"context"
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
	mu   sync.RWMutex
	data map[string]*cacheEntry
}

type cacheEntry struct {
	done   chan struct{}
	cancel context.CancelFunc

	mu    sync.RWMutex
	value string
	err   error
}

func newCache() *cache {
	return &cache{data: make(map[string]*cacheEntry)}
}

// getOrSet returns the value for a key. If the values does not exist, it is
// created by the return value of the second argument. If this method is called
// more then once at the same time, only the first calculates the result, the
// other calles get blocked until it is calculated.
func (c *cache) getOrSet(key string, set func(context.Context) (string, error)) (string, error) {
	c.mu.Lock()
	entry, ok := c.data[key]

	if !ok {
		ctx, cancel := context.WithCancel(context.Background())
		entry = &cacheEntry{
			done:   make(chan struct{}),
			cancel: cancel,
		}
		c.data[key] = entry
		c.mu.Unlock()

		value, err := set(ctx)

		// Only set enty.value and entry.err if set was not canceled.
		select {
		case <-ctx.Done():
		default:
			entry.value, entry.err = value, err
		}
		close(entry.done)
		return entry.value, entry.err
	}
	c.mu.Unlock()

	<-entry.done

	entry.mu.RLock()
	defer entry.mu.RUnlock()
	return entry.value, entry.err
}

func (c *cache) setIfExist(key string, value string) {
	c.mu.Lock()
	entry, ok := c.data[key]
	c.mu.Unlock()

	if !ok {
		return
	}

	entry.mu.Lock()
	defer entry.mu.Unlock()
	entry.value = value
	entry.err = nil
	entry.cancel()
}
