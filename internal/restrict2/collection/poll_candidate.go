package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
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

func (p PollCandidate) see(ctx context.Context, fetcher *dsfetch.Fetch, pollCandidateIDs []int) ([]attribute.Func, error) {
	return canSeeRelatedCollection(ctx, fetcher, fetcher.PollCandidate_PollCandidateListID, Collection(ctx, PollCandidateList{}).Modes("A"), pollCandidateIDs)
}
