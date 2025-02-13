package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

// Projector handels the restriction for the projector collection.
//
// The user can see a projector,
// * if the user has projector.can_see or
// * the projector is the reference projector and the user has meeting.can_see_autopilot.
//
// If the projector has internal=true, then the user needs projector.can_manage
//
// Mode A: The user can see the projector.
type Projector struct{}

// Name returns the collection name.
func (p Projector) Name() string {
	return "projector"
}

// MeetingID returns the meetingID for the object.
func (p Projector) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.Projector_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meetingID: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (p Projector) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return p.see
	}
	return nil
}

func (p Projector) see(ctx context.Context, ds *dsfetch.Fetch, projectorIDs ...int) ([]int, error) {
	return eachMeeting(ctx, ds, p, projectorIDs, func(meetingID int, projectorIDs []int) ([]int, error) {
		perms, err := perm.FromContext(ctx, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting permission: %w", err)
		}

		if perms.Has(perm.ProjectorCanManage) {
			return projectorIDs, nil
		}

		internalProjectorIDs := make([]bool, len(projectorIDs))
		usedAsReference := make([]dsfetch.Maybe[int], len(projectorIDs))
		for i, projectorID := range projectorIDs {
			ds.Projector_IsInternal(projectorID).Lazy(&internalProjectorIDs[i])
			ds.Projector_UsedAsReferenceProjectorMeetingID(projectorID).Lazy(&usedAsReference[i])
		}

		if err := ds.Execute(ctx); err != nil {
			return nil, fmt.Errorf("get internal state of projectors: %w", err)
		}

		var allowed []int
		for i, projectorID := range projectorIDs {
			if internalProjectorIDs[i] {
				continue
			}

			if !usedAsReference[i].Null() && perms.Has(perm.MeetingCanSeeAutopilot) {
				allowed = append(allowed, projectorID)
			}

			if perms.Has(perm.ProjectorCanSee) {
				allowed = append(allowed, projectorID)
			}
		}

		return allowed, nil
	})
}
