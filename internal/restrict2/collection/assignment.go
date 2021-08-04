package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Assignment handels restrictions for the assignment collection.
type Assignment struct{}

// See tells, if a user can see the agenda item.
func (a Assignment) See(ctx context.Context, fetch *datastore.Fetcher, mperms perm.MeetingPermission, assignmentID int) (bool, error) {
	meetingID, err := a.meetingID(ctx, fetch, assignmentID)
	if err != nil {
		return false, fmt.Errorf("fetching meeting id: %w", err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("fetching permissions for meeting %d: %w", meetingID, err)
	}

	return perms.Has(perm.AssignmentCanSee), nil
}

func (a Assignment) meetingID(ctx context.Context, fetch *datastore.Fetcher, id int) (int, error) {
	mid := datastore.Int(ctx, fetch.FetchIfExist, "assignment/%d/meeting_id", id)
	if err := fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching meeting_id for assignment %d: %w", id, err)
	}
	return mid, nil
}
