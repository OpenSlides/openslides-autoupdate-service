package dsmock

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
)

type getManyRequest struct {
	Keys []string `json:"requests"`
}

// DatastoreServer simulates the Datastore-Service. Only the methods required by the
// autoupdate-service are supported. This is currently only the getMany method.
//
// Has to be created with NewDatastoreServer.
type DatastoreServer struct {
	TS           *httptest.Server
	RequestCount int
	Values       *datastoreValues

	c chan map[string][]byte
}

// NewDatastoreServer creates a new DatastoreServer.
func NewDatastoreServer(closed <-chan struct{}, data map[string]string) *DatastoreServer {
	d := &DatastoreServer{
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
func (d *DatastoreServer) Update(ctx context.Context) (map[string][]byte, error) {
	select {
	case v := <-d.c:
		d.Values.set(v)
		return v, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// Send sends keys to the mock that can be received with Update().
func (d *DatastoreServer) Send(values map[string]string) {
	conv := make(map[string][]byte)
	for k, v := range values {
		conv[k] = nil
		if v != "" {
			conv[k] = []byte(v)
		}

	}
	d.c <- conv
}

func validKey(key string) bool {
	match, err := regexp.MatchString(`^([a-z]+|[a-z][a-z_]*[a-z])/[1-9][0-9]*/[a-z][a-z0-9_]*\$?[a-z0-9_]*$`, key)
	if err != nil {
		panic(err)
	}
	return match
}
