package dsmock

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"gopkg.in/yaml.v3"
)

// YAMLData creates key values from a yaml object.
//
// It is expected, that the input is a constant string. So there can not be any
// error at runtime. Therefore this function does not return an error but panics
// to get the developer a fast feetback.
func YAMLData(input string) map[string][]byte {
	input = strings.ReplaceAll(input, "\t", "  ")

	var db map[string]interface{}
	if err := yaml.Unmarshal([]byte(input), &db); err != nil {
		panic(err)
	}

	data := make(map[string][]byte)
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
					fqfield := fmt.Sprintf("%s/%d/%s", dbKey, id, fieldName)
					bs, err := json.Marshal(fieldValue)
					if err != nil {
						panic(fmt.Errorf("creating test db. Key %s: %w", fqfield, err))
					}
					data[fqfield] = bs
				}

				idField := fmt.Sprintf("%s/%d/id", dbKey, id)
				data[idField] = []byte(strconv.Itoa(id))
			}

		case 2:
			field, ok := dbValue.(map[string]interface{})
			if !ok {
				panic(fmt.Errorf("invalid object type: got %T, expected map[string]interface{}", dbValue))
			}

			for fieldName, fieldValue := range field {
				fqfield := fmt.Sprintf("%s/%s/%s", parts[0], parts[1], fieldName)
				bs, err := json.Marshal(fieldValue)
				if err != nil {
					panic(fmt.Errorf("creating test db. Key %s: %w", fqfield, err))
				}
				data[fqfield] = bs
			}

			idField := fmt.Sprintf("%s/%s/id", parts[0], parts[1])
			data[idField] = []byte(parts[1])

		case 3:
			bs, err := json.Marshal(dbValue)
			if err != nil {
				panic(fmt.Errorf("creating test db. Key %s: %w", dbKey, err))
			}
			data[dbKey] = bs

			idField := fmt.Sprintf("%s/%s/id", parts[0], parts[1])
			data[idField] = []byte(parts[1])
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
// The function starts a mock datastore server in the background. It gets
// closed, when the closed channel is closed.
func NewMockDatastore(closed <-chan struct{}, data map[string][]byte) *MockDatastore {
	source := NewStubWithUpdate(data, NewCounter)
	ds := &MockDatastore{
		source:    source,
		Datastore: datastore.New(source),
	}

	ds.counter = source.Middlewares()[0].(*Counter)

	return ds
}

// Get calls the Get() method of the datastore.
func (d *MockDatastore) Get(ctx context.Context, keys ...string) (map[string][]byte, error) {
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
func (d *MockDatastore) Requests() [][]string {
	return d.counter.Requests()
}

// ResetRequests resets the list returned by Requests().
func (d *MockDatastore) ResetRequests() {
	d.counter.Reset()
}

// KeysRequested returns true, if all given keys where requested.
func (d *MockDatastore) KeysRequested(keys ...string) bool {
	requestedKeys := make(map[string]bool)
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
func (d *MockDatastore) Send(data map[string][]byte) {
	d.source.Send(data)
}

// Update implements the datastore.Updater interface.
func (d *MockDatastore) Update(ctx context.Context) (map[string][]byte, error) {
	return d.source.Update(ctx)
}
