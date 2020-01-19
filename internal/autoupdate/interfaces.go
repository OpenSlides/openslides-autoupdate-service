package autoupdate

// KeysChangedReceiver returns keys that have changes.
// Blocks for some time until there are changed data.
// An implementation should not block forever but return
// empty data after some time to be called again.
type KeysChangedReceiver interface {
	KeysChanged() (KeyChanges, error)
}

// KeyChanges holds the information about changed keys
type KeyChanges struct {
	Created []string
	Updated []string
	Deleted []string
}

// PermChangedReceiver returns keys that have changes.
// Blocks until there are changed data.
type PermChangedReceiver interface {
	PermChanged() ([]string, error)
}

// PermChanges holds the information about changed permissions
type PermChanges struct {
	FullQualifiedIds  []string
	FullQualifiedKeys []string
	CollectionKeys    []string
	UserIds           []int
}
