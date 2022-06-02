package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
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

// MeetingID returns the meetingID for the object.
func (a Committee) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	return 0, false, nil
}

// Modes returns a map from all known modes to there restricter.
func (a Committee) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return todoToSingle(a.see)
	case "B":
		return todoToSingle(a.modeB)
	}
	return nil
}

func (a Committee) see(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, committeeID int) (bool, error) {
	userIDs, err := ds.Committee_UserIDs(committeeID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting committee users: %w", err)
	}

	for _, uid := range userIDs {
		if uid == mperms.UserID() {
			return true, nil
		}
	}

	hasOMLPerm, err := perm.HasOrganizationManagementLevel(ctx, ds, mperms.UserID(), perm.OMLCanManageUsers)
	if err != nil {
		return false, fmt.Errorf("checking oml perm: %w", err)
	}

	return hasOMLPerm, nil
}

func (a Committee) modeB(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, committeeID int) (bool, error) {
	hasOMLPerm, err := perm.HasOrganizationManagementLevel(ctx, ds, mperms.UserID(), perm.OMLCanManageOrganization)
	if err != nil {
		return false, fmt.Errorf("checking oml: %w", err)
	}

	if hasOMLPerm {
		return true, nil
	}

	cmlCanManage, err := perm.HasCommitteeManagementLevel(ctx, ds, mperms.UserID(), committeeID)
	if err != nil {
		return false, fmt.Errorf("checking committee management level: %w", err)
	}

	return cmlCanManage, nil
}
