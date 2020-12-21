package collection

//go:generate  sh -c "go run gen_derivate/main.go > derivate.go && go fmt derivate.go"

import (
	"context"
	"fmt"
	"strconv"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
)

// Permissions holds the information which permissions and groups a user has.
type Permissions struct {
	isSuperadmin bool
	groupIds     []int
	permissions  map[string]bool
}

// Perms returns a Permissions object for an user in a meeting.
func Perms(ctx context.Context, userID, meetingID int, dp dataprovider.DataProvider) (*Permissions, error) {
	// Fetch user group ids for the meeting.
	userGroupIDs := []int{}
	userGroupIdsFqfield := "user/" + strconv.Itoa(userID) + "/group_$" + strconv.Itoa(meetingID) + "_ids"
	if err := dp.GetIfExist(ctx, userGroupIdsFqfield, &userGroupIDs); err != nil {
		return nil, fmt.Errorf("get group ids: %w", err)
	}

	// Get superadmin_group_id.
	var superadminGroupID int
	fqfield := "meeting/" + strconv.Itoa(meetingID) + "/superadmin_group_id"
	if err := dp.Get(ctx, fqfield, &superadminGroupID); err != nil {
		return nil, fmt.Errorf("check for superadmin group: %w", err)
	}

	// Direct check: is the user a superadmin?
	for _, id := range userGroupIDs {
		if id == superadminGroupID {
			return &Permissions{isSuperadmin: true, groupIds: userGroupIDs, permissions: map[string]bool{}}, nil
		}
	}

	// Get default group id.
	var defaultGroupID int
	fqfield = "meeting/" + strconv.Itoa(meetingID) + "/default_group_id"
	if err := dp.Get(ctx, fqfield, &defaultGroupID); err != nil {
		return nil, fmt.Errorf("getting default group: %w", err)
	}

	// Get group ids.
	var groupIDs []int
	fqfield = "meeting/" + strconv.Itoa(meetingID) + "/group_ids"
	if err := dp.Get(ctx, fqfield, &groupIDs); err != nil {
		return nil, fmt.Errorf("getting group ids: %w", err)
	}

	// Fetch group permissions: A map from group id <-> permission array.
	groupPermissions := make(map[int][]string)
	for _, id := range groupIDs {
		fqfield := "group/" + strconv.Itoa(id) + "/permissions"
		singleGroupPermissions := []string{}
		if err := dp.GetIfExist(ctx, fqfield, &singleGroupPermissions); err != nil {
			return nil, fmt.Errorf("getting %s: %w", fqfield, err)
		}
		groupPermissions[id] = singleGroupPermissions
	}

	// Collect perms for the user.
	effectiveGroupIds := userGroupIDs
	if len(effectiveGroupIds) == 0 {
		effectiveGroupIds = []int{defaultGroupID}
	}

	permissions := make(map[string]bool, len(effectiveGroupIds))
	for _, id := range effectiveGroupIds {
		for _, perm := range groupPermissions[id] {
			permissions[perm] = true
			for _, p := range derivatePerms[perm] {
				permissions[p] = true
			}
		}
	}

	return &Permissions{isSuperadmin: false, groupIds: userGroupIDs, permissions: permissions}, nil
}

// HasOne returns true, if the permission object contains at least one of the given permissions.
func (p *Permissions) HasOne(perms ...string) bool {
	if p.isSuperadmin {
		return true
	}

	for _, perm := range perms {
		if p.permissions[perm] {
			return true
		}
	}
	return false
}
