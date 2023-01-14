package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
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

func (m MotionSubmitter) see(ctx context.Context, ds *dsfetch.Fetch, motionSubmitterIDs ...int) ([]int, error) {
	submitterToMotion := make(map[int]int, len(motionSubmitterIDs))
	motionIDs := set.New[int]()
	for _, submitterID := range motionSubmitterIDs {
		motionID := ds.MotionSubmitter_MotionID(submitterID).ErrorLater(ctx)
		submitterToMotion[submitterID] = motionID
		motionIDs.Add(motionID)
	}
	if err := ds.Err(); err != nil {
		return nil, fmt.Errorf("getting motionIDs: %w", err)
	}

	allowedMotionIDs, err := Collection(ctx, "motion").Modes("C")(ctx, ds, motionIDs.List()...)
	if err != nil {
		return nil, fmt.Errorf("checking motion.see: %w", err)
	}

	if len(allowedMotionIDs) == motionIDs.Len() {
		return motionSubmitterIDs, nil
	}

	if len(allowedMotionIDs) == 0 {
		return nil, nil
	}

	allowedMotionSet := set.New(allowedMotionIDs...)
	allowed := make([]int, 0, len(motionSubmitterIDs))
	for _, submitterID := range motionSubmitterIDs {
		if allowedMotionSet.Has(submitterToMotion[submitterID]) {
			allowed = append(allowed, submitterID)
		}
	}

	return allowed, nil
}
