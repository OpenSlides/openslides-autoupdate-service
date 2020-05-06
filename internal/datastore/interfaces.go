package datastore

// KeysChangedReceiver returns keys that have changes. Blocks until there is
// changed data.
type KeysChangedReceiver interface {
	KeysChanged() ([]string, error)
}
