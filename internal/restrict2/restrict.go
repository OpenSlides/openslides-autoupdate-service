package restrict

import (
	"context"
	"fmt"

	oldRestrict "github.com/OpenSlides/openslides-autoupdate-service/internal/restrict"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/flow"
)

// Restricter TODO
type Restricter struct {
	flow flow.Flow
	// TODO: Probably needs a mutex
	attributes attribute.Map
}

// New initializes a restricter
func New(flow flow.Flow) *Restricter {
	return &Restricter{
		flow:       flow,
		attributes: *attribute.NewMap(),
	}
}

// Get returns the full (unrestricted) data.
func (r *Restricter) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	return r.flow.Get(ctx, keys...)
}

// Update updates the precalculation.
func (r *Restricter) Update(ctx context.Context, updateFn func(map[dskey.Key][]byte, error)) {
	r.flow.Update(ctx, updateFn)
}

// ForUser returns a getter that returns restricted data for an user id.
//
// Fetches keys from the flow and pre calculates the restriction for each key.
//
// TODO: Remove the ctx here and add it on every Get() call in the restricter
func (r *Restricter) ForUser(ctx context.Context, userID int) (context.Context, flow.Getter) {
	ctx, todoOldRestricter := oldRestrict.Middleware(ctx, r.flow, userID)
	return ctx, &restrictedGetter{
		todoOldRestricter: todoOldRestricter,
		userID:            userID,
		restricter:        r,
	}
}

func (r *Restricter) precalculate(ctx context.Context, keys []dskey.Key) error {
	// TODO: Make concurency save
	if len(keys) == 0 {
		return nil
	}

	// Group by collection
	// Call collection restricter for each collection with an attribute map
	return nil
}

type restrictedGetter struct {
	todoOldRestricter flow.Getter
	userID            int
	restricter        *Restricter
}

func (r *restrictedGetter) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	needPrecalculate := r.restricter.attributes.NeedCalc(keys)

	if err := r.restricter.precalculate(ctx, needPrecalculate); err != nil {
		return nil, fmt.Errorf("restricter precalculate: %w", err)
	}

	// TODO: only send not implemented collections
	return r.todoOldRestricter.Get(ctx, keys...)
}
