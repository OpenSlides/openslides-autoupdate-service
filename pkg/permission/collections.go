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

		//assignment.NewCandidate(dp),
	}
}
