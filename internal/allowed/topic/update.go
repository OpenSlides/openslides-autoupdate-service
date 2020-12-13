package topic

import (
	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
)

// Update TODO
var Update = allowed.BuildModify([]string{
	"id",
	"title",
	"text",
	"attachment_ids",
	"tag_ids",
}, "topic", "agenda.can_manage")
