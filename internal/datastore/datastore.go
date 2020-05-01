package datastore

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const urlPath = "/internal/datastore/reader/getKeys"

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
		cache: newCache(),
		url:   url + urlPath,
	}
}

// Get returns the value for one or many keys.
func (d *Datastore) Get(ctx context.Context, keys ...string) ([]string, error) {
	values, err := d.cache.getOrSet(ctx, keys, func(keys []string) ([]string, error) {
		return d.requestKeys(keys)
	})
	if err != nil {
		return nil, fmt.Errorf("getOrSet for keys `%s`: %w", keys, err)
	}

	return values, nil
}

// KeysChanged blocks until some key have changed. Then, it returns the keys.
func (d *Datastore) KeysChanged() ([]string, error) {
	return d.keychanger.KeysChanged()

}

// requestKeys request a list of keys by the datastore. It returns the values in
// the same order. If an error happens, no key is returned.
func (d *Datastore) requestKeys(keys []string) ([]string, error) {
	requestData := struct {
		Keys []string
	}{
		Keys: keys,
	}

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(requestData); err != nil {
		return nil, fmt.Errorf("encoding request data for keys `%v`: %w", keys, err)
	}

	req, err := http.NewRequest("GET", d.url, buf)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

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

	var responseData map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return nil, fmt.Errorf("decoding responce: %w", err)
	}

	values := make([]string, 0, len(keys))
	for _, key := range keys {
		value, ok := responseData[key]
		if !ok {
			return nil, fmt.Errorf("key `%s` is not in responce", key)
		}
		values = append(values, value)
	}
	return values, nil
}
