package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"sync"
)

// MockRestricter implements the restricter interface. The returned values can
// be controlled with the the Data attribute.
type MockRestricter struct {
	mu       sync.RWMutex
	Data     map[string]string
	OnlyData bool
}

// Restrict returnes the values for the given keys be returning the values in
// the Data attribute.
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

		if r.OnlyData {
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
			value = `"Hello World"`
		}
		data[key] = json.RawMessage(value)
	}

	modelData, err := modelFormat(data)
	if err != nil {
		return nil, fmt.Errorf("convert keys to modelFormat: %w", err)
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(modelData); err != nil {
		return nil, fmt.Errorf("decode modelData: %w", err)
	}
	return buf, nil
}

func modelFormat(data map[string]json.RawMessage) (map[string]map[string]map[string]json.RawMessage, error) {
	modelData := make(map[string]map[string]map[string]json.RawMessage)
	for key, value := range data {
		keyParts := strings.SplitN(key, "/", 3)
		if len(keyParts) != 3 {
			return nil, fmt.Errorf("invalid key `%s`", key)
		}

		collection := keyParts[0]
		id := keyParts[1]
		field := keyParts[2]
		if modelData[collection] == nil {
			modelData[collection] = make(map[string]map[string]json.RawMessage)
		}
		if modelData[collection][id] == nil {
			modelData[collection][id] = make(map[string]json.RawMessage)
		}
		modelData[collection][id][field] = value
	}
	return modelData, nil
}

// Update updates the values from the restricter.
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
