package topic

import (
	"fmt"
	"strconv"

	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
)

var deleteFields = allowed.MakeSet([]string{"id"})

func Delete(params *allowed.IsAllowedParams) (bool, map[string]interface{}, error) {
	if err := allowed.ValidateFields(params.Data, deleteFields); err != nil {
		return false, nil, err
	}

	if !allowed.DoesUserExists(params.UserId, params.DataProvider) {
		return false, nil, fmt.Errorf("The user with id " + strconv.Itoa(params.UserId) + " does not exist!")
	}

	superadmin, err := allowed.HasUserSuperadminRole(params.UserId, params.DataProvider)
	if err != nil {
		return false, nil, err
	}
	if superadmin {
		return true, nil, nil
	}

	id, err := allowed.GetInt(params.Data, "id")
	if err != nil {
		return false, nil, err
	}

	meetingId, err := allowed.GetMeetingIdFromModel("topic/"+strconv.Itoa(id), params.DataProvider)
	if err != nil {
		return false, nil, err
	}

	canSeeMeeting, err := allowed.CanUserSeeMeeting(params.UserId, meetingId, params.DataProvider)
	if err != nil || !canSeeMeeting {
		return false, nil, err
	}

	perms, err := allowed.GetPermissionsForUserInMeeting(params.UserId, meetingId, params.DataProvider)
	if err != nil {
		return false, nil, err
	}
	return perms.HasPerm("agenda.can_manage"), nil, nil
}
