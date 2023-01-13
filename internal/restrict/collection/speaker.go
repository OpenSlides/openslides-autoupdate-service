package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// Speaker handels restrictions of the collection speaker.
//
// The user can see a speaker if the user can see the linked list of speakers.
//
// Mode A: The user can see the speaker.
type Speaker struct{}

// Name returns the collection name.
func (s Speaker) Name() string {
	return "speaker"
}

// MeetingID returns the meetingID for the object.
func (s Speaker) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.Speaker_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("get meeting id: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (s Speaker) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return s.see
	}
	return nil
}

func (s Speaker) see(ctx context.Context, ds *dsfetch.Fetch, speakerIDs ...int) ([]int, error) {
	return eachRelationField(ctx, ds.Speaker_ListOfSpeakersID, speakerIDs, func(losID int, ids []int) ([]int, error) {
		see, err := ListOfSpeakers{}.see(ctx, ds, losID)
		if err != nil {
			return nil, fmt.Errorf("checking see of los %d: %w", losID, err)
		}

		if len(see) == 1 {
			return ids, nil
		}
		return nil, nil
	})
}
