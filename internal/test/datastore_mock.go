package test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"

	"gopkg.in/yaml.v3"
)

// MockDatastore implements the autoupdate.Datastore interface.
type MockDatastore struct {
	changeListeners []func(map[string]json.RawMessage) error
	CountGetCalled  int
	DatastoreValues
}

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

// NewMockDatastore create a MockDatastore with data.
func NewMockDatastore(data map[string]string) *MockDatastore {
	conv := make(map[string]json.RawMessage)
	for k, v := range data {
		conv[k] = []byte(v)
	}
	return &MockDatastore{
		DatastoreValues: DatastoreValues{
			Data: conv,
		},
	}
}

// Get returnes the values for the given keys. If the keys exist in the Data
// attribute, the values are returned.
//
// If a key is not present in Data, a default value is returned.
//
// If the key starts with "error", an error it thrown.
//
// If the key ends with "_id", "1" is returned.
//
// If the key ends with "_ids", "[1,2]" is returned.
//
// In any other case, "some value" is returned.
func (d *MockDatastore) Get(ctx context.Context, keys ...string) ([]json.RawMessage, error) {
	data := make(map[string]json.RawMessage, len(keys))
	for _, key := range keys {
		value, err := d.DatastoreValues.Value(key)
		if err != nil {
			return nil, err
		}

		data[key] = value
	}

	values := make([]json.RawMessage, len(keys))
	for i, key := range keys {
		values[i] = data[key]
	}
	d.CountGetCalled++
	return values, nil
}

// RegisterChangeListener registers a change listener.
func (d *MockDatastore) RegisterChangeListener(f func(map[string]json.RawMessage) error) {
	d.changeListeners = append(d.changeListeners, f)
}

// Send sends keys to the mock that can be received with a listerer registered
// by RegisterChangeListener().
func (d *MockDatastore) Send(keys []string) {
	data := make(map[string]json.RawMessage, len(keys))
	for _, key := range keys {
		data[key] = nil
	}

	for _, f := range d.changeListeners {
		f(data)
	}
}

// SendValues updates the mock and calls Send afterwards.
func (d *MockDatastore) SendValues(data map[string]string) {
	conv := make(map[string]json.RawMessage, len(data))
	keys := make([]string, 0, len(data))
	for k, v := range data {
		conv[k] = []byte(v)
		keys = append(keys, k)
	}

	d.Update(conv)
	d.Send(keys)
}

// StartServer starts a httptest.Server and returns the url.
//
// The server returns the Date from the MockDatastore as the real datastore
// server would return them.
func (d *MockDatastore) StartServer(closed <-chan struct{}) string {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data struct {
			Keys []string `json:"requests"`
		}
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, fmt.Sprintf("Invalid json input: %v", err), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		values, err := d.Get(r.Context(), data.Keys...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		responceData := make(map[string]map[string]map[string]json.RawMessage)
		for i, key := range data.Keys {
			value := values[i]

			if value == nil {
				continue
			}

			keyParts := strings.SplitN(key, "/", 3)

			if _, ok := responceData[keyParts[0]]; !ok {
				responceData[keyParts[0]] = make(map[string]map[string]json.RawMessage)
			}

			if _, ok := responceData[keyParts[0]][keyParts[1]]; !ok {
				responceData[keyParts[0]][keyParts[1]] = make(map[string]json.RawMessage)
			}
			responceData[keyParts[0]][keyParts[1]][keyParts[2]] = json.RawMessage(value)
		}

		json.NewEncoder(w).Encode(responceData)
	}))
	return ts.URL
}

// DatastoreValues returns data for the test.MockDatastore and the
// test.DatastoreServer.
//
// If OnlyData is false, fake data is generated.
type DatastoreValues struct {
	mu   sync.RWMutex
	Data map[string]json.RawMessage
}

// Value returns a value for a key. If the value does not exist, the second
// return value is false.
func (d *DatastoreValues) Value(key string) (json.RawMessage, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	v, ok := d.Data[key]
	if ok {
		return v, nil
	}

	return nil, nil
}

// Update updates the values from the Datastore.
//
// This does not send a signal to the listeners.
func (d *DatastoreValues) Update(data map[string]json.RawMessage) {
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
