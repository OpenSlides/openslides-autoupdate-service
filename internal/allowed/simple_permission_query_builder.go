package allowed

import (
	"github.com/OpenSlides/openslides-permission-service/internal/definitions"
)

// CheckUser returns an error, if it is not allowed due to invalid data If the
// data is valid and the first return value is true, the user is a superadmin
// and it is allowed.
func CheckUser(params *IsAllowedParams) (bool, error) {
	exists, err := DoesUserExists(params.UserID, params.DataProvider)
	if err != nil {
		return false, err
	}
	if !exists {
		return false, NotAllowedf("The user with id %d does not exist", params.UserID)
	}

	superadmin, err := HasUserSuperadminRole(params.UserID, params.DataProvider)
	if err != nil {
		return false, err
	}
	return superadmin, nil
}

// CheckCommitteeMeetingPermissions TODO
func CheckCommitteeMeetingPermissions(params *IsAllowedParams, meetingID int, permissions ...string) error {
	committeeID, err := GetCommitteeIDFromMeeting(meetingID, params.DataProvider)
	if err != nil {
		return err
	}
	committeeManager, err := IsUserCommitteeManager(params.UserID, committeeID, params.DataProvider)
	if err != nil {
		return err
	}
	if committeeManager {
		return nil
	}

	canSeeMeeting, err := CanUserSeeMeeting(params.UserID, meetingID, params.DataProvider)
	if err != nil {
		return err
	}
	if !canSeeMeeting {
		return NotAllowedf("User %d is not in meeting %d", params.UserID, meetingID)
	}

	perms, err := GetPermissionsForUserInMeeting(params.UserID, meetingID, params.DataProvider)
	if err != nil {
		return err
	}
	hasPerms, missingPerm := perms.HasAllPerms(permissions...)
	if !hasPerms {
		return NotAllowedf("User %d has not %s in meeting %d", params.UserID, missingPerm, meetingID)
	}

	return nil
}

// BuildCreate TODO
func BuildCreate(allowedFields []string, permissions ...string) IsAllowed {
	allowedFieldsSet := MakeSet(allowedFields)
	return func(params *IsAllowedParams) (map[string]interface{}, error) {
		if err := ValidateFields(params.Data, allowedFieldsSet); err != nil {
			return nil, err
		}

		allowed, err := CheckUser(params)
		if err != nil {
			return nil, err
		}
		if allowed {
			return nil, nil
		}

		meetingID, err := GetID(params.Data, "meeting_id")
		if err != nil {
			return nil, err
		}
		exists, err := DoesModelExists(definitions.FqidFromCollectionAndID("meeting", meetingID), params.DataProvider)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, NotAllowedf("The meeting with id %d does not exist", meetingID)
		}

		err = CheckCommitteeMeetingPermissions(params, meetingID, permissions...)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}

// BuildCreateThroughID TODO
func BuildCreateThroughID(allowedFields []string, throughCollection definitions.Collection, throughField definitions.Field, permissions ...string) IsAllowed {
	allowedFieldsSet := MakeSet(allowedFields)
	return func(params *IsAllowedParams) (map[string]interface{}, error) {
		if err := ValidateFields(params.Data, allowedFieldsSet); err != nil {
			return nil, err
		}

		allowed, err := CheckUser(params)
		if err != nil {
			return nil, err
		}
		if allowed {
			return nil, nil
		}

		throughID, err := GetID(params.Data, throughField)
		if err != nil {
			return nil, err
		}
		exists, err := DoesModelExists(definitions.FqidFromCollectionAndID(throughCollection, throughID), params.DataProvider)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, NotAllowedf("The %s with id %d (field %s) does not exist", throughCollection, throughID, throughField)
		}

		meetingID, err := GetMeetingIDFromModel(definitions.FqidFromCollectionAndID(throughCollection, throughID), params.DataProvider)
		if err != nil {
			return nil, err
		}

		err = CheckCommitteeMeetingPermissions(params, meetingID, permissions...)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}

// BuildCreateThroughFqid TODO
func BuildCreateThroughFqid(allowedFields []string, throughField definitions.Field, permissions ...string) IsAllowed {
	allowedFieldsSet := MakeSet(allowedFields)
	return func(params *IsAllowedParams) (map[string]interface{}, error) {
		if err := ValidateFields(params.Data, allowedFieldsSet); err != nil {
			return nil, err
		}

		allowed, err := CheckUser(params)
		if err != nil {
			return nil, err
		}
		if allowed {
			return nil, nil
		}

		throughFqid, err := GetFqid(params.Data, throughField)
		if err != nil {
			return nil, err
		}
		exists, err := DoesModelExists(throughFqid, params.DataProvider)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, NotAllowedf("%s (field %s) does not exist", throughFqid, throughField)
		}

		meetingID, err := GetMeetingIDFromModel(throughFqid, params.DataProvider)
		if err != nil {
			return nil, err
		}

		err = CheckCommitteeMeetingPermissions(params, meetingID, permissions...)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}

// BuildModify TODO
func BuildModify(allowedFields []string, collection definitions.Collection, permissions ...string) IsAllowed {
	allowedFieldsSet := MakeSet(allowedFields)
	return func(params *IsAllowedParams) (map[string]interface{}, error) {
		if err := ValidateFields(params.Data, allowedFieldsSet); err != nil {
			return nil, err
		}

		allowed, err := CheckUser(params)
		if err != nil {
			return nil, err
		}
		if allowed {
			return nil, nil
		}

		id, err := GetID(params.Data, "id")
		if err != nil {
			return nil, err
		}
		exists, err := DoesModelExists(definitions.FqidFromCollectionAndID(collection, id), params.DataProvider)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, NotAllowedf("The %s with id %d does not exist", collection, id)
		}

		meetingID, err := GetMeetingIDFromModel(definitions.FqidFromCollectionAndID(collection, id), params.DataProvider)
		if err != nil {
			return nil, err
		}

		err = CheckCommitteeMeetingPermissions(params, meetingID, permissions...)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}

// BuildModifyThroughID TODO
func BuildModifyThroughID(allowedFields []string, collection definitions.Collection, throughCollection definitions.Collection, throughField definitions.Field, permissions ...string) IsAllowed {
	allowedFieldsSet := MakeSet(allowedFields)
	return func(params *IsAllowedParams) (map[string]interface{}, error) {
		if err := ValidateFields(params.Data, allowedFieldsSet); err != nil {
			return nil, err
		}

		allowed, err := CheckUser(params)
		if err != nil {
			return nil, err
		}
		if allowed {
			return nil, nil
		}

		throughID, err := GetID(params.Data, throughField)
		if err != nil {
			return nil, err
		}
		throughFqid := definitions.FqidFromCollectionAndID(throughCollection, throughID)
		exists, err := DoesModelExists(throughFqid, params.DataProvider)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, NotAllowedf("The %s with id %d does not exist", throughCollection, throughID)
		}

		meetingID, err := GetMeetingIDFromModel(throughFqid, params.DataProvider)
		if err != nil {
			return nil, err
		}

		err = CheckCommitteeMeetingPermissions(params, meetingID, permissions...)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}
