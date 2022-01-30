// Package datastore connects to the openslies-datastore-service to receive
// values.
//
// The Datastore object uses a cache to only request keys once. If a key in the
// cache gets an update via the keychanger, the cache gets updated.
package datastore

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
	"sync"
	"time"
)

const urlPath = "/internal/datastore/reader/get_many"

const (
	messageBusReconnectPause = time.Second
	httpTimeout              = 3 * time.Second
)

// Getter can get values from keys.
//
// The Datastore object implements this interface.
type Getter interface {
	Get(ctx context.Context, keys ...string) (map[string][]byte, error)
}

// Source gives the data for keys.
type Source interface {
	// Get is called when a key is not in the cache.
	Get(ctx context.Context, key ...string) (map[string][]byte, error)

	// Update is called frequently and should block until there is new data.
	Update(ctx context.Context) (map[string][]byte, error)
}

// Datastore can be used to get values from the datastore-service.
//
// Has to be created with datastore.New().
type Datastore struct {
	cache *cache

	defaultSource Source
	keySource     map[string]Source

	changeListeners  []func(map[string][]byte) error
	calculatedFields map[string]func(ctx context.Context, key string, changed map[string][]byte) ([]byte, error)
	calculatedKeys   map[string]string

	resetMu sync.Mutex
}

// New returns a new Datastore object.
func New(defaultSource Source, keySource map[string]Source) *Datastore {
	if keySource == nil {
		keySource = make(map[string]Source)
	}

	d := &Datastore{
		cache: newCache(),

		defaultSource: defaultSource,
		keySource:     keySource,

		calculatedFields: make(map[string]func(ctx context.Context, key string, changed map[string][]byte) ([]byte, error)),
		calculatedKeys:   make(map[string]string),
	}

	return d
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
		return nil, fmt.Errorf("getOrSet`: %w", err)
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

// ResetCache clears the internal cache.
func (d *Datastore) ResetCache() {
	d.resetMu.Lock()
	d.cache = newCache()
	d.resetMu.Unlock()
}

// ListenOnUpdates listens for updates and informs all listeners.
func (d *Datastore) ListenOnUpdates(ctx context.Context, errHandler func(error)) {
	if errHandler == nil {
		errHandler = func(error) {}
	}

	updatedValues := make(chan map[string][]byte)
	d.keySource["default"] = d.defaultSource
	var wg sync.WaitGroup
	wg.Add(len(d.keySource))
	for _, source := range d.keySource {
		go func(source Source) {
			defer wg.Done()
			for {
				data, err := source.Update(ctx)
				if err != nil {
					if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
						return
					}

					errHandler(fmt.Errorf("update data: %w", err))
					time.Sleep(messageBusReconnectPause)
					continue
				}
				updatedValues <- data
			}
		}(source)
	}

	go func() {
		wg.Wait()
		close(updatedValues)
	}()

	for data := range updatedValues {
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

// splitCalculatedKeys splits a list of keys in calculated keys and "normal"
// keys. The calculated keys are returned as map that point to the field name.
func (d *Datastore) splitCalculatedKeys(keys []string) (map[string]string, map[Source][]string) {
	normal := make(map[Source][]string)
	calculated := make(map[string]string)
	for _, k := range keys {
		parts := strings.SplitN(k, "/", 3)
		if len(parts) != 3 {
			continue
		}

		field := parts[0] + "/" + parts[2]
		_, ok := d.calculatedFields[field]
		if !ok {
			source := d.defaultSource
			if s := d.keySource[field]; s != nil {
				source = s
			}
			normal[source] = append(normal[source], k)
			continue
		}
		calculated[k] = field
	}
	return calculated, normal
}

func (d *Datastore) loadKeys(keys []string, set func(string, []byte)) error {
	calculatedKeys, normalKeys := d.splitCalculatedKeys(keys)
	if len(normalKeys) > 0 {
		for source, keys := range normalKeys {
			data, err := source.Get(context.Background(), keys...)
			if err != nil {
				return fmt.Errorf("requesting keys from datastore: %w", err)
			}
			for k, v := range data {
				set(k, v)
			}
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
		log.Printf("Error calculating key %s: %v", key, err)

		msg := fmt.Sprintf("calculating key %s", key)
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			msg = fmt.Sprintf("calculating key %s timed out", key)
		}

		calculated = []byte(fmt.Sprintf(`{"error": "%s"}`, msg))
	}
	return calculated

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
