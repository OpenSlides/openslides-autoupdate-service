package dsfetch

import (
	"context"
	"encoding/json"
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

	for fqfield, value := range data {
		if fqfield.Field == "id" && value == nil {
			r.err = DoesNotExistError(fqfield)
			return r.err
		}

		exec := r.requested[fqfield]
		if exec == nil {
			continue
		}
		if err := exec.execute(value); err != nil {
			r.err = fmt.Errorf("executing field %q: %w", fqfield, err)
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

// ValueRequiredInt is a lazy value from the datastore.
type ValueRequiredInt struct {
	value    int
	isNull   bool
	executed bool

	lazies []*int

	request *Fetch
}

// Value returns the value.
func (v *ValueRequiredInt) Value(ctx context.Context) (int, bool, error) {
	if v.request.err != nil {
		return 0, false, v.request.err
	}

	if v.executed {
		return v.value, !v.isNull, nil
	}

	if err := v.request.Execute(ctx); err != nil {
		return 0, false, fmt.Errorf("executing request: %w", err)
	}

	return v.value, !v.isNull, nil
}

// ErrorLater is like Value but does not return an error.
//
// If an error happs, it is saved internaly. Make sure to call request.Err() later to
// access it.
func (v *ValueRequiredInt) ErrorLater(ctx context.Context) (int, bool) {
	if v.request.err != nil {
		return 0, false
	}

	if v.executed {
		return v.value, !v.isNull
	}

	if err := v.request.Execute(ctx); err != nil {
		return 0, false
	}

	return v.value, !v.isNull
}

// execute will be called from request.
func (v *ValueRequiredInt) execute(p []byte) error {
	if p == nil {
		v.isNull = true
	} else {
		if err := json.Unmarshal(p, &v.value); err != nil {
			return fmt.Errorf("decoding value %q: %v", p, err)
		}
	}

	for i := 0; i < len(v.lazies); i++ {
		*v.lazies[i] = v.value
	}

	v.executed = true
	return nil
}

// DoesNotExistError is thrown when an object does not exist.
type DoesNotExistError datastore.Key

func (e DoesNotExistError) Error() string {
	return fmt.Sprintf("%s does not exist.", datastore.Key(e))
}
