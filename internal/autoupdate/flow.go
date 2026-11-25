package autoupdate

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/metric"
	"github.com/OpenSlides/openslides-go/datastore"
	"github.com/OpenSlides/openslides-go/datastore/cache"
	"github.com/OpenSlides/openslides-go/datastore/flow"
	"github.com/OpenSlides/openslides-go/environment"
)

// Flow is the connection to the database for the autoupdate service.
type Flow struct {
	flow.Flow
	cache *cache.Cache
}

// NewFlow initializes a flow for the autoupdate service.
func NewFlow(lookup environment.Environmenter) (*Flow, func(context.Context) error, error) {
	postgres, initPostgres, err := datastore.NewFlowPostgres(lookup)
	if err != nil {
		return nil, nil, fmt.Errorf("init postgres: %w", err)
	}

	// TODO: Add the option "cache.WithFullMessagebus" as soon as this is fixed:
	// https://github.com/OpenSlides/openslides-meta/pull/629
	cache := cache.New(postgres)

	flow := Flow{
		Flow:  cache,
		cache: cache,
	}

	metric.Register(flow.metric)

	return &flow, initPostgres, nil
}

// Snapshot retuns an immutable getter that will not change.
func (f *Flow) Snapshot(notFoundHandler flow.Getter) flow.Getter {
	return f.cache.Snapshot(notFoundHandler)
}

// ResetCache clears the cache.
func (f *Flow) ResetCache() {
	f.cache.Reset()
}

func (f *Flow) metric(values metric.Container) {
	values.Add("datastore_cache_key_len", f.cache.Len())
	values.Add("datastore_cache_size", f.cache.Size())
}
