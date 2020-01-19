package keysbuilder

import (
	"sync"
)

type cache struct {
	mu   sync.Mutex
	data map[string]*cacheEntry
}

type cacheEntry struct {
	done chan struct{}
	ids  []int
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
func (c *cache) getOrSet(key string, set func() []int) []int {
	c.mu.Lock()
	entry, ok := c.data[key]

	if !ok {
		entry = &cacheEntry{}
		c.data[key] = entry
		c.data[key].done = make(chan struct{})
		c.mu.Unlock()
		entry.ids = set()
		close(entry.done)
		return entry.ids
	}
	c.mu.Unlock()
	<-entry.done
	return entry.ids
}
