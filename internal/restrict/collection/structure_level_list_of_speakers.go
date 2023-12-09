package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// StructureLevelListOfSpeakers handels restrictions of the collection structure_level_list_of_speakers.
//
// The user can see a structure level if he has `list_of_speakers.can_see`
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

func (s StructureLevelListOfSpeakers) see(ctx context.Context, ds *dsfetch.Fetch, structureLevelIDs ...int) ([]int, error) {
	return meetingPerm(ctx, ds, s, structureLevelIDs, perm.ListOfSpeakersCanSee)
}
