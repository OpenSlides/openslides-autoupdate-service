package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// MotionSubmitter handels restrictions of the collection motion_submitter.
//
// The user can see a motion submitter if the user can see the linked motion.
//
// Mode A: The user can see the motion submitter.
type MotionSubmitter struct {
	name string
}

// Name returns the collection name.
func (m MotionSubmitter) Name() string {
	return m.name
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

func (m MotionSubmitter) see(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, attrMap AttributeMap, motionSubmitterIDs ...int) error {
	return eachRelationField(ctx, ds.MotionSubmitter_MotionID, motionSubmitterIDs, func(motionID int, ids []int) error {
		// TODO: This only works if motion is calculated before motion_submitter
		for _, id := range ids {
			attrMap.Add(m.name, id, "A", attrMap.Get("motion", motionID, "C"))
		}
		return nil
	})
}
