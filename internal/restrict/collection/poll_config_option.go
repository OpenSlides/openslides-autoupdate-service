package collection

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

// PollConfigOption handles permission its collection.
//
// A user can see it, if he can see the corresponding poll.
type PollConfigOption struct{}

// Name returns the collection name.
func (a PollConfigOption) Name() string {
	return "poll_config_approval"
}

// MeetingID returns the meeting of the poll
func (a PollConfigOption) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	pollID, err := a.pollID(ctx, ds, id)
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
		pollID, err := a.pollID(ctx, ds, pollConfigApprovalID)
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

func (a PollConfigOption) pollID(ctx context.Context, ds *dsfetch.Fetch, pollConfigOptionID int) (int, error) {
	config, err := ds.PollConfigOption_PollConfigID(pollConfigOptionID).Value(ctx)
	if err != nil {
		return 0, fmt.Errorf("getting poll config: %w", err)
	}

	collection, configIDStr, found := strings.Cut(config, "/")
	if !found {
		return 0, fmt.Errorf("invalid config %s", config)
	}

	configID, err := strconv.Atoi(configIDStr)
	if err != nil {
		return 0, fmt.Errorf("invalid config id %s", config)
	}

	var field func(int) *dsfetch.ValueInt
	switch collection {
	case "poll_config_approval":
		field = ds.PollConfigApproval_PollID
	case "poll_config_selection":
		field = ds.PollConfigSelection_PollID
	case "poll_config_rating_score":
		field = ds.PollConfigRatingScore_PollID
	case "poll_config_rating_approval":
		field = ds.PollConfigRatingApproval_PollID
	default:
		return 0, fmt.Errorf("unknown poll config collection %s", collection)
	}

	pollID, err := field(configID).Value(ctx)
	if err != nil {
		return 0, fmt.Errorf("getting poll id: %w", err)
	}

	return pollID, nil
}
