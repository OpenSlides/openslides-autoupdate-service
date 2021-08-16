package autoupdate

import (
	"context"
)

// Datastore gets values for keys and informs, if they change.
type Datastore interface {
	Get(ctx context.Context, keys ...string) (map[string][]byte, error)
	RegisterChangeListener(f func(map[string][]byte) error)
	ResetCache()
}

// KeysBuilder holds the keys that are requested by a user.
type KeysBuilder interface {
	Update(ctx context.Context) error
	Keys() []string
}
