package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// MeetingUser handels permissions for the collection meeting_user.
//
// A User can see a MeetingUser if TODO...
//
// Mode A: The user can see the mediafile.
type MeetingUser struct{}

// Name returns the collection name.
func (m MeetingUser) Name() string {
	return "meeting_user"
}

// MeetingID returns the meetingID for the object.
func (m MeetingUser) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	mid, err := ds.MeetingUser_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meeting_user id: %w", err)
	}
	return mid, true, nil
}

// Modes returns the field modes for the collection mediafile.
func (m MeetingUser) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		// TODO: Fix me
		return Allways
	case "B":
		// TODO: Fix me
		return Allways
		// TODO: Fix the models.yml. There are probably different modes.

	}
	return nil
}
