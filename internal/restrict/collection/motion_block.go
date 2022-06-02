package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

// MotionBlock handels restrictions of the collection motion_block.
//
// The user can see a motion block if any of:
//     The user has motion.can_manage.
//     The user has motion.can_see and the motion block has internal set to false.
//
// Mode A: Mode B restrictions or the user can see the agenda item (motion_block/agenda_item_id).
//
// Mode B: The user can see the motion block.
type MotionBlock struct{}

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
		return m.modeA
	case "B":
		return m.see
	}
	return nil
}

func (m MotionBlock) see(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, motionBlockIDs ...int) ([]int, error) {
	return eachMeeting(ctx, ds, m, motionBlockIDs, func(meetingID int, ids []int) ([]int, error) {
		perms, err := mperms.Meeting(ctx, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting permissions: %w", err)
		}

		if perms.Has(perm.MotionCanManage) {
			return ids, nil
		}

		if !perms.Has(perm.MotionCanSee) {
			return nil, nil
		}

		var allowed []int
		for _, motionBlockID := range ids {
			internal, err := ds.MotionBlock_Internal(motionBlockID).Value(ctx)
			if err != nil {
				return nil, fmt.Errorf("getting internal: %w", err)
			}

			if !internal {
				allowed = append(allowed, motionBlockID)
			}
		}
		return allowed, nil
	})
}

func (m MotionBlock) modeA(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, motionBlockIDs ...int) ([]int, error) {
	allowed, err := m.see(ctx, ds, mperms, motionBlockIDs...)
	if err != nil {
		return nil, fmt.Errorf("checking see: %w", err)
	}

	if len(allowed) == len(motionBlockIDs) {
		return motionBlockIDs, nil
	}

	notAllowed := set.New(motionBlockIDs...)
	notAllowed.Remove(allowed...)

	for _, motionBlockID := range notAllowed.List() {
		agendaItemID, exist, err := ds.MotionBlock_AgendaItemID(motionBlockID).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting agendaItem: %w", err)
		}

		if exist {
			see, err := AgendaItem{}.see(ctx, ds, mperms, agendaItemID)
			if err != nil {
				return nil, fmt.Errorf("checking agendaItem %d: %w", agendaItemID, err)
			}

			if len(see) == 1 {
				allowed = append(allowed, motionBlockID)
			}
		}
	}

	return allowed, nil
}
