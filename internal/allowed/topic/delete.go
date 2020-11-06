package topic

import (
	"strconv"

	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
)

var deleteFields = allowed.MakeSet([]string{"id"})

func Delete(params *allowed.IsAllowedParams) (map[string]interface{}, error) {
	if err := allowed.ValidateFields(params.Data, deleteFields); err != nil {
		return nil, err
	}

	exists, err := allowed.DoesUserExists(params.UserID, params.DataProvider)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, allowed.NotAllowed("The user with id " + strconv.Itoa(params.UserID) + " does not exist!")
	}

	superadmin, err := allowed.HasUserSuperadminRole(params.UserID, params.DataProvider)
	if err != nil {
		return nil, err
	}
	if superadmin {
		return nil, nil
	}

	id, err := allowed.GetInt(params.Data, "id")
	if err != nil {
		return nil, err
	}

	meetingID, err := allowed.GetMeetingIDFromModel("topic/"+strconv.Itoa(id), params.DataProvider)
	if err != nil {
		return nil, err
	}

	canSeeMeeting, err := allowed.CanUserSeeMeeting(params.UserID, meetingID, params.DataProvider)
	if err != nil {
		return nil, err
	}
	if !canSeeMeeting {
		return nil, allowed.NotAllowedf("User %d is not in meeting %d", params.UserID, meetingID)
	}

	perms, err := allowed.GetPermissionsForUserInMeeting(params.UserID, meetingID, params.DataProvider)
	if err != nil {
		return nil, err
	}
	if !perms.HasPerm("agenda.can_manage") {
		return nil, allowed.NotAllowedf("User %d has not agenda.can_manage in meeting %d", params.UserID, meetingID)
	}

	return nil, nil
}
