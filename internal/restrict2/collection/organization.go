package collection

import (
	"context"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// Organization handels restrictions of the collection organization.
//
// The user can always see an organization.
//
// Mode A: The user can see the organization (always).
//
// Mode B: The user must be logged in (no anonymous).
//
// Mode C: The user has the OML can_manage_users or higher.
type Organization struct{}

// Name returns the collection name.
func (o Organization) Name() string {
	return "organization"
}

// MeetingID returns the meetingID for the object.
func (o Organization) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	return 0, false, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (o Organization) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return Allways
	case "B":
		return loggedIn
	case "C":
		return o.modeC
	}
	return nil
}

func (o Organization) modeC(ctx context.Context, fetcher *dsfetch.Fetch, organizationIDs []int) ([]attribute.Func, error) {
	return attributeFuncList(len(organizationIDs), attribute.FuncGlobalLevel(perm.OMLCanManageUsers)), nil
}
