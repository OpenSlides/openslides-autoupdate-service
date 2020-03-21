package autoupdate

import (
	"hash/maphash"
)

type filter struct {
	hash    maphash.Hash
	history map[string]uint64
}

func (f *filter) filter(data map[string]string, err error) (map[string]string, error) {
	if err != nil {
		return nil, err
	}

	if f.history == nil {
		f.history = make(map[string]uint64)
	}

	for key, value := range data {
		f.hash.Reset()
		f.hash.WriteString(value)
		new := f.hash.Sum64()
		if old, ok := f.history[key]; ok && new == old {
			delete(data, key)
			continue
		}
		f.history[key] = new
	}
	return data, nil
}
