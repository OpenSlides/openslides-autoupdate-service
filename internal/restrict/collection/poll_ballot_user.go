package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
	"github.com/OpenSlides/openslides-go/perm"
)

// PollBallotUser handels restrictions of the collection poll_ballot_user.
//
// The user can see a ballot_user if
//
//	he can see the related poll and any of:
//
//		The associated poll is published.
//		The user has poll.can_see_progress.
//		The request user has the permission to vote for the meeting_user in represented_meeting_user_id.
//
// Group A: The user can see the the ballot_user.
type PollBallotUser struct{}

// Name returns the collection name.
func (b PollBallotUser) Name() string {
	return "poll_ballot_user"
}

// MeetingID returns the meetingID for the object.
func (b PollBallotUser) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	pollID, err := ds.PollBallotUser_PollID(id).Value(ctx)
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
func (b PollBallotUser) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return b.see
	}
	return nil
}

func (b PollBallotUser) see(ctx context.Context, ds *dsfetch.Fetch, ballotUserIDs ...int) ([]int, error) {
	requestUser, err := perm.RequestUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting request user: %w", err)
	}

	return eachRelationField(ctx, ds.PollBallotUser_PollID, ballotUserIDs, func(pollID int, ballotUserIDs []int) ([]int, error) {
		canSeePoll, err := Collection(ctx, Poll{}.Name()).Modes("A")(ctx, ds, pollID)
		if err != nil {
			return nil, fmt.Errorf("check poll: %w", err)
		}

		if len(canSeePoll) == 0 {
			return nil, nil
		}

		published, err := ds.Poll_Published(pollID).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting poll publihed: %w", err)
		}

		if published {
			return ballotUserIDs, nil
		}

		meetingID, err := ds.Poll_MeetingID(pollID).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting meeting id: %w", err)
		}

		perms, err := perm.FromContext(ctx, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting perms for meeting %d: %w", meetingID, err)
		}

		if canSee := perms.Has(perm.PollCanSeeProgress); canSee {
			return ballotUserIDs, nil
		}

		allowed := make([]int, 0, len(ballotUserIDs))
		for _, ballotUserID := range ballotUserIDs {
			representedMeetingUser, err := ds.PollBallotUser_RepresentedMeetingUserID(ballotUserID).Value(ctx)
			if err != nil {
				return nil, fmt.Errorf("getting represented user: %w", err)
			}

			representedUser, err := ds.MeetingUser_UserID(representedMeetingUser).Value(ctx)
			if err != nil {
				return nil, fmt.Errorf("getting represented user: %w", err)
			}

			if representedUser == requestUser {
				allowed = append(allowed, ballotUserID)
				continue
			}

			delegation, err := ds.MeetingUser_VoteDelegatedToID(representedMeetingUser).Value(ctx)
			if err != nil {
				return nil, fmt.Errorf("getting delegation from represented user: %w", err)
			}

			if v, set := delegation.Value(); set {
				delegatedUser, err := ds.MeetingUser_UserID(v).Value(ctx)
				if err != nil {
					return nil, fmt.Errorf("getting represented user: %w", err)
				}

				if delegatedUser == requestUser {
					allowed = append(allowed, ballotUserID)
					continue
				}
			}
		}

		return allowed, nil
	})
}
