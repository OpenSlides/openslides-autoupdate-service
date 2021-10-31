package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// ProjectorCountdown handels the restriction for the projector_countdown collection.
type ProjectorCountdown struct{}

// Modes returns the restrictions modes for the meeting collection.
func (p ProjectorCountdown) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return p.see
	}
	return nil
}

func (p ProjectorCountdown) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, projectorCountdownID int) (bool, error) {
	meetingID, err := ds.ProjectorCountdown_ID(projectorCountdownID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("fetching meeting_id %d: %w", projectorCountdownID, err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting perms for meeting %d: %w", meetingID, err)
	}

	return perms.Has(perm.ProjectorCanSee), nil
}
