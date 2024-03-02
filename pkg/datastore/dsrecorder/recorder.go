package dsrecorder

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/flow"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

// Recorder implements the datastore.Getter interface. It records all requested
// keys. They can be get with Recorder.Keys().
type Recorder struct {
	getter flow.Getter
	keys   set.Set[dskey.Key]
}

// New initializes a Recorder.
func New(g flow.Getter) *Recorder {
	return &Recorder{
		getter: g,
		keys:   set.New[dskey.Key](),
	}
}

// Get fetches the keys from the datastore.
func (r *Recorder) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	r.keys.Add(keys...)
	return r.getter.Get(ctx, keys...)
}

// Keys returns all datastore keys that where fetched in the process.
func (r *Recorder) Keys() set.Set[dskey.Key] {
	return r.keys
}

// KeysAsMap returns all keys as map structure.
//
// Only used for the projector. Can be removed when the projector is refactored.
func (r *Recorder) KeysAsMap() map[dskey.Key]struct{} {
	out := make(map[dskey.Key]struct{}, r.keys.Len())
	for _, k := range r.keys.List() {
		out[k] = struct{}{}
	}
	return out
}

// DB creates a json database that contains all values from the recorder.
func (r *Recorder) DB() ([]byte, error) {
	keys := r.keys.List()

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
