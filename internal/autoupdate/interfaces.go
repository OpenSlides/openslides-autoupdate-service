package autoupdate

import (
	"context"
	"encoding/json"
)

// Datastore gets values for keys and informs, if they change.
type Datastore interface {
	Get(ctx context.Context, keys ...string) ([]json.RawMessage, error)
	KeysChanged() ([]string, error)
}

// Restricter restricts keys.
type Restricter interface {
	// Restrict manipulates the values for the user with the given id.
	Restrict(uid int, data map[string]json.RawMessage)
}

// KeysBuilder holds the keys that are requested by a user.
type KeysBuilder interface {
	Update([]string) error
	Keys() []string
}
