package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
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

func (m MotionChangeRecommendation) see(ctx context.Context, fetcher *dsfetch.Fetch, motionChangeRecommendationIDs []int) ([]attribute.Func, error) {
	meetingID := make([]int, len(motionChangeRecommendationIDs))
	recommendationInternal := make([]bool, len(motionChangeRecommendationIDs))
	motionID := make([]int, len(motionChangeRecommendationIDs))
	for i, id := range motionChangeRecommendationIDs {
		if id == 0 {
			continue
		}
		fetcher.MotionChangeRecommendation_MeetingID(id).Lazy(&meetingID[i])
		fetcher.MotionChangeRecommendation_Internal(id).Lazy(&recommendationInternal[i])
		fetcher.MotionChangeRecommendation_MotionID(id).Lazy(&motionID[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching motion block data: %w", err)
	}

	motionStateID := make([]int, len(motionChangeRecommendationIDs))
	for i, id := range motionID {
		if id == 0 {
			continue
		}
		fetcher.Motion_StateID(id).Lazy(&motionStateID[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching motion state ids: %w", err)
	}

	motionStateInternal := make([]bool, len(motionChangeRecommendationIDs))
	for i, id := range motionStateID {
		if id == 0 {
			continue
		}
		fetcher.MotionState_IsInternal(id).Lazy(&motionStateInternal[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching motion state internal: %w", err)
	}

	attr := make([]attribute.Func, len(motionChangeRecommendationIDs))
	for i, id := range motionChangeRecommendationIDs {
		if id == 0 {
			continue
		}

		groupMap, err := perm.GroupMapFromContext(ctx, fetcher, meetingID[i])
		if err != nil {
			return nil, fmt.Errorf("getting group map: %w", err)
		}

		canPerm := perm.MotionCanSee
		if recommendationInternal[i] || motionStateInternal[i] {
			canPerm = perm.MotionCanManageMetadata
		}

		attr[i] = attribute.FuncInGroup(groupMap[canPerm])

	}
	return attr, nil
}
