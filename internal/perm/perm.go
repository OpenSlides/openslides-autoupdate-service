package perm

//go:generate  sh -c "go run build_derivate/main.go > derivate.go && go fmt derivate.go"

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

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
	// TODO: With this code, a Committe-Manager is superadmin in the meeting.
	// Discuss if we should do it like this.
	committeeID, err := dp.CommitteeID(ctx, meetingID)
	if err != nil {
		return nil, fmt.Errorf("getting committee id for meeting %d: %w", meetingID, err)
	}

	committeeManager, err := dp.IsManager(ctx, userID, committeeID)
	if err != nil {
		return nil, fmt.Errorf("check for manager in committee %d: %w", committeeID, err)
	}
	if committeeManager {
		return &Permission{admin: true}, nil
	}

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
	if err := dp.Get(ctx, fqfield, &superadminGroupID); err != nil {
		return nil, fmt.Errorf("check for superadmin group: %w", err)
	}

	for _, id := range groupIDs {
		if id == superadminGroupID {
			return &Permission{admin: true}, nil
		}
	}

	// effectiveGroupIDs are all ids the user is in. If the user is in no group,
	// it is the id of the default group.
	effectiveGroupIDs := groupIDs
	if len(effectiveGroupIDs) == 0 {
		var defaultGroupID int
		fqfield := fmt.Sprintf("meeting/%d/default_group_id", meetingID)
		if err := dp.Get(ctx, fqfield, &defaultGroupID); err != nil {
			return nil, fmt.Errorf("getting default group: %w", err)
		}
		effectiveGroupIDs = []int{defaultGroupID}
	}

	permissions := make(map[string]bool)
	for _, gid := range effectiveGroupIDs {
		fqfield := fmt.Sprintf("group/%d/permissions", gid)
		var perms []string
		if err := dp.Get(ctx, fqfield, &perms); err != nil {
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
func Create(dp dataprovider.DataProvider, perm, collection string) WriteChecker {
	return WriteCheckerFunc(func(ctx context.Context, userID int, payload map[string]json.RawMessage) (map[string]interface{}, error) {
		var meetingID int
		if err := json.Unmarshal(payload["meeting_id"], &meetingID); err != nil {
			return nil, fmt.Errorf("no valid meeting id: %w", err)
		}

		return check(ctx, dp, perm, meetingID, userID, payload)
	})
}

// Modify checks for the permissions to alter an existing object.
func Modify(dp dataprovider.DataProvider, perm, collection string) WriteChecker {
	return WriteCheckerFunc(func(ctx context.Context, userID int, payload map[string]json.RawMessage) (map[string]interface{}, error) {
		id, err := modelID(payload)
		if err != nil {
			return nil, fmt.Errorf("getting model id from payload: %w", err)
		}

		fqid := fmt.Sprintf("%s/%d", collection, id)
		meetingID, err := dp.MeetingFromModel(ctx, fqid)
		if err != nil {
			return nil, fmt.Errorf("getting meeting id for model %s: %w", fqid, err)
		}

		return check(ctx, dp, perm, meetingID, userID, payload)
	})
}

func check(ctx context.Context, dp dataprovider.DataProvider, managePerm string, meetingID int, userID int, payload map[string]json.RawMessage) (map[string]interface{}, error) {
	superUser, err := dp.IsSuperuser(ctx, userID)
	if err != nil {
		return nil, err
	}
	if superUser {
		return nil, nil
	}

	if err := EnsurePerm(ctx, dp, userID, meetingID, managePerm); err != nil {
		return nil, fmt.Errorf("ensure manage permission: %w", err)
	}
	return nil, nil
}

func modelID(data map[string]json.RawMessage) (int, error) {
	var id int
	if err := json.Unmarshal(data["id"], &id); err != nil {
		return 0, fmt.Errorf("no valid meeting id: %w", err)
	}
	return id, nil
}

// Restrict tells, if the user has the permission to see the requested
// fields.
func Restrict(dp dataprovider.DataProvider, perm, collection string) ReadeChecker {
	return ReadeCheckerFunc(func(ctx context.Context, userID int, fqfields []FQField, result map[string]bool) error {
		if len(fqfields) == 0 {
			return nil
		}

		for _, fqfield := range fqfields {
			meetingID, err := dp.MeetingFromModel(ctx, collection+"/"+strconv.Itoa(fqfield.ID))
			if err != nil {
				return fmt.Errorf("getting meeting from model: %w", err)
			}

			if err := EnsurePerm(ctx, dp, userID, meetingID, perm); err != nil {
				return nil
			}

			for _, fqfield := range fqfields {
				result[fqfield.String()] = true
			}
		}

		return nil
	})
}

// EnsurePerm makes sure the user has at the given permission.
//
// If the user has the permission, EnsurePerm does not return an error.
//
// If the returned error object is unwrapped to type NotAllowedError, it means,
// that the user does not have the permission. Other errors means, that a reald
// error happend.
func EnsurePerm(ctx context.Context, dp dataprovider.DataProvider, userID int, meetingID int, permission string) error {
	perm, err := New(ctx, dp, userID, meetingID)
	if err != nil {
		return fmt.Errorf("collecting perms: %w", err)
	}

	hasPerms := perm.Has(permission)
	if !hasPerms {
		return NotAllowedf("User %d does not have the permission %s int meeting %d", userID, permission, meetingID)
	}

	return nil
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
