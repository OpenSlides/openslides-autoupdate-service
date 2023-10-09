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
	return canSeeRelatedCollection(ctx, fetcher, fetcher.Tag_MeetingID, Collection(ctx, Meeting{}.Name()).Modes("B"), tagIDs)
}
