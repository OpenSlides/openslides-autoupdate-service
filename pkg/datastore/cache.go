package datastore

import (
	"context"
	"errors"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/pendingmap"
)

// cacheSetFunc is a function to update cache keys.
type cacheSetFunc func(keys []Key, set func(map[Key][]byte)) error

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
	data *pendingmap.PendingMap[Key, []byte]
}

// newCache creates an initialized cache instance.
func newCache() *cache {
	return &cache{
		data: pendingmap.New[Key, []byte](),
	}
}

// GetOrSet returns the values for a list of keys. If one or more keys do not
// exist in the cache, then the missing values are fetched with the given set
// function. If this method is called more then once at the same time, only the
// first call fetches the result, the other calles get blocked until it the
// answer was fetched.
//
// If a key is not returned by set, GetOrSet returns nil for it. But if a
// parallel call gets an error, GetOrSet also returns an error.
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
	// Blocks until all missing (but not pending) keys are fetched.
	//
	// After this call, all keys are either pending (from another parallel call)
	// or in the c.data.
	if err := c.fetchMissing(ctx, keys, set); err != nil {
		return nil, fmt.Errorf("fetching missing keys: %w", err)
	}

	got, err := c.data.Get(ctx, keys...)
	if err != nil {
		if errors.Is(err, pendingmap.ErrNotExist) {
			return nil, fmt.Errorf("fetching data in a parallel call failed")
		}
		return nil, err
	}

	return got, nil
}

// fetchMissing loads the given keys with the set method. Does not update keys
// that are already in the cache.
//
// Possible Errors: context.Canceled or context.DeadlineExeeded or the return
// value from the set func.
func (c *cache) fetchMissing(ctx context.Context, keys []Key, set cacheSetFunc) error {
	missingKeys := c.data.MarkPending(keys...)

	if len(missingKeys) == 0 {
		return nil
	}

	// Fetch missing keys in the background. Do not stop the fetching. Even
	// when the context is done. Other calls could also request it.
	errChan := make(chan error, 1)
	go func() {
		err := set(keys, func(data map[Key][]byte) {
			for key, value := range data {
				if string(value) == "null" {
					data[key] = nil
				}
			}

			c.data.SetIfPending(data)
		})

		if err != nil {
			c.data.UnMarkPending(keys...)
			errChan <- fmt.Errorf("fetching missing keys: %w", err)
			return
		}

		// Make sure all pending keys are closed. Make also sure, that
		// missing keys are set to nil.
		c.data.SetEmptyIfPending(keys...)

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
	if string(value) == "null" {
		value = nil
	}
	c.data.SetIfPendingOrExists(map[Key][]byte{key: value})
}

// SetIfExistMany is like SetIfExist but with many keys.
func (c *cache) SetIfExistMany(data map[Key][]byte) {
	for k, v := range data {
		if string(v) == "null" {
			data[k] = nil
		}
	}
	c.data.SetIfPendingOrExists(data)
}

func (c *cache) len() int {
	return c.data.Len()
}

func (c *cache) size() int {
	return -1 // TODO
	//return c.data.size()
}
