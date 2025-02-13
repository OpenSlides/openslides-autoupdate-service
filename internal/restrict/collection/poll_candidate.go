package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

// PollCandidate handels restriction for the poll_candidate collection.
//
// A user can see a poll candidate list, if he can see the linked poll_candidate_list.
//
// Mode A: The user can see the poll candidate.
type PollCandidate struct{}

// Name returns the collection name.
func (p PollCandidate) Name() string {
	return "poll_candidate"
}

// MeetingID returns the meetingID for the object.
func (p PollCandidate) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.PollCandidate_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meetingID: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (p PollCandidate) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return p.see
	}
	return nil
}

func (p PollCandidate) see(ctx context.Context, ds *dsfetch.Fetch, candidateIDs ...int) ([]int, error) {
	candidateListIDs := make([]int, len(candidateIDs))
	for i, id := range candidateIDs {
		ds.PollCandidate_PollCandidateListID(id).Lazy(&candidateListIDs[i])
	}

	if err := ds.Execute(ctx); err != nil {
		return nil, fmt.Errorf("getting poll candidate list ids: %w", err)
	}

	candidateListToCandidate := make(map[int][]int, len(candidateIDs)) // This will allocate a to big map, but I think it is better then to initialize a zero lenth map.
	for i := 0; i < len(candidateIDs); i++ {
		candidateListID := candidateListIDs[i]
		candidateID := candidateIDs[i]
		candidateListToCandidate[candidateListID] = append(candidateListToCandidate[candidateListID], candidateID)
	}

	allowedCandidateList, err := Collection(ctx, PollCandidateList{}.Name()).Modes("A")(ctx, ds, candidateListIDs...)
	if err != nil {
		return nil, fmt.Errorf("checking restriction of poll_candidate_list: %w", err)
	}

	allowedCandidate := make([]int, 0, len(candidateIDs))
	for _, id := range allowedCandidateList {
		allowedCandidate = append(allowedCandidate, candidateListToCandidate[id]...)
	}

	return allowedCandidate, nil
}
