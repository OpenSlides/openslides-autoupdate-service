// Package keysbuilder holds a datastructure to get and update requested keys.
package keysbuilder

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

// Builder builds the keys. It is not save for concourent use. There is one
// Builder instance per client. It is not allowed to call builder.Update() more
// then once or at the same time as builder.Keys(). It is ok to call
// builder.Keys() at the same time more then once.
//
// Has to be created with keysbuilder.FromJSON() or keysbuilder.ManyFromJSON().
type Builder struct {
	mu sync.Mutex

	bodies []body
	keys   []dskey.Key
}

// FromKeys creates a keysbuilder from a list of keys.
func FromKeys(rawKeys ...string) (*Builder, error) {
	b := new(Builder)
	if len(rawKeys) == 0 || rawKeys[0] == "" {
		return b, nil
	}

	keys := make([]dskey.Key, len(rawKeys))
	for i, k := range rawKeys {
		key, err := dskey.FromString(k)
		if err != nil {
			// TODO LAST ERROR
			return nil, fmt.Errorf("invalid key: %s", k)
		}
		keys[i] = key
	}

	for _, key := range keys {
		body := body{
			ids:        []int{key.ID},
			collection: key.Collection,
			fieldsMap: fieldsMap{
				fields: map[string]fieldDescription{
					key.Field: nil,
				},
			},
		}
		b.bodies = append(b.bodies, body)
	}
	return b, nil
}

// FromBuilders creates a new keysbuilder from a list of other builders.
func FromBuilders(builders ...*Builder) *Builder {
	builder := new(Builder)
	for _, b := range builders {
		builder.bodies = append(builder.bodies, b.bodies...)
	}
	return builder
}

// Update triggers a key update. It generates the list of keys, that can be
// requested with the Keys() method. It travels the KeysRequests object like a
// tree.
//
// It is not allowed to call builder.Keys() after Update returned an error.
func (b *Builder) Update(ctx context.Context, getter datastore.Getter) (err error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	defer func() {
		// Reset keys if an error happens
		if err != nil {
			b.keys = b.keys[:0]
		}
	}()

	if len(b.bodies) == 0 {
		return nil
	}

	// Start with all keys from all the bodies.
	process := make(map[dskey.Key]fieldDescription)
	for _, body := range b.bodies {
		body.keys(process)
	}

	b.keys = b.keys[:0]
	var needed []dskey.Key
	processed := make(map[dskey.Key]fieldDescription)
	for {
		// Get all keys and descriptions
		for key, description := range process {
			b.keys = append(b.keys, key)
			if description == nil {
				continue
			}

			needed = append(needed, key)
			processed[key] = description
		}

		if len(needed) == 0 {
			break
		}

		// Get values for all special (not none) fields.
		data, err := getter.Get(ctx, needed...)
		if err != nil {
			return fmt.Errorf("load needed keys: %w", err)
		}

		// Clear process and needed without freeing the memory.
		needed = needed[:0]
		for k := range process {
			delete(process, k)
		}

		for key, description := range processed {
			// This are fields that do not exist or the user has not the
			// permission to see them.
			if data[key] == nil {
				continue
			}

			if err := description.keys(key, data[key], process); err != nil {
				var invalidErr *json.UnmarshalTypeError
				if errors.As(err, &invalidErr) {
					// value has wrong type.
					return ValueError{key: key, gotType: invalidErr.Value, expectType: invalidErr.Type, err: err}
				}
				return err
			}
		}

		// Clear processed.
		for k := range processed {
			delete(processed, k)
		}
	}
	return nil
}

// Keys returns the keys.
//
// Make sure to call Update() or Keys() will return an empty list.
func (b *Builder) Keys() []dskey.Key {
	b.mu.Lock()
	defer b.mu.Unlock()

	return append(b.keys[:0:0], b.keys...)
}

// buildGenericKey returns a valid key when the collection and id are already
// together.
//
// buildGenericKey("motion/5", "title") -> "motion/5/title".
func buildGenericKey(collectionID string, field string) dskey.Key {
	key, err := dskey.FromString(collectionID + "/" + field)
	_ = err // TODO: Can this happen?

	return key
}

func buildCollectionID(collection string, id int) string {
	return collection + "/" + strconv.Itoa(id)
}
