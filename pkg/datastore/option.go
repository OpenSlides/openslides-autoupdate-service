package datastore

import (
	"context"
	"fmt"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/oserror"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
)

// Option to configure datastore.New()
type Option func(*Datastore, environment.Getenver) ([]environment.Variable, func(context.Context), error)

// WithVoteCount adds the poll/vote_count field.
func WithVoteCount() Option {
	eventer := func() (<-chan time.Time, func() bool) {
		timer := time.NewTimer(time.Second)
		return timer.C, timer.Stop
	}

	return func(ds *Datastore, lookup environment.Getenver) ([]environment.Variable, func(context.Context), error) {
		voteCountSource, usedEnv := NewVoteCountSource(lookup)
		ds.keySource["poll/vote_count"] = voteCountSource
		background := func(ctx context.Context) {
			voteCountSource.Connect(ctx, eventer, oserror.Handle)
		}
		return usedEnv, background, nil
	}
}

// WithHistory adds the posibility to fetch history data.
func WithHistory() Option {
	return func(ds *Datastore, lookup environment.Getenver) ([]environment.Variable, func(context.Context), error) {
		datastoreSource, usedEnv, err := NewSourceDatastore(lookup)
		if err != nil {
			return nil, nil, fmt.Errorf("init datastore: %w", err)
		}
		ds.history = datastoreSource

		return usedEnv, nil, nil
	}
}

// WithDefaultSource uses a different (not postgres) source. Helpful for testing.
func WithDefaultSource(source Source) Option {
	return func(ds *Datastore, lookup environment.Getenver) ([]environment.Variable, func(context.Context), error) {
		ds.defaultSource = source
		return nil, nil, nil
	}
}

// // WithProjector adds the field projection/content
// func WithProjector() Option {
// 	return func(ds *Datastore, lookup environment.Getenver) ([]environment.Variable, func(context.Context), error) {
// 		projector.Register(ds, slide.Slides())
// 		return nil, nil, nil
// 	}

// }
