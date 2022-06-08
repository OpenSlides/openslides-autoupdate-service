package datastore

import (
	"context"
	"encoding/json"
	"fmt"
)

// Recorder implements the datastore.Getter interface. It records all requested
// keys. They can be get with Recorder.Keys().
type Recorder struct {
	getter Getter
	keys   map[Key]struct{}
}

// NewRecorder initializes a Recorder.
func NewRecorder(g Getter) *Recorder {
	return &Recorder{
		getter: g,
		keys:   map[Key]struct{}{},
	}
}

// Get fetches the keys from the datastore.
func (r *Recorder) Get(ctx context.Context, keys ...Key) (map[Key][]byte, error) {
	for _, k := range keys {
		r.keys[k] = struct{}{}
	}
	return r.getter.Get(ctx, keys...)
}

// Keys returns all datastore keys that where fetched in the process.
func (r *Recorder) Keys() map[Key]struct{} {
	return r.keys
}

// DB creates a json database that contains all values from the recorder.
func (r *Recorder) DB() ([]byte, error) {
	keys := make([]Key, 0, len(r.keys))
	for k := range r.keys {
		keys = append(keys, k)
	}

	data, err := r.getter.Get(context.Background(), keys...)
	if err != nil {
		return nil, fmt.Errorf("getting all values: %w", err)
	}

	converted := make(map[string]json.RawMessage, len(data))
	for k, v := range data {
		converted[k.String()] = v
	}

	return json.Marshal(converted)
}
