package assignment

import "github.com/OpenSlides/openslides-permission-service/internal/allowed"

var Create = allowed.BuildCreate([]string{
	"title",
	"meeting_id",

	"description",
	"open_posts",
	"phase",
	"default_poll_description",
	"number_poll_candidates",
	"tag_ids",
	"attachment_ids",
}, "assignments.can_manage")

var Update = allowed.BuildModify([]string{"id",
	"title",
	"description",
	"open_posts",
	"phase",
	"default_poll_description",
	"number_poll_candidates",

	"tag_ids",
	"attachment_ids",
}, "assignment", "assignments.can_manage")

var Delete = allowed.BuildModify([]string{"id"}, "assignment", "assignments.can_manage")
