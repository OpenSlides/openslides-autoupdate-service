package slide

import (
	"context"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/models"
)

// AgendaItem renders the agenda_item slide.
func AgendaItem(store *projector.SlideStore) {
	store.AddFunc("agenda_item", func(ctx context.Context, ds projector.Datastore, p7on *models.Projection) (encoded []byte, keys []string, err error) {
		return []byte(`"TODO"`), nil, nil
	})
}

// AgendaItemList renders the agenda_item_list slide.
func AgendaItemList(store *projector.SlideStore) {
	store.AddFunc("agenda_item_list", func(ctx context.Context, ds projector.Datastore, p7on *models.Projection) (encoded []byte, keys []string, err error) {
		return []byte(`"TODO"`), nil, nil
	})
}
