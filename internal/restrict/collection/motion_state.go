package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

// MotionState handels restrictions of the collection motion_state.
//
// The user can see a motion state if the user has motion.can_see or see a
// motion in motion_state/motion_ids.
//
// Mode A: The user can see the motion state.
type MotionState struct{}

// Name returns the collection name.
func (m MotionState) Name() string {
	return "motion_state"
}

// MeetingID returns the meetingID for the object.
func (m MotionState) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.MotionState_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meetingID: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (m MotionState) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.see
	}
	return nil
}

func (m MotionState) see(ctx context.Context, ds *dsfetch.Fetch, motionStateIDs ...int) ([]int, error) {
	return eachMeeting(ctx, ds, m, motionStateIDs, func(meetingID int, motionStateIDs []int) ([]int, error) {
		perms, err := perm.FromContext(ctx, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting permission: %w", err)
		}

		if perms.Has(perm.MotionCanSee) {
			return motionStateIDs, nil
		}

		motionIDsList := make([][]int, len(motionStateIDs))
		for i, id := range motionStateIDs {
			ds.MotionState_MotionIDs(id).Lazy(&motionIDsList[i])
		}
		if err := ds.Execute(ctx); err != nil {
			return nil, fmt.Errorf("fetching motion ids: %w", err)
		}

		allowed := make([]int, 0, len(motionStateIDs))
		for i, motionIDs := range motionIDsList {
			allowedMotions, err := Collection(ctx, Motion{}.Name()).Modes("C")(ctx, ds, motionIDs...)
			if err != nil {
				return nil, fmt.Errorf("check restriction of motions: %w", err)
			}

			if len(allowedMotions) > 0 {
				allowed = append(allowed, motionStateIDs[i])
			}
		}

		return allowed, nil
	})

}
