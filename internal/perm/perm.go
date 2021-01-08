package perm

//go:generate  sh -c "go run build_derivate/main.go > derivate.go && go fmt derivate.go"

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
)

// Permission holds the information which permissions and groups a user has.
type Permission struct {
	admin       bool
	groupIDs    []int
	permissions map[string]bool
}

// New creates a new Permission object for a user in a specific meeting.
func New(ctx context.Context, dp dataprovider.DataProvider, userID, meetingID int) (*Permission, error) {
	isMeeting, err := dp.InMeeting(ctx, userID, meetingID)
	if err != nil {
		return nil, fmt.Errorf("Looking for user %d in meeting %d: %w", userID, meetingID, err)
	}
	if !isMeeting {
		return new(Permission), nil
	}

	groupIDs := []int{}
	if err := dp.GetIfExist(ctx, fmt.Sprintf("user/%d/group_$%d_ids", userID, meetingID), &groupIDs); err != nil {
		return nil, fmt.Errorf("get group ids: %w", err)
	}

	// Get superadmin_group_id.
	var superadminGroupID int
	fqfield := fmt.Sprintf("meeting/%d/superadmin_group_id", meetingID)
	if err := dp.GetIfExist(ctx, fqfield, &superadminGroupID); err != nil {
		return nil, fmt.Errorf("check for superadmin group: %w", err)
	}

	if superadminGroupID != 0 {
		for _, id := range groupIDs {
			if id == superadminGroupID {
				return &Permission{admin: true}, nil
			}
		}
	}

	// effectiveGroupIDs are all ids the user is in. If the user is in no group,
	// it is the id of the default group.
	effectiveGroupIDs := groupIDs
	if len(effectiveGroupIDs) == 0 {
		var defaultGroupID int
		fqfield := fmt.Sprintf("meeting/%d/default_group_id", meetingID)
		if err := dp.GetIfExist(ctx, fqfield, &defaultGroupID); err != nil {
			return nil, fmt.Errorf("getting default group: %w", err)
		}
		if defaultGroupID != 0 {
			effectiveGroupIDs = []int{defaultGroupID}
		}
	}

	permissions := make(map[string]bool)
	for _, gid := range effectiveGroupIDs {
		fqfield := fmt.Sprintf("group/%d/permissions", gid)
		var perms []string
		if err := dp.GetIfExist(ctx, fqfield, &perms); err != nil {
			return nil, fmt.Errorf("getting %s: %w", fqfield, err)
		}
		for _, perm := range perms {
			permissions[perm] = true
			for _, p := range derivatePerms[perm] {
				permissions[p] = true
			}
		}
	}
	return &Permission{groupIDs: groupIDs, permissions: permissions}, nil
}

// Has returns true, if the permission object contains the given permissions.
func (p *Permission) Has(perm string) bool {
	if p.admin {
		return true
	}

	return p.permissions[perm]
}

// Create checks for the mermission to create a new object.
func Create(dp dataprovider.DataProvider, managePerm, collection string) WriteChecker {
	return WriteCheckerFunc(func(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
		var meetingID int
		if err := json.Unmarshal(payload["meeting_id"], &meetingID); err != nil {
			return false, fmt.Errorf("no valid meeting id: %w", err)
		}

		ok, err := HasPerm(ctx, dp, userID, meetingID, managePerm)
		if err != nil {
			return false, fmt.Errorf("ensure manage permission: %w", err)
		}

		return ok, nil
	})
}

// Modify checks for the permissions to alter an existing object.
func Modify(dp dataprovider.DataProvider, managePerm, collection string) WriteChecker {
	return WriteCheckerFunc(func(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
		id, err := modelID(payload)
		if err != nil {
			return false, fmt.Errorf("getting model id from payload: %w", err)
		}

		fqid := fmt.Sprintf("%s/%d", collection, id)
		meetingID, err := dp.MeetingFromModel(ctx, fqid)
		if err != nil {
			return false, fmt.Errorf("getting meeting id for model %s: %w", fqid, err)
		}

		ok, err := HasPerm(ctx, dp, userID, meetingID, managePerm)
		if err != nil {
			return false, fmt.Errorf("ensure manage permission: %w", err)
		}

		return ok, nil
	})
}

func modelID(data map[string]json.RawMessage) (int, error) {
	var id int
	if err := json.Unmarshal(data["id"], &id); err != nil {
		return 0, fmt.Errorf("no valid meeting id: %w", err)
	}
	return id, nil
}

// HasPerm returns, if the user has the permission in the meeting.
func HasPerm(ctx context.Context, dp dataprovider.DataProvider, userID int, meetingID int, permission string) (bool, error) {
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
	var err error
	for _, fqfield := range fqfields {
		if lastID != fqfield.ID {
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
	log.Printf(format, a...)
}
