package projector

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// dataGet returns a value from the datastore and unpacks it in to the argument value.
//
// The argument value has to be an non nil pointer.
func dataGet(ctx context.Context, ds Datastore, fqfield string, value interface{}) error {
	fields, err := ds.Get(ctx, fqfield)
	if err != nil {
		return fmt.Errorf("getting data from datastore: %w", err)
	}

	if fields[0] == nil {
		return doesNotExistError(fqfield)
	}

	if err := json.Unmarshal(fields[0], value); err != nil {
		return fmt.Errorf("unpacking value: %w", err)
	}
	return nil
}

// dataGetIfExist behaves like Get() but does not throw an error if the fqfield does
// not exist.
func dataGetIfExist(ctx context.Context, ds Datastore, fqfield string, value interface{}) error {
	if err := dataGet(ctx, ds, fqfield, value); err != nil {
		var errDoesNotExist doesNotExistError
		if !errors.As(err, &errDoesNotExist) {
			return err
		}
	}
	return nil
}

// dataGetObject fetches an object at once.
//
// The argument value has to be a struct with ...TODO
func dataGetObject(ctx context.Context, ds Datastore, fqid string, value interface{}) error {
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
		return fmt.Errorf("fetching data: %w", err)
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
			return fmt.Errorf("decoding %dth field, fqfield `%s`, value `%s`: %w", i+1, keys[idToKey[i]], dbValue, err)
		}
	}
	return nil
}

// doesNotExistError is thowen when an field does not exist.
type doesNotExistError string

func (e doesNotExistError) Error() string {
	return fmt.Sprintf("%s does not exist.", string(e))
}
