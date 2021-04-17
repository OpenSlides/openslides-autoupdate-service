package datastore

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// Getter can get values from keys.
type Getter interface {
	Get(ctx context.Context, keys ...string) ([]json.RawMessage, error)
}

// Fetcher is a helper to fetch many keys from the datastore.
//
// This object is meant to be called like a function. Do not keep it around.
//
// The methods do not return an error. It is saved internaly. As soon, as an
// error happens, all later calls are noops.
//
// Make sure to call Fetcher.Error() to see, if an error happend.
type Fetcher struct {
	ds   Getter
	keys []string
	err  error
}

// NewFetcher initializes a Fetcher object. Make sure to check the error in the
// end of your function.
func NewFetcher(ds Getter) *Fetcher {
	return &Fetcher{ds: ds}
}

// Object fetches a struct from the datastore.
func (f *Fetcher) Object(ctx context.Context, value interface{}, fqIDFmt string, a ...interface{}) {
	if f.err != nil {
		return
	}

	fqID := fmt.Sprintf(fqIDFmt, a...)
	keys, err := GetObject(ctx, f.ds, fqID, value)
	if err != nil {
		f.err = fmt.Errorf("fetching object %s: %w", fqID, err)
		return
	}
	f.keys = append(f.keys, keys...)

}

// Value fetches a value from the datastore.
func (f *Fetcher) Value(ctx context.Context, value interface{}, keyFmt string, a ...interface{}) {
	if f.err != nil {
		return
	}

	key := fmt.Sprintf(keyFmt, a...)
	if err := Get(ctx, f.ds, key, value); err != nil {
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

func (f *Fetcher) String(ctx context.Context, keyFmt string, a ...interface{}) string {
	var s string
	f.Value(ctx, &s, keyFmt, a...)
	return s
}

// Keys returns all datastore keys that where fetched in the process.
func (f *Fetcher) Keys() []string {
	return f.keys
}

func (f *Fetcher) Error() error {
	return f.err
}

// Get returns a value from the datastore and unpacks it in to the argument value.
//
// The argument value has to be an non nil pointer.
func Get(ctx context.Context, ds Getter, fqfield string, value interface{}) error {
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

// GetIfExist behaves like Get() but does not throw an error if the fqfield does
// not exist.
func GetIfExist(ctx context.Context, ds Getter, fqfield string, value interface{}) error {
	if err := Get(ctx, ds, fqfield, value); err != nil {
		var errDoesNotExist DoesNotExistError
		if !errors.As(err, &errDoesNotExist) {
			return err
		}
	}
	return nil
}

// GetObject fetches an object at once.
//
// The argument value has to be a struct. The json-tags have to be field from
// the models.yml.
func GetObject(ctx context.Context, ds Getter, fqid string, value interface{}) ([]string, error) {
	v := reflect.ValueOf(value).Elem()
	t := reflect.TypeOf(v.Interface())
	var keys []string
	idToKey := make([]int, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)
		tag := f.Tag.Get("json")
		if tag == "" {
			idToKey[i] = -1
			continue
		}

		ci := strings.Index(tag, ",")
		if ci >= 0 {
			tag = tag[:ci]
		}
		keys = append(keys, fqid+"/"+tag)
		idToKey[i] = len(keys) - 1
	}

	fields, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("fetching data: %w", err)
	}

	for i := 0; i < v.NumField(); i++ {
		if idToKey[i] == -1 {
			continue
		}

		dbValue := fields[idToKey[i]]
		if len(dbValue) == 0 {
			// Field does not exist in db.
			continue
		}

		if err := json.Unmarshal(dbValue, v.Field(i).Addr().Interface()); err != nil {
			return nil, fmt.Errorf("decoding %dth field, fqfield `%s`, value `%s`: %w", i+1, keys[idToKey[i]], dbValue, err)
		}
	}
	return keys, nil
}

// DoesNotExistError is thowen when an field does not exist.
type DoesNotExistError string

func (e DoesNotExistError) Error() string {
	return fmt.Sprintf("%s does not exist.", string(e))
}
