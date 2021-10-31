package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
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

func (t Tag) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, tagID int) (bool, error) {
	meetingID, err := ds.Tag_MeetingID(tagID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("fetching meeting id of tag %d: %w", tagID, err)
	}
	return Meeting{}.see(ctx, ds, mperms, meetingID)
}
