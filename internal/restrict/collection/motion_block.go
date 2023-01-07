package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
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
type MotionBlock struct {
	name string
}

// Name returns the collection name.
func (m MotionBlock) Name() string {
	return m.name
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

func (m MotionBlock) see(ctx context.Context, ds *dsfetch.Fetch, mperms perm.MeetingPermission, attrMap AttributeMap, motionBlockIDs ...int) error {
	return eachMeeting(ctx, ds, m, motionBlockIDs, func(meetingID int, ids []int) error {
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

		for _, motionBlockID := range motionBlockIDs {
			internal, err := ds.MotionBlock_Internal(motionBlockID).Value(ctx)
			if err != nil {
				return fmt.Errorf("getting internal state of motion block %d: %w", motionBlockID, err)
			}

			attr := &attrPublic
			if internal {
				attr = &attrInternal
			}

			attrMap.Add(m.name, motionBlockID, "A", attr)
		}

		return nil
	})
}
