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

	var m dbTopic
	if err := json.Unmarshal(bs, &m); err != nil {
		return nil, fmt.Errorf("decoding topic data: %w", err)
	}
	return &m, nil
}

// Topic renders the topic slide.
func Topic(store *projector.SlideStore) {
	store.RegisterSlideFunc("topic", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		return []byte(`"TODO"`), nil, nil
	})
	store.RegisterTitleFunc("topic", func(ctx context.Context, fetch *datastore.Fetcher, fqid string, meeting_id int, value map[string]interface{}) (*projector.TitlerFuncResult, error) {
		data := fetch.Object(ctx, []string{"id", "title"}, fqid)
		topic, err := topicFromMap(data)
		if err != nil {
			return nil, fmt.Errorf("get topic from map: %w", err)
		}
		agenda_item_number := value["agenda_item_number"].(string)
		titleData := projector.TitlerFuncResult{
			Title:            &topic.Title,
			AgendaItemNumber: &agenda_item_number,
		}
		return &titleData, err
	})
}
