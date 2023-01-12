package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

// MotionChangeRecommendation handels restrictions of the collection motion_change_recommendation.
//
// The user can see a motion change recommendation if any of:
//
//	The user has motion.can_manage.
//	The user has motion.can_see and the motion change recommendation has internal set to false.
//
// Mode A: The user can see the motion change recommendation.
type MotionChangeRecommendation struct {
	name string
}

// Name returns the collection name.
func (m MotionChangeRecommendation) Name() string {
	return m.name
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

func (m MotionChangeRecommendation) see(ctx context.Context, ds *dsfetch.Fetch, mperms perm.MeetingPermission, attrMap AttributeMap, motionChangeRecommendationIDs ...int) error {
	return eachMeeting(ctx, ds, m, motionChangeRecommendationIDs, func(meetingID int, ids []int) error {
		groupMap, err := mperms.Meeting(ctx, ds, meetingID)
		if err != nil {
			return fmt.Errorf("groupMap: %w", err)
		}

		attrInternal := Attributes{
			GlobalPermission: byte(perm.OMLSuperadmin),
			GroupIDs:         groupMap[perm.MotionCanManage],
		}

		attrPublic := Attributes{
			GlobalPermission: byte(perm.OMLSuperadmin),
			GroupIDs:         groupMap[perm.MotionCanSee],
		}

		for _, motionChangeRecommendationID := range motionChangeRecommendationIDs {
			internal, err := ds.MotionChangeRecommendation_Internal(motionChangeRecommendationID).Value(ctx)
			if err != nil {
				return fmt.Errorf("getting internal state of motion change recommendation %d: %w", motionChangeRecommendationID, err)
			}

			attr := &attrPublic
			if internal {
				attr = &attrInternal
			}

			attrMap.Add(dskey.Key{Collection: m.name, ID: motionChangeRecommendationID, Field: "A"}, attr)
		}

		return nil
	})
}
