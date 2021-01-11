package tests

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
// permission-service are supported. This is currently only the getMany method.
//
// Has to be created with NewDatastoreServer.
type DatastoreServer struct {
	TS           *httptest.Server
	RequestCount int
}

// NewDatastoreServer creates a new DatastoreServer.
func NewDatastoreServer(data map[string]json.RawMessage) *DatastoreServer {
	ts := new(DatastoreServer)

	ts.TS = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		testData := dataProvider{data: data}
		var requestData getManyRequest
		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, fmt.Sprintf("Invalid json input: %v", err), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		responceData := make(map[string]map[string]map[string]json.RawMessage)
		for _, key := range requestData.Keys {
			result, err := testData.Get(r.Context(), key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if result[0] == nil {
				continue
			}

			keyParts := strings.SplitN(key, "/", 3)

			if _, ok := responceData[keyParts[0]]; !ok {
				responceData[keyParts[0]] = make(map[string]map[string]json.RawMessage)
			}

			if _, ok := responceData[keyParts[0]][keyParts[1]]; !ok {
				responceData[keyParts[0]][keyParts[1]] = make(map[string]json.RawMessage)
			}
			responceData[keyParts[0]][keyParts[1]][keyParts[2]] = json.RawMessage(result[0])
		}

		if err := json.NewEncoder(w).Encode(responceData); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ts.RequestCount++
	}))
	return ts
}
