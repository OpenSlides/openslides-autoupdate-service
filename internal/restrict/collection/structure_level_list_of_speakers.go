package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// StructureLevelListOfSpeakers handels restrictions of the collection structure_level_list_of_speakers.
//
// The user can see a structure level if he can see the linked list of speaker.
//
// Mode A: The user can see the speaker.
type StructureLevelListOfSpeakers struct{}

// Name returns the collection name.
func (s StructureLevelListOfSpeakers) Name() string {
	return "structure_level_list_of_speakers"
}

// MeetingID returns the meetingID for the object.
func (s StructureLevelListOfSpeakers) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.StructureLevelListOfSpeakers_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("get meeting id: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (s StructureLevelListOfSpeakers) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return s.see
	}
	return nil
}

func (s StructureLevelListOfSpeakers) see(ctx context.Context, ds *dsfetch.Fetch, structureLevelListOfSpeakersIDs ...int) ([]int, error) {
	listOfSpeakerIDs := make([]int, len(structureLevelListOfSpeakersIDs))
	for i, structureLevelID := range structureLevelListOfSpeakersIDs {
		ds.StructureLevelListOfSpeakers_ListOfSpeakersID(structureLevelID).Lazy(&listOfSpeakerIDs[i])
	}

	if err := ds.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching list of speakers: %w", err)
	}

	allowedListOfSpeakers, err := Collection(ctx, ListOfSpeakers{}.Name()).Modes("A")(ctx, ds, listOfSpeakerIDs...)
	if err != nil {
		return nil, fmt.Errorf("check list of speakers: %w", err)
	}

	if len(allowedListOfSpeakers) == len(structureLevelListOfSpeakersIDs) {
		return structureLevelListOfSpeakersIDs, nil
	}

	allowed := make([]int, 0, len(structureLevelListOfSpeakersIDs))
	for i, listOfSpeakerID := range listOfSpeakerIDs {
		for _, allowedID := range allowedListOfSpeakers {
			if listOfSpeakerID == allowedID {
				allowed = append(allowed, structureLevelListOfSpeakersIDs[i])
				break
			}
		}
	}

	return allowed, nil
}
