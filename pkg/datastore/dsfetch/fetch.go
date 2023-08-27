package dsfetch

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

//go:generate sh -c "go run gen_fields/main.go > fields_generated.go"

// Getter is the same as datastore.Getter
type Getter interface {
	Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error)
}

// Fetch provides functions to access the fields of the datastore.
//
// Fetch is not save for concurent use. One Fetch object AND its value can only be
// used in one goroutine.
type Fetch struct {
	getter Getter

	requested map[dskey.Key]executer
}

// New initializes a Request object.
func New(getter Getter) *Fetch {
	r := Fetch{
		getter:    getter,
		requested: make(map[dskey.Key]executer),
	}
	return &r
}

// Execute loads all requested keys from the datastore.
func (r *Fetch) Execute(ctx context.Context) error {
	defer func() {
		// Clear all requested fields in the end. Even if errors happened.
		r.requested = make(map[dskey.Key]executer)
	}()

	keys := make([]dskey.Key, 0, len(r.requested)*2)
	for key := range r.requested {
		keys = append(keys, key, key.IDField())
	}

	if len(keys) == 0 {
		return nil
	}

	data, err := r.getter.Get(ctx, keys...)
	if err != nil {
		return fmt.Errorf("fetching all requested keys: %w", err)
	}

	for key, value := range data {
		if data[key.IDField()] == nil {
			return DoesNotExistError(key)
		}

		exec := r.requested[key]
		if exec == nil {
			continue
		}
		if err := exec.execute(value); err != nil {
			return fmt.Errorf("executing field %q: %w", key, err)
		}
	}

	return nil
}

type executer interface {
	execute([]byte) error
}

// DoesNotExistError is thrown when an object does not exist.
type DoesNotExistError dskey.Key

func (e DoesNotExistError) Error() string {
	return fmt.Sprintf("%s does not exist.", dskey.Key(e))
}
