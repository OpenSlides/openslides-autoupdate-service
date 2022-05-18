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
type MotionSubmitter struct{}

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

func (m MotionSubmitter) see(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, motionSubmitterID int) (bool, error) {
	motionID, err := ds.MotionSubmitter_MotionID(motionSubmitterID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting motion id id: %w", err)
	}

	seeMotion, err := Motion{}.see(ctx, ds, mperms, motionID)
	if err != nil {
		return false, fmt.Errorf("checking motion %d can see: %w", motionID, err)
	}

	return seeMotion, nil
}
