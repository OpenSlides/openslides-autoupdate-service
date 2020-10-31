package autoupdate

import (
	"context"
	"encoding/json"
)

// Datastore gets values for keys and informs, if they change.
type Datastore interface {
	Get(ctx context.Context, keys ...string) ([]json.RawMessage, error)
	RegisterChangeListener(f func(map[string]json.RawMessage) error)
}

// Restricter restricts keys.
type Restricter interface {
	// Restrict manipulates the values for the user with the given id.
	Restrict(uid int, data map[string]json.RawMessage) error
}

// KeysBuilder holds the keys that are requested by a user.
type KeysBuilder interface {
	Update(ctx context.Context) error
	Keys() []string
}

// UserUpdater has a function to get user_ids, that should get a full update.
type UserUpdater interface {
	AdditionalUpdate(updated map[string]string) ([]int, error)
}
