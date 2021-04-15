package slide

import (
	"context"

	"github.com/openslides/openslides-autoupdate-service/internal/projector"
)

// Topic renders the topic slide.
func Topic(store *projector.SlideStore) {
	store.AddFunc("topic", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		return []byte(`"TODO"`), nil, nil
	})
}
