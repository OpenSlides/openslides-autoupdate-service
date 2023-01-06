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
type Tag struct {
	name string
}

// Name returns the collection name.
func (t Tag) Name() string {
	return t.name
}

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
		return t.see
	}
	return nil
}

func (t Tag) see(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, attrMap AttributeMap, tagIDs ...int) error {
	return eachMeeting(ctx, ds, t, tagIDs, func(meetingID int, ids []int) error {
		// TODO: This only works if meeting is calculated before tag
		for _, id := range ids {
			attrMap.Add(t.name, id, "A", attrMap.Get("meeting", meetingID, "B"))
		}
		return nil
	})
}
