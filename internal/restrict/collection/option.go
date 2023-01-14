package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// Option handels restrictions of the collection option.
//
// The user can see an option if the user can see the linked poll.
//
// Mode A: The user can see the option.
//
// Mode B: The user can see the poll and (manage the linked poll or poll/state is published).
type Option struct{}

// Name returns the collection name.
func (o Option) Name() string {
	return "option"
}

// MeetingID returns the meetingID for the object.
func (o Option) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.Option_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meetingID: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (o Option) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return o.see
	case "B":
		return o.modeB
	}
	return nil
}

// TODO: Group by poll
func (o Option) see(ctx context.Context, ds *dsfetch.Fetch, optionIDs ...int) ([]int, error) {
	return eachCondition(optionIDs, func(optionID int) (bool, error) {
		pollID, err := pollID(ctx, ds, optionID)
		if err != nil {
			return false, fmt.Errorf("getting poll id: %w", err)
		}

		see, err := Collection(ctx, Poll{}.Name()).Modes("A")(ctx, ds, pollID)
		if err != nil {
			return false, fmt.Errorf("checking see poll %d: %w", pollID, err)
		}

		return len(see) == 1, nil
	})
}

// TODO: Group by poll
func (o Option) modeB(ctx context.Context, ds *dsfetch.Fetch, optionIDs ...int) ([]int, error) {
	return eachCondition(optionIDs, func(optionID int) (bool, error) {
		pollID, err := pollID(ctx, ds, optionID)
		if err != nil {
			return false, fmt.Errorf("getting poll id: %w", err)
		}

		see, err := Collection(ctx, Poll{}.Name()).Modes("A")(ctx, ds, pollID)
		if err != nil {
			return false, fmt.Errorf("checking see poll %d: %w", pollID, err)
		}

		if len(see) == 0 {
			return false, nil
		}

		canManage, err := Poll{}.manage(ctx, ds, pollID)
		if err != nil {
			return false, fmt.Errorf("checking see poll %d: %w", pollID, err)
		}

		if len(canManage) == 1 {
			return true, nil
		}

		pollState, err := ds.Poll_State(pollID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting poll state: %w", err)
		}

		return pollState == "published", nil
	})
}

func pollID(ctx context.Context, ds *dsfetch.Fetch, optionID int) (int, error) {
	pollID, exist, err := ds.Option_PollID(optionID).Value(ctx)
	if err != nil {
		return 0, fmt.Errorf("getting poll id from field poll_id: %w", err)
	}

	if exist {
		return pollID, nil
	}

	pollID, exist, err = ds.Option_UsedAsGlobalOptionInPollID(optionID).Value(ctx)
	if err != nil {
		return 0, fmt.Errorf("getting used as global option id in poll id: %w", err)
	}

	if exist {
		return pollID, nil
	}

	// TODO LAST ERROR
	return 0, fmt.Errorf(
		"database seems corrupted. Both fields option/%d/poll_id and option/%d/used_as_global_option_in_poll_id are empty. One of the fields is required",
		optionID,
		optionID,
	)
}
