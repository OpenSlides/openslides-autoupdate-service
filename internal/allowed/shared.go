package allowed

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/OpenSlides/openslides-permission-service/pkg/definitions"
)

// do not check, if anonymous is enabled but always returns true, if the id is 0!
func DoesUserExists(userId int, ctx *IsAllowedContext) bool {
	if userId == 0 {
		return true
	}
	fqfield := "user/" + strconv.Itoa(userId) + "/id"
	return ctx.DataProvider.Exists(fqfield)
}

func HasUserSuperadminRole(userId int, ctx *IsAllowedContext) (bool, error) {
	// The anonymous is never a superadmin
	if userId == 0 {
		return false, nil
	}

	// get superadmin role id
	superadminRoleId, err := ctx.DataProvider.GetInt("organisation/1/superadmin_role_id")
	if err != nil {
		return false, err
	}

	// Get users role id
	fqfield := "user/" + strconv.Itoa(userId) + "/role_id"
	if !ctx.DataProvider.Exists(fqfield) {
		return false, nil // the user has no role
	}
	userRoleId, err := ctx.DataProvider.GetInt(fqfield)
	if err != nil {
		return false, fmt.Errorf("Error getting role_id: %v", err)
	}

	return superadminRoleId == userRoleId, nil
}

func CanUserSeeMeeting(userId, meetingId int, ctx *IsAllowedContext) (bool, error) {
	// userId in meeting/id/user_ids OR (userId==0 and meeting/id/enable_anonymous is true)

	if userId == 0 {
		enableAnonymous, err := ctx.DataProvider.GetBoolWithDefault("meeting/"+strconv.Itoa(meetingId)+"/enable_anonymous", false)
		if err != nil {
			return false, err
		}
		return enableAnonymous, nil
	} else {
		userIds, err := ctx.DataProvider.GetIntArrayWithDefault("meeting/"+strconv.Itoa(meetingId)+"/user_ids", []int{})
		if err != nil {
			return false, err
		}
		for _, id := range userIds {
			if id == userId {
				return true, nil
			}
		}
		return false, nil
	}
}

type Permissions struct {
	isSuperadmin bool
	groupIds     []int // effective ones!
	permissions  map[string]bool
}

// Assumes, that the user is part of the meeting! (If not, it has no groups and will have the permissions of the default group)
func GetPermissionsForUserInMeeting(userId, meetingId int, ctx *IsAllowedContext) (*Permissions, error) {
	// Fetch user group ids for the meeting
	userGroupIdsFqfield := "user/" + strconv.Itoa(userId) + "/group_" + strconv.Itoa(meetingId) + "_ids"
	userGroupIds, err := ctx.DataProvider.GetIntArrayWithDefault(userGroupIdsFqfield, []int{})
	if err != nil {
		return nil, err
	}

	// get superadmin_group_id
	superadminGroupFqfield := "meeting/" + strconv.Itoa(meetingId) + "/superadmin_group_id"
	superadminGroupId, err := ctx.DataProvider.GetInt(superadminGroupFqfield)
	if err != nil {
		return nil, err
	}

	// direct check: is the user a superadmin?
	for _, id := range userGroupIds {
		if id == superadminGroupId {
			return &Permissions{isSuperadmin: true, groupIds: userGroupIds, permissions: map[string]bool{}}, nil
		}
	}

	// get default group id
	defaultGroupFqfield := "meeting/" + strconv.Itoa(meetingId) + "/default_group_id"
	defaultGroupId, err := ctx.DataProvider.GetInt(defaultGroupFqfield)
	if err != nil {
		return nil, err
	}

	// get group ids
	groupIdsFqfield := "meeting/" + strconv.Itoa(meetingId) + "/group_ids"
	groupIds, err := ctx.DataProvider.GetIntArray(groupIdsFqfield)
	if err != nil {
		return nil, err
	}

	// Fetch group permissions: A map from group id <-> permission array
	groupPermissions := make(map[int][]string)
	for _, id := range groupIds {
		fqfield := "group/" + strconv.Itoa(id) + "/permissions"
		singleGroupPermissions, err := ctx.DataProvider.GetStringArrayWithDefault(fqfield, []string{})
		if nil != err {
			return nil, err
		}
		groupPermissions[id] = singleGroupPermissions
	}

	// collect perms for the user
	effectiveGroupIds := userGroupIds
	if len(effectiveGroupIds) == 0 {
		effectiveGroupIds = []int{defaultGroupId}
	}

	permissions := make(map[string]bool)
	for _, id := range effectiveGroupIds {
		for _, perm := range groupPermissions[id] {
			permissions[perm] = true
		}
	}

	return &Permissions{isSuperadmin: false, groupIds: userGroupIds, permissions: permissions}, nil
}

func (p *Permissions) HasPerm(perm string) bool {
	if p.isSuperadmin {
		return true
	}
	return p.permissions[perm]
}

func GetInt(data definitions.FqfieldData, property string) (int, error) {
	if val, ok := data[property]; ok {
		var value int
		err := json.Unmarshal([]byte(val), &value)

		if nil != err {
			return 0, fmt.Errorf(property + " is not an int")
		}
		return value, nil
	} else {
		return 0, fmt.Errorf(property + " is not in data")
	}
}

func GetMeetingIdFromModel(fqid definitions.Fqid, ctx *IsAllowedContext) (int, error) {
	return ctx.DataProvider.GetInt(fqid + "/meeting_id")
}
