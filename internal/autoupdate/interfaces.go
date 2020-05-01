package autoupdate

import (
	"context"
)

// Datastore gets values for keys and informs the service, if some keys change.
type Datastore interface {
	Get(ctx context.Context, keys ...string) ([]string, error)
	KeysChanged() ([]string, error)
}

// Restricter restricts keys.
type Restricter interface {
	// Restrict manipulates the values for the user with the given id.
	Restrict(uid int, data map[string]string)
}

// KeysBuilder holds the keys that are requested by a user.
type KeysBuilder interface {
	Update([]string) error
	Keys() []string
}
