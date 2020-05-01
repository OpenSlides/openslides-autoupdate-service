package datastore_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
)

type testServer struct {
	ts           *httptest.Server
	requestCount int
}

func newTestServer() *testServer {
	ts := new(testServer)
	ts.ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data struct {
			Keys []string
		}
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, fmt.Sprintf("Invalid json input: %v", err), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		response := make(map[string]string, len(data.Keys))
		for _, key := range data.Keys {
			response[key] = `"value"`
		}
		json.NewEncoder(w).Encode(response)
		ts.requestCount++
	}))
	return ts
}
