package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
	"github.com/OpenSlides/openslides-go/perm"
)

// Ballot handels restrictions of the collection ballot.
//
// The user can see a ballot if
//
//	he can see the related poll and any of:
//
//		The associated poll is published.
//		The user has poll.can_see_progress.
//		The request user has a meeting_user in ballot/acting_meeting_user_id or ballot/represented_meeting_user_id.
//
// Group A: The user is allowed to know, that the poll exists. This is necessary
// to know, how many ballots where already submitted. A user can see this
// fields, if he can see the ballot.
//
// Group B: Contains fields, that show the value of the ballot. A user is
// allowed to see it, if he is also allowed to see the poll result. For example
// for managers, if the poll is published or for live voting. A user can see the
// mode, if he can see poll restriction mode B.
//
// Group C: Contains the user ids of the ballot. For secret polls, this fields
// are restricted for everybody. For other polls, its the same as Group B.
type Ballot struct{}

// Name returns the collection name.
func (v Ballot) Name() string {
	return "ballot"
}

// MeetingID returns the meetingID for the object.
func (v Ballot) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	pollID, err := ds.Ballot_PollID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("get poll id: %w", err)
	}

	meetingID, err := ds.Poll_MeetingID(pollID).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("get meeting id: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (v Ballot) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return v.see
	case "B":
		return v.modeB

	case "C":
		return v.modeC
	}
	return nil
}

func (v Ballot) see(ctx context.Context, ds *dsfetch.Fetch, voteIDs ...int) ([]int, error) {
	requestUser, err := perm.RequestUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting request user: %w", err)
	}

	return eachCondition(voteIDs, func(voteID int) (bool, error) {
		pollID, err := ds.Ballot_PollID(voteID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting poll id: %w", err)
		}

		meetingID, err := ds.Poll_MeetingID(pollID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting meeting id: %w", err)
		}

		published, err := ds.Poll_Published(pollID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting poll publihed: %w", err)
		}

		actingMeetingUser, err := ds.Ballot_ActingMeetingUserID(voteID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting vote user: %w", err)
		}

		representedMeetingUser, err := ds.Ballot_RepresentedMeetingUserID(voteID).Value(ctx)
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

		if v, ok := actingMeetingUser.Value(); ok {
			actingUser, err := ds.MeetingUser_UserID(v).Value(ctx)
			if err != nil {
				return false, fmt.Errorf("getting acting user: %w", err)
			}
			if actingUser == requestUser {
				return true, nil
			}
		}

		if v, ok := representedMeetingUser.Value(); ok && v == requestUser {
			representedUser, err := ds.MeetingUser_UserID(v).Value(ctx)
			if err != nil {
				return false, fmt.Errorf("getting acting user: %w", err)
			}
			if representedUser == requestUser {
				return true, nil
			}
		}

		return false, nil
	})
}

func (v Ballot) modeB(ctx context.Context, ds *dsfetch.Fetch, voteIDs ...int) ([]int, error) {
	return eachRelationField(ctx, ds.Ballot_PollID, voteIDs, func(pollID int, voteIDs []int) ([]int, error) {
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

// ModeC is a workarount until crypto vote is released. With crypto vote, the
// fields can be moved into modeB and modeC be removed.
func (v Ballot) modeC(ctx context.Context, ds *dsfetch.Fetch, voteIDs ...int) ([]int, error) {
	return eachRelationField(ctx, ds.Ballot_PollID, voteIDs, func(pollID int, voteIDs []int) ([]int, error) {
		visibility, err := ds.Poll_Visibility(pollID).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("fetch visibility of poll %d: %w", pollID, err)
		}
		if visibility == "secret" {
			return nil, nil
		}

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
