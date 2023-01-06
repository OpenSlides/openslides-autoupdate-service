package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// MotionWorkflow handels restrictions of the collection motion_workflow.
//
// The user can see a motion workflow if the user has motion.can_see.
//
// Mode A: The user can see the motion workflow.
type MotionWorkflow struct{}

// MeetingID returns the meetingID for the object.
func (m MotionWorkflow) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.MotionWorkflow_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meetingID: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (m MotionWorkflow) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.see
	}
	return nil
}

func (m MotionWorkflow) see(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, attrMap map[int]*Attributes, motionWorkflowIDs ...int) ([]int, error) {
	return meetingPerm(ctx, ds, m, motionWorkflowIDs, mperms, perm.MotionCanSee)
}
