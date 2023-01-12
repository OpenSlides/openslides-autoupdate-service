package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

// MotionSubmitter handels restrictions of the collection motion_submitter.
//
// The user can see a motion submitter if the user can see the linked motion.
//
// Mode A: The user can see the motion submitter.
type MotionSubmitter struct{}

// Name returns the collection name.
func (m MotionSubmitter) Name() string {
	return "motion_submitter"
}

// MeetingID returns the meetingID for the object.
func (m MotionSubmitter) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	motionID, err := ds.MotionSubmitter_MotionID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting motionID: %w", err)
	}

	return Motion{}.MeetingID(ctx, ds, motionID)
}

// Modes returns the restrictions modes for the meeting collection.
func (m MotionSubmitter) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.see
	}
	return nil
}

func (m MotionSubmitter) see(ctx context.Context, ds *dsfetch.Fetch, mperms perm.MeetingPermission, attrMap AttributeMap, motionSubmitterIDs ...int) error {
	return eachRelationField(ctx, ds.MotionSubmitter_MotionID, motionSubmitterIDs, func(motionID int, ids []int) error {
		// TODO: This only works if motion is calculated before motion_submitter
		for _, id := range ids {
			if err := attrMap.SameAs(ctx, ds, mperms, dskey.Key{Collection: m.Name(), ID: id, Field: "A"}, dskey.Key{Collection: "motion", ID: motionID, Field: "C"}); err != nil {
				return fmt.Errorf("same as: %w", err)
			}
		}
		return nil
	})
}
