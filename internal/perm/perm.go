package perm

//go:generate  sh -c "go run build_derivate/main.go > derivate.go && go fmt derivate.go"

import (
	"context"
	"encoding/json"
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
	committeeID, err := dp.CommitteeID(ctx, meetingID)
	if err != nil {
		return fmt.Errorf("getting committee id for meeting: %w", err)
	}

	committeeManager, err := dp.IsManager(ctx, userID, committeeID)
	if err != nil {
		return fmt.Errorf("check for manager: %w", err)
	}
	if committeeManager {
		return nil
	}

	isMeeting, err := dp.InMeeting(ctx, userID, meetingID)
	if err != nil {
		return fmt.Errorf("Looking for user %d in meeting %d: %w", userID, meetingID, err)
	}
	if !isMeeting {
		return NotAllowedf("User %d is not in meeting %d", userID, meetingID)
	}

	perms, err := Perms(ctx, userID, meetingID, dp)
	if err != nil {
		return fmt.Errorf("getting user permissions: %w", err)
	}

	hasPerms := perms.HasOne(permission)
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
