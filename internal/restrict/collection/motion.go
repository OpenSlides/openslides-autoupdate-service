package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Motion handels restrictions of the collection motion.
//
// The user can see a motion if:
//
//     The user has motion.can_see in the meeting, and
//     For one `restriction` in the motion's state `state/restriction` field:
//         If: `restriction` is `is_submitter`: The user needs to be a submitter of the motion
//         Else: (a permission string): The user needs the permission
//
// Mode A: Mode B restrictions or the user can see the agenda item (motion/agenda_item_id).
//
// Mode B: Mode C restrictions or can see a referenced motion in motion/all_origin_ids and motion/all_derived_motion_ids.
//
// Mode C: The user can see the motion.
//
// Mode D: Never published to any user.
type Motion struct{}

// MeetingID returns the meetingID for the object.
func (m Motion) MeetingID(ctx context.Context, ds *datastore.Request, id int) (int, bool, error) {
	meetingID, err := ds.Motion_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("fetching meeting_id: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (m Motion) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.modeA
	case "B":
		return m.modeB
	case "C":
		return m.see
	case "D":
		return m.modeD
	}
	return nil
}

func (m Motion) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, motionID int) (bool, error) {
	meetingID, err := ds.Motion_MeetingID(motionID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("fetching meeting_id: %w", err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting permissions: %w", err)
	}

	if !perms.Has(perm.MotionCanSee) {
		return false, nil
	}

	stateID, err := ds.Motion_StateID(motionID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("fetching stateID: %w", err)
	}

	restrictions, err := ds.MotionState_Restrictions(stateID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting restrictions: %w", err)
	}

	if len(restrictions) == 0 {
		return true, nil
	}

	for _, restriction := range restrictions {
		if restriction == "is_submitter" {
			submitter, err := isSubmitter(ctx, ds, mperms, motionID)
			if err != nil {
				return false, fmt.Errorf("checking for motion submitter: %w", err)
			}

			if submitter {
				return true, nil
			}
			continue
		}

		if perms.Has(perm.TPermission(restriction)) {
			return true, nil
		}
	}

	return false, nil
}

func isSubmitter(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, motionID int) (bool, error) {
	for _, submitterID := range ds.Motion_SubmitterIDs(motionID).ErrorLater(ctx) {
		if ds.MotionSubmitter_UserID(submitterID).ErrorLater(ctx) == mperms.UserID() {
			return true, nil
		}
	}
	if err := ds.Err(); err != nil {
		return false, fmt.Errorf("getting submitter: %w", err)
	}
	return false, nil
}

func (m Motion) modeA(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, motionID int) (bool, error) {
	see, err := m.modeB(ctx, ds, mperms, motionID)
	if err != nil {
		return false, fmt.Errorf("see motion: %w", err)
	}

	if see {
		return true, nil
	}

	agendaID, exist, err := ds.Motion_AgendaItemID(motionID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting agenda item: %w", err)
	}

	if exist {
		seeAgenda, err := AgendaItem{}.see(ctx, ds, mperms, agendaID)
		if err != nil {
			return false, fmt.Errorf("checking see for agenda item %d: %w", agendaID, err)
		}

		if seeAgenda {
			return true, nil
		}
	}

	return false, nil
}

func (m Motion) modeB(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, motionID int) (bool, error) {
	allOriginIDs := ds.Motion_AllOriginIDs(motionID).ErrorLater(ctx)
	allDerivedMotionIDs := ds.Motion_AllDerivedMotionIDs(motionID).ErrorLater(ctx)
	originID, hasOrigin := ds.Motion_OriginID(motionID).ErrorLater(ctx)
	derivedMotionIDs := ds.Motion_DerivedMotionIDs(motionID).ErrorLater(ctx)
	if err := ds.Err(); err != nil {
		return false, fmt.Errorf("fetching origin and derived motions: %w", err)
	}

	motionIDs := make(map[int]struct{}, len(allOriginIDs)+len(allDerivedMotionIDs)+len(derivedMotionIDs)+2)
	motionIDs[motionID] = struct{}{}
	for _, l := range [][]int{allOriginIDs, allDerivedMotionIDs, derivedMotionIDs} {
		for _, id := range l {
			motionIDs[id] = struct{}{}
		}
	}
	if hasOrigin {
		motionIDs[originID] = struct{}{}
	}

	for referenceID := range motionIDs {
		see, err := m.see(ctx, ds, mperms, referenceID)
		if err != nil {
			return false, fmt.Errorf("see motion %d: %w", referenceID, err)
		}

		if see {
			return true, nil
		}
	}
	return false, nil
}

func (m Motion) modeD(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, UserID int) (bool, error) {
	return false, nil
}
