package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// Option handels restrictions of the collection option.
//
// The user can see an option if the user can see the linked poll.
//
// Mode A: The user can see the option.
//
// Mode B: The user can manage the poll OR (The user can see the poll AND poll/state is published).
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
func (o Option) see(ctx context.Context, fetcher *dsfetch.Fetch, optionIDs []int) ([]attribute.Func, error) {
	pollIDs, err := fetchPollIDs(ctx, fetcher, optionIDs)
	if err != nil {
		return nil, fmt.Errorf("getting related poll ids: %w", err)
	}

	return Collection(ctx, Poll{}).Modes("A")(ctx, fetcher, pollIDs)
}

// TODO: Group by poll
func (o Option) modeB(ctx context.Context, fetcher *dsfetch.Fetch, optionIDs []int) ([]attribute.Func, error) {
	pollIDList := make([]int, len(optionIDs))
	for i, id := range optionIDs {
		fetcher.Option_PollID(id).Lazy(&pollIDList[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching option data: %w", err)
	}

	// pollIDList can contain the same pollID multiple times. I think it is
	// still better not to group them. It would require more logic and an index.
	// But we have a collection cache, so the gain should be minimal.
	seePollAttr, err := Collection(ctx, Poll{}).Modes("A")(ctx, fetcher, pollIDList)
	if err != nil {
		return nil, fmt.Errorf("checking poll see: %w", err)
	}

	managePollAttr, err := Collection(ctx, Poll{}).Modes("MANAGE")(ctx, fetcher, pollIDList)
	if err != nil {
		return nil, fmt.Errorf("checking poll manage: %w", err)
	}

	pollState := make([]string, len(optionIDs))
	for i, pollID := range pollIDList {
		fetcher.Poll_State(pollID).Lazy(&pollState[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching poll data: %w", err)
	}

	attrFuncs := make([]attribute.Func, len(optionIDs))
	for i := range optionIDs {
		attrFuncs[i] = managePollAttr[i]
		switch pollState[i] {
		case "published":
			attrFuncs[i] = seePollAttr[i]
		default:
			attrFuncs[i] = managePollAttr[i]
		}
	}
	return attrFuncs, nil
}

// fetchPollIDs returns the poll ids for a ist of optionIDs.
func fetchPollIDs(ctx context.Context, fetcher *dsfetch.Fetch, optionIDs []int) ([]int, error) {
	optionPollIDs := make([]int, len(optionIDs))
	usedAsGlobal := make([]int, len(optionIDs))
	for i, id := range optionIDs {
		fetcher.Option_PollID(id).Lazy(&optionPollIDs[i])
		fetcher.Option_UsedAsGlobalOptionInPollID(id).Lazy(&usedAsGlobal[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching option data: %w", err)
	}

	pollIDs := make([]int, len(optionIDs))
	for i := range optionIDs {
		if optionPollIDs[i] != 0 {
			pollIDs[i] = optionPollIDs[i]
			continue
		}

		if usedAsGlobal[i] != 0 {
			pollIDs[i] = usedAsGlobal[i]
			continue
		}

		return nil, fmt.Errorf(
			"database seems corrupted. Both fields option/%d/poll_id and option/%d/used_as_global_option_in_poll_id are empty. One of the fields is required",
			optionIDs[i],
			optionIDs[i],
		)
	}
	return pollIDs, nil
}
