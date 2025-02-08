package autoupdate

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/zeebo/xxh3"
)

type filter struct {
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

		newHash := xxh3.Hash(value)
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

func (f *filter) hashState() (string, error) {
	kvList := make([]uint64, 0, len(f.history)*2)
	for key, hash := range f.history {
		kvList = append(kvList, uint64(key))
		kvList = append(kvList, hash)
	}

	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, kvList); err != nil {
		return "", fmt.Errorf("encode hashes: %w", err)
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func (f *filter) setHashState(hashes string) error {
	if hashes == "" {
		return nil
	}

	buf, err := base64.StdEncoding.DecodeString(hashes)
	if err != nil {
		return fmt.Errorf("decode hashes: %w", err)
	}

	kv := make([]uint64, len(buf)/8)
	if err := binary.Read(bytes.NewReader(buf), binary.LittleEndian, &kv); err != nil {
		return fmt.Errorf("encode hashes: %w", err)
	}

	f.history = make(map[dskey.Key]uint64, len(kv)/2)
	for i := 0; i < len(kv); i += 2 {
		f.history[dskey.Key(kv[i])] = kv[i+1]
	}

	return nil
}
