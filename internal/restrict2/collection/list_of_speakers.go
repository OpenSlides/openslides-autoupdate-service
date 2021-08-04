package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// ListOfSpeakers handels the restriction for the list_of_speakers collection.
type ListOfSpeakers struct{}

// See defiends the see property.
func (los ListOfSpeakers) See(ctx context.Context, fetch *datastore.Fetcher, mperms perm.MeetingPermission, losID int) (bool, error) {
	mid, err := los.meetingID(ctx, fetch, losID)
	if err != nil {
		return false, fmt.Errorf("fetching meeting id for los %d: %w", losID, err)
	}

	perms, err := mperms.Meeting(ctx, mid)
	if err != nil {
		return false, fmt.Errorf("getting perms for meetind %d: %w", mid, err)
	}

	return perms.Has(perm.ListOfSpeakersCanSee), nil
}

// Modes returns a map from all known modes to there restricter.
func (los ListOfSpeakers) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return allways
	}
	return nil
}

func (los ListOfSpeakers) meetingID(ctx context.Context, fetch *datastore.Fetcher, id int) (int, error) {
	mid := datastore.Int(ctx, fetch.FetchIfExist, "list_of_speakers/%d/meeting_id", id)
	if err := fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching meeting_id for the list of speakers %d: %w", id, err)
	}
	return mid, nil
}
