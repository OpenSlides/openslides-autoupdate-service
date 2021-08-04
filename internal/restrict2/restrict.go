package restrict

//go:generate  sh -c "go run gen_field_def/main.go > field_def.go"

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
