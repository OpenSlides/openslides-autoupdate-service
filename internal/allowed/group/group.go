package group

import (
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
)

var allPermissions = allowed.MakeSet([]string{
	"agenda.can_be_speaker",
	"agenda.can_manage",
	"agenda.can_manage_list_of_speakers",
	"agenda.can_see",
	"agenda.can_see_internal_items",
	"agenda.can_see_list_of_speakers",
	"assignments.can_manage",
	"assignments.can_nominate_other",
	"assignments.can_nominate_self",
	"assignments.can_see",
	"core.can_manage_config",
	"core.can_manage_logos_and_fonts",
	"core.can_manage_projector",
	"core.can_manage_tags",
	"core.can_see_frontpage",
	"core.can_see_history",
	"core.can_see_projector",
	"core.can_see_autopilot",
	"mediafiles.can_manage",
	"mediafiles.can_see",
	"motions.can_create",
	"motions.can_create_amendments",
	"motions.can_manage",
	"motions.can_manage_metadata",
	"motions.can_see",
	"motions.can_see_internal",
	"motions.can_support",
	"users.can_change_password",
	"users.can_manage",
	"users.can_see_extra_data",
	"users.can_see_name",
})

var preCreate = allowed.BuildCreate([]string{
	"name",
	"meeting_id",
	"permissions",
}, "users.can_manage")

func Create(params *allowed.IsAllowedParams) (map[string]interface{}, error) {
	addition, err := preCreate(params)

	if err != nil {
		return addition, fmt.Errorf("preCreate failed: %w", err)
	}

	// validate permissions
	var permissions []string
	if val, ok := params.Data["permissions"]; ok {
		err := json.Unmarshal([]byte(val), &permissions)

		if err != nil {
			return nil, allowed.NotAllowed("'permissions' is not a string array")
		}
	} else {
		permissions = []string{}
	}

	for _, perm := range permissions {
		if !allPermissions[perm] {
			return nil, allowed.NotAllowedf("Permission '%s' in permissions is not valid", perm)
		}
	}

	return addition, nil
}

var Update = allowed.BuildModify([]string{"id",
	"name",
}, "group", "users.can_manage")

var Delete = allowed.BuildModify([]string{"id"}, "group", "users.can_manage")

var preSetPermission = allowed.BuildModify([]string{"id",
	"permission",
	"set",
}, "group", "users.can_manage")

func SetPermission(params *allowed.IsAllowedParams) (map[string]interface{}, error) {
	addition, err := preSetPermission(params)

	if err != nil {
		return addition, fmt.Errorf("preSetPermission failed: %w", err)
	}

	// validate permission
	var permission string
	if val, ok := params.Data["permission"]; ok {
		err := json.Unmarshal([]byte(val), &permission)

		if err != nil {
			return nil, allowed.NotAllowed("'permission' is not a string")
		}
	} else {
		return nil, allowed.NotAllowed("'permission' must be given")
	}

	if !allPermissions[permission] {
		return nil, allowed.NotAllowedf("Permission '%s' is not valid", permission)
	}

	return addition, nil
}
