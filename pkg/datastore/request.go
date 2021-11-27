package datastore

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

//go:generate sh -c "go run gen_request/main.go > request_generated.go"

// Request provides functions to access the fields of the datastore.
//
// Request is not save for concurent use. One Request object AND its value can only be
// used in one goroutine.
type Request struct {
	getter Getter
	err    error

	requested map[string]executer
}

// NewRequest initializes a Request object.
func NewRequest(getter Getter) *Request {
	r := Request{
		getter:    getter,
		requested: make(map[string]executer),
	}
	return &r
}

// Execute loads all requested keys from the datastore.
func (r *Request) Execute(ctx context.Context) error {
	defer func() {
		// Clear all requested fields in the end. Even if errors happend.
		r.requested = make(map[string]executer)
	}()

	keys := make([]string, 0, len(r.requested)*2)
	for fqfield := range r.requested {
		keyParts := strings.Split(fqfield, "/")
		if len(keyParts) != 3 {
			return fmt.Errorf("invalid key %q", fqfield)
		}

		fqid := keyParts[0] + "/" + keyParts[1]
		idField := fqid + "/id"

		keys = append(keys, fqfield, idField)
	}

	data, err := r.getter.Get(ctx, keys...)
	if err != nil {
		r.err = fmt.Errorf("fetching all requested keys: %w", err)
		return r.err
	}

	for fqfield, value := range data {
		if strings.HasSuffix(fqfield, "/id") && value == nil {
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
func (r *Request) Err() error {
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

	request *Request
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
