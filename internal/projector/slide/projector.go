package slide

import (
	"context"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
)

// ProjectorCountdown renders the projector_countdown slide.
func ProjectorCountdown(store *projector.SlideStore) {
	store.AddFunc("projector_countdown", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		return []byte(`"TODO"`), nil, nil
	})
}

// ProjectorMessage renders the projector_message slide.
func ProjectorMessage(store *projector.SlideStore) {
	store.AddFunc("projector_message", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		return []byte(`"TODO"`), nil, nil
	})
}
