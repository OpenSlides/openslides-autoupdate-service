// Package datastore fetches the data from postgres or other sources.
//
// The datastore object uses a cache to only request keys once. If a key in the
// cache gets an update via the keychanger, the cache gets updated.
package datastore

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/metric"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/oserror"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
)

const (
	messageBusReconnectPause = time.Second
)

// Getter can get values from keys.
//
// The Datastore object implements this interface.
type Getter interface {
	Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error)
}

// Updater returns keys that have changes. Blocks until there is
// changed data.
type Updater interface {
	Update(context.Context) (map[dskey.Key][]byte, error)
}

// Source gives the data for keys.
type Source interface {
	Getter
	Updater
}

// HistoryInformationer returns the history information.
type HistoryInformationer interface {
	HistoryInformation(ctx context.Context, fqid string, w io.Writer) error
	GetPosition(ctx context.Context, position int, key ...dskey.Key) (map[dskey.Key][]byte, error)
}

// Datastore can be used to get values from the datastore-service.
//
// Has to be created with datastore.New().
type Datastore struct {
	cache *cache

	defaultSource Source
	keySource     map[string]Source

	updateListeners []func(map[dskey.Key][]byte) error
	insertListeners []func(map[dskey.Key][]byte) error

	calculatedFields map[string]func(ctx context.Context, key dskey.Key, changed map[dskey.Key][]byte) ([]byte, error)
	calculatedKeys   map[dskey.Key]string

	history HistoryInformationer

	resetMu sync.Mutex

	metricGetHitCount uint64
}

// New returns a new Datastore object.
func New(lookup environment.Environmenter, mb Updater, options ...Option) (*Datastore, func(context.Context, func(error)), error) {
	ds := Datastore{
		cache: newCache(),

		keySource: make(map[string]Source),

		calculatedFields: make(map[string]func(context.Context, dskey.Key, map[dskey.Key][]byte) ([]byte, error)),
		calculatedKeys:   make(map[dskey.Key]string),
	}

	var backgroundFuncs []func(context.Context, func(error))
	for _, o := range options {
		bgFunc, err := o(&ds, lookup)
		if err != nil {
			return nil, nil, err
		}

		if bgFunc != nil {
			backgroundFuncs = append(backgroundFuncs, bgFunc)
		}
	}

	if ds.defaultSource == nil {
		sourcePostgres, err := NewSourcePostgres(lookup, mb)
		if err != nil {
			return nil, nil, fmt.Errorf("initilizing postgres source: %w", err)
		}
		ds.defaultSource = sourcePostgres
	}

	metric.Register(ds.metric)

	background := func(ctx context.Context, errorHandler func(error)) {
		go ds.listenOnUpdates(ctx, errorHandler)

		for _, f := range backgroundFuncs {
			go f(ctx, errorHandler)
		}
	}

	return &ds, background, nil
}

// Get returns the value for one or many keys.
//
// If a key does not exist, the value nil is returned for that key.
func (d *Datastore) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	atomic.AddUint64(&d.metricGetHitCount, 1)
	values, err := d.cache.GetOrSet(ctx, keys, func(keys []dskey.Key, set func(map[dskey.Key][]byte) error) error {
		return d.loadKeys(keys, d.insertValue(set))
	})
	if err != nil {
		return nil, fmt.Errorf("getOrSet`: %w", err)
	}

	return values, nil
}

// insertValue is a wrapper around the 'set'-function from the cache.
//
// It is called every time a value is fetched for the first time.
//
// TODO: This ins only needed because of calculated fields. If calculated fields are removed, then
// the set function in d.loadKeys can be replaced by a return value.
func (d *Datastore) insertValue(f func(map[dskey.Key][]byte) error) func(map[dskey.Key][]byte) error {
	return func(values map[dskey.Key][]byte) error {
		for _, listener := range d.insertListeners {
			if err := listener(values); err != nil {
				return fmt.Errorf("insertlistener: %w", err)
			}
		}

		return f(values)
	}
}

// GetPosition is like Get() but returns the data at a specific position.
func (d *Datastore) GetPosition(ctx context.Context, position int, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	if d.history == nil {
		return nil, fmt.Errorf("histroy not supported")
	}
	return d.history.GetPosition(ctx, position, keys...)
}

// RegisterChangeListener registers a function that is called whenever an
// datastore update happens.
func (d *Datastore) RegisterChangeListener(f func(map[dskey.Key][]byte) error) {
	d.updateListeners = append(d.updateListeners, f)
}

