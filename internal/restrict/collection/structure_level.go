package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
	"github.com/OpenSlides/openslides-go/perm"
)

// StructureLevel handels restrictions of the collection structure_level.
//
// The user can see a structure level if he has `list_of_speakers.can_see` OR
// can see any of the  linked `meeting_user_ids`.
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
	}
	return nil
}

func (s StructureLevel) see(ctx context.Context, ds *dsfetch.Fetch, structureLevelIDs ...int) ([]int, error) {
	return eachMeeting(ctx, ds, s, structureLevelIDs, func(meetingID int, structureLevelIDs []int) ([]int, error) {
		perms, err := perm.FromContext(ctx, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting permission: %w", err)
		}

		if perms.Has(perm.ListOfSpeakersCanSee) {
			return structureLevelIDs, nil
		}

		return eachCondition(structureLevelIDs, func(id int) (bool, error) {
			meetingUserIDs, err := ds.StructureLevel_MeetingUserIDs(id).Value(ctx)
			if err != nil {
				return false, fmt.Errorf("getting meeting_user_ids: %w", err)
			}

			canSee, err := Collection(ctx, "meeting_user").Modes("A")(ctx, ds, meetingUserIDs...)
			if err != nil {
				return false, fmt.Errorf("checking meeting_user: %w", err)
			}

			return len(canSee) > 0, nil
		})
	})
}
