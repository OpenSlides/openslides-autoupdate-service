package datastore

import "context"

// Updater returns keys that have changes. Blocks until there is
// changed data.
type Updater interface {
	Update(context.Context) (map[string][]byte, error)
}
