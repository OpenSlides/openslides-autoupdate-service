package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
	"github.com/OpenSlides/openslides-go/perm"
)

// HistoryPosition handles restrictions of the collection history_position.
//
// Only a superadmin can see it.
//
// Mode A: The user can see the the history position.
type HistoryPosition struct{}

// Name returns the collection name.
func (h HistoryPosition) Name() string {
	return "history_position"
}

// MeetingID returns false since a HistoryPosition has no meeting.
func (h HistoryPosition) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	return 0, false, nil
}

// Modes returns the field restricters for the collection.
func (h HistoryPosition) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return h.see
	}
	return nil
}

func (h HistoryPosition) see(ctx context.Context, ds *dsfetch.Fetch, historyPositionIDs ...int) ([]int, error) {
	requestUser, err := perm.RequestUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting request user: %w", err)
	}

	isSuperadmin, err := perm.HasOrganizationManagementLevel(ctx, ds, requestUser, perm.OMLSuperadmin)
	if err != nil {
		return nil, fmt.Errorf("checking for superadmin: %w", err)
	}

	if isSuperadmin {
		return historyPositionIDs, nil
	}

	return nil, nil
}
