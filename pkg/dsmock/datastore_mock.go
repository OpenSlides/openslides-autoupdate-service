package dsmock

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"gopkg.in/yaml.v3"
)

// YAMLData creates key values from a yaml object.
//
// It is expected, that the input is a constant string. So there can not be any
// error at runtime. Therefore this function does not return an error but panics
// to get the developer a fast feetback.
func YAMLData(input string) map[string]string {
	input = strings.ReplaceAll(input, "\t", "  ")

	var db map[string]interface{}
	if err := yaml.Unmarshal([]byte(input), &db); err != nil {
		panic(err)
	}

	data := make(map[string]string)
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
					data[fqfield] = string(bs)
				}

				idField := fmt.Sprintf("%s/%d/id", dbKey, id)
				data[idField] = strconv.Itoa(id)
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
				data[fqfield] = string(bs)
			}

			idField := fmt.Sprintf("%s/%s/id", parts[0], parts[1])
			data[idField] = parts[1]

		case 3:
			bs, err := json.Marshal(dbValue)
			if err != nil {
				panic(fmt.Errorf("creating test db. Key %s: %w", dbKey, err))
			}
			data[dbKey] = string(bs)

			idField := fmt.Sprintf("%s/%s/id", parts[0], parts[1])
			data[idField] = parts[1]
		default:
			panic(fmt.Errorf("invalid db key %s", dbKey))
		}
	}

	return data
}

// Stub are data that can be used as a datastore value.
type Stub map[string]string

// Get implements the Getter interface.
func (s Stub) Get(_ context.Context, keys ...string) (map[string][]byte, error) {
	converted := make(map[string][]byte, len(keys))
	for k, v := range map[string]string(s) {
		converted[k] = []byte(v)
	}
	return converted, nil
}

// MockDatastore implements the autoupdate.Datastore interface.
type MockDatastore struct {
	*datastore.Datastore
	server *DatastoreServer
	err    error
}

// NewMockDatastore create a MockDatastore with data.
func NewMockDatastore(closed <-chan struct{}, data map[string]string) *MockDatastore {
	dsServer := NewDatastoreServer(closed, data)

	s := &MockDatastore{
		server: dsServer,
	}

	s.Datastore = datastore.New(dsServer.TS.URL, closed, func(err error) { log.Println(err) }, s.server)

	return s
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

// Send updates the data.
//
// This method is unblocking. If you want to fetch data afterwards, make sure to
// block until data is processed. For example with RegisterChanceListener.
func (d *MockDatastore) Send(data map[string]string) {
	d.server.Send(data)
}

// Update implements the datastore.Updater interface.
func (d *MockDatastore) Update(close <-chan struct{}) (map[string][]byte, error) {
	return d.server.Update(close)
}

// datastoreValues returns data for the test.MockDatastore and the
// test.DatastoreServer.
//
// If OnlyData is false, fake data is generated.
type datastoreValues struct {
	mu   sync.RWMutex
	Data map[string][]byte
}

func newDatastoreValues(data map[string]string) *datastoreValues {
	conv := make(map[string][]byte)
	for k, v := range data {
		conv[k] = []byte(v)
	}

	return &datastoreValues{
		Data: conv,
	}
}

// value returns a value for a key. If the value does not exist, the second
// return value is false.
func (d *datastoreValues) value(key string) (json.RawMessage, error) {
	if d == nil {
		return nil, nil
	}

	d.mu.RLock()
	defer d.mu.RUnlock()

	v, ok := d.Data[key]
	if ok {
		return v, nil
	}

	return nil, nil
}

// set updates the values from the Datastore.
//
// This does not send a signal to the listeners.
func (d *datastoreValues) set(data map[string][]byte) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.Data == nil {
		d.Data = data
		return
	}

	for key, value := range data {
		d.Data[key] = value
	}
}
