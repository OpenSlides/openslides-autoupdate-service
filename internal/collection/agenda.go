package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/perm"
)

// AgendaItem handels the permission of agenda_item objects.
type AgendaItem struct {
	dp dataprovider.DataProvider
}

// NewAgendaItem initializes an AgendaItem.
func NewAgendaItem(dp dataprovider.DataProvider) *AgendaItem {
	return &AgendaItem{
		dp: dp,
	}
}

// Connect registers the AgendaItem.
func (a *AgendaItem) Connect(s perm.HandlerStore) {
	s.RegisterReadHandler("agenda_item", perm.ReadCheckerFunc(a.read))
}

func (a *AgendaItem) read(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
	grouped, err := groupByMeeting(ctx, a.dp, userID, fqfields)
	if err != nil {
		return fmt.Errorf("grouping fqfields: %w", err)
	}

	var lastID int
	var hasPerm bool
	for _, g := range grouped {
		for _, fqfield := range g.fqfields {
			if lastID != fqfield.ID {
				fqid := fmt.Sprintf("agenda_item/%d", fqfield.ID)
				var isInternal bool
				if err := a.dp.Get(ctx, fqid+"/is_internal", &isInternal); err != nil {
					return fmt.Errorf("getting is_internal field: %w", err)
				}

				var isHidden bool
				if err := a.dp.Get(ctx, fqid+"/is_hidden", &isHidden); err != nil {
					return fmt.Errorf("getting is_hidden field: %w", err)
				}

				requiredPerm := "agenda.can_see"
				if isInternal {
					requiredPerm = "agenda.can_see_internal_items"
				}
				if isHidden {
					requiredPerm = "agenda.can_manage"
				}
				hasPerm = false
				if g.perm.Has(requiredPerm) {
					hasPerm = true
				}
			}

			if !hasPerm {
				continue
			}

			if fqfield.Field == "duration" && !g.perm.Has("agenda.can_see_internal_items") {
				continue
			}

			if fqfield.Field == "comment" && !g.perm.Has("agenda.can_manage") {
				continue
			}

			result[fqfield.String()] = true
		}
	}

	return nil
}

type meetingFields struct {
	meetingID int
	perm      *perm.Permission
	fqfields  []perm.FQField
}

func groupByMeeting(ctx context.Context, dp dataprovider.DataProvider, userID int, fqfields []perm.FQField) ([]meetingFields, error) {
	var grouped []meetingFields
	var lastID int
	var meetingID int
	var err error
	for _, fqfield := range fqfields {
		if lastID != fqfield.ID {
			meetingID, err = dp.MeetingFromModel(ctx, fqfield.FQID())
			if err != nil {
				return nil, fmt.Errorf("getting meeting id for %s: %w", fqfield.String(), err)
			}

			p, err := perm.New(ctx, dp, userID, meetingID)
			if err != nil {
				return nil, fmt.Errorf("getting perms for meeting %d: %w", meetingID, err)
			}

			grouped = append(grouped, meetingFields{
				meetingID: meetingID,
				perm:      p,
			})
		}
		grouped[len(grouped)-1].fqfields = append(grouped[len(grouped)-1].fqfields, fqfield)
	}
	return grouped, nil
}
