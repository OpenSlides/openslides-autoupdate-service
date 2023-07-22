package restrict

import (
	"sync"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

// attrMap holds attributes for each restriction mod
type attrMap struct {
	mu   sync.RWMutex // TODO: This is a bad place for the mutex.
	data map[dskey.Key]attribute.Func
}

func newAttrMap() *attrMap {
	return &attrMap{
		data: make(map[dskey.Key]attribute.Func),
	}
}

// Add adds a value to the map.
func (am *attrMap) Add(tuples ...collection.Tuple) {
	am.mu.Lock()
	defer am.mu.Unlock()

	for _, tuple := range tuples {
		am.data[tuple.Key] = tuple.Value
	}
}

// Keys returns all keys from the Map.
func (am *attrMap) Keys() []dskey.Key {
	am.mu.RLock()
	defer am.mu.RUnlock()

	keys := make([]dskey.Key, 0, len(am.data))
	for key := range am.data {
		keys = append(keys, key)
	}

	return keys
}

// Get returns the attributes for keys.
func (am *attrMap) Get(modeKeys ...dskey.Key) map[dskey.Key]attribute.Func {
	am.mu.RLock()
	defer am.mu.RUnlock()

	out := make(map[dskey.Key]attribute.Func, len(modeKeys))
	for _, key := range modeKeys {
		out[key] = am.data[key]
	}

	return out
}

// NeedCalc returns a list of keys, that are not in the map.
func (am *attrMap) NeedCalc(keys []dskey.Key) []dskey.Key {
	am.mu.RLock()
	defer am.mu.RUnlock()

	var needPrecalculate []dskey.Key
	for _, key := range keys {
		if _, ok := am.data[key]; !ok {
			needPrecalculate = append(needPrecalculate, key)
		}
	}

	return needPrecalculate
}
