package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

// Committee handels permission for committees.
//
// See user can see a committee, if he is in committee/user_ids or have the OML
// can_manage_users or higher.
//
// Mode A: The user can see the committee.
//
// Mode B: The user must have the OML `can_manage_organization` or higher or the
// CML `can_manage` in the committee or an parent committee.
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
		userIDs, err := ds.Committee_UserIDs(committeeID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting committee users: %w", err)
		}

		for _, uid := range userIDs {
			if uid == requestUser {
				return true, nil
			}
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

	allowed, err := eachCondition(committeeIDs, func(committeeID int) (bool, error) {
		cmlCanManage, err := perm.HasCommitteeManagementLevel(ctx, ds, requestUser, committeeID)
		if err != nil {
			return false, fmt.Errorf("checking committee management level: %w", err)
		}

		return cmlCanManage, nil
	})
	if err != nil {
		return nil, fmt.Errorf("checking has committee managemement level: %w", err)
	}

	return allowed, nil
}
