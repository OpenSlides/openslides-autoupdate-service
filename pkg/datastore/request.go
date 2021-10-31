package datastore

import (
	"context"
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
		getter: getter,
	}
	return &r
}

// Execute loads all requested keys from the datastore.
func (r *Request) Execute(ctx context.Context) error {
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
		delete(r.requested, fqfield)
	}

	r.err = nil
	return nil
}

type executer interface {
	execute([]byte) error
}
