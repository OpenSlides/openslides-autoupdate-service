package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// MotionChangeRecommendation handels restrictions of the collection motion_change_recommendation.
//
// The user can see a motion change recommendation if any of:
//
//	The user has motion.can_manage.
//	The user has motion.can_see and the motion change recommendation has internal set to false.
//
// Mode A: The user can see the motion change recommendation.
type MotionChangeRecommendation struct{}

// MeetingID returns the meetingID for the object.
func (m MotionChangeRecommendation) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.MotionChangeRecommendation_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meetingID: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (m MotionChangeRecommendation) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.see
	}
	return nil
}

func (m MotionChangeRecommendation) see(ctx context.Context, ds *dsfetch.Fetch, motionChangeRecommendationIDs ...int) ([]int, error) {
	return eachMeeting(ctx, ds, m, motionChangeRecommendationIDs, func(meetingID int, ids []int) ([]int, error) {
		perms, err := perm.FromContext(ctx, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting permissions: %w", err)
		}

		if perms.Has(perm.MotionCanManage) {
			return ids, nil
		}

		if !perms.Has(perm.MotionCanSee) {
			return nil, nil
		}

		allowed, err := eachCondition(ids, func(motionChangeRecommendationID int) (bool, error) {
			internal, err := ds.MotionChangeRecommendation_Internal(motionChangeRecommendationID).Value(ctx)
			if err != nil {
				return false, fmt.Errorf("getting internal: %w", err)
			}

			return !internal, nil
		})

		if err != nil {
			return nil, fmt.Errorf("checking internal state: %w", err)
		}
		return allowed, nil
	})
}
