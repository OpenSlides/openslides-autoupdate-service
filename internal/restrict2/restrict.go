package restrict

import (
	"context"
	"errors"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Restrict changes the keys and values in data for the user with the given user
// id.
func Restrict(ctx context.Context, fetch *datastore.Fetcher, uid int, data map[string][]byte) error {
	return errors.New("todo")
}
