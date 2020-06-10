package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
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
	DatastoreValues
}

// NewDatastoreServer creates a new DatastoreServer.
func NewDatastoreServer() *DatastoreServer {
	ts := new(DatastoreServer)
	ts.TS = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data getManyRequest
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, fmt.Sprintf("Invalid json input: %v", err), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		responceData := make(map[string]map[string]map[string]json.RawMessage)
		for _, key := range data.Keys {
			value, exist, err := ts.DatastoreValues.Value(key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if !exist {
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
		ts.RequestCount++
	}))
	return ts
}
