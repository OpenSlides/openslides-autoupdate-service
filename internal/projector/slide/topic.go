package slide

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

type dbTopic struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Text         string `json:"text"`
	AgendaItemID int    `json:"agenda_item_id"`
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
		fetch := datastore.NewFetcher(ds)
		defer func() {
			if err == nil {
				err = fetch.Error()
			}
		}()

		data := fetch.Object(
			ctx,
			[]string{
				"id",
				"title",
				"text",
				"agenda_item_id",
			},
			p7on.ContentObjectID,
		)

		topic, err := topicFromMap(data)
		if err != nil {
			return nil, nil, fmt.Errorf("get topic: %w", err)
		}

		var itemNumber string
		if topic.AgendaItemID > 0 {
			itemNumber = fetch.String(ctx, "agenda_item/%d/item_number", topic.AgendaItemID)
		}
		out := struct {
			Title            string `json:"title"`
			Text             string `json:"text"`
			AgendaItemNumber string `json:"item_number"`
		}{
			Title:            topic.Title,
			Text:             topic.Text,
			AgendaItemNumber: itemNumber,
		}

		responseValue, err := json.Marshal(out)
		if err != nil {
			return nil, nil, fmt.Errorf("encoding response slide topic: %w", err)
		}
		return responseValue, fetch.Keys(), err
	})

	store.RegisterGetTitleInformationFunc("topic", func(ctx context.Context, fetch *datastore.Fetcher, fqid string, itemNumber string, meetingID int) (json.RawMessage, error) {
		data := fetch.Object(ctx, []string{"id", "title", "agenda_item_id"}, fqid)
		topic, err := topicFromMap(data)
		if err != nil {
			return nil, fmt.Errorf("get topic from map: %w", err)
		}

		if itemNumber == "" && topic.AgendaItemID > 0 {
			itemNumber = fetch.String(ctx, "agenda_item/%d/item_number", topic.AgendaItemID)
		}
		title := struct {
			Collection       string `json:"collection"`
			ContentObjectID  string `json:"content_object_id"`
			Title            string `json:"title"`
			AgendaItemNumber string `json:"agenda_item_number"`
		}{
			"topic",
			fqid,
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
