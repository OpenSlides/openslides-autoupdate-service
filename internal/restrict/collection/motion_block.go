package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
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

func (m MotionBlock) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, motionBlockID int) (bool, error) {
	meetingID, err := ds.MotionBlock_MeetingID(motionBlockID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting meetingID: %w", err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting permissions: %w", err)
	}

	if perms.Has(perm.MotionCanManage) {
		return true, nil
	}

	if !perms.Has(perm.MotionCanSee) {
		return false, nil
	}

	internal, err := ds.MotionBlock_Internal(motionBlockID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting internal: %w", err)
	}

	if !internal {
		return true, nil
	}

	return false, nil
}

func (m MotionBlock) modeA(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, motionBlockID int) (bool, error) {
	see, err := m.see(ctx, ds, mperms, motionBlockID)
	if err != nil {
		return false, fmt.Errorf("checking see: %w", err)
	}

	if see {
		return true, nil
	}

	agendaItemID, exist, err := ds.MotionBlock_AgendaItemID(motionBlockID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting agendaItem: %w", err)
	}

	if exist {
		see, err = AgendaItem{}.see(ctx, ds, mperms, agendaItemID)
		if err != nil {
			return false, fmt.Errorf("checking agendaItem %d: %w", agendaItemID, err)
		}

		if see {
			return true, nil
		}
	}

	return false, nil
}
