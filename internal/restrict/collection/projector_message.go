package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// ProjectorMessage handels the restriction for the projector_message collection.
//
// The user can see a projector message, if the user has projector.can_see.
//
// Mode A: The user can see the projector message.
type ProjectorMessage struct{}

// MeetingID returns the meetingID for the object.
func (p ProjectorMessage) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.ProjectorMessage_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meetingID: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (p ProjectorMessage) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return p.see
	}
	return nil
}

func (p ProjectorMessage) see(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, projectorMessageIDs ...int) ([]int, error) {
	return meetingPerm(ctx, ds, p, projectorMessageIDs, mperms, perm.ProjectorCanSee)
}
