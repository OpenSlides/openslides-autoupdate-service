package collection

import (
	"context"
	"fmt"
	"slices"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// MotionChangeRecommendation handels restrictions of the collection motion_change_recommendation.
//
// The user can see a motion change recommendation if any of:
//
//	The user has motion.can_manage_metadata.
//	The user has motion.can_see AND
//		motion_change_recommendation.internal is set to false AND
//		motion_change_recommendation.motion_id.state_id.is_internal is set to false.
//
// Mode A: The user can see the motion change recommendation.
type MotionChangeRecommendation struct{}

// Name returns the collection name.
func (m MotionChangeRecommendation) Name() string {
	return "motion_change_recommendation"
}

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
	motionIDs := make([]int, len(motionChangeRecommendationIDs))
	for i, id := range motionChangeRecommendationIDs {
		ds.MotionChangeRecommendation_MotionID(id).Lazy(&motionIDs[i])
	}

	if err := ds.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fechting motion ids: %w", err)
	}

	allowedMotionIDs, err := Collection(ctx, Motion{}.Name()).Modes("C")(ctx, ds, motionIDs...)
	if err != nil {
		return nil, fmt.Errorf("checking motions: %w", err)
	}

	return eachMeeting(ctx, ds, m, motionChangeRecommendationIDs, func(meetingID int, ids []int) ([]int, error) {
		perms, err := perm.FromContext(ctx, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting permissions: %w", err)
		}

		if perms.Has(perm.MotionCanManageMetadata) {
			return ids, nil
		}

		allowed, err := eachCondition(ids, func(motionChangeRecommendationID int) (bool, error) {
			motionID, err := ds.MotionChangeRecommendation_MotionID(motionChangeRecommendationID).Value(ctx)
			if err != nil {
				return false, fmt.Errorf("getting motion_id")
			}

			if !slices.Contains(allowedMotionIDs, motionID) {
				return false, nil
			}

			internalChangeRecommendation, err := ds.MotionChangeRecommendation_Internal(motionChangeRecommendationID).Value(ctx)
			if err != nil {
				return false, fmt.Errorf("getting internal: %w", err)
			}

			if internalChangeRecommendation {
				return false, nil
			}

			stateID, err := ds.Motion_StateID(motionID).Value(ctx)
			if err != nil {
				return false, fmt.Errorf("getting state id: %w", err)
			}

			internalState, err := ds.MotionState_IsInternal(stateID).Value(ctx)
			if err != nil {
				return false, fmt.Errorf("getting state internal: %w", err)
			}

			return !internalState, nil
		})
		if err != nil {
			return nil, fmt.Errorf("checking internal state: %w", err)
		}
		return allowed, nil
	})
}
