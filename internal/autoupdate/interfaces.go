package autoupdate

import (
	"context"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Datastore gets values for keys and informs, if they change.
type Datastore interface {
	Get(ctx context.Context, keys ...string) (map[string][]byte, error)
	RegisterChangeListener(f func(map[string][]byte) error)
	ResetCache()
	RegisterCalculatedField(
		field string,
		f func(ctx context.Context, key string, changed map[string][]byte) ([]byte, error),
	)
}

// KeysBuilder holds the keys that are requested by a user.
type KeysBuilder interface {
	Update(ctx context.Context, ds datastore.Getter) error
	Keys() []string
}
