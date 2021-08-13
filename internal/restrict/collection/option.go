package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Option handels restrictions of the collection option.
type Option struct{}

// Modes returns the restrictions modes for the meeting collection.
func (o Option) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return o.see
	case "B":
		return o.modeB
	}
	return nil
}

func (o Option) see(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, optionID int) (bool, error) {
	pollID := fetch.Field().Option_PollID(ctx, optionID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("getting poll id: %w", err)
	}

	see, err := Poll{}.see(ctx, fetch, mperms, pollID)
	if err != nil {
		return false, fmt.Errorf("checking see poll %d: %w", pollID, err)
	}

	return see, nil
}

func (o Option) modeB(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, optionID int) (bool, error) {
	pollID := fetch.Field().Option_PollID(ctx, optionID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("getting poll id: %w", err)
	}

	see, err := Poll{}.manage(ctx, fetch, mperms, pollID)
	if err != nil {
		return false, fmt.Errorf("checking see poll %d: %w", pollID, err)
	}

	return see, nil
}
