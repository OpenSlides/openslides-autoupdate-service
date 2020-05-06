package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
)

type getManyRequest struct {
	Requests []struct {
		Collection   string   `json:"collection"`
		IDs          []int    `json:"ids"`
		MappedFields []string `json:"mapped_fields"`
	} `json:"requests"`
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

		responceData := make(map[string]map[string]json.RawMessage)
		for _, r := range data.Requests {
			fqid := r.Collection + "/" + strconv.Itoa(r.IDs[0])
			if _, ok := responceData[fqid]; !ok {
				responceData[fqid] = make(map[string]json.RawMessage)
			}

			key := r.Collection + "/" + strconv.Itoa(r.IDs[0]) + "/" + r.MappedFields[0]
			value, exist, err := ts.DatastoreValues.Value(key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if !exist {
				continue
			}

			responceData[fqid][r.MappedFields[0]] = json.RawMessage(value)
		}

		json.NewEncoder(w).Encode(responceData)
		ts.RequestCount++
	}))
	return ts
}
