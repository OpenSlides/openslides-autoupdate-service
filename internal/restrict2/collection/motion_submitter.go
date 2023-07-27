package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
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

func (m MotionSubmitter) see(ctx context.Context, fetcher *dsfetch.Fetch, motionSubmitterIDs []int) ([]Tuple, error) {
	submitterToMotion := make(map[int]int, len(motionSubmitterIDs))
	motionIDs := set.New[int]()
	for _, submitterID := range motionSubmitterIDs {
		motionID := fetcher.MotionSubmitter_MotionID(submitterID).ErrorLater(ctx)
		submitterToMotion[submitterID] = motionID
		motionIDs.Add(motionID)
	}
	if err := fetcher.Err(); err != nil {
		return nil, fmt.Errorf("getting motionIDs: %w", err)
	}

	motionAttrList, err := Collection(ctx, "motion").Modes("C")(ctx, fetcher, motionIDs.List())
	if err != nil {
		return nil, fmt.Errorf("checking motion.see: %w", err)
	}

	indexMotionToAttr := make(map[int]attribute.Func, len(motionAttrList))
	for _, motionAttr := range motionAttrList {
		indexMotionToAttr[motionAttr.Key.ID] = motionAttr.Value
	}

	results := make([]Tuple, len(motionSubmitterIDs))
	for i, id := range motionSubmitterIDs {
		results[i].Key = modeKey(m, id, "A")
		motionID := submitterToMotion[id]
		results[i].Value = indexMotionToAttr[motionID]
	}

	return results, nil
}
