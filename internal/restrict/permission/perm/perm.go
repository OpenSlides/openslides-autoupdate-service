package perm

//go:generate  sh -c "go run generate/main.go > generated.go && go fmt generated.go"

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/dataprovider"
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
func New(ctx context.Context, dp dataprovider.DataProvider, userID, meetingID int) (*Permission, error) {
	if userID == 0 {
		return newAnonymous(ctx, dp, meetingID)
	}

	var groupIDs []int
	if err := dp.GetIfExist(ctx, fmt.Sprintf("user/%d/group_$%d_ids", userID, meetingID), &groupIDs); err != nil {
		return nil, fmt.Errorf("get group ids: %w", err)
	}

	if len(groupIDs) == 0 {
		// User is not in the meeting
		return nil, nil
	}

	admin, err := isAdmin(ctx, dp, meetingID, groupIDs)
	if err != nil {
		return nil, fmt.Errorf("checking if user is admin: %w", err)
	}
	if admin {
		return &Permission{admin: true}, nil
	}

	perms, err := permissionsFromGroups(ctx, dp, groupIDs...)
	if err != nil {
		return nil, fmt.Errorf("getting permissions from all groups: %w", err)
	}

	return &Permission{groupIDs: groupIDs, permissions: perms}, nil
}

func newAnonymous(ctx context.Context, dp dataprovider.DataProvider, meetingID int) (*Permission, error) {
	var enableAnonymous bool
	fqfield := fmt.Sprintf("meeting/%d/enable_anonymous", meetingID)
	if err := dp.GetIfExist(ctx, fqfield, &enableAnonymous); err != nil {
		return nil, fmt.Errorf("checking anonymous enabled: %w", err)
	}
	if !enableAnonymous {
		return nil, nil
	}

	var defaultGroupID int
	fqfield = fmt.Sprintf("meeting/%d/default_group_id", meetingID)
	if err := dp.GetIfExist(ctx, fqfield, &defaultGroupID); err != nil {
		return nil, fmt.Errorf("getting default group: %w", err)
	}

	perms, err := permissionsFromGroups(ctx, dp, defaultGroupID)
	if err != nil {
		return nil, fmt.Errorf("getting permissions of default group: %w", err)
	}

	return &Permission{groupIDs: []int{defaultGroupID}, permissions: perms}, nil
}

func isAdmin(ctx context.Context, dp dataprovider.DataProvider, meetingID int, groupIDs []int) (bool, error) {
	var adminGroupID int
	fqfield := fmt.Sprintf("meeting/%d/admin_group_id", meetingID)
	if err := dp.GetIfExist(ctx, fqfield, &adminGroupID); err != nil {
		return false, fmt.Errorf("check for admin group: %w", err)
	}

	if adminGroupID != 0 {
		for _, id := range groupIDs {
			if id == adminGroupID {
				return true, nil
			}
		}
	}
	return false, nil
}

func permissionsFromGroups(ctx context.Context, dp dataprovider.DataProvider, groupIDs ...int) (map[TPermission]bool, error) {
	permissions := make(map[TPermission]bool)
	for _, gid := range groupIDs {
		fqfield := fmt.Sprintf("group/%d/permissions", gid)
		var perms []string
		if err := dp.GetIfExist(ctx, fqfield, &perms); err != nil {
			return nil, fmt.Errorf("getting %s: %w", fqfield, err)
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
	return p.admin
}

// InGroup returns true, if the user is in the given group (by group_id).
func (p *Permission) InGroup(gid int) bool {
	if p == nil {
		return false
	}
	
	for _, id := range p.groupIDs {
		if id == gid {
			return true
		}
	}
	return false
}

// HasPerm tells if the given user has a speficic permission in the meeting.
//
// It is a shortcut for calling p := perm.New(...);p.Has(...).
func HasPerm(ctx context.Context, dp dataprovider.DataProvider, userID int, meetingID int, permission TPermission) (bool, error) {
	perm, err := New(ctx, dp, userID, meetingID)
	if err != nil {
		return false, fmt.Errorf("collecting perms: %w", err)
	}

	hasPerms := perm.Has(permission)
	if !hasPerms {
		LogNotAllowedf("User %d does not have the permission %s in meeting %d", userID, permission, meetingID)
		return false, nil
	}

	return true, nil
}

// AllFields checks all fqfields by the given function f.
//
// It asumes, that if a user can see one field of the object, he can see all
// fields. So the check is only called once per fqid.
func AllFields(fqfields []FQField, result map[string]bool, f func(FQField) (bool, error)) error {
	var hasPerm bool
	var lastID int
	for _, fqfield := range fqfields {
		if lastID != fqfield.ID {
			lastID = fqfield.ID
			var err error
			hasPerm, err = f(fqfield)
			if err != nil {
				return fmt.Errorf("checking %s: %w", fqfield, err)
			}
		}
		if hasPerm {
			result[fqfield.String()] = true
		}
	}
	return nil
}

// LogNotAllowedf logs the permission failer.
func LogNotAllowedf(format string, a ...interface{}) {
	// log.Printf(format, a...)
}
