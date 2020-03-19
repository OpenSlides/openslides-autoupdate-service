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

func (f *filter) filter(r io.Reader, err error) (io.Reader, error) {
	if err != nil {
		return nil, err
	}

	if f.history == nil {
		f.history = make(map[string]uint64)
	}
	var data map[string]json.RawMessage
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return nil, fmt.Errorf("can not unpack data: %w", err)
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

	out := decodeMap(data)
	return out, nil
}

func decodeMap(m map[string]json.RawMessage) io.Reader {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "{")
	for k, v := range m {
		fmt.Fprintf(buf, `"%s":%s,`, k, v)
	}
	buf.Truncate(buf.Len() - 1)
	fmt.Fprintf(buf, "}\n")
	return buf
}
