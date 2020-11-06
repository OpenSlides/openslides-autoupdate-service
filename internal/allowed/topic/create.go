package topic

import (
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

func Create(params *allowed.IsAllowedParams) (map[string]interface{}, error) {
	if err := allowed.ValidateFields(params.Data, createFields); err != nil {
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

	meetingID, err := allowed.GetInt(params.Data, "meeting_id")
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
