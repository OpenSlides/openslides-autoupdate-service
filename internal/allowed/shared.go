package allowed

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/definitions"
)

// DoesUserExists does not check, if anonymous is enabled but always returns true, if the id is 0!
func DoesUserExists(userID int, dp dataprovider.DataProvider) bool {
	if userID == 0 {
		return true
	}

	fqfield := "user/" + strconv.Itoa(userID) + "/id"
	return dp.Exists(fqfield)
}

// HasUserSuperadminRole does ....
func HasUserSuperadminRole(userID int, dp dataprovider.DataProvider) (bool, error) {
	// The anonymous is never a superadmin
	if userID == 0 {
		return false, nil
	}

	// get superadmin role id
	superadminRoleID, err := dp.GetInt("organisation/1/superadmin_role_id")
	if err != nil {
		return false, err
	}

	// Get users role id
	fqfield := "user/" + strconv.Itoa(userID) + "/role_id"
	if !dp.Exists(fqfield) {
		return false, nil // the user has no role
	}
	userRoleID, err := dp.GetInt(fqfield)
	if err != nil {
		return false, fmt.Errorf("Error getting role_id: %v", err)
	}

	return superadminRoleID == userRoleID, nil
}

// CanUserSeeMeeting does ...
func CanUserSeeMeeting(userID, meetingID int, dp dataprovider.DataProvider) (bool, error) {
	// userId in meeting/id/user_ids OR (userId==0 and meeting/id/enable_anonymous is true)

	if userID == 0 {
		enableAnonymous, err := dp.GetBoolWithDefault("meeting/"+strconv.Itoa(meetingID)+"/enable_anonymous", false)
		if err != nil {
			return false, err
		}
		return enableAnonymous, nil
	} else {
		userIds, err := dp.GetIntArrayWithDefault("meeting/"+strconv.Itoa(meetingID)+"/user_ids", []int{})
		if err != nil {
			return false, err
		}
		for _, id := range userIds {
			if id == userID {
				return true, nil
			}
		}
		return false, nil
	}
}

// Permissions does ...
type Permissions struct {
	isSuperadmin bool
	groupIds     []int // effective ones!
	permissions  map[string]bool
}

// GetPermissionsForUserInMeeting assumes, that the user is part of the meeting!
// (If not, it has no groups and will have the permissions of the default group)
func GetPermissionsForUserInMeeting(userID, meetingID int, dp dataprovider.DataProvider) (*Permissions, error) {
	// Fetch user group ids for the meeting
	userGroupIdsFqfield := "user/" + strconv.Itoa(userID) + "/group_" + strconv.Itoa(meetingID) + "_ids"
	userGroupIds, err := dp.GetIntArrayWithDefault(userGroupIdsFqfield, []int{})
	if err != nil {
		return nil, err
	}

	// get superadmin_group_id
	superadminGroupFqfield := "meeting/" + strconv.Itoa(meetingID) + "/superadmin_group_id"
	superadminGroupID, err := dp.GetInt(superadminGroupFqfield)
	if err != nil {
		return nil, err
	}

	// direct check: is the user a superadmin?
	for _, id := range userGroupIds {
		if id == superadminGroupID {
			return &Permissions{isSuperadmin: true, groupIds: userGroupIds, permissions: map[string]bool{}}, nil
		}
	}

	// get default group id
	defaultGroupFqfield := "meeting/" + strconv.Itoa(meetingID) + "/default_group_id"
	defaultGroupID, err := dp.GetInt(defaultGroupFqfield)
	if err != nil {
		return nil, err
	}

	// get group ids
	groupIdsFqfield := "meeting/" + strconv.Itoa(meetingID) + "/group_ids"
	groupIds, err := dp.GetIntArray(groupIdsFqfield)
	if err != nil {
		return nil, err
	}

	// Fetch group permissions: A map from group id <-> permission array
	groupPermissions := make(map[int][]string)
	for _, id := range groupIds {
		fqfield := "group/" + strconv.Itoa(id) + "/permissions"
		singleGroupPermissions, err := dp.GetStringArrayWithDefault(fqfield, []string{})
		if nil != err {
			return nil, err
		}
		groupPermissions[id] = singleGroupPermissions
	}

	// collect perms for the user
	effectiveGroupIds := userGroupIds
	if len(effectiveGroupIds) == 0 {
		effectiveGroupIds = []int{defaultGroupID}
	}

	permissions := make(map[string]bool)
	for _, id := range effectiveGroupIds {
		for _, perm := range groupPermissions[id] {
			permissions[perm] = true
		}
	}

	return &Permissions{isSuperadmin: false, groupIds: userGroupIds, permissions: permissions}, nil
}

// HasPerm does ...
func (p *Permissions) HasPerm(perm string) bool {
	if p.isSuperadmin {
		return true
	}
	return p.permissions[perm]
}

// GetInt does ...
func GetInt(data definitions.FqfieldData, property string) (int, error) {
	if val, ok := data[property]; ok {
		var value int
		err := json.Unmarshal([]byte(val), &value)

		if nil != err {
			return 0, fmt.Errorf(property + " is not an int")
		}
		return value, nil
	}

	return 0, fmt.Errorf(property + " is not in data")
}

// GetMeetingIDFromModel does ...
func GetMeetingIDFromModel(FQID definitions.Fqid, dp dataprovider.DataProvider) (int, error) {
	return dp.GetInt(FQID + "/meeting_id")
}
