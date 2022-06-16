package datastore

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

// Recorder implements the datastore.Getter interface. It records all requested
// keys. They can be get with Recorder.Keys().
type Recorder struct {
	getter Getter
	keys   *set.Set[Key]
}

// NewRecorder initializes a Recorder.
func NewRecorder(g Getter) *Recorder {
	return &Recorder{
		getter: g,
		keys:   set.New[Key](),
	}
}

// Get fetches the keys from the datastore.
func (r *Recorder) Get(ctx context.Context, keys ...Key) (map[Key][]byte, error) {
	r.keys.Add(keys...)
	return r.getter.Get(ctx, keys...)
}

// Keys returns all datastore keys that where fetched in the process.
func (r *Recorder) Keys() *set.Set[Key] {
	return r.keys
}

// DB creates a json database that contains all values from the recorder.
func (r *Recorder) DB() ([]byte, error) {
	data, err := r.getter.Get(context.Background(), r.keys.List()...)
	if err != nil {
		return nil, fmt.Errorf("getting all values: %w", err)
	}

	converted := make(map[string]json.RawMessage, len(data))
	for k, v := range data {
		converted[k.String()] = v
	}

	return json.Marshal(converted)
}
