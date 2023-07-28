package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
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

func (m MotionSubmitter) see(ctx context.Context, fetcher *dsfetch.Fetch, motionSubmitterIDs []int) ([]attribute.Func, error) {
	// TODO: It would be faster to call Collection().Mode(C) with all motionIDs at once.
	return byRelationField(ctx, fetcher.MotionSubmitter_MotionID, motionSubmitterIDs, func(motionID int, motionSubmitterIDs []int) ([]attribute.Func, error) {
		motionAttr, err := Collection(ctx, "motion").Modes("C")(ctx, fetcher, []int{motionID})
		if err != nil {
			return nil, fmt.Errorf("checking motion.see: %w", err)
		}

		return attributeFuncList(len(motionSubmitterIDs), motionAttr[0]), nil
	})
}
