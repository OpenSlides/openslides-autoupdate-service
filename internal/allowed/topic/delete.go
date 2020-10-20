package topic

import (
	"fmt"
	"strconv"

	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
)

var deleteFields = allowed.MakeSet([]string{"id"})

func Delete(ctx *allowed.IsAllowedContext) (bool, map[string]interface{}, error) {
	if err := allowed.ValidateFields(ctx.Data, deleteFields); err != nil {
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

	id, err := allowed.GetInt(ctx.Data, "id")
	if err != nil {
		return false, nil, err
	}

	meetingId, err := allowed.GetMeetingIdFromModel("topic/"+strconv.Itoa(id), ctx)
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
