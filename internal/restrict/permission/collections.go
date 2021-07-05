package permission

import (
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/dataprovider"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/perm"
)

func openSlidesCollections(dp dataprovider.DataProvider) []perm.Connecter {
	return []perm.Connecter{
		collection.AgendaItem(dp),
		collection.ListOfSpeaker(dp),
		collection.Mediafile(dp),
		collection.Motion(dp),
		collection.Poll(dp),
		collection.PersonalNote(dp),
		collection.User(dp),
		collection.Meeting(dp),
		collection.Committee(dp),

		collection.Public(dp, "resource", "organization"),
		collection.LoggedIn(dp, "organization_tag"),
		collection.ReadInMeeting(dp, "tag", "group"),
		collection.ReadPerm(dp, perm.AssignmentCanSee, "assignment", "assignment_candidate"),
		collection.ReadPerm(dp, perm.AgendaItemCanSee, "topic"),
		collection.ReadPerm(
			dp,
			perm.ProjectorCanSee,
			"projector",
			"projection",
			"projectiondefault",
			"projector_message",
			"projector_countdown",
		),
		collection.ReadPerm(
			dp,
			perm.MotionCanSee,
			"motion_workflow",
			"motion_category",
			"motion_state",
			"motion_statute_paragraph",
		),
	}
}
