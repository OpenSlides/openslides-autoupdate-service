// Package datastore connects to the openslies-datastore-service to receive
// values.
//
// The Datastore object uses a cache to only request keys once. If a key in the
// cache gets an update via the keychanger, the cache gets updated.
package datastore

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

const urlPath = "/internal/datastore/reader/get_many"

// Datastore can be used to get values from the datastore-service.
//
// Has to be created with datastore.New().
type Datastore struct {
	url              string
	cache            *cache
	keychanger       Updater
	changeListeners  []func(map[string][]byte) error
	calculatedFields map[string]func(ctx context.Context, key string, changed map[string][]byte) ([]byte, error)
	calculatedKeys   map[string]string
	errHandler       func(error)

	resetMu sync.Mutex
}

// New returns a new Datastore object.
func New(ctx context.Context, url string, errHandler func(error), keychanger Updater) *Datastore {
	d := &Datastore{
		cache:            newCache(),
		url:              url + urlPath,
		keychanger:       keychanger,
		calculatedFields: make(map[string]func(ctx context.Context, key string, changed map[string][]byte) ([]byte, error)),
		calculatedKeys:   make(map[string]string),
		errHandler:       errHandler,
	}

	go d.receiveKeyChanges(ctx, errHandler)

	return d
}

var reValidKeys = regexp.MustCompile(`^([a-z]+|[a-z][a-z_]*[a-z])/[1-9][0-9]*/[a-z][a-z0-9_]*\$?[a-z0-9_]*$`)

// InvalidKeys checks if all of the given keys are valid. Invalid keys are
// returned.
//
// A return value of nil means, that all keys are valid.
func InvalidKeys(keys ...string) []string {
	var invalid []string
	for _, key := range keys {
		if ok := reValidKeys.MatchString(key); !ok {
			invalid = append(invalid, key)
		}
	}
	return invalid
}

// Get returns the value for one or many keys.
//
// If a key does not exist, the value nil is returned for that key.
func (d *Datastore) Get(ctx context.Context, keys ...string) (map[string][]byte, error) {
	values, err := d.cache.GetOrSet(ctx, keys, func(keys []string, set func(key string, value []byte)) error {
		if invalid := InvalidKeys(keys...); invalid != nil {
			return invalidKeyError{keys: invalid}
		}
		return d.loadKeys(keys, set)
	})
	if err != nil {
		return nil, fmt.Errorf("getOrSet for keys `%s`: %w", keys, err)
	}

	return values, nil
}

// RegisterChangeListener registers a function that is called whenever an
// datastore update happens.
func (d *Datastore) RegisterChangeListener(f func(map[string][]byte) error) {
	d.changeListeners = append(d.changeListeners, f)
}

// RegisterCalculatedField creates a virtual field that is not in the datastore
// but is created at runtime.
//
// `field` has to be in the form `collection/field`. The field is created for
// every full qualified field that matches that field.
//
// When a fqfield, that matches the field, is fetched for the first time, then f
// is called with `changed==nil`. On every ds-update, `f` is called again with the
// data, that has changed.
func (d *Datastore) RegisterCalculatedField(field string, f func(ctx context.Context, key string, changed map[string][]byte) ([]byte, error)) {
	d.calculatedFields[field] = f
}

// splitCalculatedKeys splits a list of keys in calculated keys and "normal"
// keys. The calculated keys are returned as map that point to the field name.
func (d *Datastore) splitCalculatedKeys(keys []string) (map[string]string, []string) {
	var normal []string
	calculated := make(map[string]string)
	for _, k := range keys {
		parts := strings.SplitN(k, "/", 3)
		if len(parts) != 3 {
			normal = append(normal, k)
			continue
		}

		field := parts[0] + "/" + parts[2]
		_, ok := d.calculatedFields[field]
		if !ok {
			normal = append(normal, k)
			continue
		}
		calculated[k] = field
	}
	return calculated, normal
}

// ResetCache clears the internal cache.
func (d *Datastore) ResetCache() {
	d.resetMu.Lock()
	d.cache = newCache()
	d.resetMu.Unlock()
}

// receiveKeyChanges listens for updates and saves then into the topic. This
// function blocks until the service is closed.
func (d *Datastore) receiveKeyChanges(ctx context.Context, errHandler func(error)) {
	if d.keychanger == nil {
		return
	}

	for {
		data, err := d.keychanger.Update(ctx)
		if err != nil {
			errHandler(fmt.Errorf("update data: %w", err))
			time.Sleep(time.Second)
			continue
		}

		// The lock prefents a cache reset while data is updating.
		d.resetMu.Lock()
		d.cache.SetIfExistMany(data)

		for key, field := range d.calculatedKeys {
			bs := d.calculateField(field, key, data)

			// Update the cache and also update the data-map. The data-map is
			// used later in this function to inform the changeListeners.
			d.cache.SetIfExist(key, bs)
			data[key] = bs
		}

		for _, f := range d.changeListeners {
			if err := f(data); err != nil {
				errHandler(err)
			}
		}
		d.resetMu.Unlock()
	}
}

func (d *Datastore) loadKeys(keys []string, set func(string, []byte)) error {
	calculatedKeys, normalKeys := d.splitCalculatedKeys(keys)
	if len(normalKeys) > 0 {
		data, err := d.requestKeys(normalKeys)
		if err != nil {
			return fmt.Errorf("requesting keys from datastore: %w", err)
		}
		for k, v := range data {
			set(k, v)
		}
	}

	for key, field := range calculatedKeys {
		calculated := d.calculateField(field, key, nil)
		d.calculatedKeys[key] = field
		set(key, calculated)
	}
	return nil
}

func (d *Datastore) calculateField(field string, key string, updated map[string][]byte) []byte {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	calculated, err := d.calculatedFields[field](ctx, key, updated)
	if err != nil {
		d.errHandler(fmt.Errorf("calculating key %s: %w", key, err))

		msg := fmt.Sprintf("calculating key %s", key)
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			msg = fmt.Sprintf("calculating key %s timed out", key)
		}

		calculated = []byte(fmt.Sprintf(`{"error": "%s"}`, msg))
	}
	return calculated

}

// requestKeys request a list of keys by the datastore.
//
// If an error happens, no key is returned.
//
// The returned map contains exacply the given keys. If a key does not exist in
// the datastore, then the value of this key is <nil>.
func (d *Datastore) requestKeys(keys []string) (map[string][]byte, error) {
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
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("datastore returned status %s", resp.Status)
		}
		return nil, fmt.Errorf("datastore returned status %s: %s", resp.Status, body)
	}

	responseData, err := getManyResponceToKeyValue(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parse responce: %w", err)
	}

	// Add keys that where not returned.
	for _, k := range keys {
		if _, ok := responseData[k]; ok {
			continue
		}
		responseData[k] = nil
	}

	return responseData, nil
}

// keysToGetManyRequest a json envoding of the get_many request.
func keysToGetManyRequest(keys []string) ([]byte, error) {
	request := struct {
		Requests []string `json:"requests"`
	}{keys}
	return json.Marshal(request)
}

// getManyResponceToKeyValue reads the responce from the getMany request and
// returns the content as key-values.
func getManyResponceToKeyValue(r io.Reader) (map[string][]byte, error) {
	var data map[string]map[string]map[string]json.RawMessage
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return nil, fmt.Errorf("decoding responce: %w", err)
	}

	keyValue := make(map[string][]byte)
	for collection, idField := range data {
		for id, fieldValue := range idField {
			for field, value := range fieldValue {
				keyValue[collection+"/"+id+"/"+field] = value
			}
		}
	}
	return keyValue, nil
}
