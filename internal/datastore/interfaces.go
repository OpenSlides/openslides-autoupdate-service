package datastore

import "encoding/json"

// Updater returns keys that have changes. Blocks until there is
// changed data.
type Updater interface {
	Update() (map[string]json.RawMessage, error)
}
