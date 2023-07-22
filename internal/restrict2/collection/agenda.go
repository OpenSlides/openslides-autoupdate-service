package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
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
func (a AgendaItem) MeetingID(ctx context.Context, fetcher *dsfetch.Fetch, id int) (int, bool, error) {
	mid, err := fetcher.AgendaItem_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meeting id: %w", err)
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

func (a AgendaItem) see(ctx context.Context, fetcher *dsfetch.Fetch, agendaIDs []int) ([]Tuple, error) {
	return byMeeting(ctx, fetcher, a, agendaIDs, func(meetingID int, agendaIDs []int) ([]Tuple, error) {
		attrSuperadmin := attribute.FuncGlobalLevel(perm.OMLSuperadmin)
		attrCanManage := attribute.FuncPerm(meetingID, perm.AgendaItemCanManage)
		attrCanSeeInternal := attribute.FuncPerm(meetingID, perm.AgendaItemCanSeeInternal)
		attrCanSee := attribute.FuncPerm(meetingID, perm.AgendaItemCanSee)

		result := make([]Tuple, len(agendaIDs))
		for i, agendaID := range agendaIDs {
			isHidden := fetcher.AgendaItem_IsHidden(agendaID).ErrorLater(ctx)
			isInternal := fetcher.AgendaItem_IsInternal(agendaID).ErrorLater(ctx)
			if err := fetcher.Err(); err != nil {
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

			attr = attribute.FuncOr(attrSuperadmin, attr)

			result[i] = Tuple{modeKey(a, agendaID, "A"), attr}
		}

		return result, nil
	})
}

func (a AgendaItem) modeB(ctx context.Context, fetcher *dsfetch.Fetch, agendaIDs []int) ([]Tuple, error) {
	return meetingPerm(ctx, fetcher, a, "B", agendaIDs, perm.AgendaItemCanSeeInternal)
}

func (a AgendaItem) modeC(ctx context.Context, fetcher *dsfetch.Fetch, agendaIDs []int) ([]Tuple, error) {
	return meetingPerm(ctx, fetcher, a, "C", agendaIDs, perm.AgendaItemCanManage)
}
