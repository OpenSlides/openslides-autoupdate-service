package allowed

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/definitions"
)

// MakeSet does ...
func MakeSet(fields []definitions.Field) map[definitions.Field]bool {
	fieldMap := make(map[definitions.Field]bool)
	for _, field := range fields {
		fieldMap[field] = true
	}
	return fieldMap
}

// ValidateFields returns an error, if there are fields in data, that are not in
// allowedFields.
func ValidateFields(data definitions.FqfieldData, allowedFields map[definitions.Field]bool) error {
	invalidFields := make([]definitions.Field, 0)
	for field := range data {
		if !allowedFields[field] {
			invalidFields = append(invalidFields, field)
		}
	}

	if len(invalidFields) > 0 {
		return NotAllowedf("Invalid fields: %v", invalidFields)
	}
	return nil
}

// DoesUserExists does not check, if anonymous is enabled but always returns true, if the id is 0!
func DoesUserExists(userID int, dp dataprovider.DataProvider) (bool, error) {
	if userID == 0 {
		return true, nil
	}

	exists, err := DoesModelExists(definitions.FqidFromCollectionAndId("user", userID), dp)
	if err != nil {
		err = fmt.Errorf("DoesUserExists: %w", err)
	}
	return exists, err
}

func DoesModelExists(fqid definitions.Fqid, dp dataprovider.DataProvider) (bool, error) {
	fqfield := definitions.FqfieldFromFqidAndField(fqid, "id")
	exists, err := dp.Exists(fqfield)
	if err != nil {
		err = fmt.Errorf("DoesModelExists: %w", err)
	}
	return exists, err
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
		return false, fmt.Errorf("HasUserSuperadminRole: %w", err)
	}

	// Get users role id
	fqfield := "user/" + strconv.Itoa(userID) + "/role_id"
	if exists, err := dp.Exists(fqfield); !exists || err != nil {
		return false, err // the user has no role
	}
	userRoleID, err := dp.GetInt(fqfield)
	if err != nil {
		return false, fmt.Errorf("Error getting role_id: %w", err)
	}

	return superadminRoleID == userRoleID, nil
}

func GetCommitteeIdFromMeeting(meetingID int, dp dataprovider.DataProvider) (int, error) {
	committeeID, err := dp.GetInt("meeting/" + strconv.Itoa(meetingID) + "/committee_id")
	if err != nil {
		return 0, fmt.Errorf("GetCommitteeIdFromMeeting: %w", err)
	}
	return committeeID, nil
}

func IsUserCommitteeManager(userID, committeeID int, dp dataprovider.DataProvider) (bool, error) {
	// The anonymous is never a manager
	if userID == 0 {
		return false, nil
	}

	// get committee manager_ids
	managerIDs, err := dp.GetIntArrayWithDefault("committee/"+strconv.Itoa(committeeID)+"/manager_ids", []int{})
	if err != nil {
		return false, fmt.Errorf("IsUserCommitteeManager: %w", err)
	}

	for _, id := range managerIDs {
		if userID == id {
			return true, nil
		}
	}
	return false, nil
}

// CanUserSeeMeeting does ...
func CanUserSeeMeeting(userID, meetingID int, dp dataprovider.DataProvider) (bool, error) {
	// userId in meeting/id/user_ids OR (userId==0 and meeting/id/enable_anonymous is true)

	if userID == 0 {
		enableAnonymous, err := dp.GetBoolWithDefault("meeting/"+strconv.Itoa(meetingID)+"/enable_anonymous", false)
		if err != nil {
			return false, fmt.Errorf("CanUserSeeMeeting: %w", err)
		}
		return enableAnonymous, nil
	} else {
		userIds, err := dp.GetIntArrayWithDefault("meeting/"+strconv.Itoa(meetingID)+"/user_ids", []int{})
		if err != nil {
			return false, fmt.Errorf("CanUserSeeMeeting: %w", err)
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
		return nil, fmt.Errorf("GetPermissionsForUserInMeeting: %w", err)
	}

	// get superadmin_group_id
	superadminGroupFqfield := "meeting/" + strconv.Itoa(meetingID) + "/superadmin_group_id"
	superadminGroupID, err := dp.GetInt(superadminGroupFqfield)
	if err != nil {
		return nil, fmt.Errorf("GetPermissionsForUserInMeeting: %w", err)
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
		return nil, fmt.Errorf("GetPermissionsForUserInMeeting: %w", err)
	}

	// get group ids
	groupIdsFqfield := "meeting/" + strconv.Itoa(meetingID) + "/group_ids"
	groupIds, err := dp.GetIntArray(groupIdsFqfield)
	if err != nil {
		return nil, fmt.Errorf("GetPermissionsForUserInMeeting: %w", err)
	}

	// Fetch group permissions: A map from group id <-> permission array
	groupPermissions := make(map[int][]string)
	for _, id := range groupIds {
		fqfield := "group/" + strconv.Itoa(id) + "/permissions"
		singleGroupPermissions, err := dp.GetStringArrayWithDefault(fqfield, []string{})
		if nil != err {
			return nil, fmt.Errorf("GetPermissionsForUserInMeeting: %w", err)
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

// HasAllPerms does ...
func (p *Permissions) HasAllPerms(permissions ...string) (bool, string) {
	if p.isSuperadmin {
		return true, ""
	}
	for _, perm := range permissions {
		if !p.permissions[perm] {
			return false, perm
		}
	}
	return true, ""
}

// GetInt does ...
func GetId(data definitions.FqfieldData, property definitions.Field) (definitions.Id, error) {
	if val, ok := data[property]; ok {
		var value int
		if err := json.Unmarshal([]byte(val), &value); nil != err {
			return 0, NotAllowedf("'%s' is not an int", property)
		}
		if err := definitions.IsValidId(value); err != nil {
			return 0, NotAllowed(err.Error())
		}
		return value, nil
	}

	return 0, NotAllowedf("'%s' is not in data", property)
}

func GetFqid(data definitions.FqfieldData, property definitions.Field) (definitions.Fqid, error) {
	if val, ok := data[property]; ok {
		var value string
		if err := json.Unmarshal([]byte(val), &value); nil != err {
			return "", NotAllowedf("'%s' is not a string", property)
		}
		if err := definitions.IsValidFqid(value); err != nil {
			return "", NotAllowed(err.Error())
		}
		return value, nil
	}

	return "", NotAllowedf("'%s' is not in data", property)
}

// GetMeetingIDFromModel does ...
func GetMeetingIDFromModel(FQID definitions.Fqid, dp dataprovider.DataProvider) (int, error) {
	id, err := dp.GetInt(FQID + "/meeting_id")
	if err != nil {
		err = fmt.Errorf("GetMeetingIDFromModel: %w", err)
	}
	return id, err
}
