package slide

import (
	"context"

	"github.com/openslides/openslides-autoupdate-service/internal/projector"
)

// AgendaItem renders the agenda_item slide.
func AgendaItem(store *projector.SlideStore) {
	store.AddFunc("agenda_item", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		return []byte(`"TODO"`), nil, nil
	})
}

// AgendaItemList renders the agenda_item_list slide.
func AgendaItemList(store *projector.SlideStore) {
	store.AddFunc("agenda_item_list", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		return []byte(`"TODO"`), nil, nil
	})
}
