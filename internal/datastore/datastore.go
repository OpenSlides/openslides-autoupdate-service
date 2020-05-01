package datastore

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const urlPath = "/internal/datastore/reader/get_many"

// Datastore can be used to get values from the datastore-service.
//
// Has to be created with datastore.New().
type Datastore struct {
	url        string
	cache      *cache
	keychanger KeysChangedReceiver
}

// New returns a new Datastore object.
func New(url string, keychanger KeysChangedReceiver) *Datastore {
	return &Datastore{
		cache:      newCache(),
		url:        url + urlPath,
		keychanger: keychanger,
	}
}

// Get returns the value for one or many keys.
func (d *Datastore) Get(ctx context.Context, keys ...string) ([]string, error) {
	values, err := d.cache.getOrSet(ctx, keys, func(keys []string) (map[string]string, error) {
		return d.requestKeys(keys)
	})
	if err != nil {
		return nil, fmt.Errorf("getOrSet for keys `%s`: %w", keys, err)
	}

	return values, nil
}

// KeysChanged blocks until some key have changed. Then, it returns the keys.
func (d *Datastore) KeysChanged() ([]string, error) {
	keys, err := d.keychanger.KeysChanged()
	if err != nil {
		return nil, err
	}

	// TODO: only request keys that exist in the cache.
	data, err := d.requestKeys(keys)
	if err != nil {
		return nil, fmt.Errorf("request values for keys: %w", err)
	}

	d.cache.setIfExist(data)

	return keys, nil
}

// requestKeys request a list of keys by the datastore. It returns the values in
// the same order. If an error happens, no key is returned.
func (d *Datastore) requestKeys(keys []string) (map[string]string, error) {
	requestData, err := keysToGetManyRequest(keys)
	if err != nil {
		return nil, fmt.Errorf("creating GetManyRequest: %w", err)
	}

	req, err := http.NewRequest("POST", d.url, bytes.NewReader(requestData))
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

// keysToGetManyRequest returns an list of datastore GetManyRequests encoded as
// json.
func keysToGetManyRequest(keys []string) ([]byte, error) {
	type getManyRequest struct {
		Collection   string   `json:"collection"`
		IDs          []int    `json:"ids"`
		MappedFields []string `json:"mapped_fields"`
	}

	requestsPart := make([]getManyRequest, len(keys))

	for i, key := range keys {
		keyParts := strings.SplitN(key, "/", 3)
		if len(keyParts) != 3 {
			return nil, fmt.Errorf("invalid key %s", key)
		}

		id, err := strconv.Atoi(keyParts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid key %s", key)
		}
		requestsPart[i] = getManyRequest{Collection: keyParts[0], IDs: []int{id}, MappedFields: []string{keyParts[2]}}
	}
	request := struct {
		Requests []getManyRequest `json:"requests"`
	}{requestsPart}
	return json.Marshal(request)
}

func getManyResponceToKeyValue(r io.Reader) (map[string]string, error) {
	var data map[string]map[string]json.RawMessage
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return nil, fmt.Errorf("decoding responce: %w", err)
	}

	keyValue := make(map[string]string)
	for fqid, inner := range data {
		for field, value := range inner {
			keyValue[fqid+"/"+field] = string(value)
		}
	}
	return keyValue, nil
}
