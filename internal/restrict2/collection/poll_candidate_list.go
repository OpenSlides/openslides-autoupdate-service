package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
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

func (p PollCandidateList) see(ctx context.Context, fetcher *dsfetch.Fetch, pollCandidateListIDs []int) ([]attribute.Func, error) {
	return canSeeRelatedCollection(ctx, fetcher, fetcher.PollCandidateList_OptionID, Collection(ctx, Option{}).Modes("A"), pollCandidateListIDs)
}
