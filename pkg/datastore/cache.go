package datastore

import (
	"context"
	"errors"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/flow"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/pendingmap"
)

// cacheSetFunc is a function to update cache keys.
type cacheSetFunc func(keys []dskey.Key, set func(map[dskey.Key][]byte)) error

// cache stores the values to the datastore.
//
// It is impelemented as a flow middleware.
//
// Each value of the cache has three states. Either it exists, it does not
// exist, or it is pending. Pending means, that there is a current request to
// the datastore. An existing key can have the value `nil` which means, that the
// cache knows, that the key does not exist in the datastore. Each value
// []byte("null") is changed to nil.
//
// A new cache instance has to be created with newCache().
type cache struct {
	data *pendingmap.PendingMap
	flow flow.Flow
}

// newCache creates an initialized cache instance.
func newCache(flow flow.Flow) *cache {
	return &cache{
		data: pendingmap.New(),
		flow: flow,
	}
}

// Get returns the values for a list of keys. If one or more keys do not exist
// in the cache, then the missing values are fetched. If this method is called
// more then once at the same time, only the first call fetches the result, the
// other calles get blocked until it the answer was fetched.
//
// If a parallel call gets an error, GetOrSet also returns an error.
//
// All values get returned together. If only one key is missing, this function
// blocks, until all values are retrieved.
//
// If a value can not be fetched from the flow, it is saved in the cache as nil
// to prevent a second call for the same key.
//
// If the context is done, GetOrSet returns. But the call to the flow is not stopped.
// Other calls to GetOrSet may wait for its result.
//
// Possible Errors: context.Canceled or context.DeadlineExeeded or the return
// value from hte set func.
func (c *cache) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	// Blocks until all missing (but not pending) keys are fetched.
	//
	// After this call, all keys are either pending (from another parallel call)
	// or in the c.data.
	if err := c.fetchMissing(ctx, keys); err != nil {
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
func (c *cache) fetchMissing(ctx context.Context, keys []dskey.Key) error {
	missingKeys := c.data.MarkPending(keys...)

	if len(missingKeys) == 0 {
		return nil
	}

	// Fetch missing keys in the background. Do not stop the fetching. Even
	// when the context is done. Other calls could also request it.
	errChan := make(chan error, 1)
	go func() {
		data, err := c.flow.Get(ctx, missingKeys...)
		if err != nil {
			c.data.UnMarkPending(missingKeys...)
			errChan <- fmt.Errorf("getting data from flow: %w", err)
			return
		}

		for key, value := range data {
			if string(value) == "null" {
				data[key] = nil
			}
		}
		// TODO: with this implementation, SetIfPending and SetEmptyIfPending could be one function to remove one lock.
		c.data.SetIfPending(data)

		// Make sure all pending keys are closed. Make also sure, that
		// missing keys are set to nil.
		c.data.SetEmptyIfPending(missingKeys...)

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

// Update gets values from the flow to update the cached values.
func (c *cache) Update(ctx context.Context, updateFn func(map[dskey.Key][]byte, error)) {
	if updateFn == nil {
		updateFn = func(m map[dskey.Key][]byte, err error) {}
	}

	c.flow.Update(ctx, func(data map[dskey.Key][]byte, err error) {
		if err != nil {
			updateFn(nil, err)
			return
		}

		for key, value := range data {
			if string(value) == "null" {
				data[key] = nil
			}
		}

		c.data.SetIfPendingOrExists(data)
		updateFn(data, nil)
	})
}

func (c *cache) len() int {
	return c.data.Len()
}

func (c *cache) size() int {
	return c.data.Size()
}

func (c *cache) Reset() {
	c.data.Reset()
}
