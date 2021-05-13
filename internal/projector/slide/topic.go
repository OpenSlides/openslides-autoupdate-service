package slide

import (
	"context"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/models"
)

// Topic renders the topic slide.
func Topic(store *projector.SlideStore) {
	store.AddFunc("topic", func(ctx context.Context, ds projector.Datastore, p7on *models.Projection) (encoded []byte, keys []string, err error) {
		return []byte(`"TODO"`), nil, nil
	})
}
