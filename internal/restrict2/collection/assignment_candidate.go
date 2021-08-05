package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// AssignmentCandidate handels the permissions for assignment_candiate collections.
type AssignmentCandidate struct{}

// Modes returns the restrictions modes for assignment_candidate.
func (a AssignmentCandidate) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return a.see
	}
	return nil
}

func (a AssignmentCandidate) see(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, assignmentCandidateID int) (bool, error) {
	assignmentID := datastore.Int(ctx, fetch.FetchIfExist, "assignment_candidate/%d/assignment_id", assignmentCandidateID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching assignment id: %w", err)
	}

	canSeeAssignment, err := Assignment{}.see(ctx, fetch, mperms, assignmentID)
	if err != nil {
		return false, fmt.Errorf("can see assignment: %w", err)
	}

	return canSeeAssignment, nil
}
