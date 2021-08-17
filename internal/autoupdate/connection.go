package autoupdate

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
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
// When Next is called for the first time, it does not block. In this case, it
// is possible, that it returns an empty map.
//
// On every other call, it blocks until there is new data. In this case, the map
// is never empty.
func (c *Connection) Next(ctx context.Context) (map[string][]byte, error) {
	firstTime := c.filter.empty()
	var data map[string][]byte

	for len(data) == 0 {
		recorder := datastore.NewRecorder(c.autoupdate.datastore)
		restricter := c.autoupdate.restricter(recorder, c.uid)

		keys, err := c.keys(ctx, restricter)
		if err != nil {
			return nil, fmt.Errorf("getting keys: %w", err)
		}

		data, err = restricter.Get(ctx, keys...)
		if err != nil {
			return nil, fmt.Errorf("get first time restricted data: %w", err)
		}

		c.filter.filter(data)

		if firstTime {
			// On firstTime return the data, even when it is empty.
			return data, nil
		}
	}

	return data, nil
}

func (c *Connection) keys(ctx context.Context, getter datastore.Getter) ([]string, error) {
	if c.filter.empty() {
		keys, err := c.allKeys(ctx, getter)
		if err != nil {
			return nil, fmt.Errorf("get all keys: %w", err)
		}
		return keys, nil
	}

	keys, err := c.nextKeys(ctx, getter)
	if err != nil {
		return nil, fmt.Errorf("get next keys: %w", err)
	}
	return keys, nil
}

func (c *Connection) allKeys(ctx context.Context, getter datastore.Getter) ([]string, error) {
	if c.tid == 0 {
		c.tid = c.autoupdate.topic.LastID()
	}

	if err := c.kb.Update(ctx, getter); err != nil {
		return nil, fmt.Errorf("create keys for keysbuilder: %w", err)
	}

	return c.kb.Keys(), nil
}

// nextKeys blocks until there are new keys for the user.
func (c *Connection) nextKeys(ctx context.Context, getter datastore.Getter) ([]string, error) {
	var keys []string
	for len(keys) == 0 {
		// Blocks until the topic is closed (on server exit) or the context is done.
		tid, changedKeys, err := c.autoupdate.topic.Receive(ctx, c.tid)
		if err != nil {
			return nil, fmt.Errorf("get updated keys: %w", err)
		}
		c.tid = tid

		changedSlice := make(map[string]bool, len(changedKeys))
		for _, key := range changedKeys {
			var uid int
			if _, err := fmt.Sscanf(key, fullUpdateFormat, &uid); err == nil {
				// The key is a fullUpdate key. Do not use it, exept of a full
				// update.
				if uid == -1 || uid == c.uid {
					return c.allKeys(ctx, getter)
				}
				continue
			}

			changedSlice[key] = true
		}

		oldKeys := c.kb.Keys()

		// Update keysbuilder get new list of keys.
		if err := c.kb.Update(ctx, getter); err != nil {
			return nil, fmt.Errorf("update keysbuilder: %w", err)
		}

		// Start with keys hat are new for the user.
		keys = keysDiff(oldKeys, c.kb.Keys())

		// Append keys that are old but have been changed.
		for _, key := range oldKeys {
			if !changedSlice[key] {
				continue
			}
			keys = append(keys, key)
		}
	}

	return keys, nil
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
