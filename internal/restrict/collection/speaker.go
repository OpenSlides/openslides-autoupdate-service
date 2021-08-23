package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Speaker handels restrictions of the collection speaker.
type Speaker struct{}

// Modes returns the restrictions modes for the meeting collection.
func (s Speaker) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return s.see
	}
	return nil
}

func (s Speaker) see(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, speakerID int) (bool, error) {
	los := fetch.Field().Speaker_ListOfSpeakersID(ctx, speakerID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("fetch los: %w", err)
	}

	see, err := ListOfSpeakers{}.see(ctx, fetch, mperms, los)
	if err != nil {
		return false, fmt.Errorf("checking see of los %d: %w", los, err)
	}

	return see, nil
}
