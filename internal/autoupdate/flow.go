package autoupdate

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/keysbuilder"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/metric"
	"github.com/OpenSlides/openslides-go/datastore"
	"github.com/OpenSlides/openslides-go/datastore/cache"
	"github.com/OpenSlides/openslides-go/datastore/flow"
	"github.com/OpenSlides/openslides-go/environment"
	"github.com/OpenSlides/openslides-go/metagen"
)

// Flow is the connection to the database for the autoupdate service.
//
// It connects to postgres and the vote-service. The values get combined and
// cached.
//
//	postgres     <->
//	                  cache
//	vote-service <->
type Flow struct {
	flow.Flow
	cache       *cache.Cache
	initRequest *keysbuilder.Builder
}

// NewFlow initializes a flow for the autoupdate service.
//
// Returns an init function, to initialize the service and a background
// function, that should be called in the background.
func NewFlow(lookup environment.Environmenter, skipVoteService bool) (*Flow, func(context.Context) error, func(context.Context, func(error)), error) {
	postgres, initPostgres, err := datastore.NewFlowPostgres(lookup)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("init postgres: %w", err)
	}

	vote := datastore.NewFlowVoteCount(lookup)

	var dataFlow flow.Flow = postgres
	background := func(context.Context, func(error)) {}
	if !skipVoteService {
		dataFlow = flow.Combine(
			postgres,
			map[string]flow.Flow{"poll/live_votes": vote},
		)

		eventer := func() (<-chan time.Time, func() bool) {
			timer := time.NewTimer(time.Second)
			return timer.C, timer.Stop
		}

		background = func(ctx context.Context, errorHandler func(error)) {
			vote.Connect(ctx, eventer, errorHandler)
		}
	}

	// TODO: Add the option "cache.WithFullMessagebus" as soon as this is fixed:
	// https://github.com/OpenSlides/openslides-meta/pull/629
	cache := cache.New(dataFlow)

	initRequest, err := keysbuilder.ManyFromJSON(bytes.NewReader(metagen.InitAutoupdateRequest))
	if err != nil {
		return nil, nil, nil, fmt.Errorf("building initial request: %w", err)
	}

	flow := Flow{
		Flow:        cache,
		cache:       cache,
		initRequest: initRequest,
	}

	metric.Register(flow.metric)

	init := func(ctx context.Context) error {
		if err := initPostgres(ctx); err != nil {
			return fmt.Errorf("init postgres: %w", err)
		}

		if err := flow.WarmupCache(ctx); err != nil {
			return fmt.Errorf("warmup cache: %w", err)
		}

		return nil
	}

	return &flow, init, background, nil
}

// Snapshot retuns an immutable getter that will not change.
func (f *Flow) Snapshot(notFoundHandler flow.Getter) flow.Getter {
	return f.cache.Snapshot(notFoundHandler)
}

// ResetCache clears the cache.
func (f *Flow) ResetCache() {
	f.cache.Reset()
}

// WarmupCache fetches all keys from metagen.InitAutoupdateRequest
func (f *Flow) WarmupCache(ctx context.Context) error {
	keys, err := f.initRequest.Update(ctx, f)
	if err != nil {
		return fmt.Errorf("build keys: %w", err)
	}

	if _, err := f.Get(ctx, keys...); err != nil {
		return fmt.Errorf("fetch keys: %w", err)
	}

	return nil
}

func (f *Flow) metric(values metric.Container) {
	values.Add("datastore_cache_key_len", f.cache.Len())
	values.Add("datastore_cache_size", f.cache.Size())
}
