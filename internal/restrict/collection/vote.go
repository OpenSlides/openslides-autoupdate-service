package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// Vote handels restrictions of the collection vote.
//
// The user can see a vote if
//
//	he can see the related poll and any of:
//
//		The associated poll/state is published.
//		The user can manage the associated poll.
//		The user's id is equal to vote/user_id.
//		The user's id is equal to vote/delegated_user_id.
//
// Group A: The user can see the vote.
//
// Group B: Depends on poll/state:
//
//	published: Accessible if the user can see the vote.
//	finished: Accessible if the user can manage the associated poll.
//	others: Not accessible for anyone.
type Vote struct{}

// Name returns the collection name.
func (v Vote) Name() string {
	return "vote"
}

// MeetingID returns the meetingID for the object.
func (v Vote) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.Vote_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("get meeting id: %w", err)
	}

	return meetingID, true, nil
}

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

// TODO: Group by poll or option
func (v Vote) see(ctx context.Context, ds *dsfetch.Fetch, voteIDs ...int) ([]int, error) {
	requestUser, err := perm.RequestUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting request user: %w", err)
	}

	return eachCondition(voteIDs, func(voteID int) (bool, error) {
		optionID, err := ds.Vote_OptionID(voteID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("fetching option_id: %w", err)
		}

		pollID, err := pollID(ctx, ds, optionID)
		if err != nil {
			return false, fmt.Errorf("getting poll id: %w", err)
		}

		canSeePoll, err := Collection(ctx, Poll{}.Name()).Modes("A")(ctx, ds, pollID)
		if err != nil {
			return false, fmt.Errorf("check poll: %w", err)
		}

		if len(canSeePoll) == 0 {
			return false, nil
		}

		state, err := ds.Poll_State(pollID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting poll id and state: %w", err)
		}

		if state == "published" {
			return true, nil
		}

		manage, err := Poll{}.manage(ctx, ds, pollID)
		if err != nil {
			return false, fmt.Errorf("checking manage poll %d: %w", pollID, err)
		}

		if len(manage) == 1 {
			return true, nil
		}

		voteUser, err := ds.Vote_UserID(voteID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting vote user: %w", err)
		}

		if v, ok := voteUser.Value(); ok && v == requestUser {
			return true, nil
		}

		delegatedUser, err := ds.Vote_DelegatedUserID(voteID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting delegated user: %w", err)
		}

		if v, ok := delegatedUser.Value(); ok && v == requestUser {
			return true, nil
		}

		return false, nil
	})
}

// TODO: Group by poll or option
func (v Vote) modeB(ctx context.Context, ds *dsfetch.Fetch, voteIDs ...int) ([]int, error) {
	return eachCondition(voteIDs, func(voteID int) (bool, error) {
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
			see, err := v.see(ctx, ds, voteID)
			if err != nil {
				return false, fmt.Errorf("checking see vote: %w", err)
			}

			return len(see) == 1, nil

		case "finished":
			manage, err := Poll{}.manage(ctx, ds, pollID)
			if err != nil {
				return false, fmt.Errorf("checking manage poll %d: %w", pollID, err)
			}

			return len(manage) == 1, nil

		default:
			return false, nil
		}
	})
}
