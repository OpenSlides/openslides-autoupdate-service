package datastore

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/ristretto"
)

// cacheSetFunc is a function to update cache keys.
type cacheSetFunc func(keys []string, set func(key string, value json.RawMessage)) error

// cache stores the values to the datastore.
//
// An existing key can have the value `nil` which means, that the cache knows,
// that the key does not exist in the datastore. Each value []byte("null") is
// changed to nil.
//
// A new cache instance has to be created with newCache().
type cache struct {
	r *ristretto.Cache
}

// newCache creates an initialized cache instance.
func newCache() (*cache, error) {
	r, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		return nil, fmt.Errorf("creating ristretto cache: %w", err)
	}

	c := cache{r: r}

	return &c, nil
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
func (c *cache) GetOrSet(ctx context.Context, keys []string, set cacheSetFunc) ([]json.RawMessage, error) {
	values := make([]json.RawMessage, len(keys))
	var missing []string
	keyToID := make(map[string]int, len(keys))
	for i := 0; i < len(keys); i++ {
		value, found := c.r.Get(keys[i])
		if !found {
			keyToID[keys[i]] = i
			missing = append(missing, keys[i])
			continue
		}
		values[i] = value.(json.RawMessage)
	}

	if len(missing) == 0 {
		return values, nil
	}

	newValues := make(map[string]json.RawMessage)
	setfunc := func(key string, value json.RawMessage) {
		newValues[key] = value
		c.Set(key, value)
	}

	if err := set(missing, setfunc); err != nil {
		return nil, fmt.Errorf("fetching missing values: %w", err)
	}

	for k, v := range newValues {
		values[keyToID[k]] = v
	}

	return values, nil
}

func (c *cache) Set(key string, value json.RawMessage) {
	c.r.Set(key, value, 1)
}

func (c *cache) SetIfExist(data map[string]json.RawMessage) {
	for key, value := range data {
		_, found := c.r.Get(key)
		// TODO: Could this be a race condition if a value is inserded at this point?
		if found {
			c.Set(key, value)
		}
	}
}
