// Package keysbuilder holds a datastructure to get and update requested keys.
package keysbuilder

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

const keySep = "/"

// Builder builds the keys. It is not save for concourent use. There is one
// Builder instance per client. It is not allowed to call builder.Update() more
// then once or at the same time as builder.Keys(). It is ok to call
// builder.Keys() at the same time more then once.
//
// Has to be created with keysbuilder.FromJSON() or keysbuilder.ManyFromJSON().
type Builder struct {
	bodies []body
	keys   []string
}

// FromKeys creates a keysbuilder from a list of keys.
func FromKeys(keys []string) (*Builder, error) {
	b := new(Builder)
	if len(keys) == 0 || keys[0] == "" {
		return b, nil
	}

	if invalid := datastore.InvalidKeys(keys...); len(invalid) != 0 {
		return nil, InvalidError{msg: fmt.Sprintf("Invalid keys: %v", invalid)}
	}

	for _, key := range keys {
		parts := strings.Split(key, "/")
		id, _ := strconv.Atoi(parts[1])
		body := body{
			ids:        []int{id},
			collection: parts[0],
			fieldsMap: fieldsMap{
				fields: map[string]fieldDescription{
					parts[2]: nil,
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
	process := make(map[string]fieldDescription)
	for _, body := range b.bodies {
		body.keys(process)
	}

	b.keys = b.keys[:0]
	var needed []string
	processed := make(map[string]fieldDescription)
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
func (b *Builder) Keys() []string {
	return append(b.keys[:0:0], b.keys...)
}

// buildGenericKey returns a valid key when the collection and id are already
// together.
//
// buildGenericKey("motion/5", "title") -> "motion/5/title".
func buildGenericKey(collectionID string, field string) string {
	return collectionID + keySep + field
}

func buildCollectionID(collection string, id int) string {
	return collection + keySep + strconv.Itoa(id)
}
