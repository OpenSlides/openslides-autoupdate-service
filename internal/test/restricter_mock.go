package test

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

// MockRestricter implements the restricter interface. The returned values can be controlled
// with the the Data attribute
type MockRestricter struct {
	mu   sync.Mutex
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
func (r *MockRestricter) Restrict(ctx context.Context, uid int, keys []string) (map[string]string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	out := make(map[string]string, len(keys))
	for _, key := range keys {
		v, ok := r.Data[key]
		if ok {
			out[key] = v
			continue
		}

		switch {
		case strings.HasPrefix(key, "error"):
			return nil, fmt.Errorf("Restricter got an error")
		case strings.HasSuffix(key, "_id"):
			out[key] = `1`
		case strings.HasSuffix(key, "_ids"):
			out[key] = `[1,2]`
		default:
			out[key] = fmt.Sprintf(`"The time is: %s"`, time.Now())
		}
	}
	return out, nil
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
