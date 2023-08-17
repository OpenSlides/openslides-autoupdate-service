package perm

//go:generate  sh -c "go run generate/main.go > generated.go"

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
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
		return newAnonymous(ctx, ds, meetingID)
	}

	isSuperAdmin, err := HasOrganizationManagementLevel(ctx, ds, userID, OMLSuperadmin)
	if err != nil {
		return nil, fmt.Errorf("getting organization management level: %w", err)
	}
	if isSuperAdmin {
		return &Permission{admin: true}, nil
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
		// User is not in the meeting. Do not just return nil. Nil would be an
		// nil interface. But what we want is a interface of type permission
		// with value nil.
		var p *Permission
		return p, nil
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

func newAnonymous(ctx context.Context, ds *dsfetch.Fetch, meetingID int) (*Permission, error) {
	enableAnonymous, err := ds.Meeting_EnableAnonymous(meetingID).Value(ctx)
	if err != nil {
		return nil, fmt.Errorf("checking anonymous enabled: %w", err)
	}
	if !enableAnonymous {
		return nil, nil
	}

	defaultGroupID, err := ds.Meeting_DefaultGroupID(meetingID).Value(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting default group: %w", err)
	}

	perms, err := permissionsFromGroups(ctx, ds, defaultGroupID)
	if err != nil {
		return nil, fmt.Errorf("getting permissions for default group: %w", err)
	}

	return &Permission{groupIDs: []int{defaultGroupID}, permissions: perms}, nil
}

func isAdmin(ctx context.Context, ds *dsfetch.Fetch, meetingID int, groupIDs []int) (bool, error) {
	adminGroupID, exist, err := ds.Meeting_AdminGroupID(meetingID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("check for admin group: %w", err)
	}

	if !exist {
		return false, nil
	}

	for _, id := range groupIDs {
		if id == adminGroupID {
			return true, nil
		}
	}
	return false, nil
}

func permissionsFromGroups(ctx context.Context, ds *dsfetch.Fetch, groupIDs ...int) (map[TPermission]bool, error) {
	permissions := make(map[TPermission]bool)
	for _, gid := range groupIDs {
		perms := ds.Group_Permissions(gid).ErrorLater(ctx)
		for _, perm := range perms {
			permissions[TPermission(perm)] = true
			for _, p := range derivatePerms[TPermission(perm)] {
				permissions[p] = true
			}
		}
	}

	if err := ds.Err(); err != nil {
		return nil, fmt.Errorf("getting permissions for groups %v: %w", groupIDs, err)
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
func (p *Permission) InGroup(gid int) bool {
	if p == nil {
		return false
	}

	if p.admin {
		return true
	}

	for _, id := range p.groupIDs {
		if id == gid {
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

	for _, id := range ids {
		if id == committeeID {
			return true, nil
		}
	}
	return false, nil
}

// ManagementLevelCommittees returns all committee-ids where the given user has
// the management level.
func ManagementLevelCommittees(ctx context.Context, ds *dsfetch.Fetch, userID int) ([]int, error) {
	if userID == 0 {
		return nil, nil
	}

	commiteeIDs, err := ds.User_CommitteeManagementIDs(userID).Value(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetching user/%d/committee_management_ids: %w", userID, err)
	}
	return commiteeIDs, nil
}
