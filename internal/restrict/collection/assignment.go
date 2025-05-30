package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
	"github.com/OpenSlides/openslides-go/perm"
)

// Assignment handels restrictions for the assignment collection.
//
// The user can see an assignment, if the user has assignment.can_see
//
// Mode A: User can see the assignment.
type Assignment struct{}

// Name returns the collection name.
func (a Assignment) Name() string {
	return "assignment"
}

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

func (a Assignment) see(ctx context.Context, ds *dsfetch.Fetch, assignmentIDs ...int) ([]int, error) {
	return meetingPerm(ctx, ds, a, assignmentIDs, perm.AssignmentCanSee)
}
