package datastore

import (
	"context"
	"encoding/json"
	"fmt"
)

// Getter can get values from keys.
//
// The Datastore object implements this interface.
type Getter interface {
	Get(ctx context.Context, keys ...string) ([]json.RawMessage, error)
}

// Fetcher is a helper to fetch many keys from the datastore.
//
// This object is meant to be called like a function. Do not store it in a
// struct.
//
// The methods do not return an error. If an error happens, it is saved
// internaly. As soon, as an error happens, all later calls to methods of that
// fetcher are noops.
//
// Make sure to call Fetcher.Error() at the end to see, if an error happend.
type Fetcher struct {
	ds   Getter
	keys []string
	err  error
}

// NewFetcher initializes a Fetcher object.
func NewFetcher(ds Getter) *Fetcher {
	return &Fetcher{ds: ds}
}

// Object fetches a struct from the datastore.
func (f *Fetcher) Object(ctx context.Context, fields []string, fqIDFmt string, a ...interface{}) map[string]json.RawMessage {
	if f.err != nil {
		return nil
	}

	fqID := fmt.Sprintf(fqIDFmt, a...)
	object, keys, err := Object(ctx, f.ds, fqID, fields)
	if err != nil {
		f.err = fmt.Errorf("fetching object %s: %w", fqID, err)
		return nil
	}
	f.keys = append(f.keys, keys...)
	return object
}

// Value fetches a value from the datastore.
func (f *Fetcher) Value(ctx context.Context, value interface{}, keyFmt string, a ...interface{}) {
	if f.err != nil {
		return
	}

	key := fmt.Sprintf(keyFmt, a...)
	if err := get(ctx, f.ds, key, value); err != nil {
		f.err = fmt.Errorf("fetching %s: %w", key, err)
		return
	}
	f.keys = append(f.keys, key)
}

// Int fetches an integer from the datastore.
func (f *Fetcher) Int(ctx context.Context, keyFmt string, a ...interface{}) int {
	var i int
	f.Value(ctx, &i, keyFmt, a...)
	return i
}

// Ints fetches an int slice from the datastore.
func (f *Fetcher) Ints(ctx context.Context, keyFmt string, a ...interface{}) []int {
	var iSlice []int
	f.Value(ctx, &iSlice, keyFmt, a...)
	return iSlice
}

// String fetches a string from the datastore.
func (f *Fetcher) String(ctx context.Context, keyFmt string, a ...interface{}) string {
	var s string
	f.Value(ctx, &s, keyFmt, a...)
	return s
}

// Keys returns all datastore keys that where fetched in the process.
func (f *Fetcher) Keys() []string {
	return f.keys
}

// Error returns the error that happend at a method call. If no error happend,
// then Error() returns nil.
func (f *Fetcher) Error() error {
	return f.err
}

// get returns a value from the datastore and unpacks it in to the argument
// value.
//
// The argument value has to be an non nil pointer.
//
// get returns a DoesNotExistError if the value des not exist in the datastore.
func get(ctx context.Context, ds Getter, fqfield string, value interface{}) error {
	fields, err := ds.Get(ctx, fqfield)
	if err != nil {
		return fmt.Errorf("getting data from datastore: %w", err)
	}

	if fields[0] == nil {
		return DoesNotExistError(fqfield)
	}

	if err := json.Unmarshal(fields[0], value); err != nil {
		return fmt.Errorf("unpacking value: %w", err)
	}
	return nil
}

// Object returns a json object for the given fqid with all given fields.
//
// If one field does not exist in the datastore, then it is returned as nil.
func Object(ctx context.Context, ds Getter, fqid string, fields []string) (map[string]json.RawMessage, []string, error) {
	keys := make([]string, len(fields))
	for i := 0; i < len(fields); i++ {
		keys[i] = fqid + "/" + fields[i]
	}

	vals, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, nil, fmt.Errorf("fetching data: %w", err)
	}

	object := make(map[string]json.RawMessage, len(fields))
	for i := 0; i < len(fields); i++ {
		object[fields[i]] = vals[i]
	}

	return object, keys, nil
}

// DoesNotExistError is thowen by the methods of a Fether when an field does not
// exist.
type DoesNotExistError string

func (e DoesNotExistError) Error() string {
	return fmt.Sprintf("%s does not exist.", string(e))
}
