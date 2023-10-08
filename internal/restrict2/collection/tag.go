package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// Tag handels the restrictions for the tag collection.
//
// The user can see a tag, if the user can see the tag's meeting.
//
// Mode A: The user can see the tag.
type Tag struct{}

// Name returns the collection name.
func (t Tag) Name() string {
	return "tag"
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

func (t Tag) see(ctx context.Context, fetcher *dsfetch.Fetch, tagIDs []int) ([]attribute.Func, error) {
	return byMeeting(ctx, fetcher, t, tagIDs, func(meetingID int, tagIDs []int) ([]attribute.Func, error) {
		attrFuncs, err := Collection(ctx, Meeting{}.Name()).Modes("B")(ctx, fetcher, []int{meetingID})
		if err != nil {
			return nil, fmt.Errorf("checking meeting can see: %w", err)
		}

		return attributeFuncList(len(tagIDs), attrFuncs[0]), nil
	})
}
