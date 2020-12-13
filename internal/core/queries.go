package core

import (
	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/agenda"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/assignment"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/assignmentcandidate"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/group"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/listofspeakers"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/motionblock"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/motioncategory"
	motion_change_recommendation "github.com/OpenSlides/openslides-permission-service/internal/allowed/motionchangerecommendation"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/motioncommentsection"
	motion_statute_paragraph "github.com/OpenSlides/openslides-permission-service/internal/allowed/motionstatuteparagraph"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/motionworkflow"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/tag"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/topic"
)

// Queries is a list of all possible queries.
var Queries = map[string]allowed.IsAllowed{
	"agenda_item.update": agenda.Update,
	"agenda_item.delete": agenda.Delete,
	// TODO: create, assign, sort

	"assignment.create": assignment.Create,
	"assignment.update": assignment.Update,
	"assignment.delete": assignment.Delete,

	"assignment_candidate.create": assignmentcandidate.Create,
	"assignment_candidate.sort":   assignmentcandidate.Sort,
	"assignment_candidate.delete": assignmentcandidate.Delete,

	// TODO: assignment_poll
	// TODO: committee

	"group.create":         group.Create,
	"group.update":         group.Update,
	"group.delete":         group.Delete,
	"group.set_permission": group.SetPermission,

	"list_of_speakers.update":              listofspeakers.Update,
	"list_of_speakers.delete_all_speakers": listofspeakers.DeleteAllSpeakers,
	"list_of_speakers.re_add_last":         listofspeakers.ReAddLast,

	// TODO: mediafile
	// TODO: meeting
	// TODO: motion

	"motion_block.create": motionblock.Create,
	"motion_block.update": motionblock.Update,
	"motion_block.delete": motionblock.Delete,

	"motion_category.create":                   motioncategory.Create,
	"motion_category.update":                   motioncategory.Update,
	"motion_category.delete":                   motioncategory.Delete,
	"motion_category.sort":                     motioncategory.Sort,
	"motion_category.sort_motions_in_category": motioncategory.SortMotionsInCategory,
	"motion_category.number_motions":           motioncategory.NumberMotions,

	"motion_change_recommendation.create": motion_change_recommendation.Create,
	"motion_change_recommendation.update": motion_change_recommendation.Update,
	"motion_change_recommendation.delete": motion_change_recommendation.Delete,

	// TODO: motion_comment

	"motion_comment_section.create": motioncommentsection.Create,
	"motion_comment_section.update": motioncommentsection.Update,
	"motion_comment_section.delete": motioncommentsection.Delete,
	// TODO: sort

	// TODO: motion_poll
	// TODO: motion_state

	"motion_statute_paragraph.create": motion_statute_paragraph.Create,
	"motion_statute_paragraph.update": motion_statute_paragraph.Update,
	"motion_statute_paragraph.delete": motion_statute_paragraph.Delete,
	// TODO: sort

	// TODO: motion_submitter

	"motion_workflow.create": motionworkflow.Create,
	"motion_workflow.update": motionworkflow.Update,
	"motion_workflow.delete": motionworkflow.Delete,

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
