package autoupdate

import (
	"context"
	"fmt"
	"hash/maphash"
)

// Connection is a generator-like object that holds the state of an
// open connection to the autoupdate system.
// It has to be created by colling Connect() on a autoupdate.Service instance.
type Connection struct {
	autoupdate *Service
	ctx        context.Context
	user       int
	tid        uint64
	kb         KeysBuilder
	histroy    map[string]uint64
	data       map[string]string
	err        error
	hash       maphash.Hash
}

// Next listens for data changes and blocks until then. When data has changed,
// it returns with the new data.
// When the given context is done, it returns immediately with nil data.
// If an error happens, Next returns with nil. You have to check if an error happend
// with the Err() method.
func (c *Connection) Next() bool {
	if c.err != nil {
		return false
	}

	var keys []string
	if c.histroy == nil {
		// First time Next() is called.
		c.tid = c.autoupdate.topic.LastID()
		keys = c.kb.Keys()
	} else {
		var err error
		c.data = nil
		keys, err = c.update()
		if err != nil {
			c.err = err
			return false
		}
		if keys == nil {
			return false
		}
	}

	data, err := c.autoupdate.restricter.Restrict(c.ctx, c.user, keys)
	if err != nil {
		c.err = fmt.Errorf("can not restrict data: %v", err)
		return false
	}
	c.filter(data)
	c.data = data
	return true
}

func (c *Connection) update() ([]string, error) {
	tid, changedKeys, err := c.autoupdate.topic.Get(c.ctx, c.tid)
	if err != nil {
		return nil, fmt.Errorf("can not get new data: %w", err)
	}
	c.tid = tid

	if len(changedKeys) == 0 {
		// Exit early
		return nil, nil
	}

	oldKeys := c.kb.Keys()

	// Update keysbuilder get new list of keys
	if err := c.kb.Update(changedKeys); err != nil {
		return nil, fmt.Errorf("can not update keysbuilder: %w", err)
	}

	// Start with keys hat are new for the user
	keys := keysDiff(oldKeys, c.kb.Keys())

	changedSlice := make(map[string]bool, len(changedKeys))
	for _, key := range changedKeys {
		changedSlice[key] = true
	}

	// Append keys that are old but have been changed.
	for _, key := range c.kb.Keys() {
		if !changedSlice[key] {
			continue
		}
		keys = append(keys, key)
	}
	return keys, nil
}

// Data returns the data gathered by Next().
func (c *Connection) Data() map[string]string {
	return c.data
}

// Err returns an error if some happen when calling Next().
func (c *Connection) Err() error {
	return c.err
}

// filter removes values from data, that are the same as before.
func (c *Connection) filter(data map[string]string) {
	if c.histroy == nil {
		c.histroy = make(map[string]uint64)
	}
	for key, value := range data {
		c.hash.Reset()
		c.hash.WriteString(value)
		new := c.hash.Sum64()
		old, ok := c.histroy[key]
		if ok && old == new {
			delete(data, key)
			continue
		}

		c.histroy[key] = new
	}
}

func keysDiff(old []string, new []string) []string {
	slice := make(map[string]bool, len(old))
	for _, key := range old {
		slice[key] = true
	}
	added := []string{}
	for _, key := range new {
		if slice[key] {
			continue
		}
		added = append(added, key)
	}
	return added
}
