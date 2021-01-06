package permission

import (
	"github.com/OpenSlides/openslides-permission-service/internal/collection"
	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/perm"
)

func openSlidesCollections(dp dataprovider.DataProvider) []perm.Connecter {
	return []perm.Connecter{
		collection.NewAutogen(dp),

		collection.NewAgendaItem(dp),
		collection.ListOfSpeaker(dp),

		collection.Assignment(dp),

		collection.Mediafile(dp),

		collection.NewMotion(dp),
		collection.Poll(dp),

		collection.NewPersonalNote(dp),
		collection.NewGroup(dp),

		collection.ReadPerm(dp, "assingment.can_see", "assignment", "assignment_candidate"),
		collection.ReadInMeeting(dp, "tag", "meeting"),
		collection.ReadPerm(dp, "agenda.can_see", "topic"),
		collection.ReadPerm(
			dp,
			"meeting.can_see_projector",
			"projector",
			"projection",
			"projectiondefault",
			"projector_message",
			"projector_countdown",
		),
		collection.ReadPerm(
			dp,
			"motion.can_see",
			"motion_workflow",
			"motion_category",
			"motion_state",
			"motion_statute_paragraph",
		),
	}
}
