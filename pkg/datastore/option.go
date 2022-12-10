package datastore

import (
	"context"
	"fmt"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/slide"
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
		voteCountSource := newVoteCountSource(lookup)
		ds.keySource["poll/vote_count"] = voteCountSource
		background := func(ctx context.Context, errorHandler func(error)) {
			voteCountSource.Connect(ctx, eventer, errorHandler)
		}
		return background, nil
	}
}

// WithHistory adds the posibility to fetch history data.
func WithHistory() Option {
	return func(ds *Datastore, lookup environment.Environmenter) (func(context.Context, func(error)), error) {
		datastoreSource, err := newSourceDatastore(lookup)
		if err != nil {
			return nil, fmt.Errorf("init datastore: %w", err)
		}
		ds.history = datastoreSource

		return nil, nil
	}
}

// WithDefaultSource uses a different (not postgres) source. Helpful for testing.
func WithDefaultSource(source Source) Option {
	return func(ds *Datastore, lookup environment.Environmenter) (func(context.Context, func(error)), error) {
		ds.defaultSource = source
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

// WithVoteDecryptPubKeySource adds the field organization/vote_decrypt_public_main_key
func WithVoteDecryptPubKeySource() Option {
	return func(ds *Datastore, lookup environment.Environmenter) (func(context.Context, func(error)), error) {
		source := NewVoteDecryptPubKeySource(lookup)
		ds.keySource["organization/vote_decrypt_public_main_key"] = source
		return nil, nil
	}
}
