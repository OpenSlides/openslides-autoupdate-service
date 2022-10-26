package dsfetch

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

//go:generate sh -c "go run gen_fields/main.go > fields_generated.go"

// Fetch provides functions to access the fields of the datastore.
//
// Fetch is not save for concurent use. One Fetch object AND its value can only be
// used in one goroutine.
type Fetch struct {
	getter datastore.Getter
	err    error

	requested map[datastore.Key]executer
}

// New initializes a Request object.
func New(getter datastore.Getter) *Fetch {
	r := Fetch{
		getter:    getter,
		requested: make(map[datastore.Key]executer),
	}
	return &r
}

// Execute loads all requested keys from the datastore.
func (r *Fetch) Execute(ctx context.Context) error {
	defer func() {
		// Clear all requested fields in the end. Even if errors happened.
		r.requested = make(map[datastore.Key]executer)
	}()

	keys := make([]datastore.Key, 0, len(r.requested)*2)
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
func (r *Fetch) Err() error {
	return r.err
}

type executer interface {
	execute([]byte) error
}

// DoesNotExistError is thrown when an object does not exist.
type DoesNotExistError datastore.Key

func (e DoesNotExistError) Error() string {
	return fmt.Sprintf("%s does not exist.", datastore.Key(e))
}
