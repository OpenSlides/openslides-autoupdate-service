package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Tag handels the restrictions for the tag collection.
type Tag struct{}

// Modes returns the field restriction for each mode.
func (t Tag) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return t.see
	}
	return nil
}

func (t Tag) see(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, tagID int) (bool, error) {
	meetingID := fetch.Field().Tag_MeetingID(ctx, tagID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching meeting id of tag %d: %w", tagID, err)
	}
	return Meeting{}.see(ctx, fetch, mperms, meetingID)
}
