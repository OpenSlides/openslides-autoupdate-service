package attribute

import (
	"sync"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

// Map holds attributes for each restriction mod
type Map struct {
	mu   sync.RWMutex // TODO: This is a bad place for the lock.
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

func (am *Map) Get(keys ...dskey.Key) map[dskey.Key]*Attribute {
	am.mu.RLock()
	defer am.mu.Unlock()

	out := make(map[dskey.Key]*Attribute, len(keys))
	for _, k := range keys {
	}
}

// NeedCalc returns a list of keys, that are not in the map.
func (am *Map) NeedCalc(keys []dskey.Key) []dskey.Key {
	am.mu.RLock()
	defer am.mu.RUnlock()

	var needPrecalculate []dskey.Key
	for _, k := range keys {
		if _, ok := am.data[k]; !ok {
			needPrecalculate = append(needPrecalculate, k)
		}
	}

	return needPrecalculate
}
