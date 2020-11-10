package motion_workflow

import (
	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
)

var Create = allowed.BuildCreate([]string{
	"name",
	"meeting_id",
}, "motions.can_manage")

var Update = allowed.BuildModify([]string{
	"id",
	"name",
	"first_state_id",
}, "motion_workflow", "motions.can_manage")

var Delete = allowed.BuildModify([]string{"id"}, "motion_workflow", "motions.can_manage")
