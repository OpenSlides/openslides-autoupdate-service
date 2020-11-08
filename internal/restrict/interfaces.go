package restrict

import (
	"context"
	"encoding/json"
)

// Permissioner tells the restricter, if a user has the required permissions.
type Permissioner interface {
	RestrictFQIDs(ctx context.Context, uid int, fqids []string) (map[string]bool, error)
	RestrictFQFields(ctx context.Context, uid int, fqfields []string) (map[string]bool, error)
}

// Datastore informs the restricter about changed data.
type Datastore interface {
	Get(ctx context.Context, keys ...string) ([]json.RawMessage, error)
	RegisterChangeListener(f func(map[string]json.RawMessage) error)
}

// Checker checks, if a user has the permission for a key value pair. The value
// gets replaced with the returned value. Check has to return nil, if the user
// is not allowed to see the key.
type Checker interface {
	Check(ctx context.Context, uid int, key string, value json.RawMessage) (json.RawMessage, error)
}

// CheckerFunc is a function that implements the Checker interface.
type CheckerFunc func(ctx context.Context, uid int, key string, value json.RawMessage) (json.RawMessage, error)

// Check calls the function.
func (f CheckerFunc) Check(ctx context.Context, uid int, key string, value json.RawMessage) (json.RawMessage, error) {
	return f(ctx, uid, key, value)
}
