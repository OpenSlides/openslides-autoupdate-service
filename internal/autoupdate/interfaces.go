package autoupdate

import "context"

// KeysChangedReceiver returns keys that have changes.
// Blocks for some time until there are changed data.
// An implementation should not block forever but handle the
// server shutdown.
type KeysChangedReceiver interface {
	KeysChanged() ([]string, error)
}

// Restricter restricts keys. See autoupdate.Restricter
type Restricter interface {
	Restrict(ctx context.Context, uid int, keys []string) (map[string]string, error)
}
