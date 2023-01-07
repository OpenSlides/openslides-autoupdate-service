package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// Speaker handels restrictions of the collection speaker.
//
// The user can see a speaker if the user can see the linked list of speakers.
//
// Mode A: The user can see the speaker.
type Speaker struct {
	name string
}

// Name returns the collection name.
func (s Speaker) Name() string {
	return s.name
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

func (s Speaker) see(ctx context.Context, ds *dsfetch.Fetch, mperms perm.MeetingPermission, attrMap AttributeMap, speakerIDs ...int) error {
	return eachRelationField(ctx, ds.Speaker_ListOfSpeakersID, speakerIDs, func(losID int, ids []int) error {
		// TODO: This only works if los is calculated before speaker
		for _, id := range ids {
			attrMap.Add(s.name, id, "A", attrMap.Get("list_of_speakers", losID, "A"))
		}
		return nil
	})
}
