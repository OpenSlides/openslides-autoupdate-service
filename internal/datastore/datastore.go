// Package datastore connects to the datastore-reader-service.
//
// On the one end, it sends http requests to the datastore-service, on the other
// end, it impelements the permission.ExternalDataProvider interface.
//
// At least for now, there is no caching.
package datastore

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const urlPath = "/internal/datastore/reader/get_many"

// Datastore connects to the datastore service. It implements the
// permission.ExternalDataProvider interface.
type Datastore struct {
	Addr string
}

// Get fetches a list of fqfields from the datastore.
func (db *Datastore) Get(ctx context.Context, fqfields ...string) ([]json.RawMessage, error) {
	keyValues, err := db.requestKeys(ctx, fqfields)
	if err != nil {
		return nil, fmt.Errorf("request keys: %w", err)
	}

	values := make([]json.RawMessage, len(fqfields))
	for i, key := range fqfields {
		v, ok := keyValues[key]
		if !ok {
			values[i] = nil
			continue
		}

		values[i] = v
	}
	return values, nil
}

func (db *Datastore) url() string {
	addr := "http://localhost:9010"
	if db.Addr != "" {
		addr = db.Addr
	}

	return addr + urlPath
}

// requestKeys request a list of keys by the datastore. If an error happens, no
// key is returned.
func (db *Datastore) requestKeys(ctx context.Context, keys []string) (map[string]json.RawMessage, error) {
	requestData, err := keysToGetManyRequest(keys)
	if err != nil {
		return nil, fmt.Errorf("creating GetManyRequest: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", db.url(), bytes.NewReader(requestData))
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("requesting keys `%v`: %w", keys, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("datastore returned status %s", resp.Status)
		}
		return nil, fmt.Errorf("datastore returned status %s: %s", resp.Status, body)
	}

	responseData, err := getManyResponceToKeyValue(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parse responce: %w", err)
	}

	return responseData, nil
}

// keysToGetManyRequest a json envoding of the get_many request.
func keysToGetManyRequest(keys []string) (json.RawMessage, error) {
	request := struct {
		Requests []string `json:"requests"`
	}{keys}
	return json.Marshal(request)
}

// getManyResponceToKeyValue reads the responce from the getMany request and
// returns the content as key-values.
func getManyResponceToKeyValue(r io.Reader) (map[string]json.RawMessage, error) {
	var data map[string]map[string]map[string]json.RawMessage
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return nil, fmt.Errorf("decoding responce: %w", err)
	}

	keyValue := make(map[string]json.RawMessage)
	for collection, idField := range data {
		for id, fieldValue := range idField {
			for field, value := range fieldValue {
				keyValue[collection+"/"+id+"/"+field] = value
			}
		}
	}
	return keyValue, nil
}
