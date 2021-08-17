package datastore

import (
	"context"
	"fmt"
	"sync"
)

// cacheSetFunc is a function to update cache keys.
type cacheSetFunc func(keys []string, set func(key string, value []byte)) error

// cache stores the values to the datastore.
//
// Each value of the cache has three states. Either it exists, it does not
// exist, or it is pending. Pending means, that there is a current request to
// the datastore. An existing key can have the value `nil` which means, that the
// cache knows, that the key does not exist in the datastore. Each value
// []byte("null") is changed to nil.
//
// A new cache instance has to be created with newCache().
type cache struct {
	data *pendingMap
}

// newCache creates an initialized cache instance.
func newCache() *cache {
	return &cache{
		data: newPendingMap(),
	}
}

// GetOrSet returns the values for a list of keys. If one or more keys do not
// exist in the cache, then the missing values are fetched with the given set
// function. If this method is called more then once at the same time, only the
// first call fetches the result, the other calles get blocked until it the
// answer was fetched.
//
// A non existing value is returned as nil.
//
// All values get returned together. If only one key is missing, this function
// blocks, until all values are retrieved.
//
// The set function is used to create the cache values. It is called only with
// the missing keys.
//
// If a value is not returned by the set function, it is saved in the cache as
// nil to prevent a second call for the same key.
//
// If the context is done, GetOrSet returns. But the set() call is not stopped.
// Other calls to GetOrSet may wait for its result.
func (c *cache) GetOrSet(ctx context.Context, keys []string, set cacheSetFunc) (map[string][]byte, error) {
	missingKeys := c.data.markPending(keys...)

	// Fetch missing keys.
	if len(missingKeys) > 0 {
		// Fetch missing keys in the background. Do not stop the fetching. Even
		// when the context is done. Other calls could also request it.
		errChan := make(chan error)
		go func() {
			errChan <- c.fetchMissing(missingKeys, set)
		}()

		select {
		case err := <-errChan:
			if err != nil {
				return nil, fmt.Errorf("fetching key: %w", err)
			}
		case <-ctx.Done():
			return nil, fmt.Errorf("waiting for fetch missing: %w", ctx.Err())
		}
	}

	// Build return values. Blocks until pending keys are fetched.
	values := make(map[string][]byte, len(keys))
	for _, key := range keys {
		// Gets a value and waits until a pending value is ready.
		v, err := c.data.get(ctx, key)
		if err != nil {
			return nil, fmt.Errorf("waiting for key %s: %w", key, err)
		}

		values[key] = v
	}
	return values, nil
}

// fetchMissing loads the given keys with the set method. Does not update keys
// that are already in the cache.
//
// Sets all keys to not pending, even when an error happens.
func (c *cache) fetchMissing(keys []string, set cacheSetFunc) error {
	// Make sure all pending keys are closed. Make also sure, that
	// missing keys are set to nil.
	defer func() {
		c.data.setEmptyIfPending(keys...)
	}()

	err := set(keys, func(key string, value []byte) {
		c.data.setIfPending(key, value)
	})

	if err != nil {
		return fmt.Errorf("fetching missing keys: %w", err)
	}
	return nil
}

// SetIfExist updates the cache if the key exists or is pending.
func (c *cache) SetIfExist(key string, value []byte) {
	c.data.setIfExist(key, value)
}

// SetIfExistMany is like SetIfExist but with many keys.
func (c *cache) SetIfExistMany(data map[string][]byte) {
	c.data.setIfExistMany(data)
}

// pendingMap is like a map but values are returned as pendingValues.
type pendingMap struct {
	mu      sync.RWMutex
	data    map[string][]byte
	pending map[string]chan struct{}
}

// newPendingMap initializes a pendingDict.
func newPendingMap() *pendingMap {
	return &pendingMap{
		data:    make(map[string][]byte),
		pending: make(map[string]chan struct{}),
	}
}

// get returns a value from the pendingMap.
//
// If the value is pending, the returned value will block until the value is not
// pending anymore.
//
// Returns nil for a value that does not exist.
func (d *pendingMap) get(ctx context.Context, key string) ([]byte, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	pending, isPending := d.pending[key]

	if !isPending {
		return d.data[key], nil
	}

	d.mu.RUnlock()

	select {
	case <-pending:
	case <-ctx.Done():
		return nil, fmt.Errorf("waiting for value: %w", ctx.Err())
	}

	d.mu.RLock()

	return d.data[key], nil
}

// markPending marks one or more keys as pending.
//
// Skips keys that are already pending or are already in the database.
//
// Returns all keys that where marked as pending (did not exist).
func (d *pendingMap) markPending(keys ...string) []string {
	d.mu.Lock()
	defer d.mu.Unlock()

	var marked []string
	for _, key := range keys {
		if _, ok := d.data[key]; ok {
			continue
		}
		if _, ok := d.pending[key]; ok {
			continue
		}

		d.pending[key] = make(chan struct{})
		marked = append(marked, key)
	}
	return marked
}

// setIfExiists is like setIfExist but without setting a lock. Should not be
// used directly.
func (d *pendingMap) setIfExistUnlocked(key string, value []byte) {
	pending, isPending := d.pending[key]
	_, exists := d.data[key]

	if !isPending && !exists {
		return
	}

	d.data[key] = value

	if isPending {
		close(pending)
		delete(d.pending, key)
	}
}

// setIfExist updates a value but only if the key already exists or is
// pending.
//
// If the key is pending, informs all listeners.
func (d *pendingMap) setIfExist(key string, value []byte) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.setIfExistUnlocked(key, value)
}

// setIfExistMany is like setIfExists but for many values
func (d *pendingMap) setIfExistMany(data map[string][]byte) {
	d.mu.Lock()
	defer d.mu.Unlock()

	for k, v := range data {
		d.setIfExistUnlocked(k, v)
	}
}

// setIfPending updates values but only if the key is pending.
//
// Informs all listeners.
func (d *pendingMap) setIfPending(key string, value []byte) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if pending, isPending := d.pending[key]; isPending {
		d.data[key] = value
		close(pending)
		delete(d.pending, key)
	}
}

// setIfPendingMany like setIfPending but with many keys.
func (d *pendingMap) setEmptyIfPending(keys ...string) {
	d.mu.Lock()
	defer d.mu.Unlock()

	for _, key := range keys {
		if pending, isPending := d.pending[key]; isPending {
			d.data[key] = nil
			close(pending)
			delete(d.pending, key)
		}
	}
}
