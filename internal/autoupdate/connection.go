package autoupdate

import (
	"context"
	"fmt"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// connection holds the state of a client. It has to be created by colling
// Connect() on a autoupdate.Service instance.
type connection struct {
	autoupdate *Autoupdate
	uid        int
	kb         KeysBuilder
	tid        uint64
	filter     filter
	hotkeys    map[datastore.Key]struct{}
}

// Next returns the next data for the user.
//
// When Next is called for the first time, it does not block. In this case, it
// is possible, that it returns an empty map.
//
// On every other call, it blocks until there is new data. In this case, the map
// is never empty.
func (c *connection) Next(ctx context.Context) (map[datastore.Key][]byte, error) {
	if c.filter.empty() {
		c.tid = c.autoupdate.topic.LastID()
		data, err := c.updatedData(ctx)
		if err != nil {
			return nil, fmt.Errorf("creating first time data: %w", err)
		}

		return data, nil
	}

	for {
		// Blocks until new data or the context is done.
		tid, changedKeys, err := c.autoupdate.topic.Receive(ctx, c.tid)
		if err != nil {
			// TODO EXTERMAL ERROR
			return nil, fmt.Errorf("get updated keys: %w", err)
		}
		c.tid = tid

		foundKey := false
		for _, key := range changedKeys {
			if _, ok := c.hotkeys[key]; ok {
				foundKey = true
				break
			}
		}

		if foundKey {
			data, err := c.updatedData(ctx)
			if err != nil {
				return nil, fmt.Errorf("creating later data: %w", err)
			}

			if len(data) > 0 {
				return data, nil
			}
		}
		time.Sleep(5 * time.Second)
	}
}

// updatedData returns all values from the datastore.getter.
func (c *connection) updatedData(ctx context.Context) (map[datastore.Key][]byte, error) {
	recorder := datastore.NewRecorder(c.autoupdate.datastore)
	restricter := c.autoupdate.restricter(recorder, c.uid)

	oldKeys := c.kb.Keys()
	if err := c.kb.Update(ctx, restricter); err != nil {
		return nil, fmt.Errorf("create keys for keysbuilder: %w", err)
	}

	newKeys := c.kb.Keys()
	removedKeys := notInSlice(oldKeys, newKeys)
	for _, key := range removedKeys {
		c.filter.delete(key)
	}

	data, err := restricter.Get(ctx, newKeys...)
	if err != nil {
		return nil, fmt.Errorf("get restricted data: %w", err)
	}
	c.hotkeys = recorder.Keys()

	c.filter.filter(data)

	return data, nil
}

// notInSlice returns elements that are in slice a but not in b.
func notInSlice(a, b []datastore.Key) []datastore.Key {
	bSet := make(map[datastore.Key]struct{}, len(b))
	for _, k := range b {
		bSet[k] = struct{}{}
	}

	var missing []datastore.Key
	for _, k := range a {
		if _, ok := bSet[k]; !ok {
			missing = append(missing, k)
		}
	}
	return missing
}
