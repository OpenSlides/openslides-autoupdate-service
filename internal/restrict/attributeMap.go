package restrict

import (
	"context"
	"fmt"
	"sync"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

// AttributeMap holds attributes for each restriction mod
type AttributeMap struct {
	mu   sync.RWMutex // TODO: This is a bad place for the lock.
	data map[dskey.Key]*collection.Attributes
}

// NewAttributeMap initializes an AttributeMap
func NewAttributeMap() *AttributeMap {
	return &AttributeMap{
		data: make(map[dskey.Key]*collection.Attributes),
	}
}

// Add adds a value to the map.
func (am *AttributeMap) Add(modeKey dskey.Key, value *collection.Attributes) {
	am.mu.Lock()
	defer am.mu.Unlock()

	am.data[modeKey] = value
}

// Get returns an attribute pointer to a restriction mod field. Do not modify it.
func (am *AttributeMap) Get(ctx context.Context, fetch *dsfetch.Fetch, mperms perm.MeetingPermission, modeKey dskey.Key) (*collection.Attributes, error) {
	am.mu.RLock()
	value := am.data[modeKey]
	am.mu.RUnlock()

	if value != nil {
		return value, nil
	}

	restricter, err := restrictModefunc(modeKey.Collection, modeKey.Field)
	if err != nil {
		return nil, fmt.Errorf("restricter for %s, %s: %w", modeKey.Collection, modeKey.Field, err)
	}

	if err := restricter(ctx, fetch, mperms, am, modeKey.ID); err != nil {
		return nil, fmt.Errorf("restrict: %w", err)
	}

	am.mu.RLock()
	value = am.data[modeKey]
	am.mu.RUnlock()

	if value == nil {
		return nil, fmt.Errorf("how can this happen?????")
	}

	return value, nil
}

// SameAs sets the attribute for 'toModeKey' to the same value as 'fromModeKey'.
func (am *AttributeMap) SameAs(ctx context.Context, fetch *dsfetch.Fetch, mperms perm.MeetingPermission, toModeKey, fromModeKey dskey.Key) error {
	v, err := am.Get(ctx, fetch, mperms, fromModeKey)
	if err != nil {
		return fmt.Errorf("get other mode: %w", err)
	}

	am.Add(toModeKey, v)
	return nil
}

// RestrictModeIDs returns a map from collection/mode to a set of ids.
func (am *AttributeMap) RestrictModeIDs() map[collection.CM]set.Set[int] {
	result := make(map[collection.CM]set.Set[int])
	for key := range am.data {
		cm := collection.CM{Collection: key.Collection, Mode: key.Field}
		if result[cm].IsNotInitialized() {
			result[cm] = set.New[int]()
		}
		result[cm].Add(key.ID)
	}
	return result
}
