package collection

import (
	"context"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Organization handels restrictions of the collection organization.
//
// The user can always see an organization.
//
// Mode A: The user can see the organization (always).
//
// Mode B: The user must be logged in (no anonymous).
type Organization struct{}

// MeetingID returns the meetingID for the object.
func (o Organization) MeetingID(ctx context.Context, ds *datastore.Request, id int) (int, bool, error) {
	return 0, false, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (o Organization) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return Allways
	case "B":
		return loggedIn
	}
	return nil
}
