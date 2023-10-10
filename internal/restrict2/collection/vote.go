package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// Vote handels restrictions of the collection vote.
// The user can see a vote if any of:
//
//	The associated poll/state is published.
//	The user can manage the associated poll.
//	The user's id is equal to vote/user_id.
//	The user's id is equal to vote/delegated_user_id.
//
// Group A: The user can see the vote.
//
// Group B: Depends on poll/state:
//
//	published: Accessible if the user can see the vote.
//	finished: Accessible if the user can manage the associated poll.
//	others: Not accessible for anyone.
type Vote struct{}

// Name returns the collection name.
func (v Vote) Name() string {
	return "vote"
}

// MeetingID returns the meetingID for the object.
func (v Vote) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.Vote_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("get meeting id: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (v Vote) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return v.see
	case "B":
		return v.modeB
	}
	return nil
}

func (v Vote) see(ctx context.Context, fetcher *dsfetch.Fetch, voteIDs []int) ([]attribute.Func, error) {
	optionIDs := make([]int, len(voteIDs))
	voteUserList := make([]int, len(voteIDs))
	voteDelegatedUserList := make([]int, len(voteIDs))
	for i, id := range voteIDs {
		if id == 0 {
			continue
		}

		fetcher.Vote_OptionID(id).Lazy(&optionIDs[i])
		fetcher.Vote_UserID(id).Lazy(&voteUserList[i])
		fetcher.Vote_DelegatedUserID(id).Lazy(&voteDelegatedUserList[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching vote data: %w", err)
	}

	pollIDs, err := fetchPollIDs(ctx, fetcher, optionIDs)
	if err != nil {
		return nil, fmt.Errorf("fetching poll ids: %w", err)
	}

	stateList := make([]string, len(voteIDs))
	for i, id := range pollIDs {
		if id == 0 {
			continue
		}

		fetcher.Poll_State(id).Lazy(&stateList[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching poll state: %w", err)
	}

	canManage, err := Collection(ctx, Poll{}).Modes("MANAGE")(ctx, fetcher, pollIDs)
	if err != nil {
		return nil, fmt.Errorf("checking can manage poll: %w", err)
	}

	attr := make([]attribute.Func, len(voteIDs))
	for i, id := range voteIDs {
		if id == 0 {
			continue
		}

		if stateList[i] == "published" {
			attr[i] = attribute.FuncAllowed
			continue
		}

		if voteUserList[i] == 0 && voteDelegatedUserList[i] == 0 {
			return nil, fmt.Errorf("database is invalid. vote/%d has no vote_user and no vote_delegate", id)
		}

		if voteUserList[i] != 0 && voteDelegatedUserList[i] != 0 {
			return nil, fmt.Errorf("database is invalid. vote/%d has vote_user and vote_delegate", id)
		}

		attr[i] = attribute.FuncOr(
			canManage[i],
			attribute.FuncUserIDs([]int{voteUserList[i] + voteDelegatedUserList[i]}),
		)
	}
	return attr, nil
}

// TODO: Group by poll or option
func (v Vote) modeB(ctx context.Context, fetcher *dsfetch.Fetch, voteIDs []int) ([]attribute.Func, error) {
	// Group B: Depends on poll/state:
	//
	//	published: Accessible if the user can see the vote.
	//	finished: Accessible if the user can manage the associated poll.
	//	others: Not accessible for anyone.
	optionIDs := make([]int, len(voteIDs))
	for i, id := range voteIDs {
		if id == 0 {
			continue
		}
		fetcher.Vote_OptionID(id).Lazy(&optionIDs[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching vote data: %w", err)
	}

	pollIDs, err := fetchPollIDs(ctx, fetcher, optionIDs)
	if err != nil {
		return nil, fmt.Errorf("fetching poll ids: %w", err)
	}

	stateList := make([]string, len(voteIDs))
	for i, id := range pollIDs {
		if id == 0 {
			continue
		}
		fetcher.Poll_State(id).Lazy(&stateList[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching poll state: %w", err)
	}

	canSee, err := Collection(ctx, v).Modes("A")(ctx, fetcher, voteIDs)
	if err != nil {
		return nil, fmt.Errorf("checking vote see: %w", err)
	}

	canManage, err := Collection(ctx, Poll{}).Modes("MANAGE")(ctx, fetcher, pollIDs)
	if err != nil {
		return nil, fmt.Errorf("checking poll manage: %w", err)
	}

	attr := make([]attribute.Func, len(voteIDs))
	for i, id := range voteIDs {
		if id == 0 {
			continue
		}

		switch stateList[i] {
		case "published":
			attr[i] = canSee[i]
		case "finished":
			attr[i] = canManage[i]
		default:
			attr[i] = attribute.FuncNotAllowed
		}
	}
	return attr, nil
}
