package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// MotionState handels restrictions of the collection motion_state.
type MotionState struct{}

// Modes returns the restrictions modes for the meeting collection.
func (m MotionState) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.see
	}
	return nil
}

func (m MotionState) see(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, motionStateID int) (bool, error) {
	meetingID := fetch.Field().MotionState_MeetingID(ctx, motionStateID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("getting meetingID: %w", err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting permission: %w", err)
	}

	return perms.Has(perm.MotionCanSee), nil
}
