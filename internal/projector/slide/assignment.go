package slide

import (
	"context"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/models"
)

// Assignment renders the assignment slide.
func Assignment(store *projector.SlideStore) {
	store.AddFunc("assignment", func(ctx context.Context, ds projector.Datastore, p7on *models.Projection) (encoded []byte, keys []string, err error) {
		return []byte(`"TODO"`), nil, nil
	})
}
