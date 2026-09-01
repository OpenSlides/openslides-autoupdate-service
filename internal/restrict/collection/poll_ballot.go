package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

// PollBallot handels restrictions of the collection poll_ballot.
//
// The user can see a ballot if he can see the associated poll and the poll is
// published.
//
// Group A: The user can see the ballot.
type PollBallot struct{}

// Name returns the collection name.
func (b PollBallot) Name() string {
	return "poll_ballot"
}

// MeetingID returns the meetingID for the object.
func (b PollBallot) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	pollID, err := ds.PollBallot_PollID(id).Value(ctx)
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
func (b PollBallot) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return b.see
	}
	return nil
}

func (b PollBallot) see(ctx context.Context, ds *dsfetch.Fetch, ballotIDs ...int) ([]int, error) {
	return eachRelationField(ctx, ds.PollBallot_PollID, ballotIDs, func(pollID int, ballotIDs []int) ([]int, error) {
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
			return ballotIDs, nil
		}

		return nil, nil
	})
}
