package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// Tag handels the restrictions for the tag collection.
//
// The user can see a tag, if the user can see the tag's meeting.
//
// Mode A: The user can see the tag.
type Tag struct{}

// MeetingID returns the meetingID for the object.
func (t Tag) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.Tag_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("get meeting id: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the field restriction for each mode.
func (t Tag) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return todoToSingle(t.see)
	}
	return nil
}

func (t Tag) see(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, tagID int) (bool, error) {
	meetingID, err := ds.Tag_MeetingID(tagID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("fetching meeting id of tag %d: %w", tagID, err)
	}
	return Meeting{}.see(ctx, ds, mperms, meetingID)
}
