package collection

import (
	"context"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// OrganizationTag handels restrictions of the collection organization_tag.
//
// A logged in user can always see an organization tag.
//
// Mode A: The user can see the organization tag.
type OrganizationTag struct {
	name string
}

// Name returns the collection name.
func (o OrganizationTag) Name() string {
	return o.name
}

// MeetingID returns the meetingID for the object.
func (o OrganizationTag) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	return 0, false, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (o OrganizationTag) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return loggedIn(o.name, mode)
	}
	return nil
}
