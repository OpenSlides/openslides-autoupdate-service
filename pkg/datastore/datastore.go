// Package datastore fetches the data from postgres or other sources.
//
// The datastore object uses a cache to only request keys once. If a key in the
// cache gets an update via the keychanger, the cache gets updated.
package datastore

import (
	"context"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/metric"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/oserror"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/flow"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
)

const (
	messageBusReconnectPause = time.Second
)

// Datastore can be used to get values from the datastore-service.
//
// Has to be created with datastore.New().
type Datastore struct {
	cache *cache

	flow            flow.Flow
	additionalFlows map[string]flow.Flow

	changeListeners  []func(map[dskey.Key][]byte) error
	calculatedFields map[string]func(ctx context.Context, key dskey.Key, changed map[dskey.Key][]byte) ([]byte, bool, error)
	calculatedKeys   map[dskey.Key]string
	calculatedKeysMu sync.RWMutex

	resetMu sync.Mutex

	metricGetHitCount uint64
}

// New returns a new Datastore object.
func New(lookup environment.Environmenter, messageBus flow.Updater, options ...Option) (*Datastore, func(context.Context, func(error)), error) {
	ds := Datastore{
		cache: newCache(),

		additionalFlows: make(map[string]flow.Flow),

		calculatedFields: make(map[string]func(context.Context, dskey.Key, map[dskey.Key][]byte) ([]byte, bool, error)),
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

	if ds.flow == nil {
		flowPostgres, err := NewFlowPostgres(lookup, messageBus)
		if err != nil {
			return nil, nil, fmt.Errorf("initilizing postgres flow: %w", err)
		}

		ds.flow = flowPostgres
	}

	if len(ds.additionalFlows) > 0 {
		ds.flow = flow.Combine(ds.flow, ds.additionalFlows)
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
	values, err := d.cache.GetOrSet(ctx, keys, func(keys []dskey.Key, set func(map[dskey.Key][]byte)) error {
		return d.loadKeys(keys, set)
	})
	if err != nil {
		return nil, fmt.Errorf("getOrSet`: %w", err)
	}

	return values, nil
}

// RegisterChangeListener registers a function that is called whenever an
// datastore update happens.
func (d *Datastore) RegisterChangeListener(f func(map[dskey.Key][]byte) error) {
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
//
// depricated. Do not use this, instead add an option to register a field.
func (d *Datastore) RegisterCalculatedField(
	field string,
	f func(ctx context.Context, key dskey.Key, changed map[dskey.Key][]byte) ([]byte, bool, error),
) {
	d.calculatedFields[field] = f
}

// ResetCache clears the internal cache.
func (d *Datastore) ResetCache() {
	d.resetMu.Lock()
	d.cache = newCache()
	d.resetMu.Unlock()
}

// listenOnUpdates listens for updates and informs all listeners.
func (d *Datastore) listenOnUpdates(ctx context.Context, errHandler func(error)) {
	if errHandler == nil {
		errHandler = func(error) {}
	}

	d.flow.Update(ctx, func(data map[dskey.Key][]byte, err error) {
		if err != nil {
			errHandler(err)
			return
		}

		d.resetMu.Lock()
		defer d.resetMu.Unlock()

		d.cache.SetIfExistMany(data)

		d.calculatedKeysMu.RLock()
		for key, field := range d.calculatedKeys {
			bs, changed := d.calculateField(field, key, data)
			if !changed {
				continue
			}

			// Update the cache and also update the data-map. The data-map is
			// used later in this function to inform the changeListeners.
			d.cache.SetIfExist(key, bs)
			data[key] = bs
		}
		d.calculatedKeysMu.RUnlock()

		for _, f := range d.changeListeners {
			if err := f(data); err != nil {
				errHandler(err)
			}
		}
	})
}

// splitCalculateddskey.Key splits a list of keys in calculated keys and "normal"
// keys. The calculated keys are returned as map that point to the field name.
func (d *Datastore) splitCalculatedKeys(keys []dskey.Key) (map[dskey.Key]string, []dskey.Key) {
	var normal []dskey.Key
	calculated := make(map[dskey.Key]string)
	for _, k := range keys {
		field := k.Collection + "/" + k.Field
		_, ok := d.calculatedFields[field]
		if !ok {
			normal = append(normal, k)
			continue
		}
		calculated[k] = field
	}
	return calculated, normal
}

func (d *Datastore) loadKeys(keys []dskey.Key, set func(map[dskey.Key][]byte)) error {
	calculatedKeys, normalKeys := d.splitCalculatedKeys(keys)

	if len(normalKeys) > 0 {
		data, err := d.flow.Get(context.Background(), normalKeys...)
		if err != nil {
			return fmt.Errorf("requesting keys: %w", err)
		}
		set(data)
	}

	for key, field := range calculatedKeys {
		// Since calculatedField is called with no updated data, it not
		// necessary to check the second return value.
		calculated, _ := d.calculateField(field, key, nil)
		d.calculatedKeysMu.Lock()
		d.calculatedKeys[key] = field
		d.calculatedKeysMu.Unlock()
		set(map[dskey.Key][]byte{key: calculated})
	}
	return nil
}

// calculateField calculates a calculated field.
//
// Returns the new value and true, if the value was changed.
//
// Returns nil, false, if the value was not changed.
func (d *Datastore) calculateField(field string, key dskey.Key, updated map[dskey.Key][]byte) ([]byte, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	calculated, changed, err := d.calculatedFields[field](ctx, key, updated)
	if err != nil {
		log.Printf("Error calculating key %s: %v", key, err)

		msg := fmt.Sprintf("calculating key %s", key)
		if oserror.ContextDone(err) {
			msg = fmt.Sprintf("calculating key %s timed out", key)
		}

		return []byte(fmt.Sprintf(`{"error": "%s"}`, msg)), true
	}

	return calculated, changed
}
