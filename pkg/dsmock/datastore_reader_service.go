package dsmock

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"sync"
)

type getManyRequest struct {
	Keys []string `json:"requests"`
}

// DatastoreReader simulates the datastore-reader-Service. Only the methods
// required by the autoupdate-service are supported. This is currently only the
// getMany method.
//
// Has to be created with NewDatastoreReader.
type DatastoreReader struct {
	TS            *httptest.Server
	RequestCount  int
	RequestedKeys [][]string
	Values        *datastoreValues

	c chan map[string][]byte
}

// NewDatastoreReader creates a new fake DatastoreServer.
//
// It creates a webserver that handels get_many requests like the reald
// datastore-reader.
//
// If the given channel is closed, the server shuts down.
func NewDatastoreReader(closed <-chan struct{}, data map[string][]byte) *DatastoreReader {
	d := &DatastoreReader{
		Values: newDatastoreValues(data),
		c:      make(chan map[string][]byte),
	}

	d.TS = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data getManyRequest
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, fmt.Sprintf("Invalid json input: %v", err), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		d.RequestedKeys = append(d.RequestedKeys, data.Keys)

		responceData := make(map[string]map[string]map[string]json.RawMessage)
		for _, key := range data.Keys {
			if !validKey(key) {
				http.Error(w, "Key is invalid: "+key, 400)
			}
			value, err := d.Values.value(key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if value == nil {
				continue
			}

			keyParts := strings.SplitN(key, "/", 3)
			if len(keyParts) != 3 {
				http.Error(w, fmt.Sprintf("invalid key %s", key), 500)
				return
			}

			if _, ok := responceData[keyParts[0]]; !ok {
				responceData[keyParts[0]] = make(map[string]map[string]json.RawMessage)
			}

			if _, ok := responceData[keyParts[0]][keyParts[1]]; !ok {
				responceData[keyParts[0]][keyParts[1]] = make(map[string]json.RawMessage)
			}
			responceData[keyParts[0]][keyParts[1]][keyParts[2]] = value
		}

		if err := json.NewEncoder(w).Encode(responceData); err != nil {
			http.Error(w, fmt.Sprintf("Error encoding responceData `%s`: %v", responceData, err), 500)
			return
		}
		d.RequestCount++
	}))

	go func() {
		<-closed
		d.TS.Close()
	}()
	return d
}

// Update returnes keys that have changed. Blocks until keys are send with
// the Send-method.
func (d *DatastoreReader) Update(ctx context.Context) (map[string][]byte, error) {
	select {
	case v := <-d.c:
		d.Values.set(v)
		return v, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// Send sends keys to the mock that can be received with Update().
func (d *DatastoreReader) Send(values map[string][]byte) {
	d.c <- values
}

// Requests returns all keys that where requested.
func (d *DatastoreReader) Requests() [][]string {
	return d.RequestedKeys
}

// ResetRequests resets the returnvalue of Requests().
func (d *DatastoreReader) ResetRequests() {
	d.RequestedKeys = make([][]string, 0)
}

func validKey(key string) bool {
	match, err := regexp.MatchString(`^([a-z]+|[a-z][a-z_]*[a-z])/[1-9][0-9]*/[a-z][a-z0-9_]*\$?[a-z0-9_]*$`, key)
	if err != nil {
		panic(err)
	}
	return match
}

// datastoreValues returns data for the test.MockDatastore and the
// test.DatastoreServer.
//
// If OnlyData is false, fake data is generated.
type datastoreValues struct {
	mu   sync.RWMutex
	Data map[string][]byte
}

func newDatastoreValues(data map[string][]byte) *datastoreValues {
	conv := make(map[string][]byte)
	for k, v := range data {
		if bytes.Equal(v, []byte("null")) {
			conv[k] = nil
			continue
		}
		conv[k] = []byte(v)
	}

	return &datastoreValues{
		Data: conv,
	}
}

// value returns a value for a key. If the value does not exist, the second
// return value is false.
func (d *datastoreValues) value(key string) ([]byte, error) {
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
