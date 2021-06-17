package slide

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

type dbTopic struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func topicFromMap(in map[string]json.RawMessage) (*dbTopic, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding topic data: %w", err)
	}

	var t dbTopic
	if err := json.Unmarshal(bs, &t); err != nil {
		return nil, fmt.Errorf("decoding topic data: %w", err)
	}
	return &t, nil
}

// Topic renders the topic slide.
func Topic(store *projector.SlideStore) {
	store.RegisterSliderFunc("topic", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		return []byte(`"TODO"`), nil, nil
	})

	store.RegisterAgendaTitlerFunc("topic", func(ctx context.Context, fetch *datastore.Fetcher, fqid string, itemNumber string) (json.RawMessage, error) {
		data := fetch.Object(ctx, []string{"id", "title"}, fqid)
		topic, err := topicFromMap(data)
		if err != nil {
			return nil, fmt.Errorf("get topic from map: %w", err)
		}

		title := struct {
			Title  string `json:"title"`
			Number string `json:"agenda_item_number"`
		}{
			topic.Title,
			itemNumber,
		}

		bs, err := json.Marshal(title)
		if err != nil {
			return nil, fmt.Errorf("encoding title: %w", err)
		}
		return bs, err
	})
}
