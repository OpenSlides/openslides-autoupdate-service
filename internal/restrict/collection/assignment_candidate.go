package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// AssignmentCandidate handels the permissions for assignment_candiate collections.
//
// The user can see an assignment candidate, if the user can see the linked assignment.
//
// Mode A: The user can see the assignment candidate.
type AssignmentCandidate struct{}

// MeetingID returns the meetingID for the object.
func (a AssignmentCandidate) MeetingID(ctx context.Context, ds *datastore.Request, id int) (int, bool, error) {
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

func (a AssignmentCandidate) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, assignmentCandidateID int) (bool, error) {
	assignmentID, err := ds.AssignmentCandidate_AssignmentID(assignmentCandidateID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("fetching assignment id: %w", err)
	}

	canSeeAssignment, err := Assignment{}.see(ctx, ds, mperms, assignmentID)
	if err != nil {
		return false, fmt.Errorf("can see assignment: %w", err)
	}

	return canSeeAssignment, nil
}
