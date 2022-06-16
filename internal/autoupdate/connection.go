package autoupdate

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

// connection holds the state of a client. It has to be created by colling
// Connect() on a autoupdate.Service instance.
type connection struct {
	autoupdate *Autoupdate
	uid        int
	kb         KeysBuilder
	tid        uint64
	filter     filter

	restrictHotkeys    *set.Set[datastore.Key]
	keysbuilderHotKeys *set.Set[datastore.Key]
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
		data, err := c.updatedData(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("creating first time data: %w", err)
		}

		return data, nil
	}

	requestedKeys := set.New(c.kb.Keys()...)
	for {
		// Blocks until new data or the context is done.
		tid, changedKeys, err := c.autoupdate.topic.Receive(ctx, c.tid)
		if err != nil {
			// TODO EXTERMAL ERROR
			return nil, fmt.Errorf("get updated keys: %w", err)
		}
		c.tid = tid

		changedKeysSet := set.New(changedKeys...)
		if !(set.Intersect(changedKeysSet, c.restrictHotkeys) || set.Intersect(changedKeysSet, requestedKeys)) {
			continue
		}

		data, err := c.updatedData(ctx, changedKeysSet)
		if err != nil {
			return nil, fmt.Errorf("creating later data: %w", err)
		}

		if len(data) == 0 {
			continue
		}

		return data, nil
	}
}

// updatedData returns all values from the datastore.getter.
func (c *connection) updatedData(ctx context.Context, changedKeys *set.Set[datastore.Key]) (map[datastore.Key][]byte, error) {
	var fullUpdate bool
	var requestedKeys []datastore.Key

	// Check if keysbuilder has to be updated.
	if changedKeys == nil || set.Intersect(c.keysbuilderHotKeys, changedKeys) {
		kbRecorder := datastore.NewRecorder(c.autoupdate.datastore)
		oldKeys := c.kb.Keys()
		if err := c.kb.Update(ctx, c.autoupdate.restricter(kbRecorder, c.uid)); err != nil {
			return nil, fmt.Errorf("create keys for keysbuilder: %w", err)
		}
		c.keysbuilderHotKeys = kbRecorder.Keys()

		newKeys := c.kb.Keys()
		removedKeys := notInSlice(oldKeys, newKeys)
		for _, key := range removedKeys {
			c.filter.delete(key)
		}

		requestedKeys = newKeys
		fullUpdate = true
	} else if set.Intersect(c.restrictHotkeys, changedKeys) {
		requestedKeys = c.kb.Keys()
	} else {
		requestedKeys = changedKeys.List()
	}

	data, err := c.autoupdate.datastore.Get(ctx, requestedKeys...)
	if err != nil {
		return nil, fmt.Errorf("getting full data: %w", err)
	}

	restrictHotKeys, err := restrict.Restrict(ctx, c.autoupdate.datastore, c.uid, data)
	if err != nil {
		return nil, fmt.Errorf("restricting keys: %w", err)
	}

	if fullUpdate {
		c.restrictHotkeys = restrictHotKeys
	} else {
		c.restrictHotkeys.AddOther(restrictHotKeys)
	}

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
