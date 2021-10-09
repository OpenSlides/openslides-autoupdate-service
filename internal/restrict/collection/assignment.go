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

func (a Assignment) see(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, assignmentID int) (bool, error) {
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

	losID := fetch.Field().Assignment_ListOfSpeakersID(ctx, assignmentID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching losID: %w", err)
	}

	canSeeLOS, err := ListOfSpeakers{}.see(ctx, fetch, mperms, losID)
	if err != nil {
		return false, fmt.Errorf("calculating los see: %w", err)
	}

	if canSeeLOS {
		return true, nil
	}

	agendaID, exist := fetch.Field().Assignment_AgendaItemID(ctx, assignmentID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching agendaID: %w", err)
	}

	if exist {
		canSeeAgenda, err := AgendaItem{}.see(ctx, fetch, mperms, agendaID)
		if err != nil {
			return false, fmt.Errorf("calculating agendaItem see: %w", err)
		}

		if canSeeAgenda {
			return true, nil
		}
	}

	return false, nil
}

func (a Assignment) modeB(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, assignmentID int) (bool, error) {
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
	mid := fetch.Field().Assignment_MeetingID(ctx, id)
	if err := fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching meeting_id for assignment %d: %w", id, err)
	}
	return mid, nil
}
