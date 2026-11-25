package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

// PollEntitledUser handels the restriction for poll_entitled_user.
//
// The user can see poll_entitled_user, if he can see the poll.
//
// Mode A: The user can see the poll_entitled_user.
type PollEntitledUser struct{}

// Name returns the collection name.
func (a PollEntitledUser) Name() string {
	return "poll_entitled_user"
}

// MeetingID returns the meetingID for the object.
func (a PollEntitledUser) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	pollID, err := ds.PollEntitledUser_PollID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("fetching poll id: %w", err)
	}

	return Poll{}.MeetingID(ctx, ds, pollID)
}

// Modes returns the restrictions modes.
func (a PollEntitledUser) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return a.see
	}
	return nil
}

func (a PollEntitledUser) see(ctx context.Context, ds *dsfetch.Fetch, pollEntitledUserIDs ...int) ([]int, error) {
	return eachRelationField(ctx, ds.PollEntitledUser_PollID, pollEntitledUserIDs, func(pollID int, ids []int) ([]int, error) {
		canSeePoll, err := Collection(ctx, Poll{}.Name()).Modes("A")(ctx, ds, pollID)
		if err != nil {
			return nil, fmt.Errorf("can see poll: %w", err)
		}

		if len(canSeePoll) == 1 {
			return ids, nil
		}

		return nil, nil
	})
}
