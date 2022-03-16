package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Projector handels the restriction for the projector collection.
//
// The user can see a projector, if the user has projector.can_see.
//
// Mode A: The user can see the projector.
type Projector struct{}

// Modes returns the restrictions modes for the meeting collection.
func (p Projector) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return p.see
	}
	return nil
}

func (p Projector) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, projectorID int) (bool, error) {
	meetingID, err := ds.Projector_MeetingID(projectorID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("fetching meeting_id %d: %w", projectorID, err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting perms for meeting %d: %w", meetingID, err)
	}

	return perms.Has(perm.ProjectorCanSee), nil
}
