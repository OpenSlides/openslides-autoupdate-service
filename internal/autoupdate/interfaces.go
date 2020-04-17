package autoupdate

import (
	"context"
	"io"
)

// KeysChangedReceiver returns keys that have changes. Blocks until there is
// changed data.
type KeysChangedReceiver interface {
	KeysChanged() ([]string, error)
}

// Restricter restricts keys.
type Restricter interface {
	// Restrict returns an io.Reader which returns a json encoded object with
	// the requested keys with the values.
	Restrict(ctx context.Context, uid int, keys []string) (io.Reader, error)
}

// KeysBuilder holds the keys that are requested by a user.
type KeysBuilder interface {
	Update([]string) error
	Keys() []string
}
