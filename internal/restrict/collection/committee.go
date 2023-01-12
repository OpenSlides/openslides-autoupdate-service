package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

// Committee handels permission for committees.
//
// See user can see a committee, if he is in committee/user_ids or have the OML
// can_manage_users or higher.
//
// Mode A: The user can see the committee.
//
// Mode B: The user must have the OML `can_manage_organization` or higher or the
// CML `can_manage` in the committee.
type Committee struct{}

// Name returns the collection name.
func (c Committee) Name() string {
	return "committee"
}

// MeetingID returns the meetingID for the object.
func (c Committee) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	return 0, false, nil
}

// Modes returns a map from all known modes to there restricter.
func (c Committee) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return c.see
	case "B":
		return c.modeB
	}
	return nil
}

func (c Committee) see(ctx context.Context, ds *dsfetch.Fetch, mperms perm.MeetingPermission, attrMap AttributeMap, committeeIDs ...int) error {
	for _, committeeID := range committeeIDs {
		userIDs, err := ds.Committee_UserIDs(committeeID).Value(ctx)
		if err != nil {
			return fmt.Errorf("getting committee users: %w", err)
		}

		attrMap.Add(dskey.Key{Collection: c.Name(), ID: committeeID, Field: "A"}, &Attributes{
			GlobalPermission: byte(perm.OMLCanManageUsers),
			UserIDs:          set.New(userIDs...),
		})
	}

	return nil
}

func (c Committee) modeB(ctx context.Context, ds *dsfetch.Fetch, mperms perm.MeetingPermission, attrMap AttributeMap, committeeIDs ...int) error {
	for _, committeeID := range committeeIDs {
		committeeManager, err := ds.Committee_UserManagementLevel(committeeID, "can_manage").Value(ctx)
		if err != nil {
			return fmt.Errorf("getting committee managers: %w", err)
		}

		attrMap.Add(dskey.Key{Collection: c.Name(), ID: committeeID, Field: "B"}, &Attributes{
			GlobalPermission: byte(perm.OMLCanManageOrganization),
			UserIDs:          set.New(committeeManager...),
		})
	}

	return nil
}
