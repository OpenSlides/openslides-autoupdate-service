package autoupdate

import (
	"encoding/json"
	"hash/maphash"
)

type filter struct {
	hash    maphash.Hash
	history map[string]uint64
}

// filter has to be called on a reader that contains a decoded json object.
// Filter is called multiple times it removes values from the json object, that
// did not chance. If the given error is not nil, it is returned immediately.
func (f *filter) filter(data map[string]json.RawMessage) error {
	if f.history == nil {
		f.history = make(map[string]uint64)
	}

	for key, value := range data {
		if len(value) == 0 {
			// Delete empty data
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
	return nil
}
