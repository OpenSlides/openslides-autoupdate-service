package tag

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
}, "tag", "motions.can_manage")

var Delete = allowed.BuildModify([]string{"id"}, "tag", "motions.can_manage")
