package collection

import (
	"context"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Organization handels restrictions of the collection organization.
type Organization struct{}

// Modes returns the restrictions modes for the meeting collection.
func (o Organization) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return o.see
	case "B":
		return o.modeB
	}
	return nil
}

func (o Organization) see(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, organizationID int) (bool, error) {
	return true, nil
}

func (o Organization) modeB(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, organizationID int) (bool, error) {
	return mperms.UserID() != 0, nil
}
