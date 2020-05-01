package datastore_test

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

type testServer struct {
	ts           *httptest.Server
	requestCount int
}

func newTestServer() *testServer {
	ts := new(testServer)
	ts.ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data getManyRequest
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, fmt.Sprintf("Invalid json input: %v", err), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		responceData := make(map[string]map[string]string)
		for _, r := range data.Requests {
			fqid := r.Collection + "/" + strconv.Itoa(r.IDs[0])
			if _, ok := responceData[fqid]; !ok {
				responceData[fqid] = make(map[string]string)
			}
			responceData[fqid][r.MappedFields[0]] = `"value"`
		}

		json.NewEncoder(w).Encode(responceData)
		ts.requestCount++
	}))
	return ts
}
