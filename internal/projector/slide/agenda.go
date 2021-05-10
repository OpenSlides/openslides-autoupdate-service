package slide

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

type dbAgendaItem struct {
	TestField       string `json:"item_$_num" replacement:"MeetingID"`
	ItemNumber      string `json:"item_number"`
	ContentObjectID string `json:"content_object_id"`
	MeetingID       int    //`json:"meeting_id"`
	IsHidden        bool   `json:"is_hidden"`
	IsInternal      bool   `json:"is_internal"`
	Depth           int    `json:"level"`
}

type dbAgendaItemList struct {
	AgendaItemIds []int `json:"agenda_item_ids"`
}

// AgendaItem renders the agenda_item slide.
func AgendaItem(store *projector.SlideStore) {
	store.AddFunc("agenda_item", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		var agendaItem dbAgendaItem
		keys, err = datastore.Object(ctx, ds, p7on.ContentObjectID, &agendaItem)
		if err != nil {
			return nil, nil, fmt.Errorf("getting user object: %w", err)
		}

		//return []byte(fmt.Sprintf(`{"title_information":"%s","depth":"%d"}`, get_title_information(agendaItem.ContentObjectID, agendaItem.MeetingID), agendaItem.Depth)), keys, nil
		return []byte(fmt.Sprintf(`{"ContentObjectID": "%s", "MeetingID": "%d", "Depth": "%d"}`, agendaItem.ContentObjectID, agendaItem.MeetingID, agendaItem.Depth)), keys, nil

	})
}

// AgendaItemList renders the agenda_item_list slide.
func AgendaItemList(store *projector.SlideStore) {
	store.AddFunc("agenda_item_list", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		var agendaItemList dbAgendaItemList
		keys, err = datastore.Object(ctx, ds, p7on.ContentObjectID, &agendaItemList)
		if err != nil {
			return nil, nil, fmt.Errorf("getting user object: %w", err)
		}

		agendaItem := dbAgendaItem{MeetingID: p7on.MeetingID}
		records, keys, err := datastore.Objects(ctx, ds, "agenda_item", agendaItemList.AgendaItemIds, agendaItem)
		var _ = records
		fieldNames, _, err := datastore.GetStructJsonNames(agendaItem)
		if err != nil {
			return nil, nil, slidesError{err.Error(), "agenda_item", p7on.ID, p7on.Type, p7on.ContentObjectID, p7on.MeetingID}
		}

		var aiKeys []string
		for _, aiid := range agendaItemList.AgendaItemIds {
			for _, fieldName := range fieldNames {
				aiKeys = append(aiKeys, fmt.Sprintf("agenda_item/%d/%s", aiid, fieldName))
			}
		}
		dbValues, err := ds.Get(ctx, aiKeys...)
		if err != nil {
			return nil, nil, fmt.Errorf("fetching data: %w", err)
		}

		for nr, x := range dbValues {
			fmt.Printf("%15s: %s\n", aiKeys[nr], string(x))
		}
		return []byte(`"AgendaItemList"`), keys, nil
	})
}
