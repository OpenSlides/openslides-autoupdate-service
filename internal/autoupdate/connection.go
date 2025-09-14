package autoupdate

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/OpenSlides/openslides-go/datastore/dsrecorder"
	"github.com/OpenSlides/openslides-go/datastore/flow"
)

// Connection holds the connection to data and has the ability to return the next.
type Connection interface {
	Next() (func(context.Context) (map[dskey.Key][]byte, error), bool)
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

// Next returns a function to fetch the next data.
//
// Next is a pull function as described in
// https://github.com/golang/go/discussions/56413
//
// With the current version of go, it has to be called like this:
//
//	next := conn.Next(
//	for f, ok := next(); ok; f, ok = next() {
//		data, err := f(ctx)
//		if err != nil {
//			break
//		}
//		...
//	}
//
// In a future version of go, it meight be called as:
//
//	for f := range conn.Next {
//		data, err := f(ctx)
//		if err != nil {
//			break
//		}
//		...
//	}
//
// When the returned function is called for the first time, it does not block.
// In this case, it is possible, that it returns an empty map.
//
// On every other call, it blocks until there is new data. In this case, the map
// is never empty.
func (c *connection) Next() (func(context.Context) (map[dskey.Key][]byte, error), bool) {
	return func(ctx context.Context) (map[dskey.Key][]byte, error) {
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

			if !foundKey {
				continue
			}

			data, err := c.updatedData(ctx)
			if err != nil {
				return nil, fmt.Errorf("creating later data: %w", err)
			}

			if len(data) > 0 {
				return data, nil
			}
		}
	}, true
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
		fn, _ := c.Next()
		dd, err := fn(ctx)
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

	type snapshotter interface {
		Snapshot(notFoundHandler flow.Getter) flow.Getter
	}

	if snapy, ok := c.autoupdate.flow.(snapshotter); ok {
		data, err := c.updatedDataWithGetter(ctx, snapy.Snapshot(c.autoupdate.flow))
		if err != nil {
			return nil, fmt.Errorf("update data with snapshot: %w", err)
		}

		return data, nil
	}

	data, err := c.updatedDataWithGetter(ctx, c.autoupdate.flow)
	if err != nil {
		return nil, fmt.Errorf("update with no snapshotter: %w", err)
	}

	return data, nil

}

func (c *connection) updatedDataWithGetter(ctx context.Context, getter flow.Getter) (map[dskey.Key][]byte, error) {
	recorder := dsrecorder.New(getter)
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
