package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// Assignment handels restrictions for the assignment collection.
//
// The user can see an assignment, if the user has assignment.can_see
//
// Mode A: User can see the assignment.
type Assignment struct{}

// MeetingID returns the meetingID for the object.
func (a Assignment) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	mid, err := ds.Assignment_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meeting id: %w", err)
	}
	return mid, true, nil
}

// Modes returns the restricter for the a restriction mode.
func (a Assignment) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return a.see
	}
	return nil
}

func (a Assignment) see(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, attrMap map[int]*Attributes, assignmentIDs ...int) error {
	return meetingPerm(ctx, ds, a, assignmentIDs, mperms, perm.AssignmentCanSee, attrMap)
}
