package tag

import (
	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
)

// Create TODO
var Create = allowed.BuildCreate([]string{
	"name",
	"meeting_id",
}, "motions.can_manage")

// Update TODO
var Update = allowed.BuildModify([]string{
	"id",
	"name",
}, "tag", "motions.can_manage")

// Delete TODO
var Delete = allowed.BuildModify([]string{"id"}, "tag", "motions.can_manage")
