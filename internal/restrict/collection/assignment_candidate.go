package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// AssignmentCandidate handels the permissions for assignment_candiate collections.
//
// The user can see an assignment candidate, if the user can see the linked assignment.
//
// Mode A: The user can see the assignment candidate.
type AssignmentCandidate struct {
	name string
}

// Name returns the collection name.
func (a AssignmentCandidate) Name() string {
	return a.name
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

func (a AssignmentCandidate) see(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, attrMap AttributeMap, assignmentCandidateIDs ...int) error {
	return eachRelationField(ctx, ds.AssignmentCandidate_AssignmentID, assignmentCandidateIDs, func(assignmentID int, ids []int) error {
		// TODO: This only works if assignment is calculated before assignment_candidate
		for _, id := range ids {
			attrMap.Add(a.name, id, "A", attrMap.Get("assignment", assignmentID, "A"))
		}
		return nil
	})
}
