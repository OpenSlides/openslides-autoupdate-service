package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Assignment handels restrictions for the assignment collection.
type Assignment struct{}

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

	losID, err := ds.Assignment_ListOfSpeakersID(assignmentID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("fetching losID: %w", err)
	}

	canSeeLOS, err := ListOfSpeakers{}.see(ctx, ds, mperms, losID)
	if err != nil {
		return false, fmt.Errorf("calculating los see: %w", err)
	}

	if canSeeLOS {
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
		return 0, fmt.Errorf("fetching meeting_id for assignment %d: %w", id, err)
	}
	return mid, nil
}
