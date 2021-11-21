package collection

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// ListOfSpeakers handels the restriction for the list_of_speakers collection.
type ListOfSpeakers struct{}

// Modes returns the restrictions modes for the meeting collection.
func (los ListOfSpeakers) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return los.see
	}
	return nil
}

func (los ListOfSpeakers) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, losID int) (bool, error) {
	mid, err := los.meetingID(ctx, ds, losID)
	if err != nil {
		return false, fmt.Errorf("fetching meeting id for los %d: %w", losID, err)
	}

	perms, err := mperms.Meeting(ctx, mid)
	if err != nil {
		return false, fmt.Errorf("getting perms for meetind %d: %w", mid, err)
	}

	if canSee := perms.Has(perm.ListOfSpeakersCanSee); !canSee {
		return false, nil
	}

	contentObjectID, err := ds.ListOfSpeakers_ContentObjectID(losID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting content object id: %w", err)
	}

	parts := strings.Split(contentObjectID, "/")
	if len(parts) != 2 {
		return false, fmt.Errorf("content object_id has to have exacly one /, got %q", contentObjectID)
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		return false, fmt.Errorf("second part of content_object_id has to be int, got %q", parts[1])
	}

	switch parts[0] {
	case "motion":
		return Motion{}.see(ctx, ds, mperms, id)
	case "motion_block":
		return MotionBlock{}.see(ctx, ds, mperms, id)
	case "assignment":
		return Assignment{}.see(ctx, ds, mperms, id)
	case "topic":
		return Topic{}.see(ctx, ds, mperms, id)
	case "mediafile":
		return Mediafile{}.see(ctx, ds, mperms, id)
	default:
		return false, fmt.Errorf("unknown content_object collection %q", parts[0])
	}
}

func (los ListOfSpeakers) meetingID(ctx context.Context, ds *datastore.Request, id int) (int, error) {
	mid, err := ds.ListOfSpeakers_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, fmt.Errorf("fetching meeting_id for the list of speakers %d: %w", id, err)
	}
	return mid, nil
}