// RegisterInsertListener registers a function that is called whenever an
// value is fetched for the first time.
func (d *Datastore) RegisterInsertListener(f func(map[dskey.Key][]byte) error) {
	d.insertListeners = append(d.insertListeners, f)
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
//
// depricated. Do not use this, instead add an option to register a field.
func (d *Datastore) RegisterCalculatedField(
	field string,
	f func(ctx context.Context, key dskey.Key, changed map[dskey.Key][]byte) ([]byte, error),
) {
	d.calculatedFields[field] = f
}

// ResetCache clears the internal cache.
func (d *Datastore) ResetCache() {
	d.resetMu.Lock()
	d.cache = newCache()
	d.resetMu.Unlock()
}

// HistoryInformation writes the history information for a fqid.
func (d *Datastore) HistoryInformation(ctx context.Context, fqid string, w io.Writer) error {
	return d.history.HistoryInformation(ctx, fqid, w)
}

// listenOnUpdates listens for updates and informs all listeners.
func (d *Datastore) listenOnUpdates(ctx context.Context, errHandler func(error)) {
	if errHandler == nil {
		errHandler = func(error) {}
	}

	updatedValues := make(chan map[dskey.Key][]byte)
	sources := make([]Source, 0, len(d.keySource)+1)
	sources = append(sources, d.defaultSource)
	for _, s := range d.keySource {
		sources = append(sources, s)
	}

	var wg sync.WaitGroup
	wg.Add(len(sources))
	for _, source := range sources {
		go func(source Source) {
			defer wg.Done()
			for {
				data, err := source.Update(ctx)
				if err != nil {
					if oserror.ContextDone(err) {
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

		for _, f := range d.updateListeners {
			if err := f(data); err != nil {
				errHandler(err)
			}
		}
		d.resetMu.Unlock()
	}
}

// splitCalculateddskey.Key splits a list of keys in calculated keys and "normal"
// keys. The calculated keys are returned as map that point to the field name.
func (d *Datastore) splitCalculatedKeys(keys []dskey.Key) (map[dskey.Key]string, map[Source][]dskey.Key) {
	normal := make(map[Source][]dskey.Key)
	calculated := make(map[dskey.Key]string)
	for _, k := range keys {
		field := k.Collection + "/" + k.Field
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

func (d *Datastore) loadKeys(keys []dskey.Key, set func(map[dskey.Key][]byte) error) error {
	calculatedKeys, normalKeys := d.splitCalculatedKeys(keys)
	for source, keys := range normalKeys {
		data, err := source.Get(context.Background(), keys...)
		if err != nil {
			return fmt.Errorf("requesting keys from datastore: %w", err)
		}
		set(data)
	}

	for key, field := range calculatedKeys {
		calculated := d.calculateField(field, key, nil)
		d.calculatedKeys[key] = field
		set(map[dskey.Key][]byte{key: calculated})
	}
	return nil
}

func (d *Datastore) calculateField(field string, key dskey.Key, updated map[dskey.Key][]byte) []byte {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	calculated, err := d.calculatedFields[field](ctx, key, updated)
	if err != nil {
		log.Printf("Error calculating key %s: %v", key, err)

		msg := fmt.Sprintf("calculating key %s", key)
		if oserror.ContextDone(err) {
			msg = fmt.Sprintf("calculating key %s timed out", key)
		}

		calculated = []byte(fmt.Sprintf(`{"error": "%s"}`, msg))
	}
	return calculated

}

// keysToGetManyRequest a json envoding of the get_many request.
func keysToGetManyRequest(keys []dskey.Key, position int) ([]byte, error) {
	request := struct {
		Requests []dskey.Key `json:"requests"`
		Position int         `json:"position,omitempty"`
	}{keys, position}
	return json.Marshal(request)
}

// parseGetManyResponse reads the response from the getMany request and
// returns the content as key-values.
func parseGetManyResponse(r io.Reader) (map[dskey.Key][]byte, error) {
	var data map[string]map[string]map[string]json.RawMessage
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	keyValue := make(map[dskey.Key][]byte)
	for collection, idField := range data {
		for idstr, fieldValue := range idField {
			id, err := strconv.Atoi(idstr)
			if err != nil {
				// TODO LAST ERROR
				return nil, fmt.Errorf("invalid key. Id is no number: %s", idstr)
			}
			for field, value := range fieldValue {
				keyValue[dskey.Key{Collection: collection, ID: id, Field: field}] = value
			}
		}
	}
	return keyValue, nil
}
