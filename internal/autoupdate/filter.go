package autoupdate

import (
	"encoding/json"
	"hash/maphash"
)

type filter struct {
	hash    maphash.Hash
	history map[string]uint64
}

// filter has to be called on a reader that contains a decoded json object. It
// removes nil values from a map. Filter is called multiple times it removes
// values from the map, that did not chance.
func (f *filter) filter(data map[string]json.RawMessage) {
	if f.history == nil {
		f.history = make(map[string]uint64)
	}

	for k, v := range data {
		// Filter empty values that where empty before.
		if len(v) == 0 && f.history[k] == 0 {
			delete(data, k)
		}
	}

	for key, value := range data {
		if len(value) == 0 {
			// Delete empty data
			if f.history[key] == 0 {
				// Data was empty before
				delete(data, key)
			}
			f.history[key] = 0
			continue
		}

		f.hash.Reset()
		f.hash.Write(value)
		new := f.hash.Sum64()
		if old, ok := f.history[key]; ok && new == old {
			delete(data, key)
			continue
		}
		f.history[key] = new
	}
	return
}

// empty returns true, if the filter was not called before.
func (f *filter) empty() bool {
	return f.history == nil
}

// reset sets the filter to its original state.
func (f *filter) reset() {
	f.history = nil
}
