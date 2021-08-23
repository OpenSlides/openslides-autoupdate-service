package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// MotionCategory handels restrictions of the collection motion_category.
type MotionCategory struct{}

// Modes returns the restrictions modes for the meeting collection.
func (m MotionCategory) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.see
	}
	return nil
}

func (m MotionCategory) see(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, motionCategoryID int) (bool, error) {
	meetingID := fetch.Field().MotionCategory_MeetingID(ctx, motionCategoryID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("getting meetingID: %w", err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting permission: %w", err)
	}

	return perms.Has(perm.MotionCanSee), nil
}
