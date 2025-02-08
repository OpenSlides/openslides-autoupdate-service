package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
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

func (a AssignmentCandidate) see(ctx context.Context, ds *dsfetch.Fetch, assignmentCandidateIDs ...int) ([]int, error) {
	return eachRelationField(ctx, ds.AssignmentCandidate_AssignmentID, assignmentCandidateIDs, func(assignmentID int, ids []int) ([]int, error) {
		canSeeAssignment, err := Collection(ctx, Assignment{}.Name()).Modes("A")(ctx, ds, assignmentID)
		if err != nil {
			return nil, fmt.Errorf("can see assignment: %w", err)
		}

		if len(canSeeAssignment) == 1 {
			return ids, nil
		}

		return nil, nil
	})
}
