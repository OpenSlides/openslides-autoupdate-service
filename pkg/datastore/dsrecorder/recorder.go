package dsrecorder

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

// Getter is the same as datastore.Getter
type Getter interface {
	Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error)
}

type getkver interface {
	GetKV(ctx context.Context, keys ...dskey.Key) ([][]byte, error)
}

// Recorder implements the datastore.Getter interface. It records all requested
// keys. They can be get with Recorder.Keys().
type Recorder struct {
	getter Getter
	keys   map[dskey.Key]struct{}
}

// New initializes a Recorder.
func New(g Getter) *Recorder {
	return &Recorder{
		getter: g,
		keys:   map[dskey.Key]struct{}{},
	}
}

// Get fetches the keys from the datastore.
func (r *Recorder) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	for _, k := range keys {
		r.keys[k] = struct{}{}
	}
	return r.getter.Get(ctx, keys...)
}

// GetKV fetches the keys from the datastore.
func (r *Recorder) GetKV(ctx context.Context, keys ...dskey.Key) ([][]byte, error) {
	for _, k := range keys {
		r.keys[k] = struct{}{}
	}

	kver, ok := r.getter.(getkver)
	if ok {
		return kver.GetKV(ctx, keys...)
	}

	data, err := r.getter.Get(ctx, keys...)

	converted := make([][]byte, len(keys))
	for i, key := range keys {
		converted[i] = data[key]
	}

	return converted, err
}

// Keys returns all datastore keys that where fetched in the process.
func (r *Recorder) Keys() map[dskey.Key]struct{} {
	return r.keys
}

// DB creates a json database that contains all values from the recorder.
func (r *Recorder) DB() ([]byte, error) {
	keys := make([]dskey.Key, 0, len(r.keys))
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
