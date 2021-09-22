package autoupdate

import (
	"hash/maphash"
)

func createHash(hasher *maphash.Hash, value []byte) uint64 {
	hasher.Reset()
	hasher.Write(value)
	return hasher.Sum64()
}

type filter struct {
	hasher  maphash.Hash
	history map[string]uint64
}

// filter has to be called on a reader that contains a decoded json object. It
// removes nil values from a map. Filter is called multiple times it removes
// values from the map, that did not chance.
func (f *filter) filter(data map[string][]byte) {
	if f.history == nil {
		f.history = make(map[string]uint64)
	}

	// Keys that are in oldHashes but not in data.
	//
	// They have to be added to data as nil, so the client gets informed, that
	// they are removed.
	for key := range f.history {
		if _, ok := data[key]; ok {
			continue
		}

		data[key] = nil
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

		newHash := createHash(&f.hasher, value)
		if old, inHistory := f.history[key]; inHistory && newHash == old {
			delete(data, key)
			continue
		}
		f.history[key] = newHash
	}

	return
}

// empty returns true, if the filter was not called before.
func (f *filter) empty() bool {
	return f.history == nil
}
