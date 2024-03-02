package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
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
//
// Mode D: The user has agenda_item.can_see_notes
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
	case "D":
		return a.modeD
	}
	return nil
}

func (a AgendaItem) see(ctx context.Context, fetcher *dsfetch.Fetch, agendaIDs []int) ([]attribute.Func, error) {
	return byMeeting(ctx, fetcher, a, agendaIDs, func(meetingID int, agendaIDs []int) ([]attribute.Func, error) {
		groupMap, err := perm.GroupMapFromContext(ctx, fetcher, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting group map: %w", err)
		}

		attrSuperadmin := attribute.FuncOrgaLevel(perm.OMLSuperadmin)
		attrCanManage := attribute.FuncInGroup(groupMap[perm.AgendaItemCanManage])
		attrCanSeeInternal := attribute.FuncInGroup(groupMap[perm.AgendaItemCanSeeInternal])
		attrCanSee := attribute.FuncInGroup(groupMap[perm.AgendaItemCanSee])

		result := make([]attribute.Func, len(agendaIDs))
		for i, agendaID := range agendaIDs {
			if agendaID == 0 {
				continue
			}

			var isHidden bool
			var isInternal bool
			fetcher.AgendaItem_IsHidden(agendaID).Lazy(&isHidden)
			fetcher.AgendaItem_IsInternal(agendaID).Lazy(&isInternal)
			if err := fetcher.Execute(ctx); err != nil {
				return nil, fmt.Errorf("fetching isHidden and isInternal: %w", err)
			}

			var attr attribute.Func
			switch {
			case isHidden:
				attr = attrCanManage
			case isInternal:
				attr = attrCanSeeInternal
			default:
				attr = attrCanSee
			}

			result[i] = attribute.FuncOr(attrSuperadmin, attr)
		}

		return result, nil
	})
}

func (a AgendaItem) modeB(ctx context.Context, fetcher *dsfetch.Fetch, agendaIDs []int) ([]attribute.Func, error) {
	return meetingPerm(ctx, fetcher, a, agendaIDs, perm.AgendaItemCanSeeInternal)
}

func (a AgendaItem) modeC(ctx context.Context, fetcher *dsfetch.Fetch, agendaIDs []int) ([]attribute.Func, error) {
	return meetingPerm(ctx, fetcher, a, agendaIDs, perm.AgendaItemCanManage)
}

func (a AgendaItem) modeD(ctx context.Context, fetcher *dsfetch.Fetch, agendaIDs []int) ([]attribute.Func, error) {
	return meetingPerm(ctx, fetcher, a, agendaIDs, perm.AgendaItemCanSeeModeratorNotes)
}
