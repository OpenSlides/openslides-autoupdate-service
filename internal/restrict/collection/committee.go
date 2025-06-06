package collection

import (
	"context"
	"fmt"
	"slices"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
	"github.com/OpenSlides/openslides-go/perm"
)

// Committee handels permission for committees.
//
// See user can see a committee, if he is in committee/user_ids, have the OML
// can_manage_users or higher or the CML `can_manage` in the committee or an
// parent committee.
//
// Mode A: The user can see the committee.
//
// Mode B:
//   - The user must have the OML `can_manage_organization` or higher or the
//     CML `can_manage` in the committee or an parent committee.
//   - The user has the CML `can_manage` in one of the committees linked in
//     `forward_to_committee_ids` or `receive_forwardings_from_committee_ids`.
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

func (c Committee) see(ctx context.Context, ds *dsfetch.Fetch, committeeIDs ...int) ([]int, error) {
	requestUser, err := perm.RequestUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting request user: %w", err)
	}

	hasOMLPerm, err := perm.HasOrganizationManagementLevel(ctx, ds, requestUser, perm.OMLCanManageUsers)
	if err != nil {
		return nil, fmt.Errorf("checking oml perm: %w", err)
	}

	if hasOMLPerm {
		return committeeIDs, nil
	}

	allowed, err := eachCondition(committeeIDs, func(committeeID int) (bool, error) {
		cmlCanManage, err := perm.HasCommitteeManagementLevel(ctx, ds, requestUser, committeeID)
		if err != nil {
			return false, fmt.Errorf("checking committee management level: %w", err)
		}

		if cmlCanManage {
			return true, nil
		}

		userIDs, err := ds.Committee_UserIDs(committeeID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting committee users: %w", err)
		}

		if slices.Contains(userIDs, requestUser) {
			return true, nil
		}
		return false, nil
	})
	if err != nil {
		return nil, fmt.Errorf("checking if user is in committee: %w", err)
	}

	return allowed, nil
}

func (c Committee) modeB(ctx context.Context, ds *dsfetch.Fetch, committeeIDs ...int) ([]int, error) {
	requestUser, err := perm.RequestUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting request user: %w", err)
	}

	hasOMLPerm, err := perm.HasOrganizationManagementLevel(ctx, ds, requestUser, perm.OMLCanManageOrganization)
	if err != nil {
		return nil, fmt.Errorf("checking oml: %w", err)
	}

	if hasOMLPerm {
		return committeeIDs, nil
	}

	committeeManager, err := perm.ManagementLevelCommittees(ctx, ds, requestUser)
	if err != nil {
		return nil, fmt.Errorf("getting all committees where the request user is manager: %w", err)
	}

	allowed, err := eachCondition(committeeIDs, func(committeeID int) (bool, error) {
		// No error check here. If the fields do not exist, then an empty list ist correct.
		inForwardTo, _ := ds.Committee_ForwardToCommitteeIDs(committeeID).Value(ctx)
		inForwardReceive, _ := ds.Committee_ReceiveForwardingsFromCommitteeIDs(committeeID).Value(ctx)

		combined := append(append([]int{committeeID}, inForwardTo...), inForwardReceive...)

		for _, manageID := range committeeManager {
			if slices.Contains(combined, manageID) {
				return true, nil
			}
		}

		return false, nil
	})
	if err != nil {
		return nil, fmt.Errorf("checking has committee managemement level: %w", err)
	}

	return allowed, nil
}
