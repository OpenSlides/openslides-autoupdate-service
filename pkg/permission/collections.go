package permission

import (
	"github.com/OpenSlides/openslides-permission-service/internal/collection"
	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/perm"
)

func openSlidesCollections(edp DataProvider) []perm.Connecter {
	dp := dataprovider.DataProvider{External: edp}

	return []perm.Connecter{
		collection.NewAutogen(dp),

		collection.NewSpeaker(dp),
		collection.NewPersonalNote(dp),
		collection.NewGroup(dp),
		collection.NewSimpleRead(dp, "list_of_speakers", "agenda.can_see_list_of_speakers"),
		collection.NewSimpleRead(dp, "assignment", "assingment.can_see"),
		collection.NewSimpleRead(dp, "assignment_candidate", "assingment.can_see"),
	}
}
