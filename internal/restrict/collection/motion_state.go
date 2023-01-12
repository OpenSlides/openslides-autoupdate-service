package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// MotionState handels restrictions of the collection motion_state.
//
// The user can see a motion state if the user has motion.can_see.
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

func (m MotionState) see(ctx context.Context, ds *dsfetch.Fetch, mperms perm.MeetingPermission, attrMap AttributeMap, motionStateIDs ...int) error {
	return meetingPerm(ctx, ds, m, "A", motionStateIDs, mperms, perm.MotionCanSee, attrMap)
}
