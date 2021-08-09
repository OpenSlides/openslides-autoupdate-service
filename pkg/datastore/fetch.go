package datastore

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
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
// The method Fetcher.Err() can be used to get the error. After it is called
// once, the error is cleared. So the next call to Fether after Err() is not a
// noop.
//
// Make sure to call Fetcher.Err() at the end to see, if an error happend.
type Fetcher struct {
	ds   Getter
	keys []string
	err  error
}

// NewFetcher initializes a Fetcher object.
func NewFetcher(ds Getter) *Fetcher {
	return &Fetcher{ds: ds}
}

// Fetch gets a value from the datastore and saves it into the argument `value`.
//
// If the object, that the key belongs to does not exist, no error is thrown.
//
// To get the error, call f.Err().
func (f *Fetcher) Fetch(ctx context.Context, value interface{}, keyFmt string, a ...interface{}) {
	if f.err != nil {
		return
	}

	fqfield := fmt.Sprintf(keyFmt, a...)
	f.keys = append(f.keys, fqfield)

	fields, err := f.ds.Get(ctx, fqfield)
	if err != nil {
		f.err = fmt.Errorf("getting data from datastore: %w", err)
		return
	}

	if fields[0] == nil {
		return
	}

	if err := json.Unmarshal(fields[0], value); err != nil {
		f.err = fmt.Errorf("unpacking value of %q: %w", fqfield, err)
	}
	return
}

// FetchIfExist is like Fetch but if the element that the key belongs to does
// not exist, then a DoesNotExistError is returned.
func (f *Fetcher) FetchIfExist(ctx context.Context, value interface{}, keyFmt string, a ...interface{}) {
	if f.err != nil {
		return
	}

	fqfield := fmt.Sprintf(keyFmt, a...)
	keyParts := strings.Split(fqfield, "/")
	if len(keyParts) != 3 {
		f.err = fmt.Errorf("invalid key %q", fqfield)
		return
	}

	fqid := keyParts[0] + "/" + keyParts[1]
	idField := fqid + "/id"

	f.keys = append(f.keys, idField, fqfield)
	fields, err := f.ds.Get(ctx, idField, fqfield)
	if err != nil {
		f.err = fmt.Errorf("getting data from datastore: %w", err)
		return
	}

	if fields[0] == nil {
		f.err = DoesNotExistError(fqid)
		return
	}
	if fields[1] == nil {
		return
	}

	if err := json.Unmarshal(fields[1], value); err != nil {
		f.err = fmt.Errorf("unpacking value of %q: %w", fqfield, err)
	}
	return
}

// Object returns a json object for the given fqid with all given fields.
//
// If one field does not exist in the datastore, then it is returned as nil.
//
// If the object does not exist, then a DoesNotExistError is thrown.
func (f *Fetcher) Object(ctx context.Context, fqID string, fields ...string) map[string]json.RawMessage {
	if f.err != nil {
		return nil
	}

	keys := make([]string, len(fields)+1)
	keys[0] = fqID + "/id"
	for i := 0; i < len(fields); i++ {
		keys[i+1] = fqID + "/" + fields[i]
	}

	f.keys = append(f.keys, keys...)
	vals, err := f.ds.Get(ctx, keys...)
	if err != nil {
		f.err = fmt.Errorf("fetching data: %w", err)
		return nil
	}

	if vals[0] == nil {
		f.err = DoesNotExistError(fqID)
		return nil
	}

	object := make(map[string]json.RawMessage, len(fields))
	for i := 0; i < len(fields); i++ {
		object[fields[i]] = vals[i+1]
	}

	return object
}

// Field returns function to fetch all existing field.
func (f *Fetcher) Field() Fields {
	return Fields{f}
}

// Keys returns all datastore keys that where fetched in the process.
func (f *Fetcher) Keys() []string {
	return f.keys
}

// Err returns the error that happend at a method call. If no error happend,
// then Err() returns nil.
//
// Calling this method clears the error. So a second call to Err() does not
// return the error anymore.
func (f *Fetcher) Err() error {
	err := f.err
	f.err = nil
	return err
}

// FetchFunc is a function that fetches a value. It has the signature of
// fetch.Fetch() or fetch.FetchIfExist().
type FetchFunc func(ctx context.Context, value interface{}, keyFmt string, a ...interface{})

// Int fetches an integer from the datastore.
func Int(ctx context.Context, fetch FetchFunc, keyFmt string, a ...interface{}) int {
	var value int
	fetch(ctx, &value, keyFmt, a...)
	return value
}

// Ints fetches an int slice from the datastore.
func Ints(ctx context.Context, fetch FetchFunc, keyFmt string, a ...interface{}) []int {
	var value []int
	fetch(ctx, &value, keyFmt, a...)
	return value
}

// String fetches a string from the datastore.
func String(ctx context.Context, fetch FetchFunc, keyFmt string, a ...interface{}) string {
	var value string
	fetch(ctx, &value, keyFmt, a...)
	return value
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
		return fmt.Errorf("getting %q from datastore: %w", fqfield, err)
	}

	if fields[0] == nil {
		return DoesNotExistError(fqfield)
	}

	if err := json.Unmarshal(fields[0], value); err != nil {
		return fmt.Errorf("unpacking value: %w", err)
	}
	return nil
}

// object returns a json object for the given fqid with all given fields.
//
// If one field does not exist in the datastore, then it is returned as nil.
func object(ctx context.Context, ds Getter, fqid string, fields []string) (map[string]json.RawMessage, []string, error) {
	keys := make([]string, len(fields))
	for i := 0; i < len(fields); i++ {
		keys[i] = fqid + "/" + fields[i]
	}

	vals, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, keys, fmt.Errorf("fetching data: %w", err)
	}

	object := make(map[string]json.RawMessage, len(fields))
	for i := 0; i < len(fields); i++ {
		object[fields[i]] = vals[i]
	}

	return object, keys, nil
}

// DoesNotExistError is thrown by the methods of a Fetcher when an object does
// not exist.
type DoesNotExistError string

func (e DoesNotExistError) Error() string {
	return fmt.Sprintf("%s does not exist.", string(e))
}
