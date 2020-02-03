package autoupdate

import (
	"context"
	"fmt"

	"github.com/cespare/xxhash/v2"
	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate/keysbuilder"
)

// Connection holds the state of an open connection to the autoupdate system.
// It has to be created by colling Connect() on a autoupdate.Service instance.
type Connection struct {
	s    *Service
	ctx  context.Context
	user int
	tid  uint64
	b    *keysbuilder.Builder
	data map[string]uint64
}

// Read listens for data changes and blocks until then. When data has changed,
// it returns with the new data.
// When the given context is done, it returns immediately with nil data
func (c *Connection) Read() (map[string]string, error) {
	tid, changedKeys, err := c.s.topic.Get(c.ctx, c.tid)
	if err != nil {
		return nil, fmt.Errorf("can not get new data: %w", err)
	}
	c.tid = tid

	if len(changedKeys) == 0 {
		// Exit early
		return nil, nil
	}

	oldKeys := c.b.Keys()

	// Update keysbuilder get new list of keys
	if err := c.b.Update(changedKeys); err != nil {
		return nil, fmt.Errorf("can not update keysbuilder: %w", err)
	}

	// Start with keys hat are new for the user
	keys := keysDiff(oldKeys, c.b.Keys())

	changedSlice := make(map[string]bool, len(changedKeys))
	for _, key := range changedKeys {
		changedSlice[key] = true
	}

	// Append keys that are old but have been changed.
	for _, key := range c.b.Keys() {
		if !changedSlice[key] {
			continue
		}
		keys = append(keys, key)
	}

	data, err := c.s.restricter.Restrict(c.ctx, c.user, keys)
	if err != nil {
		return nil, fmt.Errorf("can not restrict data: %v", err)
	}
	c.filter(data)
	return data, nil
}

// filter removes values from data, that are the same as before
func (c *Connection) filter(data map[string]string) {
	if c.data == nil {
		c.data = make(map[string]uint64)
	}
	for key, value := range data {
		new := xxhash.Sum64String(value)
		old, ok := c.data[key]
		if ok && old == new {
			delete(data, key)
			continue
		}

		c.data[key] = new
	}
}
