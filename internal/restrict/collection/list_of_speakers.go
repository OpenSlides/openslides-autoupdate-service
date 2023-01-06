package collection

import (
	"context"
	"fmt"

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

func (los ListOfSpeakers) see(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, attrMap map[int]*Attributes, losIDs ...int) ([]int, error) {
	return eachMeeting(ctx, ds, los, losIDs, func(meetingID int, ids []int) ([]int, error) {
		perms, err := mperms.Meeting(ctx, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting perms for meetind %d: %w", meetingID, err)
		}

		if canSee := perms.Has(perm.ListOfSpeakersCanSee); !canSee {
			return nil, nil
		}

		return eachContentObjectCollection(ctx, ds.ListOfSpeakers_ContentObjectID, ids, func(collection string, id int, ids []int) ([]int, error) {
			// TODO: This should return not one contentobject, but all content objects with the same collection at once. So the first argument should be objectIDs
			var restricter FieldRestricter
			switch collection {
			case "motion":
				restricter = Motion{}.see

			case "motion_block":
				restricter = MotionBlock{}.see

			case "assignment":
				restricter = Assignment{}.see

			case "topic":
				restricter = Topic{}.see
			case "mediafile":
				restricter = Mediafile{}.see
			default:
				// TODO LAST ERROR
				return nil, fmt.Errorf("unknown content_object collection %q", collection)
			}

			canSee, err := restricter(ctx, ds, mperms, id)
			if err != nil {
				return nil, fmt.Errorf("checking can see of %s: %w", collection, err)
			}

			if len(canSee) == 1 {
				return ids, nil
			}
			return nil, nil
		})
	})

}
