package cache

import (
	"context"
	"errors"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/cache/pendingmap"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/flow"
)

// Cache stores the values to the datastore.
//
// It is impelemented as a flow middleware.
//
// Each value of the Cache has three states. Either it exists, it does not
// exist, or it is pending. Pending means, that there is a current request to
// the datastore. An existing key can have the value `nil` which means, that the
// Cache knows, that the key does not exist in the datastore.
//
// A new Cache instance has to be created with newCache().
type Cache struct {
	data *pendingmap.PendingMap
	flow flow.Flow

	onlyCollectionField bool
	collectionField     dskey.Key
}

// New creates an initialized cache instance.
func New(flow flow.Flow) *Cache {
	return &Cache{
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
func (c *Cache) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
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

// fetchMissing loads all keys, that are currently not in the cache.
//
// Possible Errors: context.Canceled or context.DeadlineExeeded.
func (c *Cache) fetchMissing(ctx context.Context, keys []dskey.Key) error {
	missingKeys := c.data.MarkPending(keys...)

	if len(missingKeys) == 0 {
		return nil
	}

	// Fetch missing keys in the background. Do not stop the fetching. Even
	// when the context is done. Other calls could also request it.
	errChan := make(chan error, 1)
	go func() {
		data, err := c.flow.Get(context.Background(), missingKeys...)
		if err != nil {
			c.data.UnMarkPending(missingKeys...)
			errChan <- fmt.Errorf("getting data from flow: %w", err)
			return
		}

		if len(data) != len(missingKeys) {
			// A getter has to return the same amount of values, as keys where
			// requested. So this check should not be necessary. But there will
			// be very strange behaviour, if the getter has a but.
			c.data.UnMarkPending(missingKeys...)
			errChan <- fmt.Errorf("got %d keys from getter, but requested %d", len(data), len(missingKeys))
			return
		}

		c.data.SetIfPending(data)

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
func (c *Cache) Update(ctx context.Context, updateFn func(map[dskey.Key][]byte, error)) {
	if updateFn == nil {
		updateFn = func(m map[dskey.Key][]byte, err error) {}
	}

	c.flow.Update(ctx, func(data map[dskey.Key][]byte, err error) {
		if err != nil {
			updateFn(nil, err)
			return
		}

		c.data.SetIfPendingOrExists(data)
		updateFn(data, nil)
	})
}

// Len returns the amount of keys in the cache.
func (c *Cache) Len() int {
	return c.data.Len()
}

// Size returns the size of all values in the cache.
func (c *Cache) Size() int {
	return c.data.Size()
}

// Reset clears the cache.
func (c *Cache) Reset() {
	c.data.Reset()
}
