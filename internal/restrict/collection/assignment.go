package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
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
func (a Assignment) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	mid, err := ds.Assignment_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meeting id: %w", err)
	}
	return mid, true, nil
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

func (a Assignment) see(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, assignmentIDs ...int) ([]int, error) {
	return eachMeeting(ctx, ds, a, assignmentIDs, func(meetingID int, ids []int) ([]int, error) {
		perms, err := mperms.Meeting(ctx, meetingID)
		if err != nil {
			return nil, fmt.Errorf("fetching permissions for meeting %d: %w", meetingID, err)
		}

		if perms.Has(perm.AssignmentCanSee) {
			return assignmentIDs, nil
		}

		allowed, err := eachCondition(ids, func(assignmentID int) (bool, error) {
			agendaID, exist, err := ds.Assignment_AgendaItemID(assignmentID).Value(ctx)
			if err != nil {
				return false, fmt.Errorf("fetching agendaID: %w", err)
			}

			if exist {
				canSeeAgenda, err := AgendaItem{}.see(ctx, ds, mperms, agendaID)
				if err != nil {
					return false, fmt.Errorf("calculating agendaItem see: %w", err)
				}

				if len(canSeeAgenda) == 1 {
					return true, nil
				}
			}
			return false, nil
		})

		if err != nil {
			return nil, fmt.Errorf("checking if agenda item exist and can be seen: %w", err)
		}
		return allowed, nil
	})
}

func (a Assignment) modeB(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, assignmentIDs ...int) ([]int, error) {
	return meetingPerm(ctx, ds, a, assignmentIDs, mperms, perm.AssignmentCanSee)
}
