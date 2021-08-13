package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
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

func (t Topic) see(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, topicID int) (bool, error) {
	meetingID := fetch.Field().Topic_ID(ctx, topicID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching meeting_id %d: %w", topicID, err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting perms for meeting %d: %w", meetingID, err)
	}

	return perms.Has(perm.AgendaItemCanSee), nil
}

func (t Topic) modeA(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, topicID int) (bool, error) {
	see, err := t.see(ctx, fetch, mperms, topicID)
	if err != nil {
		return false, fmt.Errorf("checking see: %w", err)
	}

	if see {
		return true, nil
	}

	losID := fetch.Field().Topic_ListOfSpeakersID(ctx, topicID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("getting list of speakers id: %w", err)
	}

	if losID != 0 {
		see, err := ListOfSpeakers{}.see(ctx, fetch, mperms, losID)
		if err != nil {
			return false, fmt.Errorf("checking list of speakers %d: %w", losID, err)
		}

		if see {
			return true, nil
		}
	}

	return false, nil
}
