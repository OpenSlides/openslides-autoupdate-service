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
	return canSeeRelatedCollection(ctx, fetcher, fetcher.MotionSubmitter_MotionID, Collection(ctx, Motion{}).Modes("C"), motionSubmitterIDs)
}
