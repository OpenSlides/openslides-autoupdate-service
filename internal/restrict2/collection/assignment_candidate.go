package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// AssignmentCandidate handels the permissions for assignment_candiate collections.
//
// The user can see an assignment candidate, if the user can see the linked assignment.
//
// Mode A: The user can see the assignment candidate.
type AssignmentCandidate struct{}

// Name returns the collection name.
func (a AssignmentCandidate) Name() string {
	return "assignment_candidate"
}

// MeetingID returns the meetingID for the object.
func (a AssignmentCandidate) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	assignmentID, err := ds.AssignmentCandidate_AssignmentID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("fetching assignment id: %w", err)
	}

	return Assignment{}.MeetingID(ctx, ds, assignmentID)
}

// Modes returns the restrictions modes for assignment_candidate.
func (a AssignmentCandidate) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return a.see
	}
	return nil
}

func (a AssignmentCandidate) see(ctx context.Context, fetcher *dsfetch.Fetch, assignmentCandidateIDs []int) ([]attribute.Func, error) {
	return canSeeRelatedCollection(ctx, fetcher, fetcher.AssignmentCandidate_AssignmentID, Collection(ctx, Assignment{}).Modes("A"), assignmentCandidateIDs)
}
