package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Assignment handels restrictions for the assignment collection.
//
// The user can see an assignment, if the user has assignment.can_see
// or the user can see the agenda item (assignment/agenda_item_id).
//
// Mode A: User can see the assignment.
//
// Mode B: The has assignment.can_see.
type Assignment struct{}

// MeetingID returns the meetingID for the object.
func (a Assignment) MeetingID(ctx context.Context, ds *datastore.Request, id int) (int, bool, error) {
	meetingID, err := a.meetingID(ctx, ds, id)
	if err != nil {
		return 0, false, fmt.Errorf("getting meetingID: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restricter for the a restriction mode.
func (a Assignment) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return a.see
	case "B":
		return a.modeB
	}
	return nil
}

func (a Assignment) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, assignmentID int) (bool, error) {
	meetingID, err := a.meetingID(ctx, ds, assignmentID)
	if err != nil {
		return false, fmt.Errorf("fetching meeting id: %w", err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("fetching permissions for meeting %d: %w", meetingID, err)
	}

	if perms.Has(perm.AssignmentCanSee) {
		return true, nil
	}

	agendaID, exist, err := ds.Assignment_AgendaItemID(assignmentID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("fetching agendaID: %w", err)
	}

	if exist {
		canSeeAgenda, err := AgendaItem{}.see(ctx, ds, mperms, agendaID)
		if err != nil {
			return false, fmt.Errorf("calculating agendaItem see: %w", err)
		}

		if canSeeAgenda {
			return true, nil
		}
	}

	return false, nil
}

func (a Assignment) modeB(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, assignmentID int) (bool, error) {
	meetingID, err := a.meetingID(ctx, ds, assignmentID)
	if err != nil {
		return false, fmt.Errorf("fetching meeting id: %w", err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("fetching permissions for meeting %d: %w", meetingID, err)
	}

	return perms.Has(perm.AssignmentCanSee), nil
}

func (a Assignment) meetingID(ctx context.Context, ds *datastore.Request, id int) (int, error) {
	mid, err := ds.Assignment_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, fmt.Errorf("fetching meeting_id: %w", err)
	}
	return mid, nil
}
