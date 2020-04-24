package autoupdate

import (
	"context"
	"fmt"
	"io"
)

// Connection holds the state of a client. It has to be created by colling
// Connect() on a autoupdate.Service instance.
type Connection struct {
	autoupdate *Service
	ctx        context.Context
	uid        int
	kb         KeysBuilder
	tid        uint64
	next       bool
}

// Next returns the next data in form of an reader.
//
// Next blocks until there are new data or the context or the server closes. In
// this case, the returned io.Reader is nil.
func (c *Connection) Next() (io.Reader, error) {
	if !c.next {
		// First time called
		c.next = true
		c.tid = c.autoupdate.topic.LastID()
		return c.autoupdate.restricter.Restrict(c.ctx, c.uid, c.kb.Keys())
	}

	var err error
	var changedKeys []string

	// Blocks until the topic is closed (on server exit) or the context is done.
	c.tid, changedKeys, err = c.autoupdate.topic.Receive(c.ctx, c.tid)
	if err != nil {
		return nil, fmt.Errorf("can not get new keys: %w", err)
	}

	// When changedKeys is empty, then the service or the connection is closed.
	if len(changedKeys) == 0 {
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

	if len(keys) == 0 {
		// No data. Try again.
		return c.Next()
	}

	return c.autoupdate.restricter.Restrict(c.ctx, c.uid, keys)
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
