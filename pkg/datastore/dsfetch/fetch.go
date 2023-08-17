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
	err    error

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
	if err := r.err; err != nil {
		r.err = nil
		return err
	}

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
		r.err = fmt.Errorf("fetching all requested keys: %w", err)
		return r.err
	}

	for key, value := range data {
		if data[key.IDField()] == nil {
			r.err = DoesNotExistError(key)
			return r.err
		}

		exec := r.requested[key]
		if exec == nil {
			continue
		}
		if err := exec.execute(value); err != nil {
			r.err = fmt.Errorf("executing field %q: %w", key, err)
			return r.err
		}
	}

	r.err = nil
	return nil
}

// Err returns an error from a previous call.
//
// Resets the error
func (r *Fetch) Err() error {
	err := r.err
	r.err = nil
	return err
}

type executer interface {
	execute([]byte) error
}

// DoesNotExistError is thrown when an object does not exist.
type DoesNotExistError dskey.Key

func (e DoesNotExistError) Error() string {
	return fmt.Sprintf("%s does not exist.", dskey.Key(e))
}
