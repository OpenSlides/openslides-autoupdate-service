// Package datastore connects to the datastore service to receive values. It
// also connections to redis to get the keyupdates from the datastore
// connection.
//
// The Datastore object uses a cache to only request keys once. If a key in the
// cache gets an update via the keychanger, the cache gets updated.
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
func (d *Datastore) Get(ctx context.Context, keys ...string) ([]json.RawMessage, error) {
	values, err := d.cache.getOrSet(ctx, keys, func(keys []string) (map[string]json.RawMessage, error) {
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

	// Add keys to data, that where not returned by the datastore. This is not
	// neccessary if redis would send the data
	for _, key := range keys {
		if _, ok := data[key]; !ok {
			data[key] = []byte("null")
		}
	}

	d.cache.setIfExist(data)

	return keys, nil
}

// requestKeys request a list of keys by the datastore. If an error happens, no
// key is returned.
func (d *Datastore) requestKeys(keys []string) (map[string]json.RawMessage, error) {
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
	var data map[string]map[string]json.RawMessage
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return nil, fmt.Errorf("decoding responce: %w", err)
	}

	keyValue := make(map[string]json.RawMessage)
	for fqid, inner := range data {
		for field, value := range inner {
			keyValue[fqid+"/"+field] = value
		}
	}
	return keyValue, nil
}
