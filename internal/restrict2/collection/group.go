package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// Group handels restrictions of the collection group.
//
// The user can see a group, if the user can see the group's meeting.
//
// Mode A: The user can see the group.
type Group struct{}

// Name returns the collection name.
func (g Group) Name() string {
	return "group"
}

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
		return g.see
	}
	return nil
}

func (g Group) see(ctx context.Context, fetcher *dsfetch.Fetch, groupIDs []int) ([]attribute.Func, error) {
	// TODO: It would be faster to call Collection().Mode(B) with all meeting IDs at once.
	return byMeeting(ctx, fetcher, g, groupIDs, func(meetingID int, ids []int) ([]attribute.Func, error) {
		meetingAttr, err := Collection(ctx, "meeting").Modes("B")(ctx, fetcher, []int{meetingID})
		if err != nil {
			return nil, fmt.Errorf("checking motion.see: %w", err)
		}

		return attributeFuncList(len(ids), meetingAttr[0]), nil
	})
}
