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
	mid, err := ds.ListOfSpeakers_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("fetching meeting_id: %w", err)
	}
	return mid, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (los ListOfSpeakers) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return los.see
	}
	return nil
}

func (los ListOfSpeakers) see(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, losIDs ...int) ([]int, error) {
	return eachMeeting(ctx, ds, los, losIDs, func(meetingID int, ids []int) ([]int, error) {
		perms, err := mperms.Meeting(ctx, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting perms for meetind %d: %w", meetingID, err)
		}

		if canSee := perms.Has(perm.ListOfSpeakersCanSee); !canSee {
			return nil, nil
		}

		// TODO bundle with same collection
		var allowed []int
		for _, id := range ids {
			contentObjectID, err := ds.ListOfSpeakers_ContentObjectID(id).Value(ctx)
			if err != nil {
				return nil, fmt.Errorf("getting content object id: %w", err)
			}

			parts := strings.Split(contentObjectID, "/")
			if len(parts) != 2 {
				// TODO LAST ERROR
				return nil, fmt.Errorf("content object_id has to have exacly one /, got %q", contentObjectID)
			}

			id, err := strconv.Atoi(parts[1])
			if err != nil {
				// TODO LAST ERROR
				return nil, fmt.Errorf("second part of content_object_id has to be int, got %q", parts[1])
			}

			canSee := false
			switch parts[0] {
			case "motion":
				var todo []int
				todo, err = Motion{}.see(ctx, ds, mperms, id)
				canSee = len(todo) > 0
			case "motion_block":
				canSee, err = MotionBlock{}.see(ctx, ds, mperms, id)
			case "assignment":
				var todo []int
				todo, err = Assignment{}.see(ctx, ds, mperms, id)
				canSee = len(todo) > 0
			case "topic":
				var todo []int
				todo, err = Topic{}.see(ctx, ds, mperms, id)
				canSee = len(todo) > 0
			case "mediafile":
				canSee, err = Mediafile{}.see(ctx, ds, mperms, id)
			default:
				// TODO LAST ERROR
				return nil, fmt.Errorf("unknown content_object collection %q", parts[0])
			}
			if err != nil {
				return nil, fmt.Errorf("checking can see of %s: %w", parts[0], err)
			}

			if canSee {
				allowed = append(allowed, id)
			}
		}
		return allowed, nil
	})

}
