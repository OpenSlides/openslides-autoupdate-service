package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Projection handels the restriction for the projection collection.
type Projection struct{}

// Modes returns the restrictions modes for the meeting collection.
func (p Projection) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return p.see
	}
	return nil
}

func (p Projection) see(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, projectionID int) (bool, error) {
	meetingID := fetch.Field().Projection_MeetingID(ctx, projectionID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching meeting_id %d: %w", projectionID, err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting perms for meeting %d: %w", meetingID, err)
	}

	return perms.Has(perm.ProjectorCanSee), nil
}
