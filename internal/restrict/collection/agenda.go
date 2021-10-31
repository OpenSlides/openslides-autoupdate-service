package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// AgendaItem handels permission for the agenda.
type AgendaItem struct{}

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

func (a AgendaItem) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, agendaID int) (bool, error) {
	meetingID, err := a.meetingID(ctx, ds, agendaID)
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

	isHidden := ds.AgendaItem_IsHidden(agendaID).ErrorLater(ctx)
	isInternal := ds.AgendaItem_IsInternal(agendaID).ErrorLater(ctx)
	if err := ds.Err(); err != nil {
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

func (a AgendaItem) modeB(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, agendaID int) (bool, error) {
	meetingID, err := a.meetingID(ctx, ds, agendaID)
	if err != nil {
		return false, fmt.Errorf("getting meetingID: %w", err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting permissions: %w", err)
	}

	return perms.Has(perm.AgendaItemCanSeeInternal), nil
}

func (a AgendaItem) modeC(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, agendaID int) (bool, error) {
	meetingID, err := a.meetingID(ctx, ds, agendaID)
	if err != nil {
		return false, fmt.Errorf("getting meetingID: %w", err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting permissions: %w", err)
	}

	return perms.Has(perm.AgendaItemCanManage), nil
}

func (a AgendaItem) meetingID(ctx context.Context, request *datastore.Request, id int) (int, error) {
	mid, err := request.AgendaItem_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, fmt.Errorf("fetching meeting_id for agenda_item %d: %w", id, err)
	}
	return mid, nil
}
