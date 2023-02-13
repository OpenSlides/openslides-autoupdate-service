package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// PollCandidateList handels restriction for the poll_candidate_list collection.
//
// A user can see a poll candidate list, if he can see the linked option.
//
// Mode A: The user can see the poll candidate list.
type PollCandidateList struct{}

// Name returns the collection name.
func (p PollCandidateList) Name() string {
	return "poll_candidate_list"
}

// MeetingID returns the meetingID for the object.
func (p PollCandidateList) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.PollCandidateList_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meetingID: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (p PollCandidateList) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return p.see
	}
	return nil
}

func (p PollCandidateList) see(ctx context.Context, ds *dsfetch.Fetch, pollCandidateListIDs ...int) ([]int, error) {
	optionIDs := make([]int, len(pollCandidateListIDs))
	for i, id := range pollCandidateListIDs {
		ds.PollCandidateList_OptionID(id).Lazy(&optionIDs[i])
	}

	if err := ds.Execute(ctx); err != nil {
		return nil, fmt.Errorf("getting option ids: %w", err)
	}

	optionToPollCandidateList := make(map[int][]int, len(pollCandidateListIDs)) // This will allocate a to big map, but I think it is better then to initialize a zero lenth map.
	for i := 0; i < len(pollCandidateListIDs); i++ {
		optionID := optionIDs[i]
		pollCandidateListID := pollCandidateListIDs[i]
		optionToPollCandidateList[optionID] = append(optionToPollCandidateList[optionID], pollCandidateListID)
	}

	allowedOption, err := Collection(ctx, Option{}.Name()).Modes("A")(ctx, ds, optionIDs...)
	if err != nil {
		return nil, fmt.Errorf("checking restriction of options: %w", err)
	}

	allowedPollCandidateList := make([]int, 0, len(pollCandidateListIDs))
	for _, id := range allowedOption {
		allowedPollCandidateList = append(allowedPollCandidateList, optionToPollCandidateList[id]...)
	}

	return allowedPollCandidateList, nil
}
