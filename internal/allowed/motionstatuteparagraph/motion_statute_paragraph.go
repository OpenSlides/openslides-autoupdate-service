package motionstatuteparagraph

import (
	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
)

// Create TODO
var Create = allowed.BuildCreate([]string{
	"title",
	"text",
	"meeting_id",
}, "motions.can_manage")

// Update TODO
var Update = allowed.BuildModify([]string{
	"id",
	"title",
	"text",
}, "motion_statute_paragraph", "motions.can_manage")

// Delete TODO
var Delete = allowed.BuildModify([]string{"id"}, "motion_statute_paragraph", "motions.can_manage")

// TODO:
// var Sort =
