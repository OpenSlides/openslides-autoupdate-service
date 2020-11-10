package motion_category

import (
	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
)

var Create = allowed.BuildCreate([]string{
	"name",
	"prefix",
	"meeting_id",
	"parent_id",
}, "motions.can_manage")

var Update = allowed.BuildModify([]string{
	"id",
	"name",
	"prefix",
	"motion_ids",
}, "motion_category", "motions.can_manage")

var Delete = allowed.BuildModify([]string{"id"}, "motion_category", "motions.can_manage")

var Sort = allowed.BuildModify([]string{"id"}, "motion_category", "motions.can_manage")

var SortMotionsInCategory = allowed.BuildModify([]string{"id"}, "motion_category", "motions.can_manage")

var NumberMotions = allowed.BuildModify([]string{"id"}, "motion_category", "motions.can_manage")
