package slide

import (
	"context"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Mediafile renders the mediafile slide.
func Mediafile(store *projector.SlideStore) {
	store.RegisterSlideFunc("mediafile", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		return []byte(`"TODO"`), nil, nil
	})
	store.RegisterTitleFunc("mediafile", func(ctx context.Context, fetch *datastore.Fetcher, fqid string, meeting_id int, value map[string]interface{}) (*projector.TitlerFuncResult, error) {
		title := "title of mediafile"
		titleData := projector.TitlerFuncResult{
			Title: &title,
		}
		return &titleData, nil
	})
}
