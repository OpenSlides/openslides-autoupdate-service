package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
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
//
// Mode E: The user is meeting admin in at least one meeting.
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
	case "E":
		return o.modeE
	}
	return nil
}

func (Organization) modeC(ctx context.Context, ds *dsfetch.Fetch, userIDs ...int) ([]int, error) {
	requestUser, err := perm.RequestUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting request user: %w", err)
	}

	isUserManager, err := perm.HasOrganizationManagementLevel(ctx, ds, requestUser, perm.OMLCanManageUsers)
	if err != nil {
		return nil, fmt.Errorf("check organization management level: %w", err)
	}

	if isUserManager {
		return userIDs, nil
	}

	return nil, nil
}

func (Organization) modeE(ctx context.Context, ds *dsfetch.Fetch, ids ...int) ([]int, error) {
	isAdmin, err := isAdminInAnyMeeting(ctx, ds)
	if err != nil {
		return nil, fmt.Errorf("checking is user meeting admin: %w", err)
	}

	if !isAdmin {
		return nil, nil
	}

	return ids, nil
}
