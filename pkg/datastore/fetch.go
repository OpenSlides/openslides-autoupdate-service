package datastore

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
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
func (f *Fetcher) Object(ctx context.Context, value interface{}, fqIDFmt string, a ...interface{}) {
	if f.err != nil {
		return
	}

	fqID := fmt.Sprintf(fqIDFmt, a...)
	keys, err := Object(ctx, f.ds, fqID, value)
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

// Object fetches an object at once.
//
// The argument `value` has to be a pointer to a struct. The json-tags have to
// be field names from the models.yml. For example:
//
// type dbUser struct {
//  ID        int               `json:"id"`
//  Username  string            `json:"username"`
//  Title     string            `json:"title"`
//  FirstName string            `json:"first_name"`
//  LastName  string            `json:"last_name"`
//  Level     map[string]string `json:"structure_level_$"`
//  Groups    map[int][]int     `json:"group_$_ids"``
// }
//
// If one of the fields contain a $, then the field is handeled as a template
// Field. In this case the value has to be a map from string to the field type.
// As a special case it is possible to use int as the map key. This can be used
// for related-list fields.
//
// Objects writes the Attributes of the `value` struct. The first return
// value are the fqFields that where requested.
func Object(ctx context.Context, ds Getter, fqid string, value interface{}) ([]string, error) {
	v := reflect.ValueOf(value).Elem()
	t := reflect.TypeOf(v.Interface())
	var keys []string
	// unknownTemplateKeys are template keys that could not be found in the
	// database.
	var unknownTemplateKeys []string
	// idToKey is an index from the field-idx to the db key. -1 means, that the field has no
	// db key.
	var idToKey []int
	templates := make(map[int][]string)
	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)
		tag := f.Tag.Get("json")
		if tag == "" {
			idToKey = append(idToKey, -1)
			continue
		}

		commaIndex := strings.Index(tag, ",")
		if commaIndex >= 0 {
			tag = tag[:commaIndex]
		}

		keys = append(keys, fqid+"/"+tag)
		idToKey = append(idToKey, len(keys)-1)

		if strings.Contains(tag, "$") {
			templateKey := fqid + "/" + tag
			var replacements []string
			if err := get(ctx, ds, templateKey, &replacements); err != nil {
				var errNotExist DoesNotExistError
				if errors.As(err, &errNotExist) {
					// Skip fields that do not exist
					idToKey[len(idToKey)-1] = -1
					keys = keys[:len(keys)-1]
					unknownTemplateKeys = append(unknownTemplateKeys, templateKey)
					continue
				}
				return nil, fmt.Errorf("fetching template key %s: %v", templateKey, err)
			}

			templates[i] = replacements
			for _, r := range replacements {
				newKey := strings.Replace(templateKey, "$", "$"+r, 1)
				keys = append(keys, newKey)
			}
		}

	}

	dbValues, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("fetching data: %w", err)
	}

	for i := 0; i < v.NumField(); i++ {
		if idToKey[i] == -1 {
			continue
		}

		dbValue := dbValues[idToKey[i]]
		if len(dbValue) == 0 {
			// Field does not exist in db.
			continue
		}

		if templates[i] != nil {
			// The field is a template field.
			if v.Field(i).Kind() != reflect.Map {
				panic(fmt.Sprintf("%dth field is a template field and has to be represented as a map", i))
			}

			m := reflect.MakeMapWithSize(t.Field(i).Type, len(templates[i]))
			for j, key := range templates[i] {
				dbValue := dbValues[idToKey[i]+j+1]
				if dbValue == nil {
					// Field does not exist in the db.
					continue
				}

				mkey := reflect.ValueOf(key)

				if t.Field(i).Type.Key().Kind() == reflect.Int {
					num, err := strconv.Atoi(key)
					if err != nil {
						return nil, err
					}
					mkey = reflect.ValueOf(num)
				}

				val := reflect.New(t.Field(i).Type.Elem())

				if err := json.Unmarshal(dbValue, val.Interface()); err != nil {
					return nil, fmt.Errorf("decodig %dth field (template=%s): %w", i+1, key, err)
				}

				m.SetMapIndex(mkey, val.Elem())
			}

			v.Field(i).Set(m)
			continue
		}

		if err := json.Unmarshal(dbValue, v.Field(i).Addr().Interface()); err != nil {
			return nil, fmt.Errorf("decoding %dth field, fqfield `%s`, value `%s`: %w", i+1, keys[idToKey[i]], dbValue, err)
		}
	}
	keys = append(keys, unknownTemplateKeys...)
	return keys, nil
}

// DoesNotExistError is thowen by the methods of a Fether when an field does not
// exist.
type DoesNotExistError string

func (e DoesNotExistError) Error() string {
	return fmt.Sprintf("%s does not exist.", string(e))
}
