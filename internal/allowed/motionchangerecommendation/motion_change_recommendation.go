package motionchangerecommendation

import (
	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
)

// Create TODO
var Create = allowed.BuildCreate([]string{
	"line_from",
	"line_to",
	"text",
	"motion_id",
	"rejected",
	"internal",
	"type",
	"other_description",
}, "motions.can_manage")

// Update TODO
var Update = allowed.BuildModify([]string{
	"id",
	"text",
	"rejected",
	"internal",
	"type",
	"other_description",
}, "motion_change_recommendation", "motions.can_manage")

// Delete TODO
var Delete = allowed.BuildModify([]string{"id"}, "motion_change_recommendation", "motions.can_manage")
