package collection

import (
	"context"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
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
		return Allways(o.Name(), mode)
	case "B":
		return loggedIn(o.Name(), mode)
	case "C":
		return o.modeC
	}
	return nil
}

func (o Organization) modeC(ctx context.Context, ds *dsfetch.Fetch, mperms perm.MeetingPermission, attrMap AttributeMap, organizationIDs ...int) error {
	attr := Attributes{
		GlobalPermission: byte(perm.OMLCanManageUsers),
	}
	for _, id := range organizationIDs {
		attrMap.Add(dskey.Key{Collection: o.Name(), ID: id, Field: "C"}, &attr)
	}

	return nil
}
