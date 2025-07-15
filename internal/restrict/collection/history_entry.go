package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
	"github.com/OpenSlides/openslides-go/perm"
)

// HistoryEntry handles restrictions of the collection history_entry.
//
// For entries, that belong to a meeting (like motion), a user requires the
// permission `meeting.can_see_history`. For other entries (like for an user),
// only organization admins can see them.
//
// Mode A: The user can see the the history entry.
type HistoryEntry struct{}

// Name returns the collection name.
func (h HistoryEntry) Name() string {
	return "history_entry"
}

// MeetingID returns false since a HistoryEntry has no meeting.
func (h HistoryEntry) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	return 0, false, nil
}

// Modes returns the field restricters for the collection.
func (h HistoryEntry) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return h.see
	}
	return nil
}

func (h HistoryEntry) see(ctx context.Context, ds *dsfetch.Fetch, ids ...int) ([]int, error) {
	requestUser, err := perm.RequestUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting request user: %w", err)
	}

	isOrgaAdmin, err := perm.HasOrganizationManagementLevel(ctx, ds, requestUser, perm.OMLCanManageOrganization)
	if err != nil {
		return nil, fmt.Errorf("checking for superadmin: %w", err)
	}

	return eachContentObjectCollection(ctx, ds.HistoryEntry_ModelID, ids, func(collection string, modelID int, ids []int) ([]int, error) {
		meetingID, hasMeeting, err := Collection(ctx, collection).MeetingID(ctx, ds, modelID)
		if err != nil {
			return nil, fmt.Errorf("get meeting of %s/%d: %w", collection, modelID, err)
		}
		if !hasMeeting {
			if isOrgaAdmin {
				return ids, nil
			}
			return nil, nil
		}

		perms, err := perm.FromContext(ctx, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting permissions: %w", err)
		}

		if perms.Has(perm.MeetingCanSeeHistory) {
			return ids, nil
		}

		return nil, nil
	})
}
