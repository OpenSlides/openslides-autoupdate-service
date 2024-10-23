package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

// AgendaItem handels permission for the agenda.
//
//	The user can see an agenda item if any of:
//	   The user has `agenda_item.can_manage` in the meeting
//	   The user has `agenda_item.can_see_internal` in the meeting and the item has `is_hidden` set to `false`.
//	   The user has `agenda_item.can_see` in the meeting and the item has `is_hidden` and `is_internal` set to `false`.
//
// Mode A: The user can see the agenda item.
//
// Mode B: The user has agenda_item.can_see_internal.
//
// Mode C: The user has agenda_item.can_manage.
type AgendaItem struct{}

// Name returns the collection name.
func (a AgendaItem) Name() string {
	return "agenda_item"
}

// MeetingID returns the meetingID for the object.
func (a AgendaItem) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	mid, err := ds.AgendaItem_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meeting id: %w", err)
	}

	if mid == 0 {
		key, err := dskey.FromParts("agenda_item", id, "meeting_id")
		if err != nil {
			return 0, false, fmt.Errorf("building key for logging: %w", err)
		}

		value, err := ds.Get(ctx, key)
		if err != nil {
			return 0, false, fmt.Errorf("getting value from %s for logging: %w", key, err)
		}

		return 0, false, fmt.Errorf("agenda/%d/meeting_id == 0: %w", id, datastore.InvalidDataError{Key: key, Value: value[key]})
	}
	return mid, true, nil
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

func (a AgendaItem) see(ctx context.Context, ds *dsfetch.Fetch, agendaIDs ...int) ([]int, error) {
	return eachMeeting(ctx, ds, a, agendaIDs, func(meetingID int, ids []int) ([]int, error) {
		perms, err := perm.FromContext(ctx, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting permissions: %w", err)
		}

		if perms.Has(perm.AgendaItemCanManage) {
			return ids, nil
		}

		allowed, err := eachCondition(ids, func(agendaID int) (bool, error) {
			var isHidden bool
			var isInternal bool
			ds.AgendaItem_IsHidden(agendaID).Lazy(&isHidden)
			ds.AgendaItem_IsInternal(agendaID).Lazy(&isInternal)
			if err := ds.Execute(ctx); err != nil {
				return false, fmt.Errorf("fetching isHidden and isInternal: %w", err)
			}

			if perms.Has(perm.AgendaItemCanSeeInternal) && !isHidden {
				return true, nil
			}

			if perms.Has(perm.AgendaItemCanSee) && (!isHidden && !isInternal) {
				return true, nil
			}

			return false, nil
		})
		if err != nil {
			return nil, fmt.Errorf("checking agende is hidden and is internal with perm can_see_internal and can_see: %w", err)
		}
		return allowed, nil
	})
}

func (a AgendaItem) modeB(ctx context.Context, ds *dsfetch.Fetch, agendaIDs ...int) ([]int, error) {
	return meetingPerm(ctx, ds, a, agendaIDs, perm.AgendaItemCanSeeInternal)
}

func (a AgendaItem) modeC(ctx context.Context, ds *dsfetch.Fetch, agendaIDs ...int) ([]int, error) {
	return meetingPerm(ctx, ds, a, agendaIDs, perm.AgendaItemCanManage)
}
