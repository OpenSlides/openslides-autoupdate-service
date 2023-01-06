package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
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
type Committee struct {
	name string
}

// Name returns the collection name.
func (a Committee) Name() string {
	return a.name
}

// MeetingID returns the meetingID for the object.
func (a Committee) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	return 0, false, nil
}

// Modes returns a map from all known modes to there restricter.
func (a Committee) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return a.see
	case "B":
		return a.modeB
	}
	return nil
}

func (a Committee) see(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, attrMap AttributeMap, committeeIDs ...int) error {
	for _, committeeID := range committeeIDs {
		userIDs, err := ds.Committee_UserIDs(committeeID).Value(ctx)
		if err != nil {
			return fmt.Errorf("getting committee users: %w", err)
		}

		attrMap[committeeID] = &Attributes{
			GlobalPermission: byte(perm.OMLCanManageUsers),
			UserIDs:          set.New(userIDs...),
		}
	}

	return nil
}

func (a Committee) modeB(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, attrMap AttributeMap, committeeIDs ...int) error {
	for _, committeeID := range committeeIDs {
		committeeManager, err := ds.Committee_UserManagementLevel(committeeID, "can_manage").Value(ctx)
		if err != nil {
			return fmt.Errorf("getting committee managers: %w", err)
		}

		attrMap[committeeID] = &Attributes{
			GlobalPermission: byte(perm.OMLCanManageOrganization),
			UserIDs:          set.New(committeeManager...),
		}
	}

	return nil
}
