package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// ListOfSpeakers handels the restriction for the list_of_speakers collection.
//
// The user can see a list of speakers if the user has list_of_speakers.can_see
// or can_be_speaker.
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

func (los ListOfSpeakers) see(ctx context.Context, fetcher *dsfetch.Fetch, losIDs []int) ([]attribute.Func, error) {
	return byMeeting(ctx, fetcher, los, losIDs, func(meetingID int, ids []int) ([]attribute.Func, error) {
		groupMap, err := perm.GroupMapFromContext(ctx, fetcher, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting group map: %w", err)
		}

		attrFn := attribute.FuncOr(
			attribute.FuncInGroup(groupMap[perm.ListOfSpeakersCanSee]),
			attribute.FuncInGroup(groupMap[perm.ListOfSpeakersCanBeSpeaker]),
		)

		return attributeFuncList(ids, attrFn), nil
	})
}
