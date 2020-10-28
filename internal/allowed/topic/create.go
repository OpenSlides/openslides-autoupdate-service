package topic

import (
	"fmt"
	"strconv"

	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
)

var createFields = allowed.MakeSet([]string{
	"title",
	"meeting_id",

	"text",
	"attachment_ids",
	"tag_ids",

	"agenda_type",
	"agenda_parent_id",
	"agenda_comment",
	"agenda_duration",
	"agenda_weight",
})

func Create(params *allowed.IsAllowedParams) (bool, map[string]interface{}, error) {
	if err := allowed.ValidateFields(params.Data, createFields); err != nil {
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

	meetingId, err := allowed.GetInt(params.Data, "meeting_id")
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
