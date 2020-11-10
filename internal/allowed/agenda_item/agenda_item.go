package agenda_item

import "github.com/OpenSlides/openslides-permission-service/internal/allowed"

// TODO: This does not work due to the implicit meeting id via the content object
var Create = allowed.BuildCreate([]string{
	"content_object_id",

	"item_number",
	"parent_id",
	"comment",
	"closed",
	"type",
	"duration",
	"weight",
	"tag_ids",
}, "agenda.can_see_internal_items", "agenda.can_manage")

var Update = allowed.BuildModify([]string{"id",
	"item_number",
	"comment",
	"closed",
	"type",
	"duration",
	"weight",
	"tag_ids"}, "agenda_item", "agenda.can_see_internal_items", "agenda.can_manage")

var Delete = allowed.BuildModify([]string{"id"}, "agenda_item", "agenda.can_see_internal_items", "agenda.can_manage")

// TODO
// var Assign =
// ids: Id[];
// parent_id: Id | null;
// meeting_id: Id;
// needs "agenda.can_see_internal_items", "agenda.can_manage"

// TODO
// var Sort =
// meeting_id: Id;
// tree: TreeIdNode[];
// needs "agenda.can_manage"
