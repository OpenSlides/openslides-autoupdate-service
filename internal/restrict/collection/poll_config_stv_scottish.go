package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

// PollConfigStvScottish handles permission its collection.
//
// A user can see it, if he can see the corresponding poll.
type PollConfigStvScottish struct{}

// Name returns the collection name.
func (a PollConfigStvScottish) Name() string {
	return "poll_config_stv_scottish"
}

// MeetingID returns the meeting of the poll
func (a PollConfigStvScottish) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	pollID, err := ds.PollConfigStvScottish_PollID(id).Value(ctx)
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
func (a PollConfigStvScottish) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return a.see
	}
	return nil
}

func (a PollConfigStvScottish) see(ctx context.Context, ds *dsfetch.Fetch, pollConfigStvScottishIDs ...int) ([]int, error) {
	return eachRelationField(ctx, ds.PollConfigStvScottish_PollID, pollConfigStvScottishIDs, func(pollID int, pollConfigStvScottishIDs []int) ([]int, error) {
		allowed, err := Collection(ctx, Poll{}.Name()).Modes("A")(ctx, ds, pollID)
		if err != nil {
			return nil, fmt.Errorf("check permission of poll: %w", err)
		}

		if len(allowed) > 0 {
			return pollConfigStvScottishIDs, nil
		}

		return nil, nil
	})
}
