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

// get returns the ids for a key.
// if it returns nil, the ids have to be generated
// by the caller and set via cache.set()
func (c *cache) get(key string, set func() ([]int, error)) ([]int, error) {
	c.mu.Lock()
	entry, ok := c.data[key]

	if !ok {
		entry = &cacheEntry{}
		c.data[key] = entry
		c.data[key].done = make(chan struct{})
		c.mu.Unlock()
		var err error
		entry.ids, err = set()
		close(entry.done)
		return entry.ids, err
	}
	c.mu.Unlock()
	<-entry.done
	return entry.ids, nil
}
