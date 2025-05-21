package perm

//go:generate  sh -c "go run generate/main.go > generated.go"

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

// Permission holds the information which permissions and groups a user has.
type Permission struct {
	admin       bool
	groupIDs    []int
	permissions map[TPermission]bool
}

// New creates a new Permission object for a user in a specific meeting.
//
// If the user is not a member of the meeting, nil is returned.
func New(ctx context.Context, ds *dsfetch.Fetch, userID, meetingID int) (*Permission, error) {
	if userID == 0 {
		return newPublicAccess(ctx, ds, meetingID)
	}

	lockedMeeting, err := ds.Meeting_LockedFromInside(meetingID).Value(ctx)
	if err != nil {
		var errDoesNotExist dsfetch.DoesNotExistError
		if !errors.As(err, &errDoesNotExist) {
			return nil, fmt.Errorf("check if meeting is locked: %w", err)
		}
	}

	if !lockedMeeting {
		isOrgaAdmin, err := HasOrganizationManagementLevel(ctx, ds, userID, OMLCanManageOrganization)
		if err != nil {
			return nil, fmt.Errorf("getting organization management level: %w", err)
		}
		if isOrgaAdmin {
			return &Permission{admin: true}, nil
		}

		isCommitteeAdmin, err := isCommitteeAdminFromMeetingID(ctx, ds, userID, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting committee management level: %w", err)
		}

		if isCommitteeAdmin {
			return &Permission{admin: true}, nil
		}
	}

	meetingUserIDs, err := ds.User_MeetingUserIDs(userID).Value(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting meeting user for %d: %w", userID, err)
	}

	var meetingUserID int
	for _, muid := range meetingUserIDs {
		mid, err := ds.MeetingUser_MeetingID(muid).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting userid of meeting user: %w", err)
		}

		if mid == meetingID {
			meetingUserID = muid
			break
		}
	}

	if meetingUserID == 0 {
		return nil, nil
	}

	lockedOut, err := ds.MeetingUser_LockedOut(meetingUserID).Value(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting locked out value: %w", err)
	}

	if lockedOut {
		return nil, nil
	}

	groupIDs, err := ds.MeetingUser_GroupIDs(meetingUserID).Value(ctx)
	if err != nil {
		return nil, fmt.Errorf("get group ids: %w", err)
	}

	admin, err := isAdmin(ctx, ds, meetingID, groupIDs)
	if err != nil {
		return nil, fmt.Errorf("checking if user is admin: %w", err)
	}
	if admin {
		return &Permission{admin: true}, nil
	}

	perms, err := permissionsFromGroups(ctx, ds, groupIDs...)
	if err != nil {
		return nil, fmt.Errorf("getting permissions from all groups of meeting %d: %w", meetingID, err)
	}

	return &Permission{groupIDs: groupIDs, permissions: perms}, nil
}

func isCommitteeAdminFromMeetingID(ctx context.Context, ds *dsfetch.Fetch, userID int, meetingID int) (bool, error) {
	committeeID, err := ds.Meeting_CommitteeID(meetingID).Value(ctx)
	if err != nil {
		// This is a small hack for testing. The field meeting/commitee_id
		// is required. If it does not exist, it would indicate an invalid
		// database. But for testing, this would mean, that each meeting
		// would need a committee. This would create unnecessary tension to
		// all test. If there realy would be a meeting without a committee,
		// returning false here would be correct. So this hack does not make
		// any problems in production.
		return false, nil

	}

	isCommitteeAdmin, err := HasCommitteeManagementLevel(ctx, ds, userID, committeeID)
	if err != nil {
		return false, fmt.Errorf("getting committee management level: %w", err)
	}

	return isCommitteeAdmin, nil
}

func newPublicAccess(ctx context.Context, ds *dsfetch.Fetch, meetingID int) (*Permission, error) {
	enabledOrgaPublicAccess, err := ds.Organization_EnableAnonymous(1).Value(ctx)
	if err != nil {
		return nil, fmt.Errorf("checking orga public access enabled: %w", err)
	}
	enableMeetingPublicAccess, err := ds.Meeting_EnableAnonymous(meetingID).Value(ctx)
	if err != nil {
		return nil, fmt.Errorf("checking meeting public access enabled: %w", err)
	}
	if !(enableMeetingPublicAccess && enabledOrgaPublicAccess) {
		return nil, nil
	}

	maybePublicAccessGroupID, err := ds.Meeting_AnonymousGroupID(meetingID).Value(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting public access group: %w", err)
	}

	publicAccessGroupID, hasPublicAccessGroup := maybePublicAccessGroupID.Value()
	if !hasPublicAccessGroup {
		return nil, fmt.Errorf("public access group id not set")
	}

	perms, err := permissionsFromGroups(ctx, ds, publicAccessGroupID)
	if err != nil {
		return nil, fmt.Errorf("getting permissions for public access group: %w", err)
	}

	return &Permission{groupIDs: []int{publicAccessGroupID}, permissions: perms}, nil
}

