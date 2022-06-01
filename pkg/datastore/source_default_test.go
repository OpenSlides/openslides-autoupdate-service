package datastore_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

func TestSourceDefaultRequestCount(t *testing.T) {
	var mu sync.Mutex
	var count int
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		count++
		mu.Unlock()
		handleGetMany(w, r)
	}))

	for _, tt := range []struct {
		maxKeysPerRequest int
		keyCount          int
		expectCount       int
	}{
		{2, 4, 2},
		{2, 5, 3},
		{10, 10, 1},
		{1, 10, 10},
	} {
		t.Run(fmt.Sprintf("%d-%d", tt.maxKeysPerRequest, tt.keyCount), func(t *testing.T) {
			count = 0
			sd := datastore.NewSourceDatastore(ts.URL, nil, tt.maxKeysPerRequest)
			keys := make([]datastore.Key, tt.keyCount)
			for i := 0; i < len(keys); i++ {
				keys[i] = datastore.Key{Collection: "coll", ID: i + 1, Field: "field"}
			}

			got, err := sd.Get(context.Background(), keys...)
			if err != nil {
				t.Fatalf("Get: %v", err)
			}

			if count != tt.expectCount {
				t.Errorf("got %d requests, expected %d", count, tt.expectCount)
			}

			for _, k := range keys {
				if string(got[k]) != `"value"` {
					t.Errorf("got for key %s value %s, expected \"value\"", k, got[k])
				}
			}
		})
	}
}

func handleGetMany(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Keys []string `json:"requests"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, fmt.Sprintf("Invalid json input: %v", err), http.StatusBadRequest)
		return
	}

	responseData := make(map[string]map[string]map[string]json.RawMessage)
	for _, key := range data.Keys {
		value := []byte(`"value"`)

		keyParts := strings.SplitN(key, "/", 3)

		if _, ok := responseData[keyParts[0]]; !ok {
			responseData[keyParts[0]] = make(map[string]map[string]json.RawMessage)
		}

		if _, ok := responseData[keyParts[0]][keyParts[1]]; !ok {
			responseData[keyParts[0]][keyParts[1]] = make(map[string]json.RawMessage)
		}
		responseData[keyParts[0]][keyParts[1]][keyParts[2]] = value
	}

	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		http.Error(w, fmt.Sprintf("encoding response: %v", err), 400)
		return
	}
}
