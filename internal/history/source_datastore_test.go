package history

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
)

func parseURL(raw string) (host, port, protocol string) {
	parsed, err := url.Parse(raw)
	if err != nil {
		panic(fmt.Sprintf("parsing url %s: %v", raw, err))
	}

	return parsed.Hostname(), parsed.Port(), parsed.Scheme
}

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

			host, port, schema := parseURL(ts.URL)
			env := environment.ForTests(map[string]string{
				"DATASTORE_READER_HOST":       host,
				"DATASTORE_READER_PORT":       port,
				"DATASTORE_READER_PROTOCOL":   schema,
				"DATASTORE_TIMEOUT":           "1s",
				"DATASTORE_MAX_PARALLEL_KEYS": strconv.Itoa(tt.maxKeysPerRequest),
			})

			sd, err := newSourceDatastore(env)
			if err != nil {
				t.Fatalf("Initialize: %v", err)
			}

			keys := make([]dskey.Key, tt.keyCount)
			for i := 0; i < len(keys); i++ {
				keys[i], _ = dskey.FromParts("user", i+1, "username")
			}

			got, err := sd.GetPosition(context.Background(), 0, keys...)
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
