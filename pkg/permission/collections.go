package permission

import (
	"github.com/OpenSlides/openslides-permission-service/internal/collection"
	"github.com/OpenSlides/openslides-permission-service/internal/collection/agenda"
	"github.com/OpenSlides/openslides-permission-service/internal/collection/assignment"
	"github.com/OpenSlides/openslides-permission-service/internal/collection/autogen"
	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
)

func openSlidesCollections(edp DataProvider) []collection.Connecter {
	dp := dataprovider.DataProvider{External: edp}

	return []collection.Connecter{
		autogen.NewAutogen(dp),

		agenda.NewSpeaker(dp),

		// TODO: Remove unneeded collections.

		//collection.NewGeneric(dp, "agenda_item", "agenda.can_see", "agenda.can_manage"),
		//agenda.NewListOfSpeaker(dp),
		//collection.NewGeneric(dp, "topic", "agenda.can_see", "agenda.can_manage"),

		//collection.NewGeneric(dp, "assignment", "assignments.can_see", "assignments.can_manage"),
		assignment.NewCandidate(dp),

		// collection.NewGeneric(dp, "group", "users.can_see", "users.can_manage", collection.WithManageRoute(
		// 	"set_permission",
		// )),

		//collection.NewGeneric(dp, "motion_block", "motion.can_see", "motion.can_manage"),
		// collection.NewGeneric(dp, "motion_category", "motion.can_see", "motion.can_manage", collection.WithManageRoute(
		// 	"sort",
		// 	"sort_motions_in_category",
		// 	"number_motions",
		// )),
		// collection.NewGeneric(dp, "motion_change_recommendation", "motion.can_see", "motion.can_manage"),
		// collection.NewGeneric(dp, "motion_comment_section", "motion.can_see", "motion.can_manage"),
		// collection.NewGeneric(dp, "motion_statute_paragraph", "motion.can_see", "motion.can_manage"),
		// collection.NewGeneric(dp, "motion_workflow", "motion.can_see", "motion.can_manage"),
	}
}
