package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
	"github.com/OpenSlides/openslides-go/set"
)

// MotionSupporter handels restrictions of the collection motion_supporter.
//
// The user can see a motion supporter if the user can see the linked motion.
//
// Mode A: The user can see the motion supporter.
type MotionSupporter struct{}

// Name returns the collection name.
func (m MotionSupporter) Name() string {
	return "motion_supporter"
}

// MeetingID returns the meetingID for the object.
func (m MotionSupporter) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	motionID, err := ds.MotionSupporter_MotionID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting motionID: %w", err)
	}

	return Motion{}.MeetingID(ctx, ds, motionID)
}

// Modes returns the restrictions modes for the meeting collection.
func (m MotionSupporter) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.see
	}
	return nil
}

func (m MotionSupporter) see(ctx context.Context, ds *dsfetch.Fetch, motionSupporterIDs ...int) ([]int, error) {
	supporterToMotion := make(map[int]int, len(motionSupporterIDs))
	motionIDs := set.New[int]()
	for _, supporterID := range motionSupporterIDs {
		motionID, err := ds.MotionSupporter_MotionID(supporterID).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting motion id for supporter %d: %w", supporterID, err)
		}

		supporterToMotion[supporterID] = motionID
		motionIDs.Add(motionID)
	}

	allowedMotionIDs, err := Collection(ctx, "motion").Modes("C")(ctx, ds, motionIDs.List()...)
	if err != nil {
		return nil, fmt.Errorf("checking motion.see: %w", err)
	}

	if len(allowedMotionIDs) == motionIDs.Len() {
		return motionSupporterIDs, nil
	}

	if len(allowedMotionIDs) == 0 {
		return nil, nil
	}

	allowedMotionSet := set.New(allowedMotionIDs...)
	allowed := make([]int, 0, len(motionSupporterIDs))
	for _, supporterID := range motionSupporterIDs {
		if allowedMotionSet.Has(supporterToMotion[supporterID]) {
			allowed = append(allowed, supporterID)
		}
	}

	return allowed, nil
}
