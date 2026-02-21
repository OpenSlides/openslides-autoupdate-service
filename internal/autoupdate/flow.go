package autoupdate

import (
	"context"
	"fmt"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/metric"
	"github.com/OpenSlides/openslides-go/datastore"
	"github.com/OpenSlides/openslides-go/datastore/cache"
	"github.com/OpenSlides/openslides-go/datastore/flow"
	"github.com/OpenSlides/openslides-go/environment"
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

	cache    *cache.Cache
	postgres *datastore.FlowPostgres
}

// NewFlow initializes a flow for the autoupdate service.
func NewFlow(lookup environment.Environmenter, skipVoteService bool) (*Flow, func(context.Context, func(error)), error) {
	postgres, err := datastore.NewFlowPostgres(lookup)
	if err != nil {
		return nil, nil, fmt.Errorf("init postgres: %w", err)
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

	cache := cache.New(dataFlow, cache.WithFullMessagebus)

	flow := Flow{
		Flow:     cache,
		cache:    cache,
		postgres: postgres,
	}

	metric.Register(flow.metric)

	return &flow, background, nil
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
