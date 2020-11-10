package allowed

import "strconv"

func BuildCreate(allowedFields []string, permissions ...string) IsAllowed {
	allowedFieldsSet := MakeSet(allowedFields)
	return func(params *IsAllowedParams) (map[string]interface{}, error) {
		if err := ValidateFields(params.Data, allowedFieldsSet); err != nil {
			return nil, err
		}

		exists, err := DoesUserExists(params.UserID, params.DataProvider)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, NotAllowedf("The user with id %d does not exist", params.UserID)
		}

		superadmin, err := HasUserSuperadminRole(params.UserID, params.DataProvider)
		if err != nil {
			return nil, err
		}
		if superadmin {
			return nil, nil
		}

		meetingID, err := GetInt(params.Data, "meeting_id")
		if err != nil {
			return nil, err
		}

		committeeID, err := GetCommitteeIdFromMeeting(meetingID, params.DataProvider)
		if err != nil {
			return nil, err
		}
		committeeManager, err := IsUserCommitteeManager(params.UserID, committeeID, params.DataProvider)
		if err != nil {
			return nil, err
		}
		if committeeManager {
			return nil, nil
		}

		canSeeMeeting, err := CanUserSeeMeeting(params.UserID, meetingID, params.DataProvider)
		if err != nil {
			return nil, err
		}
		if !canSeeMeeting {
			return nil, NotAllowedf("User %d is not in meeting %d", params.UserID, meetingID)
		}

		perms, err := GetPermissionsForUserInMeeting(params.UserID, meetingID, params.DataProvider)
		if err != nil {
			return nil, err
		}
		hasPerms, missingPerm := perms.HasAllPerms(permissions...)
		if !hasPerms {
			return nil, NotAllowedf("User %d has not %s in meeting %d", params.UserID, missingPerm, meetingID)
		}

		return nil, nil
	}
}

func BuildModify(allowedFields []string, collection string, permissions ...string) IsAllowed {
	allowedFieldsSet := MakeSet(allowedFields)
	return func(params *IsAllowedParams) (map[string]interface{}, error) {
		if err := ValidateFields(params.Data, allowedFieldsSet); err != nil {
			return nil, err
		}

		exists, err := DoesUserExists(params.UserID, params.DataProvider)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, NotAllowedf("The user with id %d does not exist", params.UserID)
		}

		superadmin, err := HasUserSuperadminRole(params.UserID, params.DataProvider)
		if err != nil {
			return nil, err
		}
		if superadmin {
			return nil, nil
		}

		id, err := GetInt(params.Data, "id")
		if err != nil {
			return nil, err
		}

		exists, err = DoesModelExists(collection, id, params.DataProvider)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, NotAllowedf("The %s with id %d does not exist", collection, id)
		}

		meetingID, err := GetMeetingIDFromModel(collection+"/"+strconv.Itoa(id), params.DataProvider)
		if err != nil {
			return nil, err
		}

		committeeID, err := GetCommitteeIdFromMeeting(meetingID, params.DataProvider)
		if err != nil {
			return nil, err
		}
		committeeManager, err := IsUserCommitteeManager(params.UserID, committeeID, params.DataProvider)
		if err != nil {
			return nil, err
		}
		if committeeManager {
			return nil, nil
		}

		canSeeMeeting, err := CanUserSeeMeeting(params.UserID, meetingID, params.DataProvider)
		if err != nil {
			return nil, err
		}
		if !canSeeMeeting {
			return nil, NotAllowedf("User %d is not in meeting %d", params.UserID, meetingID)
		}

		perms, err := GetPermissionsForUserInMeeting(params.UserID, meetingID, params.DataProvider)
		if err != nil {
			return nil, err
		}
		hasPerms, missingPerm := perms.HasAllPerms(permissions...)
		if !hasPerms {
			return nil, NotAllowedf("User %d has not %s in meeting %d", params.UserID, missingPerm, meetingID)
		}

		return nil, nil
	}
}
