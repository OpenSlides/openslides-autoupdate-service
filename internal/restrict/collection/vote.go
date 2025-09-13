package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
	"github.com/OpenSlides/openslides-go/perm"
)

// Vote handels restrictions of the collection vote.
//
// The user can see a vote if
//
//	he can see the related poll and any of:
//
//		The associated poll is published.
//		The user has poll.can_see_progress.
//		The user's id is equal to vote/acting_user_id or vote/represented_user_id.
//
// Group A: The user is allowed to know, that the poll exists. This is necessary
// to know, how many votes where already submitted. A user can see this fields,
// if he can see the vote.
//
// Group B: Contains fields, that show the value of the vote. A user is allowed
// to see it, if he is also allowed to see the poll result. For example for
// managers, if the poll is published or for live voting. A user can see the
// mode, if he can see poll restriction mode b.
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

func (v Vote) see(ctx context.Context, ds *dsfetch.Fetch, voteIDs ...int) ([]int, error) {
	requestUser, err := perm.RequestUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting request user: %w", err)
	}

	return eachCondition(voteIDs, func(voteID int) (bool, error) {
		pollID, err := ds.Vote_PollID(voteID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting poll id: %w", err)
		}

		meetingID, err := ds.Vote_MeetingID(voteID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting meeting id: %w", err)
		}

		published, err := ds.Poll_Published(pollID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting poll publihed: %w", err)
		}

		voteUser, err := ds.Vote_ActingUserID(voteID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting vote user: %w", err)
		}

		delegatedUser, err := ds.Vote_RepresentedUserID(voteID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting delegated user: %w", err)
		}

		// TODO: get all values in one request

		canSeePoll, err := Collection(ctx, Poll{}.Name()).Modes("A")(ctx, ds, pollID)
		if err != nil {
			return false, fmt.Errorf("check poll: %w", err)
		}

		if len(canSeePoll) == 0 {
			return false, nil
		}

		if published {
			return true, nil
		}

		perms, err := perm.FromContext(ctx, meetingID)
		if err != nil {
			return false, fmt.Errorf("getting perms for meeting %d: %w", meetingID, err)
		}

		if canSee := perms.Has(perm.PollCanSeeProgress); canSee {
			return true, nil
		}

		if v, ok := voteUser.Value(); ok && v == requestUser {
			return true, nil
		}

		if v, ok := delegatedUser.Value(); ok && v == requestUser {
			return true, nil
		}

		return false, nil
	})
}

func (v Vote) modeB(ctx context.Context, ds *dsfetch.Fetch, voteIDs ...int) ([]int, error) {
	return eachRelationField(ctx, ds.Vote_PollID, voteIDs, func(pollID int, voteIDs []int) ([]int, error) {
		seePollModeB, err := Collection(ctx, Poll{}.Name()).Modes("B")(ctx, ds, pollID)
		if err != nil {
			return nil, fmt.Errorf("checking poll mode B: %w", err)
		}

		if len(seePollModeB) > 0 {
			return voteIDs, nil
		}
		return nil, nil
	})
}
