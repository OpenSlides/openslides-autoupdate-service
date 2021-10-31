package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// MotionChangeRecommendation handels restrictions of the collection motion_change_recommendation.
type MotionChangeRecommendation struct{}

// Modes returns the restrictions modes for the meeting collection.
func (m MotionChangeRecommendation) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.see
	}
	return nil
}

func (m MotionChangeRecommendation) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, motionChangeRecommendationID int) (bool, error) {
	meetingID, err := ds.MotionChangeRecommendation_MeetingID(motionChangeRecommendationID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting meetingID: %w", err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting permissions: %w", err)
	}

	if perms.Has(perm.MotionCanManage) {
		return true, nil
	}

	if !perms.Has(perm.MotionCanSee) {
		return false, nil
	}

	internal, err := ds.MotionChangeRecommendation_Internal(motionChangeRecommendationID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting internal: %w", err)
	}

	if !internal {
		return true, nil
	}

	return false, nil
}
