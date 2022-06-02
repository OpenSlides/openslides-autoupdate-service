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
		return todoToSingle(m.see)
	}
	return nil
}

func (m MotionWorkflow) see(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, motionWorkflowID int) (bool, error) {
	meetingID, err := ds.MotionWorkflow_MeetingID(motionWorkflowID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting meetingID: %w", err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting permission: %w", err)
	}

	return perms.Has(perm.MotionCanSee), nil
}