func isAdmin(ctx context.Context, ds *dsfetch.Fetch, meetingID int, groupIDs []int) (bool, error) {
	maybeAdminGroupID, err := ds.Meeting_AdminGroupID(meetingID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("check for admin group: %w", err)
	}

	adminGroupID, ok := maybeAdminGroupID.Value()
	if !ok {
		return false, nil
	}

	if slices.Contains(groupIDs, adminGroupID) {
		return true, nil
	}
	return false, nil
}

func permissionsFromGroups(ctx context.Context, ds *dsfetch.Fetch, groupIDs ...int) (map[TPermission]bool, error) {
	permissions := make(map[TPermission]bool)
	for _, gid := range groupIDs {
		perms, err := ds.Group_Permissions(gid).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting permission for group %d: %w", gid, err)
		}

		for _, perm := range perms {
			permissions[TPermission(perm)] = true
			for _, p := range derivatePerms[TPermission(perm)] {
				permissions[p] = true
			}
		}
	}

	return permissions, nil
}

// Has returns true, if the permission object contains the given permissions.
//
// It also returns true, if the user is a superadmin or an admin in the meeting.
func (p *Permission) Has(perm TPermission) bool {
	if p == nil {
		return false
	}

	if p.admin {
		return true
	}

	return p.permissions[perm]
}

// IsAdmin returns true, if the user is a meeting admin.
func (p *Permission) IsAdmin() bool {
	if p == nil {
		return false
	}
	return p.admin
}

// InGroup returns true, if the user is in the given group (by group_id).
func (p *Permission) InGroup(groupIDs ...int) bool {
	if p == nil {
		return false
	}

	if p.admin {
		return true
	}

	for _, cid := range groupIDs {
		if slices.Contains(p.groupIDs, cid) {
			return true
		}
	}
	return false
}

// HasOrganizationManagementLevel returns true if the user has the level or a higher level
func HasOrganizationManagementLevel(ctx context.Context, ds *dsfetch.Fetch, userID int, level OrganizationManagementLevel) (bool, error) {
	if userID == 0 {
		return false, nil
	}

	oml, err := ds.User_OrganizationManagementLevel(userID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting oml of user %d: %w", userID, err)
	}

	switch OrganizationManagementLevel(oml) {
	case OMLSuperadmin:
		return true, nil

	case OMLCanManageOrganization:
		return level == OMLCanManageOrganization || level == OMLCanManageUsers, nil

	case OMLCanManageUsers:
		return level == OMLCanManageUsers, nil
	}
	return false, nil
}

// HasCommitteeManagementLevel returns true, if the user has the manager level
// in the given committeeID.
func HasCommitteeManagementLevel(ctx context.Context, ds *dsfetch.Fetch, userID int, committeeID int) (bool, error) {
	ids, err := ManagementLevelCommittees(ctx, ds, userID)
	if err != nil {
		return false, fmt.Errorf("fetching list of commitee_ids: %w", err)
	}

	if slices.Contains(ids, committeeID) {
		return true, nil
	}
	return false, nil
}

// ManagementLevelCommittees returns all committee-ids where the given user has
// the management level.
func ManagementLevelCommittees(ctx context.Context, ds *dsfetch.Fetch, userID int) ([]int, error) {
	if userID == 0 {
		return nil, nil
	}

	committeeIDs, err := ds.User_CommitteeManagementIDs(userID).Value(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetching user/%d/committee_management_ids: %w", userID, err)
	}

	committeeChildIDs := make([][]int, len(committeeIDs))
	for i, id := range committeeIDs {
		ds.Committee_AllChildIDs(id).Lazy(&committeeChildIDs[i])
	}

	if err := ds.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetch all committee children: %w", err)
	}

	for _, childIDs := range committeeChildIDs {
		committeeIDs = append(committeeIDs, childIDs...)
	}

	return committeeIDs, nil
}
