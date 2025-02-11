package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

// ListOfSpeakers handels the restriction for the list_of_speakers collection.
//
// The user can see a list of speakers if the user has list_of_speakers.can_see
// or can_be_speaker.
//
// Mode A: The user can see the list of speakers.
//
// Mode B: The user has list_of_speakers.can_see_moderator_notes
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
	case "B":
		return los.modeB
	}
	return nil
}

func (los ListOfSpeakers) see(ctx context.Context, ds *dsfetch.Fetch, losIDs ...int) ([]int, error) {
	return eachMeeting(ctx, ds, los, losIDs, func(meetingID int, ids []int) ([]int, error) {
		perms, err := perm.FromContext(ctx, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting perms for meeting %d: %w", meetingID, err)
		}

		canSee := perms.Has(perm.ListOfSpeakersCanSee) || perms.Has(perm.ListOfSpeakersCanBeSpeaker)

		if !canSee {
			return nil, nil
		}
		return losIDs, nil
	})
}

func (los ListOfSpeakers) modeB(ctx context.Context, ds *dsfetch.Fetch, losIDs ...int) ([]int, error) {
	return meetingPerm(ctx, ds, los, losIDs, perm.ListOfSpeakersCanSeeModeratorNotes)
}
