package core

import (
	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/agenda_item"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/assignment"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/assignment_candidate"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/group"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/list_of_speakers"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/motion_block"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/motion_category"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/motion_change_recommendation"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/motion_comment_section"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/motion_statute_paragraph"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/motion_workflow"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/tag"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/topic"
)

// Queries is a list of all possible queries.
var Queries = map[string]allowed.IsAllowed{
	"agenda_item.update": agenda_item.Update,
	"agenda_item.delete": agenda_item.Delete,
	// TODO: create, assign, sort

	"assignment.create": assignment.Create,
	"assignment.update": assignment.Update,
	"assignment.delete": assignment.Delete,

	"assignment_candidate.create": assignment_candidate.Create,
	"assignment_candidate.sort":   assignment_candidate.Sort,
	"assignment_candidate.delete": assignment_candidate.Delete,

	// TODO: assignment_poll
	// TODO: committee

	"group.create":         group.Create,
	"group.update":         group.Update,
	"group.delete":         group.Delete,
	"group.set_permission": group.SetPermission,

	"list_of_speakers.update":              list_of_speakers.Update,
	"list_of_speakers.delete_all_speakers": list_of_speakers.DeleteAllSpeakers,
	"list_of_speakers.re_add_last":         list_of_speakers.ReAddLast,

	// TODO: mediafile
	// TODO: meeting
	// TODO: motion

	"motion_block.create": motion_block.Create,
	"motion_block.update": motion_block.Update,
	"motion_block.delete": motion_block.Delete,

	"motion_category.create":                   motion_category.Create,
	"motion_category.update":                   motion_category.Update,
	"motion_category.delete":                   motion_category.Delete,
	"motion_category.sort":                     motion_category.Sort,
	"motion_category.sort_motions_in_category": motion_category.SortMotionsInCategory,
	"motion_category.number_motions":           motion_category.NumberMotions,

	"motion_change_recommendation.create": motion_change_recommendation.Create,
	"motion_change_recommendation.update": motion_change_recommendation.Update,
	"motion_change_recommendation.delete": motion_change_recommendation.Delete,

	// TODO: motion_comment

	"motion_comment_section.create": motion_comment_section.Create,
	"motion_comment_section.update": motion_comment_section.Update,
	"motion_comment_section.delete": motion_comment_section.Delete,
	// TODO: sort

	// TODO: motion_poll
	// TODO: motion_state

	"motion_statute_paragraph.create": motion_statute_paragraph.Create,
	"motion_statute_paragraph.update": motion_statute_paragraph.Update,
	"motion_statute_paragraph.delete": motion_statute_paragraph.Delete,
	// TODO: sort

	// TODO: motion_submitter

	"motion_workflow.create": motion_workflow.Create,
	"motion_workflow.update": motion_workflow.Update,
	"motion_workflow.delete": motion_workflow.Delete,

	// TODO: personal_note
	// TODO: speaker

	"tag.create": tag.Create,
	"tag.update": tag.Update,
	"tag.delete": tag.Delete,

	"topic.create": topic.Create,
	"topic.update": topic.Update,
	"topic.delete": topic.Delete,

	// TODO: users
}
