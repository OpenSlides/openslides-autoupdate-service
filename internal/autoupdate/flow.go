package autoupdate

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/metric"
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
//	                  cache <-> projector
//	vote-service <->
type Flow struct {
	flow.Flow

	cache     *cache.Cache
	projector *projector.Projector
	postgres  *datastore.FlowPostgres
}

// NewFlow initializes a flow for the autoupdate service.
func NewFlow(lookup environment.Environmenter, messageBus flow.Updater, skipVoteService bool) (*Flow, func(context.Context, func(error)), error) {
	postgres, err := datastore.NewFlowPostgres(lookup, messageBus)
	if err != nil {
		return nil, nil, fmt.Errorf("init postgres: %w", err)
	}

	vote := datastore.NewFlowVoteCount(lookup)

	var dataFlow flow.Flow = postgres
	background := func(context.Context, func(error)) {}
	if !skipVoteService {
		dataFlow = flow.Combine(
			postgres,
			map[string]flow.Flow{"poll/has_voted_user_ids": vote},
		)

		eventer := func() (<-chan time.Time, func() bool) {
			timer := time.NewTimer(time.Second)
			return timer.C, timer.Stop
		}

		background = func(ctx context.Context, errorHandler func(error)) {
			vote.Connect(ctx, eventer, errorHandler)
		}
	}

	cache := cache.New(dataFlow)
	projector := projector.NewProjector(cache, slide.Slides())

	flow := Flow{
		Flow:      projector,
		cache:     cache,
		projector: projector,
		postgres:  postgres,
	}

	metric.Register(flow.metric)

	return &flow, background, nil
}

// ResetCache clears the cache.
func (f *Flow) ResetCache() {
	f.cache.Reset()
	f.projector.Reset()
}

func (f *Flow) metric(values metric.Container) {
	values.Add("datastore_cache_key_len", f.cache.Len())
	values.Add("datastore_cache_size", f.cache.Size())
}

func (f *Flow) historyInformation(ctx context.Context, fqid string, w io.Writer) error {
	return f.postgres.HistoryInformation(ctx, fqid, w)
}
