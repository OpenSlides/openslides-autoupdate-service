package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Projection handels the restriction for the projection collection.
//
// The user can see a projection, if the user has projector.can_see.
//
// Mode A: The user can see the projection.
type Projection struct{}

// Modes returns the restrictions modes for the meeting collection.
func (p Projection) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return p.see
	}
	return nil
}

func (p Projection) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, projectionID int) (bool, error) {
	meetingID, err := ds.Projection_MeetingID(projectionID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("fetching meeting_id %d: %w", projectionID, err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting perms for meeting %d: %w", meetingID, err)
	}

	return perms.Has(perm.ProjectorCanSee), nil
}
