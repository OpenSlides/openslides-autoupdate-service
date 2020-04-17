package keysbuilder

import (
	"sync"
)

// cache caches the requests to the ider. If many requests to the same key are
// send at the same time, only one request is send and the other request block
// until the first one gets an answer.
//
// A cache object has to be initialized with newCache().
type cache struct {
	mu   sync.Mutex
	data map[string]*cacheEntry
}

type cacheEntry struct {
	done  chan struct{}
	value interface{}
	err   error
}

func newCache() *cache {
	return &cache{data: make(map[string]*cacheEntry)}
}

// getOrSet returns the ids for a key. If the key does not exist, it is created
// by the return value of the second argument. If this method is called more
// then once at the same time, only the first calculates the result, the other
// calles get blocked until it is calculated.
func (c *cache) getOrSet(key string, set func() (interface{}, error)) (interface{}, error) {
	c.mu.Lock()
	entry, ok := c.data[key]

	if !ok {
		entry = &cacheEntry{}
		c.data[key] = entry
		c.data[key].done = make(chan struct{})
		c.mu.Unlock()

		entry.value, entry.err = set()
		close(entry.done)
		return entry.value, entry.err
	}
	c.mu.Unlock()

	<-entry.done
	return entry.value, entry.err
}

// delete removes some keys from the cache.
func (c *cache) delete(keys []string) {
	c.mu.Lock()
	for _, key := range keys {
		delete(c.data, key)
	}
	c.mu.Unlock()
}
