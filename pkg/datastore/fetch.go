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

// Objects fetch mor ethan one object, but all from same type, at once.
//
// The argument `value` has to be a pointer to a struct. This struct will be copied for the objects, created during reading.
// Example:
//
// type dbUser struct {
// 	Username     string     `json:"username"`
// 	Title        string     `json:"title"`
// 	FirstName    string     `json:"first_name"`
// 	LastName     string     `json:"last_name"`
// 	DefaultLevel string     `json:"default_structure_level"`
// 	Level        string     `json:"structure_level_$" replacement: "ReplMeetingID"`
//  Groups       []int      `json:"group_$_ids" replacement: "ReplMeetingID"`
//  ReplMeetingID int
// }
//
// If one of the fields contain a $, then the field is handeled as a template
// Field. Then there are 2 cases:
// 1. A known replacement: The templated field tag, used for the json-field-name
// get one more key replacement with the name of a field thar will hold the value,
// in the example above the ReplMeetingID.
// The ReplMeetingID will contain the concrete value and without json it will not be
// read from database and will stay unchanged.
// 2. Without replacement-key in field tag we will do like in Object and read all the keys.
//    This case in't implemented here and should be copied
//
func Objects(ctx context.Context, ds Getter, collection string, ids []int, value interface{}) (records []interface{}, keys []string, err error) {
	var dbValues []json.RawMessage
	fieldNames, indices, err := GetStructJsonNames(value)
	if err != nil {
		return nil, nil, err
	}

	for _, id := range ids {
		for _, fieldName := range fieldNames {
			keys = append(keys, fmt.Sprintf("%s/%d/%s", collection, id, fieldName))
		}
	}

	dbValues, err = ds.Get(ctx, keys...)
	if err != nil {
		return nil, nil, fmt.Errorf("fetching data: %w", err)
	}

	totalIndex := 0
	for i := 0; i < len(ids); i++ {
		record := value                             // shallow copy, should suffice, otherwise use deep copy lib
		v := reflect.ValueOf(&record).Elem().Elem() // interface needs 2 Elem()
		for _, index := range indices {
			dbValue := dbValues[totalIndex]
			if len(dbValue) == 0 {
				// Field does not exist in db.
				totalIndex++
				continue
			}

			f := v.Field(index)
			if err := json.Unmarshal(dbValue, f.Addr().Interface()); err != nil {
				return nil, nil, fmt.Errorf("decoding error") // "decoding %dth field, fqfield `%s`, value `%s`: %w", i+1, keys[idToKey[i]], dbValue, err)
			}

			fmt.Printf("%03d: record:%02d no:%d: %s\n", totalIndex, i, index, string(dbValues[totalIndex]))
			totalIndex++
		}
		records = append(records, record)
	}

	for no, _ := range records {
		fmt.Printf("%02d", no)
	}

	return records, keys, nil
}

// DoesNotExistError is thowen by the methods of a Fether when an field does not
// exist.
type DoesNotExistError string

func (e DoesNotExistError) Error() string {
	return fmt.Sprintf("%s does not exist.", string(e))
}
