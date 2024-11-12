package dsfetch

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/flow"
)

//go:generate sh -c "go run gen_fields/main.go > fields_generated.go"

// Fetch provides functions to access the fields of the datastore.
//
// Fetch is not save for concurent use. One Fetch object AND its value can only be
// used in one goroutine.
type Fetch struct {
	getter flow.Getter

	requested map[dskey.Key][]setLazyer
}

// New initializes a Request object.
func New(getter flow.Getter) *Fetch {
	r := Fetch{
		getter:    getter,
		requested: make(map[dskey.Key][]setLazyer),
	}
	return &r
}

func (f *Fetch) getOneKey(ctx context.Context, key dskey.Key) ([]byte, error) {
	idKey := key.IDField()

	data, err := f.getter.Get(ctx, key, idKey)
	if err != nil {
		return nil, fmt.Errorf("fetching key: %w", err)
	}

	if data[idKey] == nil {
		return nil, DoesNotExistError(key)
	}

	return data[key], nil
}

// Execute loads all requested keys from the datastore.
func (f *Fetch) Execute(ctx context.Context) error {
	if len(f.requested) == 0 {
		return nil
	}

	defer func() {
		// Clear all requested fields in the end. Even if errors happened.
		clear(f.requested)
	}()

	keys := make([]dskey.Key, 0, len(f.requested)*2)
	for key := range f.requested {
		keys = append(keys, key, key.IDField())
	}

	data, err := f.getter.Get(ctx, keys...)
	if err != nil {
		return fmt.Errorf("fetching all requested keys: %w", err)
	}

	for key, value := range data {
		if data[key.IDField()] == nil {
			return fmt.Errorf("key has no _id field. Executing %d keys: %w", len(f.requested), DoesNotExistError(key))
		}

		for _, exec := range f.requested[key] {
			if err := exec.setLazy(value); err != nil {
				return fmt.Errorf("executing field %s: %w", key, err)
			}
		}
	}

	return nil
}

// Get calls the getter the flow was created with.
func (f *Fetch) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	return f.getter.Get(ctx, keys...)
}

type setLazyer interface {
	setLazy([]byte) error
}

// DoesNotExistError is thrown when an object does not exist.
type DoesNotExistError dskey.Key

func (e DoesNotExistError) Error() string {
	return fmt.Sprintf("%s does not exist.", dskey.Key(e))
}
