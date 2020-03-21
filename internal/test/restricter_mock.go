package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"
)

// MockRestricter implements the restricter interface. The returned values can be controlled
// with the the Data attribute
type MockRestricter struct {
	mu   sync.RWMutex
	Data map[string]string
}

// Restrict returnes the values for the given keys be returning the values in the Data attribute.
//
// If a key is not present in Data, a default value is returned.
//
// If the key starts with "error", an error it thrown.
//
// If the key ends with "_id", "1" is returned.
//
// If the key ends with "_ids", "[1,2]" is returned.
//
// In any other case, "some value" is returned.
func (r *MockRestricter) Restrict(ctx context.Context, uid int, keys []string) (io.Reader, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	data := make(map[string]json.RawMessage, len(keys))
	for _, key := range keys {
		v, ok := r.Data[key]
		if ok {
			data[key] = json.RawMessage(v)
			continue
		}

		var value string
		switch {
		case strings.HasPrefix(key, "error"):
			return nil, fmt.Errorf("Restricter got an error")
		case strings.HasSuffix(key, "_id"):
			value = `1`
		case strings.HasSuffix(key, "_ids"):
			value = `[1,2]`
		default:
			value = fmt.Sprintf(`"The time is: %s"`, time.Now())
		}
		data[key] = json.RawMessage(value)
	}
	return encodeMap(data), nil
}

// Update updates the values from the restricter
func (r *MockRestricter) Update(data map[string]string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.Data == nil {
		r.Data = data
		return
	}

	for key, value := range data {
		r.Data[key] = value
	}
}

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
