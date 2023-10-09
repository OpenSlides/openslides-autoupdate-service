package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// MotionBlock handels restrictions of the collection motion_block.
//
// The user can see a motion block if any of:
//
//	The user has motion.can_manage.
//	The user has motion.can_see and the motion block has internal set to false.
//
// Mode A: The user can see the motion block.
type MotionBlock struct{}

// Name returns the collection name.
func (m MotionBlock) Name() string {
	return "motion_block"
}

// MeetingID returns the meetingID for the object.
func (m MotionBlock) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.MotionBlock_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meetingID: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (m MotionBlock) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.see
	}
	return nil
}

func (m MotionBlock) see(ctx context.Context, fetcher *dsfetch.Fetch, motionBlockIDs []int) ([]attribute.Func, error) {
	meetingID := make([]int, len(motionBlockIDs))
	internal := make([]bool, len(motionBlockIDs))
	for i, id := range motionBlockIDs {
		if id == 0 {
			continue
		}
		fetcher.MotionBlock_MeetingID(id).Lazy(&meetingID[i])
		fetcher.MotionBlock_Internal(id).Lazy(&internal[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching motion block data: %w", err)
	}

	attr := make([]attribute.Func, len(motionBlockIDs))
	for i, id := range motionBlockIDs {
		if id == 0 {
			continue
		}

		groupMap, err := perm.GroupMapFromContext(ctx, fetcher, meetingID[i])
		if err != nil {
			return nil, fmt.Errorf("getting group map: %w", err)
		}

		canPerm := perm.MotionCanSee
		if internal[i] {
			canPerm = perm.MotionCanManage
		}

		attr[i] = attribute.FuncInGroup(groupMap[canPerm])

	}
	return attr, nil
}
