package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// MotionCategory handels restrictions of the collection motion_category.
//
// The user can see a motion category if the user has motion.can_see.
//
// Mode A: The user can see the motion category.
type MotionCategory struct{}

// Name returns the collection name.
func (m MotionCategory) Name() string {
	return "motion_category"
}

// MeetingID returns the meetingID for the object.
func (m MotionCategory) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.MotionCategory_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meetingID: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (m MotionCategory) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.see
	}
	return nil
}

func (m MotionCategory) see(ctx context.Context, fetcher *dsfetch.Fetch, motionCategoryIDs []int) ([]attribute.Func, error) {
	return meetingPerm(ctx, fetcher, m, motionCategoryIDs, perm.MotionCanSee)
}
