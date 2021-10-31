package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
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

func (v Vote) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, voteID int) (bool, error) {
	optionID, err := ds.Vote_OptionID(voteID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("fetching option_id: %w", err)
	}

	pollID, err := pollID(ctx, ds, optionID)
	if err != nil {
		return false, fmt.Errorf("getting poll id: %w", err)
	}

	state, err := ds.Poll_State(pollID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting poll id and state: %w", err)
	}

	if state == "published" {
		return true, nil
	}

	manage, err := Poll{}.manage(ctx, ds, mperms, pollID)
	if err != nil {
		return false, fmt.Errorf("checking manage poll %d: %w", pollID, err)
	}

	if manage {
		return true, nil
	}

	voteUser, exist, err := ds.Vote_UserID(voteID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting vote user: %w", err)
	}

	if exist && voteUser == mperms.UserID() {
		return true, nil
	}

	delegatedUser, exist, err := ds.Vote_DelegatedUserID(ctx, voteID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting delegated user: %w", err)
	}

	if exist && delegatedUser == mperms.UserID() {
		return true, nil
	}

	return false, nil
}

func (v Vote) modeB(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, voteID int) (bool, error) {
	optionID, err := ds.Vote_OptionID(voteID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("fetching option_id: %w", err)
	}

	pollID, err := pollID(ctx, ds, optionID)
	if err != nil {
		return false, fmt.Errorf("getting poll id: %w", err)
	}
	state, err := ds.Poll_State(pollID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting poll id and state: %w", err)
	}

	switch state {
	case "published":
		see, err := v.see(ctx, ds, mperms, voteID)
		if err != nil {
			return false, fmt.Errorf("checking see vote: %w", err)
		}

		return see, nil

	case "finished":
		manage, err := Poll{}.manage(ctx, ds, mperms, pollID)
		if err != nil {
			return false, fmt.Errorf("checking manage poll %d: %w", pollID, err)
		}

		return manage, nil

	default:
		return false, nil

	}
}
