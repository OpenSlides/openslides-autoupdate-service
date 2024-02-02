package slide

import (
	"context"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/datastore"
)

// ListOfSpeaker renders current list of speaker slide.
func Game(store *projector.SlideStore) {
	store.RegisterSliderFunc("game", func(ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection) (encoded []byte, err error) {
		return nil, nil
	})
}
