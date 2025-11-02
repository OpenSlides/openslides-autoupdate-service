package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

// PollConfigRatingApproval handles permission its collection.
//
// A user can see it, if he can see the corresponding poll.
type PollConfigRatingApproval struct{}

// Name returns the collection name.
func (a PollConfigRatingApproval) Name() string {
	return "poll_config_rating_approval"
}

// MeetingID returns the meeting of the poll
func (a PollConfigRatingApproval) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	pollID, err := ds.PollConfigRatingApproval_PollID(id).Value(ctx)
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
func (a PollConfigRatingApproval) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return a.see
	}
	return nil
}

func (a PollConfigRatingApproval) see(ctx context.Context, ds *dsfetch.Fetch, pollConfigRatingApprovalDs ...int) ([]int, error) {
	return eachRelationField(ctx, ds.PollConfigRatingApproval_PollID, pollConfigRatingApprovalDs, func(pollID int, pollConfigRatingApprovalDs []int) ([]int, error) {
		allowed, err := Collection(ctx, Poll{}.Name()).Modes("A")(ctx, ds, pollID)
		if err != nil {
			return nil, fmt.Errorf("check permission of poll: %w", err)
		}

		if len(allowed) > 0 {
			return pollConfigRatingApprovalDs, nil
		}

		return nil, nil
	})
}
