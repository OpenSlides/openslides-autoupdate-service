package collection

import (
	"context"
	"fmt"
	"slices"

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
// are restricted for everybody. For other polls, its the same as Group B. A
// user can see this fields, if he can vote for the represented_meeting_user.
type Ballot struct{}

// Name returns the collection name.
func (b Ballot) Name() string {
	return "ballot"
}

// MeetingID returns the meetingID for the object.
func (b Ballot) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
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
func (b Ballot) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return b.see
	case "B":
		return b.modeB

	case "C":
		return b.modeC
	}
	return nil
}

func (b Ballot) see(ctx context.Context, ds *dsfetch.Fetch, voteIDs ...int) ([]int, error) {
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

		if v, ok := representedMeetingUser.Value(); ok {
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

func (b Ballot) modeB(ctx context.Context, ds *dsfetch.Fetch, ballotIDs ...int) ([]int, error) {
	return eachRelationField(ctx, ds.Ballot_PollID, ballotIDs, func(pollID int, ballotIDs []int) ([]int, error) {
		seePollModeB, err := Collection(ctx, Poll{}.Name()).Modes("B")(ctx, ds, pollID)
		if err != nil {
			return nil, fmt.Errorf("checking poll mode B: %w", err)
		}

		if len(seePollModeB) > 0 {
			return ballotIDs, nil
		}
		return nil, nil
	})
}

func (b Ballot) modeC(ctx context.Context, ds *dsfetch.Fetch, ballotIDs ...int) ([]int, error) {
	requestUser, err := perm.RequestUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting request user: %w", err)
	}
	fmt.Println(requestUser)

	requestMeetingUserIDs, err := ds.User_MeetingUserIDs(requestUser).Value(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetch meeting_users_ids for user %d: %w", requestUser, err)
	}

	delegatedFromLists := make([][]int, len(requestMeetingUserIDs))
	for i, meetingUserID := range requestMeetingUserIDs {
		ds.MeetingUser_VoteDelegationsFromIDs(meetingUserID).Lazy(&delegatedFromLists[i])
	}

	representedMeetingUserIDs := make([]dsfetch.Maybe[int], len(ballotIDs))
	for i, ballotID := range ballotIDs {
		ds.Ballot_RepresentedMeetingUserID(ballotID).Lazy(&representedMeetingUserIDs[i])
	}

	if err := ds.Execute(ctx); err != nil {
		return nil, fmt.Errorf("feting represented_meeting_user_id for each ballot and all votedelegationsfrom for request user: %w", err)
	}

	delegatedFrom := mergeLists(delegatedFromLists)

	var allowedByRepresented []int
	for i, representedUserID := range representedMeetingUserIDs {
		id, exist := representedUserID.Value()
		if !exist {
			continue
		}
		if slices.Contains(requestMeetingUserIDs, id) || slices.Contains(delegatedFrom, id) {
			allowedByRepresented = append(allowedByRepresented, ballotIDs[i])
		}
	}

	allowedByPublished, err := eachRelationField(ctx, ds.Ballot_PollID, ballotIDs, func(pollID int, ballotDs []int) ([]int, error) {
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
			return ballotDs, nil
		}
		return nil, nil
	})

	if err != nil {
		return nil, fmt.Errorf("check poll permission: %w", err)
	}

	return mergeUnique(allowedByRepresented, allowedByPublished), nil

}
