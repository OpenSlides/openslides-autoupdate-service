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

func Create(ctx *allowed.IsAllowedContext) (bool, map[string]interface{}, error) {
	if err := allowed.ValidateFields(ctx.Data, createFields); err != nil {
		return false, nil, err
	}

	if !allowed.DoesUserExists(ctx.UserId, ctx) {
		return false, nil, fmt.Errorf("The user with id " + strconv.Itoa(ctx.UserId) + " does not exist!")
	}

	superadmin, err := allowed.HasUserSuperadminRole(ctx.UserId, ctx)
	if err != nil {
		return false, nil, err
	}
	if superadmin {
		return true, nil, nil
	}

	meetingId, err := allowed.GetInt(ctx.Data, "meeting_id")
	if err != nil {
		return false, nil, err
	}

	canSeeMeeting, err := allowed.CanUserSeeMeeting(ctx.UserId, meetingId, ctx)
	if err != nil || !canSeeMeeting {
		return false, nil, err
	}

	perms, err := allowed.GetPermissionsForUserInMeeting(ctx.UserId, meetingId, ctx)
	if err != nil {
		return false, nil, err
	}
	return perms.HasPerm("agenda.can_manage"), nil, nil
}
