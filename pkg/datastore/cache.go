package datastore

import (
	"bytes"
	"context"
	"fmt"
	"sync"
)

// cacheSetFunc is a function to update cache keys.
type cacheSetFunc func(keys []Key, set func(key Key, value []byte)) error

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
//
// Possible Errors: context.Canceled or context.DeadlineExeeded or the return
// value from hte set func.
func (c *cache) GetOrSet(ctx context.Context, keys []Key, set cacheSetFunc) (map[Key][]byte, error) {
	// Blocks until all missing keys are fetched.
	//
	// After this call, all keys are either pending (from another parallel call)
	// or in the c.data. This is a requirement to call c.data.get().
	if err := c.fetchMissing(ctx, keys, set); err != nil {
		return nil, fmt.Errorf("fetching missing keys: %w", err)
	}

	return c.data.get(ctx, keys)
}

// fetchMissing loads the given keys with the set method. Does not update keys
// that are already in the cache.
//
// Possible Errors: context.Canceled or context.DeadlineExeeded or the return
// value from the set func.
func (c *cache) fetchMissing(ctx context.Context, keys []Key, set cacheSetFunc) error {
	missingKeys := c.data.markPending(keys...)

	if len(missingKeys) == 0 {
		return nil
	}

	// Fetch missing keys in the background. Do not stop the fetching. Even
	// when the context is done. Other calls could also request it.
	errChan := make(chan error, 1)
	go func() {
		err := set(keys, func(key Key, value []byte) {
			c.data.setIfPending(key, value)
		})

		if err != nil {
			c.data.unMarkPending(keys...)
			errChan <- fmt.Errorf("fetching missing keys: %w", err)
			return
		}

		// Make sure all pending keys are closed. Make also sure, that
		// missing keys are set to nil.
		c.data.setEmptyIfPending(keys...)

		errChan <- nil
	}()

	select {
	case err := <-errChan:
		if err != nil {
			return fmt.Errorf("fetching key: %w", err)
		}
	case <-ctx.Done():
		return fmt.Errorf("waiting for fetch missing: %w", ctx.Err())
	}

	return nil
}

// SetIfExist updates the cache if the key exists or is pending.
func (c *cache) SetIfExist(key Key, value []byte) {
	c.data.setIfExist(key, value)
}

// SetIfExistMany is like SetIfExist but with many keys.
func (c *cache) SetIfExistMany(data map[Key][]byte) {
	c.data.setIfExistMany(data)
}

func (c *cache) len() int {
	return c.data.len()
}

func (c *cache) size() int {
	return c.data.size()
}

// pendingMap is like a map but values are returned as pendingValues.
type pendingMap struct {
	sync.RWMutex
	data    map[Key][]byte
	pending map[Key]chan struct{}
}

// newPendingMap initializes a pendingDict.
func newPendingMap() *pendingMap {
	return &pendingMap{
		data:    map[Key][]byte{},
		pending: map[Key]chan struct{}{},
	}
}

// get returns a list o keys from the pendingMap.
//
// The function blocks, until all values are not pending anymore.
//
// Returns nil for a value that does not exist.
//
// Makes sure that all values are returned at the same version. So if setIfExist
// is called while the function is running, then all values are returned at the
// latest version.
//
// Expects, that all keys are either pending or in the data. It is not allowed,
// that a key is not pending when this starts and gets pending whil it runs.
//
// Possible Errors: context.Canceled or context.DeadlineExeeded
func (pm *pendingMap) get(ctx context.Context, keys []Key) (map[Key][]byte, error) {
	if err := pm.waitForPending(ctx, keys); err != nil {
		return nil, err
	}

	out := make(map[Key][]byte, len(keys))
	reading(pm, func() {
		for _, k := range keys {
			out[k] = pm.data[k]
		}
	})

	return out, nil
}

