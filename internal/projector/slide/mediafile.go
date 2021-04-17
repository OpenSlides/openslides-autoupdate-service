package slide

import (
	"context"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
)

// Mediafile renders the mediafile slide.
func Mediafile(store *projector.SlideStore) {
	store.AddFunc("mediafile", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		return []byte(`"TODO"`), nil, nil
	})
}
