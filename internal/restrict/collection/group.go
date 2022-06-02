package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// Group handels restrictions of the collection group.
//
// The user can see a group, if the user can see the group's meeting.
//
// Mode A: The user can see the group.
type Group struct{}

// MeetingID returns the meetingID for the object.
func (g Group) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.Group_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("fetching meeting id of group %d: %w", id, err)
	}

	return meetingID, true, nil
}

// Modes returns the field restricters for the collection group.
func (g Group) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return todoToSingle(g.see)
	}
	return nil
}

func (g Group) see(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, groupID int) (bool, error) {
	meetingID, err := ds.Group_MeetingID(groupID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("fetching meeting id of group %d: %w", groupID, err)
	}

	canSee, err := Meeting{}.see(ctx, ds, mperms, meetingID)
	if err != nil {
		return false, fmt.Errorf("can see meeting %d: %w", meetingID, err)
	}
	return canSee, nil
}
