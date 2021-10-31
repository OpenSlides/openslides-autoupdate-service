package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Group handels restrictions of the collection group.
type Group struct{}

// Modes returns the field restricters for the collection group.
func (g Group) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return g.see
	}
	return nil
}

func (g Group) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, groupID int) (bool, error) {
	meetingID, err := ds.Group_MeetingID(groupID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("fetching meeting id of group %d: %w", groupID, err)
	}

	canSee, err := Meeting{}.see(ctx, ds, mperms, meetingID)
	if err != nil {
		return false, fmt.Errorf("can see meeting %d: %w", meetingID, err)
	}
	return canSee, nil
}
