package datastore

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

// Fetcher is a helper to fetch many keys from the datastore.
//
// The methods do not return an error. If an error happens, it is saved
// internaly. As soon, as an error happens, all later calls to methods of that
// fetcher are noops.
//
// The method Fetcher.Err() can be used to get the error.
//
// Make sure to call Fetcher.Err() at the end to see, if an error happened.
type Fetcher struct {
	getter Getter
	err    error
}

// NewFetcher initializes a Fetcher object.
func NewFetcher(getter Getter) *Fetcher {
	return &Fetcher{getter: getter}
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

	fqfield, err := dskey.FromString(fmt.Sprintf(keyFmt, a...))
	if err != nil {
		f.err = err
		return
	}

	fields, err := f.getter.Get(ctx, fqfield)
	if err != nil {
		f.err = fmt.Errorf("getting data from datastore: %w", err)
		return
	}

	if fields[fqfield] == nil {
		return
	}

	if err := json.Unmarshal(fields[fqfield], value); err != nil {
		f.err = fmt.Errorf("unpacking value of %q: %w", fqfield, err)
	}
}

// FetchIfExist is like Fetch but if the element that the key belongs to does
// not exist, then a DoesNotExistError is returned.
func (f *Fetcher) FetchIfExist(ctx context.Context, value interface{}, keyFmt string, a ...interface{}) {
	if f.err != nil {
		return
	}

	fqfield, err := dskey.FromString(fmt.Sprintf(keyFmt, a...))
	if err != nil {
		f.err = err
		return
	}

	idField := fqfield.IDField()

	fields, err := f.getter.Get(ctx, idField, fqfield)
	if err != nil {
		f.err = fmt.Errorf("getting data from datastore: %w", err)
		return
	}

	if fields[idField] == nil {
		f.err = dsfetch.DoesNotExistError(idField)
		return
	}
	if fields[fqfield] == nil {
		return
	}

	if err := json.Unmarshal(fields[fqfield], value); err != nil {
		f.err = fmt.Errorf("unpacking value of %q: %w", fqfield, err)
	}
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

	keys := make([]dskey.Key, len(fields)+1)
	idKey, err := dskey.FromString(fqID + "/id")
	if err != nil {
		f.err = err
		return nil
	}
	keys[0] = idKey

	for i := 0; i < len(fields); i++ {
		k, err := dskey.FromString(fqID + "/" + fields[i])
		if err != nil {
			f.err = err
			return nil
		}
		keys[i+1] = k
	}

	vals, err := f.getter.Get(ctx, keys...)
	if err != nil {
		f.err = fmt.Errorf("fetching data: %w", err)
		return nil
	}

	if vals[idKey] == nil {
		f.err = dsfetch.DoesNotExistError(idKey)
		return nil
	}

	object := make(map[string]json.RawMessage, len(fields))
	for i := 0; i < len(fields); i++ {
		key, err := dskey.FromString(fqID + "/" + fields[i])
		if err != nil {
			f.err = err
			return nil
		}
		object[fields[i]] = vals[key]
	}
	return object
}

// Err returns the error that happened at a method call. If no error happened,
// then Err() returns nil.
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
