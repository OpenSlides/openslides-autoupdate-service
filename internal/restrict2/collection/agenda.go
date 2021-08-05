package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// AgendaItem handels permission for the agenda.
type AgendaItem struct{}

func (a AgendaItem) see(ctx context.Context, fetch *datastore.Fetcher, mperms perm.MeetingPermission, agendaID int) (bool, error) {
	meetingID, err := a.meetingID(ctx, fetch, agendaID)
	if err != nil {
		return false, fmt.Errorf("getting meetingID: %w", err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting permissions: %w", err)
	}

	if perms.Has(perm.AgendaItemCanManage) {
		return true, nil
	}

	isHidden := datastore.Bool(ctx, fetch.FetchIfExist, "agenda_item/%d/is_hidden", agendaID)
	isInternal := datastore.Bool(ctx, fetch.FetchIfExist, "agenda_item/%d/is_internal", agendaID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching isHidden and isInternal: %w", err)
	}

	if perms.Has(perm.AgendaItemCanSeeInternal) && !isHidden {
		return true, nil
	}

	if perms.Has(perm.AgendaItemCanSee) && (!isHidden && !isInternal) {
		return true, nil
	}

	return false, nil
}

// Modes returns a map from all known modes to there restricter.
func (a AgendaItem) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return a.see
	case "B":
		return a.modeB
	case "C":
		return a.modeC
	}
	return nil
}

func (a AgendaItem) modeB(ctx context.Context, fetch *datastore.Fetcher, mperms perm.MeetingPermission, agendaID int) (bool, error) {
	meetingID, err := a.meetingID(ctx, fetch, agendaID)
	if err != nil {
		return false, fmt.Errorf("getting meetingID: %w", err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting permissions: %w", err)
	}

	return perms.Has(perm.AgendaItemCanSeeInternal), nil
}

func (a AgendaItem) modeC(ctx context.Context, fetch *datastore.Fetcher, mperms perm.MeetingPermission, agendaID int) (bool, error) {
	meetingID, err := a.meetingID(ctx, fetch, agendaID)
	if err != nil {
		return false, fmt.Errorf("getting meetingID: %w", err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting permissions: %w", err)
	}

	return perms.Has(perm.AgendaItemCanManage), nil
}

func (a AgendaItem) meetingID(ctx context.Context, fetch *datastore.Fetcher, id int) (int, error) {
	mid := datastore.Int(ctx, fetch.FetchIfExist, "agenda_item/%d/meeting_id", id)
	if err := fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching meeting_id for agenda_item %d: %w", id, err)
	}
	return mid, nil
}
