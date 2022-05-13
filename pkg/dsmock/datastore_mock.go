package dsmock

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
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
func YAMLData(input string) map[datastore.Key][]byte {
	input = strings.ReplaceAll(input, "\t", "  ")

	var db map[string]interface{}
	if err := yaml.Unmarshal([]byte(input), &db); err != nil {
		panic(err)
	}

	data := make(map[datastore.Key][]byte)
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
					key, err := datastore.KeyFromString(fmt.Sprintf("%s/%d/%s", dbKey, id, fieldName))
					if err != nil {
						panic(err)
					}
					bs, err := json.Marshal(fieldValue)
					if err != nil {
						panic(fmt.Errorf("creating test db. Key %s: %w", key, err))
					}
					data[key] = bs
				}

				idKey, err := datastore.KeyFromString(fmt.Sprintf("%s/%d", dbKey, id))
				if err != nil {
					panic(err)
				}
				data[idKey] = []byte(strconv.Itoa(id))
			}

		case 2:
			field, ok := dbValue.(map[string]interface{})
			if !ok {
				panic(fmt.Errorf("invalid object type: got %T, expected map[string]interface{}", dbValue))
			}

			for fieldName, fieldValue := range field {
				fqfield, err := datastore.KeyFromString(fmt.Sprintf("%s/%s/%s", parts[0], parts[1], fieldName))
				if err != nil {
					panic(err)
				}
				bs, err := json.Marshal(fieldValue)
				if err != nil {
					panic(fmt.Errorf("creating test db. Key %s: %w", fqfield, err))
				}
				data[fqfield] = bs
			}

			idKey, err := datastore.KeyFromString(fmt.Sprintf("%s/%s", parts[0], parts[1]))
			if err != nil {
				panic(err)
			}
			data[idKey] = []byte(parts[1])

		case 3:
			key, err := datastore.KeyFromString(dbKey)
			if err != nil {
				panic(err)
			}
			bs, err := json.Marshal(dbValue)
			if err != nil {
				panic(fmt.Errorf("creating test db. Key %s: %w", dbKey, err))
			}

			data[key] = bs

			idKey, err := datastore.KeyFromString(fmt.Sprintf("%s/%s/id", parts[0], parts[1]))
			if err != nil {
				panic(err)
			}
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
// The function starts a mock datastore server in the background. It gets
// closed, when the closed channel is closed.
func NewMockDatastore(closed <-chan struct{}, data map[datastore.Key][]byte) *MockDatastore {
	source := NewStubWithUpdate(data, NewCounter)
	ds := &MockDatastore{
		source:    source,
		Datastore: datastore.New(source, nil, source),
	}

	ds.counter = source.Middlewares()[0].(*Counter)

	return ds
}

// Get calls the Get() method of the datastore.
func (d *MockDatastore) Get(ctx context.Context, keys ...datastore.Key) (map[datastore.Key][]byte, error) {
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
func (d *MockDatastore) Requests() [][]datastore.Key {
	return d.counter.Requests()
}

// ResetRequests resets the list returned by Requests().
func (d *MockDatastore) ResetRequests() {
	d.counter.Reset()
}

// HistoryInformation writes a fake history.
func (d *MockDatastore) HistoryInformation(ctx context.Context, fqid string, w io.Writer) error {
	w.Write([]byte(`[{"position":42,"user_id": 5,"information": "motion was created","timestamp: 1234567}]`))
	return nil
}

// KeysRequested returns true, if all given keys where requested.
func (d *MockDatastore) KeysRequested(keys ...datastore.Key) bool {
	requestedKeys := make(map[datastore.Key]bool)
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
func (d *MockDatastore) Send(data map[datastore.Key][]byte) {
	d.source.Send(data)
}

// Update implements the datastore.Updater interface.
func (d *MockDatastore) Update(ctx context.Context) (map[datastore.Key][]byte, error) {
	return d.source.Update(ctx)
}
