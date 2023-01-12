package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
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

func (a AgendaItem) see(ctx context.Context, ds *dsfetch.Fetch, mperms perm.MeetingPermission, attrMap AttributeMap, agendaIDs ...int) error {
	return eachMeeting(ctx, ds, a, agendaIDs, func(meetingID int, ids []int) error {
		groupMap, err := mperms.Meeting(ctx, ds, meetingID)
		if err != nil {
			return fmt.Errorf("get meeting permissions: %w", err)
		}

		attrCanManage := Attributes{
			GlobalPermission: byte(perm.OMLSuperadmin),
			GroupIDs:         groupMap[perm.AgendaItemCanManage],
		}

		attrCanSeeInternal := Attributes{
			GlobalPermission: byte(perm.OMLSuperadmin),
			GroupIDs:         groupMap[perm.AgendaItemCanSeeInternal],
		}

		attrCanSee := Attributes{
			GlobalPermission: byte(perm.OMLSuperadmin),
			GroupIDs:         groupMap[perm.AgendaItemCanSee],
		}

		for _, agendaID := range ids {
			isHidden := ds.AgendaItem_IsHidden(agendaID).ErrorLater(ctx)
			isInternal := ds.AgendaItem_IsInternal(agendaID).ErrorLater(ctx)
			if err := ds.Err(); err != nil {
				return fmt.Errorf("fetching isHidden and isInternal: %w", err)
			}

			var attr Attributes
			switch {
			case isHidden:
				attr = attrCanManage
			case isInternal:
				attr = attrCanSeeInternal
			default:
				attr = attrCanSee
			}
			attrMap.Add(dskey.Key{Collection: a.Name(), ID: agendaID, Field: "A"}, &attr)

		}

		return nil
	})
}

func (a AgendaItem) modeB(ctx context.Context, ds *dsfetch.Fetch, mperms perm.MeetingPermission, attrMap AttributeMap, agendaIDs ...int) error {
	return meetingPerm(ctx, ds, a, "B", agendaIDs, mperms, perm.AgendaItemCanSeeInternal, attrMap)
}

func (a AgendaItem) modeC(ctx context.Context, ds *dsfetch.Fetch, mperms perm.MeetingPermission, attrMap AttributeMap, agendaIDs ...int) error {
	return meetingPerm(ctx, ds, a, "C", agendaIDs, mperms, perm.AgendaItemCanManage, attrMap)
}
