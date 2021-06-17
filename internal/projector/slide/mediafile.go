package slide

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Mediafile renders the mediafile slide.
func Mediafile(store *projector.SlideStore) {
	store.RegisterSliderFunc("mediafile", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		return []byte(`"TODO"`), nil, nil
	})

	store.RegisterAgendaTitlerFunc("mediafile", func(ctx context.Context, fetch *datastore.Fetcher, fqid string, itemNumber string) (json.RawMessage, error) {
		title := "title of mediafile (TODO)"

		agendatitle := struct {
			Title string `json:"title"`
		}{
			title,
		}

		bs, err := json.Marshal(agendatitle)
		if err != nil {
			return nil, fmt.Errorf("decoding title: %w", err)
		}
		return bs, err
	})
}