// markPending marks one or more keys as pending.
//
// Skips keys that are already pending or are already in the database.
//
// Returns all keys that where marked as pending (did not exist).
func (pm *pendingMap) markPending(keys ...Key) []Key {
	var marked []Key
	reading(pm, func() {
		for _, key := range keys {
			if _, ok := pm.data[key]; ok {
				continue
			}
			if _, ok := pm.pending[key]; ok {
				continue
			}

			marked = append(marked, key)
		}
	})

	if len(marked) > 0 {
		pm.Lock()
		defer pm.Unlock()

		for _, key := range marked {
			pm.pending[key] = make(chan struct{})
		}
		return marked
	}

	return nil
}

// unMarkPending sets any key that is still pending not to be pending.
//
// Skips keys that are already pending or are already in the database.
func (pm *pendingMap) unMarkPending(keys ...Key) {
	pm.Lock()
	defer pm.Unlock()

	for _, key := range keys {
		if _, ok := pm.data[key]; ok {
			continue
		}
		pending := pm.pending[key]

		if pending == nil {
			continue
		}

		close(pending)
		delete(pm.pending, key)
	}
}

// waitForPending blocks until all the given keys are not pending anymore.
//
// Expects, that all keys are either pending or in the data. It is not allowed,
// that a key is not pending when this starts and gets pending whil it runs.
//
// Possible Errors: context.Canceled or context.DeadlineExeeded
func (pm *pendingMap) waitForPending(ctx context.Context, keys []Key) error {
	for _, k := range keys {
		var pending chan struct{}
		reading(pm, func() {
			pending = pm.pending[k]
		})

		if pending == nil {
			continue
		}

		select {
		case <-pending:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return nil
}

// setIfExiists is like setIfExist but without setting a lock. Should not be
// used directly.
func (pm *pendingMap) setIfExistUnlocked(key Key, value []byte) {
	pending := pm.pending[key]
	_, exists := pm.data[key]

	if pending == nil && !exists {
		return
	}

	if bytes.Equal(value, []byte("null")) {
		value = nil
	}

	pm.data[key] = value

	if pending != nil {
		close(pending)
		delete(pm.pending, key)
	}
}

// setIfExist updates a value but only if the key already exists or is
// pending.
//
// If the key is pending, informs all listeners.
func (pm *pendingMap) setIfExist(key Key, value []byte) {
	pm.Lock()
	defer pm.Unlock()

	pm.setIfExistUnlocked(key, value)
}

// setIfExistMany is like setIfExists but for many values.
func (pm *pendingMap) setIfExistMany(data map[Key][]byte) {
	pm.Lock()
	defer pm.Unlock()

	for k, v := range data {
		pm.setIfExistUnlocked(k, v)
	}
}

// setIfPending updates values but only if the key is pending.
//
// Informs all listeners.
func (pm *pendingMap) setIfPending(key Key, value []byte) {
	pm.Lock()
	defer pm.Unlock()

	if pending, isPending := pm.pending[key]; isPending {
		if bytes.Equal(value, []byte("null")) {
			value = nil
		}

		pm.data[key] = value
		close(pending)
		delete(pm.pending, key)
	}
}

// setIfPendingMany like setIfPending but with many keys.
func (pm *pendingMap) setEmptyIfPending(keys ...Key) {
	pm.Lock()
	defer pm.Unlock()

	for _, key := range keys {
		if pending, isPending := pm.pending[key]; isPending {
			pm.data[key] = nil
			close(pending)
			delete(pm.pending, key)
		}
	}
}

func (pm *pendingMap) len() int {
	pm.RLock()
	defer pm.RUnlock()

	return len(pm.data)
}

// size returns the size of all values in the cache in bytes.
func (pm *pendingMap) size() int {
	pm.RLock()
	defer pm.RUnlock()

	var size int
	for _, v := range pm.data {
		size += len(v)
	}
	return size
}

type rlocker interface {
	RLock()
	RUnlock()
}

func reading(l rlocker, cmd func()) {
	l.RLock()
	defer l.RUnlock()
	cmd()
}
