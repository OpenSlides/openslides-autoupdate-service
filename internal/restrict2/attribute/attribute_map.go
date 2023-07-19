package attribute

import (
	"sync"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

// Map holds attributes for each restriction mod
type Map struct {
	mu   sync.RWMutex // TODO: This is a bad place for the mutex.
	data map[dskey.Key]*Attribute
}

// NewMap initializes an AttributeMap
func NewMap() *Map {
	return &Map{
		data: make(map[dskey.Key]*Attribute),
	}
}

// Add adds a value to the map.
func (am *Map) Add(modeKey dskey.Key, value *Attribute) {
	am.mu.Lock()
	defer am.mu.Unlock()

	am.data[modeKey] = value
}

// Keys returns all keys from the Map.
func (am *Map) Keys() []dskey.Key {
	am.mu.RLock()
	defer am.mu.RUnlock()

	keys := make([]dskey.Key, 0, len(am.data))
	for key := range am.data {
		keys = append(keys, key)
	}

	return keys
}

// Get returns the attributes for keys.
func (am *Map) Get(modeKeys ...dskey.Key) map[dskey.Key]*Attribute {
	am.mu.RLock()
	defer am.mu.RUnlock()

	out := make(map[dskey.Key]*Attribute, len(modeKeys))
	for _, key := range modeKeys {
		out[key] = am.data[key]
	}

	return out
}

// NeedCalc returns a list of keys, that are not in the map.
func (am *Map) NeedCalc(keys []dskey.Key) []dskey.Key {
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
