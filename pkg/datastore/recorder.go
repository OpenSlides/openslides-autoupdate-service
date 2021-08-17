package datastore

import "context"

// Recorder implements the datastore.Getter interface. It records all requested
// keys. They can be get with Recorder.Keys()
type Recorder struct {
	getter Getter
	keys   []string
}

// NewRecorder initializes a Recorder.
func NewRecorder(g Getter) *Recorder {
	return &Recorder{getter: g}
}

// Get fetches the keys from the datastore.
func (r *Recorder) Get(ctx context.Context, keys ...string) (map[string][]byte, error) {
	r.keys = append(r.keys, keys...)
	return r.getter.Get(ctx, keys...)
}

// Keys returns all datastore keys that where fetched in the process.
func (r *Recorder) Keys() []string {
	return r.keys
}
