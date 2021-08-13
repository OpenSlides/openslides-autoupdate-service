package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// MotionWorkflow handels restrictions of the collection motion_workflow.
type MotionWorkflow struct{}

// Modes returns the restrictions modes for the meeting collection.
func (m MotionWorkflow) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.see
	}
	return nil
}

func (m MotionWorkflow) see(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, motionWorkflowID int) (bool, error) {
	meetingID := fetch.Field().MotionWorkflow_MeetingID(ctx, motionWorkflowID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("getting meetingID: %w", err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting permission: %w", err)
	}

	return perms.Has(perm.MotionCanSee), nil
}
