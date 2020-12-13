package motionblock

import (
	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
)

// Create TODO
var Create = allowed.BuildCreate([]string{
	"title",
	"internal",
	"meeting_id",
	"agenda_create",
	"agenda_type",
	"agenda_parent_id",
	"agenda_comment",
	"agenda_duration",
	"agenda_weight",
}, "motions.can_manage")

// Update TODO
var Update = allowed.BuildModify([]string{
	"id",
	"title",
	"internal",
	"motion_ids",
}, "motion_block", "motions.can_manage")

// Delete TODO
var Delete = allowed.BuildModify([]string{"id"}, "motion_block", "motions.can_manage")
