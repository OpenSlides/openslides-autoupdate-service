package autoupdate

import (
	"hash/maphash"

	"github.com/OpenSlides/openslides-go/datastore/dskey"
)

func createHash(hasher *maphash.Hash, value []byte) uint64 {
	hasher.Reset()

	// This can not return any error.
	_, _ = hasher.Write(value)
	return hasher.Sum64()
}

type filter struct {
	hasher  maphash.Hash
	history map[dskey.Key]uint64
}

// filter removes nil values from a map. If filter is called multiple times it
// removes values from the map, that did not chance.
func (f *filter) filter(data map[dskey.Key][]byte) {
	if f.history == nil {
		f.history = make(map[dskey.Key]uint64)
	}

	for k := range f.history {
		if _, ok := data[k]; !ok {
			delete(f.history, k)
		}
	}

	for key, value := range data {
		if len(value) == 0 {
			// Value does not exist or user has no permission to see it.
			if f.history[key] == 0 {
				// Data was empty before. Do not sent it to the user.
				delete(data, key)
			}
			f.history[key] = 0
			continue
		}

		newHash := createHash(&f.hasher, value)
		if oldHash, inHistory := f.history[key]; inHistory && newHash == oldHash {
			delete(data, key)
			continue
		}
		f.history[key] = newHash
	}
}

// empty returns true, if the filter was not called before.
func (f *filter) empty() bool {
	return f.history == nil
}
