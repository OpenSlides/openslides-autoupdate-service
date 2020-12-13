package motioncommentsection

import (
	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
)

// Create TODO
var Create = allowed.BuildCreate([]string{
	"name",
	"meeting_id",
	"read_group_ids",
	"write_group_ids",
}, "motions.can_manage")

// Update TODO
var Update = allowed.BuildModify([]string{
	"id",
	"name",
	"read_group_ids",
	"write_group_ids",
}, "motion_comment_section", "motions.can_manage")

// Delete TODO
var Delete = allowed.BuildModify([]string{"id"}, "motion_comment_section", "motions.can_manage")

// TODO:
// var Sort =
