package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

// PollConfigSelection handles permission its collection.
//
// A user can see it, if he can see the corresponding poll.
type PollConfigSelection struct{}

// Name returns the collection name.
func (a PollConfigSelection) Name() string {
	return "poll_config_selection"
}

// MeetingID returns the meeting of the poll
func (a PollConfigSelection) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	pollID, err := ds.PollConfigSelection_PollID(id).Value(ctx)
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
func (a PollConfigSelection) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return a.see
	}
	return nil
}

func (a PollConfigSelection) see(ctx context.Context, ds *dsfetch.Fetch, pollConfigSelectionIDs ...int) ([]int, error) {
	return eachRelationField(ctx, ds.PollConfigSelection_PollID, pollConfigSelectionIDs, func(pollID int, pollConfigSelectionIDs []int) ([]int, error) {
		allowed, err := Collection(ctx, Poll{}.Name()).Modes("A")(ctx, ds, pollID)
		if err != nil {
			return nil, fmt.Errorf("check permission of poll: %w", err)
		}

		if len(allowed) > 0 {
			return pollConfigSelectionIDs, nil
		}

		return nil, nil
	})
}
