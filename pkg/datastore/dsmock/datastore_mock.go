package dsmock

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
	"gopkg.in/yaml.v3"
)

// YAMLData creates key values from a yaml object.
//
// It is expected, that the input is a constant string. So there can not be any
// error at runtime. Therefore this function does not return an error but panics
// to get the developer a fast feetback.
func YAMLData(input string) map[dskey.Key][]byte {
	input = strings.ReplaceAll(input, "\t", "  ")

	var db map[string]interface{}
	if err := yaml.Unmarshal([]byte(input), &db); err != nil {
		panic(err)
	}

	data := make(map[dskey.Key][]byte)
	for dbKey, dbValue := range db {
		parts := strings.Split(dbKey, "/")
		switch len(parts) {
		case 1:
			map1, ok := dbValue.(map[interface{}]interface{})
			if !ok {
				panic(fmt.Errorf("invalid type in db key %s: %T", dbKey, dbValue))
			}
			for rawID, rawObject := range map1 {
				id, ok := rawID.(int)
				if !ok {
					panic(fmt.Errorf("invalid id type: got %T expected int", rawID))
				}
				field, ok := rawObject.(map[string]interface{})
				if !ok {
					panic(fmt.Errorf("invalid object type: got %T, expected map[string]interface{}", rawObject))
				}

				for fieldName, fieldValue := range field {
					key := dskey.MustKey(fmt.Sprintf("%s/%d/%s", dbKey, id, fieldName))
					bs, err := json.Marshal(fieldValue)
					if err != nil {
						panic(fmt.Errorf("creating test db. Key %s: %w", key, err))
					}
					data[key] = bs
				}

				idKey := dskey.MustKey(fmt.Sprintf("%s/%d/id", dbKey, id))
				data[idKey] = []byte(strconv.Itoa(id))
			}

		case 2:
			field, ok := dbValue.(map[string]interface{})
			if !ok {
				panic(fmt.Errorf("invalid object type: got %T, expected map[string]interface{}", dbValue))
			}

			for fieldName, fieldValue := range field {
				fqfield := dskey.MustKey(fmt.Sprintf("%s/%s/%s", parts[0], parts[1], fieldName))
				bs, err := json.Marshal(fieldValue)
				if err != nil {
					panic(fmt.Errorf("creating test db. Key %s: %w", fqfield, err))
				}
				data[fqfield] = bs
			}

			idKey := dskey.MustKey(fmt.Sprintf("%s/%s/id", parts[0], parts[1]))
			data[idKey] = []byte(parts[1])

		case 3:
			key := dskey.MustKey(dbKey)
			bs, err := json.Marshal(dbValue)
			if err != nil {
				panic(fmt.Errorf("creating test db. Key %s: %w", dbKey, err))
			}

			data[key] = bs

			idKey := dskey.MustKey(fmt.Sprintf("%s/%s/id", parts[0], parts[1]))
			data[idKey] = []byte(parts[1])
		default:
			panic(fmt.Errorf("invalid db key %s", dbKey))
		}
	}

	for k, v := range data {
		if bytes.Equal(v, []byte("null")) {
			data[k] = nil
		}
	}

	return data
}

// MockDatastore implements the autoupdate.Datastore interface.
type MockDatastore struct {
	*datastore.Datastore
	source  *StubWithUpdate
	counter *Counter
	err     error
}

// NewMockDatastore create a MockDatastore with data.
//
// It is a wrapper around the datastore.Datastore object.
func NewMockDatastore(data map[dskey.Key][]byte) (*MockDatastore, func(context.Context, func(error))) {
	source := NewStubWithUpdate(data, NewCounter)
	rawDS, bg, err := datastore.New(environment.ForTests{}, nil, datastore.WithDefaultSource(source))
	if err != nil {
		panic(err)
	}

	ds := &MockDatastore{
		source:    source,
		Datastore: rawDS,
	}

	ds.counter = source.Middlewares()[0].(*Counter)

	return ds, bg
}

// Get calls the Get() method of the datastore.
func (d *MockDatastore) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	if d.err != nil {
		return nil, d.err
	}

	return d.Datastore.Get(ctx, keys...)
}

// InjectError lets the next calls to Get() return the injected error.
func (d *MockDatastore) InjectError(err error) {
	d.err = err
}

// Requests returns a list of all requested keys.
func (d *MockDatastore) Requests() [][]dskey.Key {
	return d.counter.Requests()
}

// ResetRequests resets the list returned by Requests().
func (d *MockDatastore) ResetRequests() {
	d.counter.Reset()
}

// KeysRequested returns true, if all given keys where requested.
func (d *MockDatastore) KeysRequested(keys ...dskey.Key) bool {
	requestedKeys := make(map[dskey.Key]bool)
	for _, l := range d.Requests() {
		for _, k := range l {
			requestedKeys[k] = true
		}
	}

	for _, k := range keys {
		if !requestedKeys[k] {
			return false
		}
	}
	return true
}

// Send updates the data.
//
// This method is unblocking. If you want to fetch data afterwards, make sure to
// block until data is processed. For example with RegisterChanceListener.
func (d *MockDatastore) Send(data map[dskey.Key][]byte) {
	d.source.Send(data)
}

// Update implements the datastore.Updater interface.
func (d *MockDatastore) Update(ctx context.Context) (map[dskey.Key][]byte, error) {
	return d.source.Update(ctx)
}
