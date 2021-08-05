package collection

import (
	"context"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// FieldRestricter is a function to restrict fields of a collection.
type FieldRestricter func(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, id int) (bool, error)

func allways(ctx context.Context, fetch *datastore.Fetcher, mperms perm.MeetingPermission, agndaID int) (bool, error) {
	return true, nil
}

func never(ctx context.Context, fetch *datastore.Fetcher, mperms perm.MeetingPermission, agndaID int) (bool, error) {
	return false, nil
}
