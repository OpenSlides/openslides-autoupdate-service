package autoupdate

import (
	"context"
	"encoding/json"
	"fmt"
)

// Connection holds the state of a client. It has to be created by colling
// Connect() on a autoupdate.Service instance.
type Connection struct {
	autoupdate *Autoupdate
	uid        int
	kb         KeysBuilder
	tid        uint64
	filter     filter
}

// Next returns the next data for the user.
//
// Next blocks until there are new data or the context or the server closes. In
// this case, nil is returned.
func (c *Connection) Next(ctx context.Context) (map[string]json.RawMessage, error) {
	if c.filter.empty() {
		return c.allData(ctx)
	}

	var err error
	var changedKeys []string

	// Blocks until the topic is closed (on server exit) or the context is done.
	c.tid, changedKeys, err = c.autoupdate.topic.Receive(ctx, c.tid)
	if err != nil {
		return nil, fmt.Errorf("get updated keys: %w", err)
	}

	changedSlice := make(map[string]bool, len(changedKeys))
	for _, key := range changedKeys {
		var uid int
		if _, err := fmt.Sscanf(key, fullUpdateFormat, &uid); err == nil {
			// The key is a fullUpdate key. Do not use it, excpect of a full
			// update.
			if uid == c.uid {
				return c.allData(ctx)
			}
			continue
		}

		changedSlice[key] = true
	}

	oldKeys := c.kb.Keys()

	// Update keysbuilder get new list of keys
	if err := c.kb.Update(ctx); err != nil {
		return nil, fmt.Errorf("update keysbuilder: %w", err)
	}

	// Start with keys hat are new for the user
	keys := keysDiff(oldKeys, c.kb.Keys())

	// Append keys that are old but have been changed.
	for _, key := range oldKeys {
		if !changedSlice[key] {
			continue
		}
		keys = append(keys, key)
	}

	if len(keys) == 0 {
		// No data. Try again.
		return c.Next(ctx)
	}

	data, err := c.autoupdate.RestrictedData(ctx, c.uid, keys...)
	if err != nil {
		return nil, fmt.Errorf("restrict data: %w", err)
	}

	c.filter.filter(data)
	return data, nil
}

func (c *Connection) allData(ctx context.Context) (map[string]json.RawMessage, error) {
	c.filter.reset()
	// First time called
	if c.tid == 0 {
		c.tid = c.autoupdate.topic.LastID()
	}

	if err := c.kb.Update(ctx); err != nil {
		return nil, fmt.Errorf("create keys for keysbuilder: %w", err)
	}

	data, err := c.autoupdate.RestrictedData(ctx, c.uid, c.kb.Keys()...)
	if err != nil {
		return nil, fmt.Errorf("get first time restricted data: %w", err)
	}

	c.filter.filter(data)
	return data, nil
}

func keysDiff(old []string, new []string) []string {
	keySet := make(map[string]bool, len(old))
	for _, key := range old {
		keySet[key] = true
	}

	added := []string{}
	for _, key := range new {
		if keySet[key] {
			continue
		}
		added = append(added, key)
	}

	return added
}
