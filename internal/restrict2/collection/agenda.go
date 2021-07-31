package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// AgendaItem handels permission for the agenda.
type AgendaItem struct{}

// See tells, if a user can see the agenda item.
func (a *AgendaItem) See(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, uid int, id int) (bool, error) {
	meetingID, err := a.MeetingID(ctx, fetch, id)
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

	isHidden := datastore.Bool(ctx, fetch.FetchIfExist, "agenda/%d/is_hidden", id)

	if perms.Has(perm.AgendaItemCanSeeInternal) && !isHidden {
		return true, nil
	}

	isInternal := datastore.Bool(ctx, fetch.FetchIfExist, "agenda/%d/is_internal", id)

	if perms.Has(perm.AgendaItemCanSee) && (!isHidden || !isInternal) {
		return true, nil
	}

	return false, nil
}

// FieldRestricter is a function to restrict fields of a collection.
type FieldRestricter func(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, uid int, id int) (bool, error)

// FieldGroups returns a map from all known groups to there restricter.
func (a *AgendaItem) FieldGroups() map[string]FieldRestricter {
	return map[string]FieldRestricter{
		"group_a": a.GroupA,
		"group_b": a.GroupB,
		"group_c": a.GroupC,
	}
}

// GroupA handels restructions for fields in GroupA.
func (a *AgendaItem) GroupA(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, uid int, id int) (bool, error) {
	return true, nil
}

// GroupB handels restructions for fields in GroupB.
func (a *AgendaItem) GroupB(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, uid int, id int) (bool, error) {
	meetingID, err := a.MeetingID(ctx, fetch, id)
	if err != nil {
		return false, fmt.Errorf("getting meetingID: %w", err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting permissions: %w", err)
	}

	return perms.Has(perm.AgendaItemCanSeeInternal), nil
}

// GroupC handels restructions for fields in GroupC.
func (a *AgendaItem) GroupC(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, uid int, id int) (bool, error) {
	meetingID, err := a.MeetingID(ctx, fetch, id)
	if err != nil {
		return false, fmt.Errorf("getting meetingID: %w", err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting permissions: %w", err)
	}

	return perms.Has(perm.AgendaItemCanManage), nil
}

// MeetingID returns the meetingID for the agenda
func (a *AgendaItem) MeetingID(ctx context.Context, fetch *datastore.Fetcher, id int) (int, error) {
	mid := datastore.Int(ctx, fetch.Fetch, "meeting/%d/meeting_id", id)
	if err := fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching meeting_id for agenda_item %d: %w", id, err)
	}
	return mid, nil
}
