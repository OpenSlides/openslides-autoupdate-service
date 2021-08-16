package datastore

// Updater returns keys that have changes. Blocks until there is
// changed data.
type Updater interface {
	Update(<-chan struct{}) (map[string][]byte, error)
}
