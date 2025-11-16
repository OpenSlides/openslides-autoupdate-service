package autoupdate

import (
	"context"
	"fmt"
	"iter"

	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/OpenSlides/openslides-go/datastore/dsrecorder"
)

// Connection holds the connection to data and has the ability to return the next.
type Connection interface {
	Messages(ctx context.Context) iter.Seq2[map[dskey.Key][]byte, error]
	NextWithFilter(ctx context.Context, filterHashes string) (map[dskey.Key][]byte, string, error)
}

// connection holds the state of a client. It has to be created by colling
// Connect() on a autoupdate.Service instance.
type connection struct {
	autoupdate   *Autoupdate
	uid          int
	kb           KeysBuilder
	tid          uint64
	filter       filter
	skipWorkpool bool
	hotkeys      map[dskey.Key]struct{}
}

// Messages returns an iterator over new messages.
//
// It impelements the go 1.23 iter interface and can be called like
//
//	for data, err := range conn.Messages(ctx) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// When the returned function is called for the first time, it does not block.
// In this case, it is possible, that it returns an empty map.
//
// On every other call, it blocks until there is new data. In this case, the map
// is never empty.
func (c *connection) Messages(ctx context.Context) iter.Seq2[map[dskey.Key][]byte, error] {
	return func(yield func(key map[dskey.Key][]byte, value error) bool) {
		if c.filter.empty() {
			c.tid = c.autoupdate.topic.LastID()
			data, err := c.updatedData(ctx)
			if err != nil {
				yield(nil, fmt.Errorf("creating first time data: %w", err))
				return
			}

			if !yield(data, nil) {
				return // break was used in for-loop
			}
		}

		for {
			// Blocks until new data or the context is done.
			tid, changedKeys, err := c.autoupdate.topic.ReceiveSince(ctx, c.tid)
			if err != nil {
				yield(nil, fmt.Errorf("get updated keys: %w", err))
				return
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
					yield(nil, fmt.Errorf("creating later data: %w", err))
					return
				}

				if len(data) > 0 {
					if !yield(data, nil) {
						return // break was used in for-loop
					}
				}
			}
		}
	}
}

func (c *connection) NextWithFilter(ctx context.Context, filterHashes string) (map[dskey.Key][]byte, string, error) {
	c.tid = c.autoupdate.topic.LastID()

	if err := c.filter.setHashState(filterHashes); err != nil {
		return nil, "", fmt.Errorf("set history state: %w", err)
	}

	data, err := c.updatedData(ctx)
	if err != nil {
		return nil, "", fmt.Errorf("creating first time data: %w", err)
	}

	if len(data) == 0 {
		dd, err, _ := getFirst(c.Messages(ctx))
		if err != nil {
			return nil, "", fmt.Errorf("getting new data: %w", err)
		}
		data = dd
	}

	hashes, err := c.filter.hashState()
	if err != nil {
		return nil, "", fmt.Errorf("create new hashes: %w", err)
	}
	return data, hashes, nil
}

// updatedData returns all values from the datastore.getter.
func (c *connection) updatedData(ctx context.Context) (map[dskey.Key][]byte, error) {
	if !c.skipWorkpool {
		done, err := c.autoupdate.pool.Wait(ctx)
		if err != nil {
			return nil, err
		}
		defer done()
	}

	recorder := dsrecorder.New(c.autoupdate.flow)
	ctx, restricter := c.autoupdate.restricter(ctx, recorder, c.uid)

	keys, err := c.kb.Update(ctx, restricter)
	if err != nil {
		return nil, fmt.Errorf("create keys for keysbuilder: %w", err)
	}

	data, err := restricter.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("get restricted data: %w", err)
	}
	c.hotkeys = recorder.Keys()

	c.filter.filter(data)

	return data, nil
}

func getFirst[K, V any](seq iter.Seq2[K, V]) (K, V, bool) {
	for k, v := range seq {
		return k, v, true
	}
	var zeroK K
	var zeroV V
	return zeroK, zeroV, false
}
