package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

// ListOfSpeakers handels the restriction for the list_of_speakers collection.
//
// The user can see a list of speakers if the user has list_of_speakers.can_see
// in the meeting and can see the content_object.
//
// Mode A: The user can see the list of speakers.
type ListOfSpeakers struct{}

// Name returns the collection name.
func (los ListOfSpeakers) Name() string {
	return "list_of_speakers"
}

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

func (los ListOfSpeakers) see(ctx context.Context, ds *dsfetch.Fetch, mperms perm.MeetingPermission, attrMap AttributeMap, losIDs ...int) error {
	return eachMeeting(ctx, ds, los, losIDs, func(meetingID int, ids []int) error {
		groupMap, err := mperms.Meeting(ctx, ds, meetingID)
		if err != nil {
			return fmt.Errorf("getting perms for meetind %d: %w", meetingID, err)
		}

		groups := groupMap[perm.ListOfSpeakersCanSee]

		//TODO: make sure to be called after each restriction mode.

		return eachContentObjectCollection(ctx, ds.ListOfSpeakers_ContentObjectID, ids, func(collection string, id int, ids []int) error {
			// TODO: This should return not one contentobject, but all content objects with the same collection at once. So the first argument should be objectIDs
			var seeMode string
			switch collection {
			case "motion":
				seeMode = "C"

			case "motion_block":
				seeMode = "A"

			case "assignment":
				seeMode = "A"

			case "topic":
				seeMode = "A"

			case "mediafile":
				seeMode = "A"

			default:
				// TODO LAST ERROR
				return fmt.Errorf("unknown content_object collection %q", collection)
			}

			andAttr, err := attrMap.Get(ctx, ds, mperms, dskey.Key{Collection: collection, ID: id, Field: seeMode})
			if err != nil {
				return fmt.Errorf("andAttr: %w", err)
			}

			attr := &Attributes{
				GlobalPermission: byte(perm.OMLSuperadmin),
				GroupIDs:         groups,
				GroupAnd:         andAttr,
			}

			for _, losID := range ids {
				attrMap.Add(dskey.Key{Collection: los.Name(), ID: losID, Field: "A"}, attr)
			}

			return nil
		})
	})

}
