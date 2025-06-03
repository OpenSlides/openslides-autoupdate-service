package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
	"github.com/OpenSlides/openslides-go/perm"
)

// ProjectorCountdown handels the restriction for the projector_countdown collection.
//
// The user can see a projector countdown, if the user has projector.can_see.
//
// Group A: The user can see the projector countdown.
type ProjectorCountdown struct{}

// Name returns the collection name.
func (p ProjectorCountdown) Name() string {
	return "projector_countdown"
}

// MeetingID returns the meetingID for the object.
func (p ProjectorCountdown) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.ProjectorCountdown_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meetingID: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (p ProjectorCountdown) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return p.see
	}
	return nil
}

func (p ProjectorCountdown) see(ctx context.Context, ds *dsfetch.Fetch, projectorCountdownIDs ...int) ([]int, error) {
	return meetingPerm(ctx, ds, p, projectorCountdownIDs, perm.ProjectorCanSee)
}
