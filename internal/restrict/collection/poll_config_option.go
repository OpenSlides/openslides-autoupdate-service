package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

// PollConfigOption handles permission its collection.
//
// A user can see it, if he can see the corresponding poll.
type PollConfigOption struct{}

// Name returns the collection name.
func (a PollConfigOption) Name() string {
	return "poll_config_option"
}

// MeetingID returns the meeting of the poll
func (a PollConfigOption) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	pollID, err := ds.PollOption_PollID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting poll id: %w", err)
	}

	meetingID, err := ds.Poll_MeetingID(pollID).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meeting id: %w", err)
	}
	return meetingID, true, nil
}

// Modes returns the restrictions modes for the action_worker collection.
func (a PollConfigOption) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return a.see
	}
	return nil
}

func (a PollConfigOption) see(ctx context.Context, ds *dsfetch.Fetch, pollConfigApprovalIDs ...int) ([]int, error) {
	return eachCondition(pollConfigApprovalIDs, func(pollConfigApprovalID int) (bool, error) {
		pollID, err := ds.PollOption_PollID(pollConfigApprovalID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting poll id: %w", err)
		}

		allowed, err := Collection(ctx, Poll{}.Name()).Modes("A")(ctx, ds, pollID)
		if err != nil {
			return false, fmt.Errorf("check permission of poll: %w", err)
		}

		if len(allowed) > 0 {
			return true, nil
		}

		return false, nil
	})
}
