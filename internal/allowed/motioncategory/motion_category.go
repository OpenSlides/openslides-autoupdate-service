package motioncategory

import (
	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
)

// Create TODO
var Create = allowed.BuildCreate([]string{
	"name",
	"prefix",
	"meeting_id",
	"parent_id",
}, "motions.can_manage")

// Update TODO
var Update = allowed.BuildModify([]string{
	"id",
	"name",
	"prefix",
	"motion_ids",
}, "motion_category", "motions.can_manage")

// Delete TODO
var Delete = allowed.BuildModify([]string{"id"}, "motion_category", "motions.can_manage")

// Sort TODO
var Sort = allowed.BuildModify([]string{"id"}, "motion_category", "motions.can_manage")

// SortMotionsInCategory TODO
var SortMotionsInCategory = allowed.BuildModify([]string{"id"}, "motion_category", "motions.can_manage")

// NumberMotions TODO
var NumberMotions = allowed.BuildModify([]string{"id"}, "motion_category", "motions.can_manage")
