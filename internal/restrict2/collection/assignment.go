package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Assignment handels restrictions for the assignment collection.
type Assignment struct{}

// See tells, if a user can see the agenda item.
func (a Assignment) See(ctx context.Context, fetch *datastore.Fetcher, mperms perm.MeetingPermission, assignmentID int) (bool, error) {
	meetingID, err := a.meetingID(ctx, fetch, assignmentID)
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

	losID := datastore.Int(ctx, fetch.Fetch, "assignment/%d/list_of_speakers_id", assignmentID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching losID: %w", err)
	}

	if losID != 0 {
		canSeeLOS, err := ListOfSpeakers{}.See(ctx, fetch, mperms, losID)
		if err != nil {
			return false, fmt.Errorf("calculating los see: %w", err)
		}

		if canSeeLOS {
			return true, nil
		}
	}

	agendaID := datastore.Int(ctx, fetch.Fetch, "assignment/%d/agenda_item_id", assignmentID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching agendaID: %w", err)
	}

	if agendaID != 0 {
		canSeeAgenda, err := AgendaItem{}.See(ctx, fetch, mperms, agendaID)
		if err != nil {
			return false, fmt.Errorf("calculating agendaItem see: %w", err)
		}

		if canSeeAgenda {
			return true, nil
		}
	}

	return false, nil
}

// Modes returns the restricter for the a restriction mode.
func (a Assignment) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return allways
	case "B":
		return a.modeB
	}
	return nil
}

func (a Assignment) modeB(ctx context.Context, fetch *datastore.Fetcher, mperms perm.MeetingPermission, assignmentID int) (bool, error) {
	meetingID, err := a.meetingID(ctx, fetch, assignmentID)
	if err != nil {
		return false, fmt.Errorf("fetching meeting id: %w", err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("fetching permissions for meeting %d: %w", meetingID, err)
	}

	return perms.Has(perm.AssignmentCanSee), nil
}

func (a Assignment) meetingID(ctx context.Context, fetch *datastore.Fetcher, id int) (int, error) {
	mid := datastore.Int(ctx, fetch.FetchIfExist, "assignment/%d/meeting_id", id)
	if err := fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching meeting_id for assignment %d: %w", id, err)
	}
	return mid, nil
}
