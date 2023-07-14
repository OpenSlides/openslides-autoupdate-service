package datastore

import (
	"context"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/slide"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/flow"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
)

// Option to configure datastore.New()
type Option func(*Datastore, environment.Environmenter) (func(context.Context, func(error)), error)

// WithVoteCount adds the poll/vote_count field.
func WithVoteCount() Option {
	eventer := func() (<-chan time.Time, func() bool) {
		timer := time.NewTimer(time.Second)
		return timer.C, timer.Stop
	}

	return func(ds *Datastore, lookup environment.Environmenter) (func(context.Context, func(error)), error) {
		voteCountSource := newFlowVoteCount(lookup)
		ds.additionalFlows["poll/vote_count"] = voteCountSource
		background := func(ctx context.Context, errorHandler func(error)) {
			voteCountSource.Connect(ctx, eventer, errorHandler)
		}
		return background, nil
	}
}

// WithDefaultFlow uses a different (not postgres) flow. Helpful for testing.
func WithDefaultFlow(flow flow.Flow) Option {
	return func(ds *Datastore, lookup environment.Environmenter) (func(context.Context, func(error)), error) {
		ds.flow = flow
		return nil, nil
	}
}

// WithProjector activates the field projection/content
func WithProjector() Option {
	return func(ds *Datastore, lookup environment.Environmenter) (func(context.Context, func(error)), error) {
		projector.Register(ds, slide.Slides())
		return nil, nil
	}
}
