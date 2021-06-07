package slide

import (
	"context"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Mediafile renders the mediafile slide.
func Mediafile(store *projector.SlideStore) {
	store.AddFunc("mediafile", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		return []byte(`"TODO"`), nil, nil
	})
	store.RegisterTitleFunc("mediafile", func(ctx context.Context, fetch *datastore.Fetcher, fqid string, meeting_id int, value map[string]interface{}) (title map[string]interface{}, err error) {
		return map[string]interface{}{"title": "title of mediafile"}, nil
	})
}
