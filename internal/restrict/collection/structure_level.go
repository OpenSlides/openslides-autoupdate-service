package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// StructureLevel handels restrictions of the collection structure_level.
//
// The user can see a structure level if he has `list_of_speakers.can_see` OR `user.can_see`
//
// Mode A: The user can see the speaker.
type StructureLevel struct{}

// Name returns the collection name.
func (s StructureLevel) Name() string {
	return "structure_level"
}

// MeetingID returns the meetingID for the object.
func (s StructureLevel) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.StructureLevel_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("get meeting id: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (s StructureLevel) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return s.see
	case "B":
		return never // TODO: Remove me after the fix in the backend
	}
	return nil
}

func (s StructureLevel) see(ctx context.Context, ds *dsfetch.Fetch, structureLevelIDs ...int) ([]int, error) {
	return meetingPerm(ctx, ds, s, structureLevelIDs, perm.ListOfSpeakersCanSee, perm.UserCanSee)
}
