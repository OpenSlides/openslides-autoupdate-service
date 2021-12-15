package autoupdate

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"go.opentelemetry.io/otel"
)

// connection holds the state of a client. It has to be created by colling
// Connect() on a autoupdate.Service instance.
type connection struct {
	autoupdate *Autoupdate
	uid        int
	kb         KeysBuilder
	tid        uint64
	filter     filter
	hotkeys    map[string]bool
}

// Next returns the next data for the user.
//
// When Next is called for the first time, it does not block. In this case, it
// is possible, that it returns an empty map.
//
// On every other call, it blocks until there is new data. In this case, the map
// is never empty.
func (c *connection) Next(ctx context.Context) (map[string][]byte, error) {
	if c.filter.empty() {
		data, err := c.data(ctx)
		if err != nil {
			return nil, fmt.Errorf("creating first time data: %w", err)
		}
		return data, nil
	}

	for {
		// Blocks until the topic is closed (on server exit) or the context is done.
		tid, changedKeys, err := c.autoupdate.topic.Receive(ctx, c.tid)
		if err != nil {
			return nil, fmt.Errorf("get updated keys: %w", err)
		}

		c.tid = tid

		for _, key := range changedKeys {
			if c.hotkeys[key] {
				data, err := c.data(ctx)
				if err != nil {
					return nil, fmt.Errorf("creating later data: %w", err)
				}
				if len(data) > 0 {
					return data, nil
				}
				break
			}
		}
	}
}

// data returns all values from the datastore.getter.
func (c *connection) data(ctx context.Context) (map[string][]byte, error) {
	ctx, span := otel.Tracer("autoupdate").Start(ctx, "next")
	defer span.End()

	if c.tid == 0 {
		c.tid = c.autoupdate.topic.LastID()
	}

	recorder := datastore.NewRecorder(c.autoupdate.datastore)
	restricter := c.autoupdate.restricter(recorder, c.uid)

	if err := c.kb.Update(ctx, restricter); err != nil {
		return nil, fmt.Errorf("create keys for keysbuilder: %w", err)
	}

	data, err := restricter.Get(ctx, c.kb.Keys()...)
	if err != nil {
		return nil, fmt.Errorf("get restricted data: %w", err)
	}
	c.hotkeys = recorder.Keys()

	c.filter.filter(data)

	return data, nil
}
