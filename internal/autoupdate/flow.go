package autoupdate

import (
	"context"
	"fmt"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/slide"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/cache"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/flow"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
)

// Flow is the connection to the database for the autoupdate service.
//
// It connects to postgres and the vote-service. The values get combined and
// cached. Then the projection/content fields are calculated and the results are
// cached again.
//
//	postgres     <->
//	                  cache2 <-> projector <-> cache1
//	vote-service <->
//
// The cache1 contains most of the data. cache2 only contains the data needed
// for the projector.
//
// TODO: The cache1 needs an argument, that it only caches projector/content
// fields, or the cache has to be part of the projector.
type Flow struct {
	flow.Flow
	cacheReset func()
}

// NewFlow initializes a flow for the autoupdate service.
func NewFlow(lookup environment.Environmenter, messageBus flow.Updater) (*Flow, func(context.Context, func(error)), error) {
	postgres, err := datastore.NewFlowPostgres(lookup, messageBus)
	if err != nil {
		return nil, nil, fmt.Errorf("init postgres: %w", err)
	}

	vote := datastore.NewFlowVoteCount(lookup)

	combined := flow.Combine(
		postgres,
		map[string]flow.Flow{"poll/vote_count": vote},
	)

	cache2 := cache.New(combined)
	projector := projector.NewProjector(cache2, slide.Slides())
	cache1 := cache.New(projector)

	cacheReset := func() {
		cache1.Reset()
		cache2.Reset()
		projector.Reset()
	}

	eventer := func() (<-chan time.Time, func() bool) {
		timer := time.NewTimer(time.Second)
		return timer.C, timer.Stop
	}

	background := func(ctx context.Context, errorHandler func(error)) {
		vote.Connect(ctx, eventer, errorHandler)
	}

	return &Flow{
		Flow:       cache1,
		cacheReset: cacheReset,
	}, background, nil
}

// ResetCache clears the cache.
func (f *Flow) ResetCache() {
	f.cacheReset()
}

// TODO
// func (f *Flow) metric(values metric.Container) {
// 	values.Add("datastore_cache_key_len", f.cache.len())
// 	values.Add("datastore_cache_size", f.cache.size())

// 	hitCount := atomic.LoadUint64(&f.metricGetHitCount)
// 	values.Add("datastore_get_calls", int(hitCount))
// }
