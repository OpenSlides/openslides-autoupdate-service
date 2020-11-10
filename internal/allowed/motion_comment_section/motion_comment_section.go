package motion_comment_section

import (
	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
)

var Create = allowed.BuildCreate([]string{
	"name",
	"meeting_id",
	"read_group_ids",
	"write_group_ids",
}, "motions.can_manage")

var Update = allowed.BuildModify([]string{
	"id",
	"name",
	"read_group_ids",
	"write_group_ids",
}, "motion_comment_section", "motions.can_manage")

var Delete = allowed.BuildModify([]string{"id"}, "motion_comment_section", "motions.can_manage")

// TODO:
// var Sort =
