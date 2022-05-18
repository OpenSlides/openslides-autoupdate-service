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
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/metric"
)

const (
	messageBusReconnectPause = time.Second
	httpTimeout              = 3 * time.Second
)

// Getter can get values from keys.
//
// The Datastore object implements this interface.
type Getter interface {
	Get(ctx context.Context, keys ...Key) (map[Key][]byte, error)
}

// GetPositioner is like a Getter but also taks a position
type GetPositioner interface {
	GetPosition(ctx context.Context, position int, keys ...Key) (map[Key][]byte, error)
}

// Source gives the data for keys.
type Source interface {
	// Get is called when a key is not in the cache.
	Get(ctx context.Context, key ...Key) (map[Key][]byte, error)

	// Update is called frequently and should block until there is new data.
	Update(ctx context.Context) (map[Key][]byte, error)
}

// SourcePosition is a Source that also supports getting the data at a specific position.
type SourcePosition interface {
	Source
	GetPosition(ctx context.Context, position int, key ...Key) (map[Key][]byte, error)
}

// HistoryInformationer returns the history information.
type HistoryInformationer interface {
	HistoryInformation(ctx context.Context, fqid string, w io.Writer) error
}

// Datastore can be used to get values from the datastore-service.
//
// Has to be created with datastore.New().
type Datastore struct {
	cache *cache

	defaultSource Source
	keySource     map[string]Source

	changeListeners  []func(map[Key][]byte) error
	calculatedFields map[string]func(ctx context.Context, key Key, changed map[Key][]byte) ([]byte, error)
	calculatedKeys   map[Key]string

	history HistoryInformationer

	resetMu sync.Mutex

	metricGetHitCount uint64
}

// New returns a new Datastore object.
func New(defaultSource Source, keySource map[string]Source, history HistoryInformationer) *Datastore {
	if keySource == nil {
		keySource = make(map[string]Source)
	}

	d := &Datastore{
		cache: newCache(),

		defaultSource: defaultSource,
		keySource:     keySource,

		calculatedFields: make(map[string]func(ctx context.Context, key Key, changed map[Key][]byte) ([]byte, error)),
		calculatedKeys:   make(map[Key]string),

		history: history,
	}

	metric.Register(d.metric)

	return d
}

// Get returns the value for one or many keys.
//
// If a key does not exist, the value nil is returned for that key.
func (d *Datastore) Get(ctx context.Context, keys ...Key) (map[Key][]byte, error) {
	atomic.AddUint64(&d.metricGetHitCount, 1)
	values, err := d.cache.GetOrSet(ctx, keys, func(keys []Key, set func(key Key, value []byte)) error {
		return d.loadKeys(keys, set)
	})
	if err != nil {
		return nil, fmt.Errorf("getOrSet`: %w", err)
	}

	return values, nil
}

// GetPosition is like Get() but returns the data at a specific position.
func (d *Datastore) GetPosition(ctx context.Context, position int, keys ...Key) (map[Key][]byte, error) {
	var data map[Key][]byte
	// Ignore calculated keys. They are not supported on GetPosition.
	_, normalKeys := d.splitCalculatedKeys(keys)
	for source, keys := range normalKeys {
		sourcePosition, ok := source.(SourcePosition)
		if !ok {
			// Ignore keys from sources that do not support the history.
			continue
		}

		values, err := sourcePosition.GetPosition(ctx, position, keys...)
		if err != nil {
			return nil, fmt.Errorf("get keys: %w", err)
		}

		if data == nil {
			data = values
			continue
		}

		for k, v := range values {
			data[k] = v
		}
	}

	return data, nil
}

// RegisterChangeListener registers a function that is called whenever an
// datastore update happens.
func (d *Datastore) RegisterChangeListener(f func(map[Key][]byte) error) {
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
func (d *Datastore) RegisterCalculatedField(
	field string,
	f func(ctx context.Context, key Key, changed map[Key][]byte) ([]byte, error),
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

// ListenOnUpdates listens for updates and informs all listeners.
func (d *Datastore) ListenOnUpdates(ctx context.Context, errHandler func(error)) {
	if errHandler == nil {
		errHandler = func(error) {}
	}

	updatedValues := make(chan map[Key][]byte)
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
func (d *Datastore) splitCalculatedKeys(keys []Key) (map[Key]string, map[Source][]Key) {
	normal := make(map[Source][]Key)
	calculated := make(map[Key]string)
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

func (d *Datastore) loadKeys(keys []Key, set func(Key, []byte)) error {
	calculatedKeys, normalKeys := d.splitCalculatedKeys(keys)
	for source, keys := range normalKeys {
		data, err := source.Get(context.Background(), keys...)
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

func (d *Datastore) calculateField(field string, key Key, updated map[Key][]byte) []byte {
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
func keysToGetManyRequest(keys []Key, position int) ([]byte, error) {
	request := struct {
		Requests []Key `json:"requests"`
		Position int   `json:"position,omitempty"`
	}{keys, position}
	return json.Marshal(request)
}

// parseGetManyResponse reads the response from the getMany request and
// returns the content as key-values.
func parseGetManyResponse(r io.Reader) (map[Key][]byte, error) {
	var data map[string]map[string]map[string]json.RawMessage
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	keyValue := make(map[Key][]byte)
	for collection, idField := range data {
		for idstr, fieldValue := range idField {
			id, err := strconv.Atoi(idstr)
			if err != nil {
				return nil, fmt.Errorf("invalid key. Id is no number: %s", idstr)
			}
			for field, value := range fieldValue {
				keyValue[Key{collection, id, field}] = value
			}
		}
	}
	return keyValue, nil
}
