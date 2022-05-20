package collection

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// ListOfSpeakers handels the restriction for the list_of_speakers collection.
//
// The user can see a list of speakers if the user has list_of_speakers.can_see
// in the meeting and can see the content_object.
//
// Mode A: The user can see the list of speakers.
type ListOfSpeakers struct{}

// MeetingID returns the meetingID for the object.
func (los ListOfSpeakers) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := los.meetingID(ctx, ds, id)
	if err != nil {
		return 0, false, fmt.Errorf("fetching meeting id for los %d: %w", id, err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (los ListOfSpeakers) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return los.see
	}
	return nil
}

func (los ListOfSpeakers) see(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, losID int) (bool, error) {
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
		// TODO LAST ERROR
		return false, fmt.Errorf("content object_id has to have exacly one /, got %q", contentObjectID)
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		// TODO LAST ERROR
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
		// TODO LAST ERROR
		return false, fmt.Errorf("unknown content_object collection %q", parts[0])
	}
}

func (los ListOfSpeakers) meetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, error) {
	mid, err := ds.ListOfSpeakers_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, fmt.Errorf("fetching meeting_id: %w", err)
	}
	return mid, nil
}
