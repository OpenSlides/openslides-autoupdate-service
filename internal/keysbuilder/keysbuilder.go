// Package keysbuilder holds a datastructure to get and update requested keys.
package keysbuilder

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/flow"
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
			ids:        []int{key.ID()},
			collection: key.Collection(),
			fieldsMap: fieldsMap{
				fields: map[string]fieldDescription{
					key.Field(): nil,
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

func bodyFieldLen(bodies []body) int {
	sum := 0
	for _, body := range bodies {
		sum += len(body.fieldsMap.fields)
	}
	return sum
}

// Update triggers a key update. It generates the list of keys, that can be
// requested with the Keys() method. It travels the KeysRequests object like a
// tree.
//
// It is not allowed to call builder.Keys() after Update returned an error.
func (b *Builder) Update(ctx context.Context, getter flow.Getter) ([]dskey.Key, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if len(b.bodies) == 0 {
		return nil, nil
	}

	// Start with all keys from all the bodies.
	var err error
	queue := make([]keyDescription, 0, bodyFieldLen(b.bodies))
	for _, body := range b.bodies {
		queue, err = body.appendKeys(queue)
		if err != nil {
			return nil, fmt.Errorf("building keys from bodys: %w", err)
		}
	}

	var keys []dskey.Key

	// neededX contains keys, where the value has to be fetched from the database.
	var neededKeys []dskey.Key
	var neededDescriptions []keyDescription
	for len(queue) > 0 {
		// Get all keys and descriptions.
		for _, kd := range queue {
			keys = append(keys, kd.key)
			if kd.description == nil {
				continue
			}

			neededKeys = append(neededKeys, kd.key)
			neededDescriptions = append(neededDescriptions, keyDescription{key: kd.key, description: kd.description})
		}
		queue = queue[:0]

		if len(neededKeys) == 0 {
			continue
		}

		// Get values for all special (not none) fields.
		data, err := getter.Get(ctx, neededKeys...)
		if err != nil {
			return nil, fmt.Errorf("load needed keys: %w", err)
		}
		neededKeys = neededKeys[:0]

		for _, kd := range neededDescriptions {
			// This are fields that do not exist or the user has not the
			// permission to see them.
			if data[kd.key] == nil {
				continue
			}

			var err error
			queue, err = kd.description.appendKeys(kd.key, data[kd.key], queue)
			if err != nil {
				var invalidErr *json.UnmarshalTypeError
				if errors.As(err, &invalidErr) {
					// value has wrong type.
					return nil, ValueError{key: kd.key, gotType: invalidErr.Value, expectType: invalidErr.Type, err: err}
				}
				return nil, fmt.Errorf("appending keys for key %s: %w", kd.key, err)
			}
		}

		// Clear processed.
		neededDescriptions = neededDescriptions[:0]
	}
	return keys, nil
}
