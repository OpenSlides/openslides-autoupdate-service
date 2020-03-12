package keysbuilder

import (
	"strings"
	"sync"
)

// cache caches the requests to the ider. If many requests to the same key
// are send at the same time, only one request is send and the other request
// block until the first one gets an answer.
type cache struct {
	mu   sync.Mutex
	data map[string]*cacheEntry
}

type cacheEntry struct {
	done  chan struct{}
	value interface{}
}

func newCache() *cache {
	c := &cache{}
	c.data = make(map[string]*cacheEntry)
	return c
}

// getOrSet returns the ids for a key.
// If the key does not exist, it is created be the return value
// of the second argument. If this method is called more then once
// at the same time, only the first calculates the result, the other
// calles get blocked until it is calculated.
func (c *cache) getOrSet(key string, set func() interface{}) interface{} {
	c.mu.Lock()
	entry, ok := c.data[key]

	if !ok {
		entry = &cacheEntry{}
		c.data[key] = entry
		c.data[key].done = make(chan struct{})
		c.mu.Unlock()
		entry.value = set()
		close(entry.done)
		return entry.value
	}
	c.mu.Unlock()
	<-entry.done
	return entry.value
}

func (c *cache) delete(keys []string) {
	c.mu.Lock()
	for _, key := range keys {
		if !(strings.HasSuffix(key, "_id") || strings.HasSuffix(key, "_ids")) {
			continue
		}
		delete(c.data, key)
	}
	c.mu.Unlock()
}
