package autoupdate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hash/maphash"
	"io"
)

type filter struct {
	hash    maphash.Hash
	history map[string]uint64
}

// filter has to be called on a reader that contains a decoded json object.
// Filter is called multiple times it removes values from the json object, that
// did not chance. If the given error is not nil, it is returned immediately.
func (f *filter) filter(r io.Reader, err error) (io.Reader, error) {
	if err != nil {
		return nil, err
	}

	if f.history == nil {
		f.history = make(map[string]uint64)
	}

	var data map[string]json.RawMessage
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return nil, fmt.Errorf("can not decode data: %w", err)
	}

	for key, value := range data {
		f.hash.Reset()
		f.hash.Write(value)
		new := f.hash.Sum64()
		if old, ok := f.history[key]; ok && new == old {
			delete(data, key)
			continue
		}
		f.history[key] = new
	}

	if len(data) == 0 {
		return new(bytes.Reader), nil
	}

	out := encodeMap(data)
	return out, nil
}

// encodeMap decodes m into a json object. It does this manualy to be faster
// then the json package.
func encodeMap(m map[string]json.RawMessage) io.Reader {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "{")
	for k, v := range m {
		fmt.Fprintf(buf, `"%s":%s,`, k, v)
	}
	buf.Truncate(buf.Len() - 1)
	fmt.Fprintf(buf, "}\n")
	return buf
}
