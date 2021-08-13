package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Vote handels restrictions of the collection vote.
type Vote struct{}

// Modes returns the restrictions modes for the meeting collection.
func (v Vote) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return v.see
	case "B":
		return v.modeB
	}
	return nil
}

func (v Vote) see(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, voteID int) (bool, error) {
	optionID := fetch.Field().Vote_OptionID(ctx, voteID)
	pollID := fetch.Field().Option_PollID(ctx, optionID)
	state := fetch.Field().Poll_State(ctx, pollID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("getting poll id and state: %w", err)
	}

	if state == "published" {
		return true, nil
	}

	manage, err := Poll{}.manage(ctx, fetch, mperms, pollID)
	if err != nil {
		return false, fmt.Errorf("checking manage poll %d: %w", pollID, err)
	}

	if manage {
		return true, nil
	}

	voteUser := fetch.Field().Vote_UserID(ctx, voteID)
	if err != nil {
		return false, fmt.Errorf("getting vote user: %w", err)
	}

	if voteUser == mperms.UserID() {
		return true, nil
	}

	delegatedUser := fetch.Field().Vote_DelegatedUserID(ctx, voteID)
	if err != nil {
		return false, fmt.Errorf("getting delegated user: %w", err)
	}

	if delegatedUser == mperms.UserID() {
		return true, nil
	}

	return false, nil
}

func (v Vote) modeB(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, voteID int) (bool, error) {
	optionID := fetch.Field().Vote_OptionID(ctx, voteID)
	pollID := fetch.Field().Option_PollID(ctx, optionID)
	state := fetch.Field().Poll_State(ctx, pollID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("getting poll id and state: %w", err)
	}

	switch state {
	case "published":
		see, err := v.see(ctx, fetch, mperms, voteID)
		if err != nil {
			return false, fmt.Errorf("checking see vote: %w", err)
		}

		return see, nil

	case "finished":
		manage, err := Poll{}.manage(ctx, fetch, mperms, pollID)
		if err != nil {
			return false, fmt.Errorf("checking manage poll %d: %w", pollID, err)
		}

		return manage, nil

	default:
		return false, nil

	}
}
