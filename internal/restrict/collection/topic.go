package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Topic handels the restrictions for the topic collection.
type Topic struct{}

// Modes returns the field restriction for each mode.
func (t Topic) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return t.modeA
	case "B":
		return t.see
	}
	return nil
}

func (t Topic) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, topicID int) (bool, error) {
	meetingID, err := ds.Topic_MeetingID(topicID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("fetching meeting_id %d: %w", topicID, err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting perms for meeting %d: %w", meetingID, err)
	}

	return perms.Has(perm.AgendaItemCanSee), nil
}

func (t Topic) modeA(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, topicID int) (bool, error) {
	see, err := t.see(ctx, ds, mperms, topicID)
	if err != nil {
		return false, fmt.Errorf("checking see: %w", err)
	}

	if see {
		return true, nil
	}

	losID, err := ds.Topic_ListOfSpeakersID(topicID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting list of speakers id: %w", err)
	}

	if losID != 0 {
		see, err := ListOfSpeakers{}.see(ctx, ds, mperms, losID)
		if err != nil {
			return false, fmt.Errorf("checking list of speakers %d: %w", losID, err)
		}

		if see {
			return true, nil
		}
	}

	return false, nil
}
