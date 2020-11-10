package topic

import (
	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
)

var Create = allowed.BuildCreate([]string{
	"title",
	"meeting_id",

	"text",
	"attachment_ids",
	"tag_ids",

	"agenda_type",
	"agenda_parent_id",
	"agenda_comment",
	"agenda_duration",
	"agenda_weight",
}, "agenda.can_manage")
