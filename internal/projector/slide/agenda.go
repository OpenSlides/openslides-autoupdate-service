package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
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
	Weight          int    `json:"weight"`
	ParentID        int    `json:"parent_id"`
}

func agendaItemFromMap(in map[string]json.RawMessage) (*dbAgendaItem, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding agenda item data: %w", err)
	}

	var ai dbAgendaItem
	if err := json.Unmarshal(bs, &ai); err != nil {
		return nil, fmt.Errorf("decoding agenda item data: %w", err)
	}
	return &ai, nil
}

type dbAgendaItemList struct {
	AgendaItemIDs      []int `json:"agenda_item_ids"`
	AgendaShowInternal bool  `json:"agenda_show_internal_items_on_projector"`
}

func agendaItemListFromMap(in map[string]json.RawMessage) (*dbAgendaItemList, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding agenda item list data: %w", err)
	}

	var ail dbAgendaItemList
	if err := json.Unmarshal(bs, &ail); err != nil {
		return nil, fmt.Errorf("decoding agenda item list data: %w", err)
	}
	return &ail, nil
}

type outAgendaItem struct {
	TitleInformation json.RawMessage `json:"title_information"`
	Depth            int             `json:"depth"`
	weight           int
	parent           int
	id               int
}

// AgendaItem renders the agenda_item slide.
func AgendaItem(store *projector.SlideStore) {
	store.RegisterSliderFunc("agenda_item", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
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
				"item_number",
				"content_object_id",
				"meeting_id",
				"is_hidden",
				"is_internal",
				"weight",
				"level",
			},
			p7on.ContentObjectID,
		)

		agendaItem, err := agendaItemFromMap(data)
		if err != nil {
			return nil, nil, fmt.Errorf("get agenda item: %w", err)
		}

		collection := strings.Split(agendaItem.ContentObjectID, "/")[0]
		titler := store.GetTitleInformationFunc(collection)
		if titler == nil {
			return nil, nil, fmt.Errorf("no titler function registered for %s", collection)
		}

		titleInfo, err := titler.GetTitleInformation(ctx, fetch, agendaItem.ContentObjectID, agendaItem.ItemNumber, p7on.MeetingID)
		if err != nil {
			return nil, nil, fmt.Errorf("get title func: %w", err)
		}

		out := outAgendaItem{
			TitleInformation: titleInfo,
			Depth:            agendaItem.Depth,
		}

		responseValue, err := json.Marshal(out)
		if err != nil {
			return nil, nil, fmt.Errorf("encoding response slide agenda item: %w", err)
		}
		return responseValue, fetch.Keys(), err
	})
}

// AgendaItemList renders the agenda_item_list slide.
func AgendaItemList(store *projector.SlideStore) {
	store.RegisterSliderFunc("agenda_item_list", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		fetch := datastore.NewFetcher(ds)
		defer func() {
			if err == nil {
				err = fetch.Error()
			}
		}()

		data := fetch.Object(
			ctx,
			[]string{
				"agenda_item_ids",
				"agenda_show_internal_items_on_projector",
			},
			p7on.ContentObjectID,
		)
		agendaItemList, err := agendaItemListFromMap(data)
		if err != nil {
			return nil, nil, fmt.Errorf("get agenda item list: %w", err)
		}

		var options struct {
			OnlyMainItems bool `json:"only_main_items"`
		}
		if p7on.Options != nil {
			if err := json.Unmarshal(p7on.Options, &options); err != nil {
				return nil, nil, fmt.Errorf("decoding projection options: %w", err)
			}
		}
		var allAgendaItems []outAgendaItem
		for _, aiID := range agendaItemList.AgendaItemIDs {
			data = fetch.Object(
				ctx,
				[]string{
					"id",
					"item_number",
					"content_object_id",
					"meeting_id",
					"is_hidden",
					"is_internal",
					"level",
					"weight",
					"parent_id",
				},
				"agenda_item/%d",
				aiID,
			)
			agendaItem, err := agendaItemFromMap(data)
			if err != nil {
				return nil, nil, fmt.Errorf("get agenda item: %w", err)
			}

			if agendaItem.IsHidden || (agendaItem.IsInternal && !agendaItemList.AgendaShowInternal) {
				continue
			}

			if options.OnlyMainItems && agendaItem.Depth > 0 {
				continue
			}

			collection := strings.Split(agendaItem.ContentObjectID, "/")[0]
			titler := store.GetTitleInformationFunc(collection)
			if titler == nil {
				return nil, nil, fmt.Errorf("no titler function registered for %s", collection)
			}

			titleInfo, err := titler.GetTitleInformation(ctx, fetch, agendaItem.ContentObjectID, agendaItem.ItemNumber, p7on.MeetingID)
			if err != nil {
				return nil, nil, fmt.Errorf("get title func: %w", err)
			}

			allAgendaItems = append(
				allAgendaItems,
				outAgendaItem{
					TitleInformation: titleInfo,
					Depth:            agendaItem.Depth,
					weight:           agendaItem.Weight,
					parent:           agendaItem.ParentID,
					id:               agendaItem.ID,
				},
			)
		}

		sort.SliceStable(allAgendaItems, func(i, j int) bool {
			// sort by parent is not necessary, but helps to understand
			if allAgendaItems[i].parent == allAgendaItems[j].parent {
				if allAgendaItems[i].weight == allAgendaItems[j].weight {
					return allAgendaItems[i].id < allAgendaItems[j].id
				}
				return allAgendaItems[i].weight < allAgendaItems[j].weight
			}
			return allAgendaItems[i].parent < allAgendaItems[j].parent
		})

		out := struct {
			Items []outAgendaItem `json:"items"`
		}{getFlatTree(allAgendaItems)}

		responseValue, err := json.Marshal(out)
		if err != nil {
			return nil, nil, fmt.Errorf("encoding response for slide agenda item list: %w", err)
		}
		return responseValue, fetch.Keys(), err
	})
}

func getFlatTree(allAgendaItems []outAgendaItem) []outAgendaItem {
	children := make(map[int][]int)
	allItemMap := make(map[int]outAgendaItem)
	for _, item := range allAgendaItems {
		children[item.parent] = append(children[item.parent], item.id)
		allItemMap[item.id] = item
	}

	flatTree := make([]outAgendaItem, 0, len(allAgendaItems))
	var buildTree func(itemIDS []int, depth int)
	buildTree = func(itemIDS []int, depth int) {
		for _, itemID := range itemIDS {
			item := allItemMap[itemID]
			item.Depth = depth
			flatTree = append(flatTree, item)
			buildTree(children[itemID], depth+1)
		}
	}
	buildTree(children[0], 0)
	return flatTree
}
