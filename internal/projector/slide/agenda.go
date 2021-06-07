package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

type dbAgendaItem struct {
	ID              int    `json:"id"`
	ItemNumber      string `json:"item_number"`
	ContentObjectID string `json:"content_object_id"`
	MeetingID       int    `json:"meeting_id"`
	IsHidden        bool   `json:"is_hidden"`
	IsInternal      bool   `json:"is_internal"`
	Depth           int    `json:"level"`
}

func agendaItemFromMap(in map[string]json.RawMessage) (*dbAgendaItem, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding agenda item data")
	}

	var ai dbAgendaItem
	if err := json.Unmarshal(bs, &ai); err != nil {
		return nil, fmt.Errorf("decoding agenda item: %w", err)
	}
	return &ai, nil
}

type dbAgendaItemList struct {
	AgendaItemIds      []int `json:"agenda_item_ids"`
	AgendaShowInternal bool  `json:"agenda_show_internal_items_on_projector"`
}

func agendaItemListFromMap(in map[string]json.RawMessage) (*dbAgendaItemList, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding agenda item list data")
	}

	var ail dbAgendaItemList
	if err := json.Unmarshal(bs, &ail); err != nil {
		return nil, fmt.Errorf("decoding agenda item list: %w", err)
	}
	return &ail, nil
}

// AgendaItem renders the agenda_item slide.
func AgendaItem(store *projector.SlideStore) {
	store.AddFunc("agenda_item", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		fetch := datastore.NewFetcher(ds)
		defer func() {
			if err == nil {
				err = fetch.Error()
			}
		}()
		data := fetch.Object(ctx, []string{"id", "item_number", "content_object_id", "meeting_id", "is_hidden", "is_internal", "level"}, p7on.ContentObjectID)
		agendaItem, err := agendaItemFromMap(data)
		if err != nil {
			return nil, nil, err
		}
		collection := strings.Split(agendaItem.ContentObjectID, "/")[0]
		titleFunc := store.GetTitleFunc(collection)
		value := map[string]interface{}{"agenda_item_number": agendaItem.ItemNumber}
		titleInfo, err := titleFunc(ctx, fetch, agendaItem.ContentObjectID, agendaItem.MeetingID, value)
		if err != nil {
			return nil, nil, err
		}
		responseValue, err := json.Marshal(map[string]interface{}{"title_information": titleInfo, "depth": agendaItem.Depth})
		if err != nil {
			return nil, nil, err
		}
		return responseValue, fetch.Keys(), err
	})
}

// AgendaItemList renders the agenda_item_list slide.
func AgendaItemList(store *projector.SlideStore) {
	store.AddFunc("agenda_item_list", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		allAgendaItems := []map[string]interface{}{}
		fetch := datastore.NewFetcher(ds)
		defer func() {
			if err == nil {
				err = fetch.Error()
			}
		}()

		data := fetch.Object(ctx, []string{"agenda_item_ids", "agenda_show_internal_items_on_projector"}, p7on.ContentObjectID)
		agendaItemList, err := agendaItemListFromMap(data)
		if err != nil {
			return nil, nil, err
		}

		for _, aiID := range agendaItemList.AgendaItemIds {
			data = fetch.Object(ctx, []string{"id", "item_number", "content_object_id", "meeting_id", "is_hidden", "is_internal", "level"}, fmt.Sprintf("agenda_item/%d", aiID))
			agendaItem, err := agendaItemFromMap(data)
			if err != nil {
				return nil, nil, err
			}
			if agendaItem.IsHidden || (agendaItem.IsInternal && !agendaItemList.AgendaShowInternal) {
				continue
			}
			if val, ok := p7on.Options["only_main_items"].(bool); ok && val && agendaItem.Depth > 0 {
				continue
			}

			collection := strings.Split(agendaItem.ContentObjectID, "/")[0]
			titleFunc := store.GetTitleFunc(collection)
			value := map[string]interface{}{"agenda_item_number": agendaItem.ItemNumber}
			titleInfo, err := titleFunc(ctx, fetch, agendaItem.ContentObjectID, agendaItem.MeetingID, value)
			if err != nil {
				return nil, nil, err
			}
			allAgendaItems = append(allAgendaItems, map[string]interface{}{"title_information": titleInfo, "depth": agendaItem.Depth})
		}

		responseValue, err := json.Marshal(map[string]interface{}{"items": allAgendaItems})
		if err != nil {
			return nil, nil, err
		}
		return responseValue, fetch.Keys(), err
	})
}
