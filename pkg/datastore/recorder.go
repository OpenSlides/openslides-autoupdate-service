package datastore

import "context"

// Recorder implements the datastore.Getter interface. It records all requested
// keys. They can be get with Recorder.Keys()
type Recorder struct {
	getter Getter
	keys   map[string]bool
}

// NewRecorder initializes a Recorder.
func NewRecorder(g Getter) *Recorder {
	return &Recorder{
		getter: g,
		keys:   map[string]bool{},
	}
}

// Get fetches the keys from the datastore.
func (r *Recorder) Get(ctx context.Context, keys ...string) (map[string][]byte, error) {
	for _, k := range keys {
		r.keys[k] = true
	}
	return r.getter.Get(ctx, keys...)
}

// Keys returns all datastore keys that where fetched in the process.
func (r *Recorder) Keys() map[string]bool {
	return r.keys
}
