package slide_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/slide"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/dsmock"
	"github.com/stretchr/testify/assert"
)

func TestAgendaItemListAllContentObjectTypes(t *testing.T) {
	s := new(projector.SlideStore)
	slide.AgendaItemList(s)
	slide.Assignment(s)
	slide.Motion(s)
	slide.MotionBlock(s)
	slide.Topic(s)

	ailSlide := s.GetSlider("agenda_item_list")
	assert.NotNilf(t, ailSlide, "Slide with name `agenda_item_list` not found.")

	data := dsmock.YAMLData(`
	meeting:
		1:
			agenda_show_internal_items_on_projector: false
			agenda_item_ids: [1,2,3,4,5,6,7]
	agenda_item:
		1:
			item_number: Ino1.2
			content_object_id: assignment/1
			meeting_id: 1
			is_hidden: false
			is_internal: false
			weight: 8
			level: 0
		2:
			item_number: Ino1.1
			content_object_id: motion/1
			meeting_id: 1
			is_hidden: false
			is_internal: false
			weight: 4
			level: 0
		3:
			item_number: Ino1
			content_object_id: topic/1
			meeting_id: 1
			is_hidden: false
			is_internal: false
			weight: 2
			level: 0
		4:
			item_number: Ino1.1.1
			content_object_id: motion_block/1
			meeting_id: 1
			is_hidden: false
			is_internal: false
			weight: 6
			level: 0
		5:
			item_number: Ino5 misses because of level
			content_object_id: topic/2
			meeting_id: 1
			is_hidden: false
			is_internal: false
			weight: 10
			level: 1
		6:
			item_number: Ino6 misses because of hidden
			content_object_id: topic/3
			meeting_id: 1
			is_hidden: true
			is_internal: false
			weight: 12
			level: 0
		7:
			item_number: Ino7 misses because of internal
			content_object_id: topic/4
			meeting_id: 1
			is_hidden: false
			is_internal: True
			weight: 14
			level: 0
	motion/1:
		title:  motion title 1
		number: motion number 1
	assignment/1/title: assignment title 1
	motion_block/1/title: motion_block title 1
	topic/1/title: topic title 1
	topic/2/title: topic title 2
	topic/3/title: topic title 3
	topic/4/title: topic title 4
    `)

	for _, tt := range []struct {
		name       string
		data       map[string]string
		expect     string
		expectKeys []string
	}{
		{
			"Starter AgendaItemList",
			data,
			`{
				"items": [
					{
						"depth": 0,
						"title_information": {
							"collection": "topic",
							"agenda_item_number": "Ino1",
							"content_object_id": "topic/1",
							"title": "topic title 1"
					    }
					},
					{
						"depth": 0,
						"title_information": {
							"collection": "motion",
							"agenda_item_number": "Ino1.1",
							"content_object_id": "motion/1",
							"number": "motion number 1",
							"title": "motion title 1"
						}
					},
					{
						"depth": 0,
						"title_information": {
							"collection": "motion_block",
							"agenda_item_number": "Ino1.1.1",
							"content_object_id": "motion_block/1",
							"title": "motion_block title 1"
						}
					},
					{
						"depth": 0,
						"title_information": {
							"collection": "assignment",
							"agenda_item_number": "Ino1.2",
							"content_object_id": "assignment/1",
							"title": "assignment title 1"
					    }
					}
				]
			}
			`,
			[]string{
				"meeting/1/agenda_item_ids",
				"meeting/1/agenda_show_internal_items_on_projector",
				"agenda_item/1/id",
				"agenda_item/1/item_number",
				"agenda_item/1/content_object_id",
				"agenda_item/1/meeting_id",
				"agenda_item/1/is_hidden",
				"agenda_item/1/is_internal",
				"agenda_item/1/level",
				"agenda_item/1/weight",
				"agenda_item/1/parent_id",
				"topic/1/id",
				"topic/1/title",
				"topic/1/agenda_item_id",
				"agenda_item/2/id",
				"agenda_item/2/item_number",
				"agenda_item/2/content_object_id",
				"agenda_item/2/meeting_id",
				"agenda_item/2/is_hidden",
				"agenda_item/2/is_internal",
				"agenda_item/2/level",
				"agenda_item/2/weight",
				"agenda_item/2/parent_id",
				"motion/1/id",
				"motion/1/number",
				"motion/1/title",
				"motion/1/agenda_item_id",
				"agenda_item/3/id",
				"agenda_item/3/item_number",
				"agenda_item/3/content_object_id",
				"agenda_item/3/meeting_id",
				"agenda_item/3/is_hidden",
				"agenda_item/3/is_internal",
				"agenda_item/3/level",
				"agenda_item/3/weight",
				"agenda_item/3/parent_id",
				"motion_block/1/id",
				"motion_block/1/title",
				"motion_block/1/agenda_item_id",
				"agenda_item/4/id",
				"agenda_item/4/item_number",
				"agenda_item/4/content_object_id",
				"agenda_item/4/meeting_id",
				"agenda_item/4/is_hidden",
				"agenda_item/4/is_internal",
				"agenda_item/4/level",
				"agenda_item/4/weight",
				"agenda_item/4/parent_id",
				"assignment/1/id",
				"assignment/1/title",
				"assignment/1/agenda_item_id",
				"agenda_item/5/id",
				"agenda_item/5/item_number",
				"agenda_item/5/content_object_id",
				"agenda_item/5/meeting_id",
				"agenda_item/5/is_hidden",
				"agenda_item/5/is_internal",
				"agenda_item/5/level",
				"agenda_item/5/weight",
				"agenda_item/5/parent_id",
				"agenda_item/6/id",
				"agenda_item/6/item_number",
				"agenda_item/6/content_object_id",
				"agenda_item/6/meeting_id",
				"agenda_item/6/is_hidden",
				"agenda_item/6/is_internal",
				"agenda_item/6/level",
				"agenda_item/6/weight",
				"agenda_item/6/parent_id",
				"agenda_item/7/id",
				"agenda_item/7/item_number",
				"agenda_item/7/content_object_id",
				"agenda_item/7/meeting_id",
				"agenda_item/7/is_hidden",
				"agenda_item/7/is_internal",
				"agenda_item/7/level",
				"agenda_item/7/weight",
				"agenda_item/7/parent_id",
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			closed := make(chan struct{})
			defer close(closed)
			ds := dsmock.NewMockDatastore(closed, tt.data)
			fetch := datastore.NewFetcher(ds)

			p7on := &projector.Projection{
				ContentObjectID: "meeting/1",
				Type:            "agenda_item_list",
				MeetingID:       1,
				Options:         []byte(`{"only_main_items":true}`),
			}

			bs, keys, err := ailSlide.Slide(context.Background(), fetch, p7on)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expect, string(bs))
			assert.ElementsMatch(t, tt.expectKeys, keys)
		})
	}
}

// TestAgendaItemListWithDepthItems tests the sorting and delivery
// with weights per level
func TestAgendaItemListWithDepthItems(t *testing.T) {
	s := new(projector.SlideStore)
	slide.AgendaItemList(s)
	slide.Topic(s)

	ailSlide := s.GetSlider("agenda_item_list")
	assert.NotNilf(t, ailSlide, "Slide with name `agenda_item_list` not found.")

	data := dsmock.YAMLData(`
	meeting:
		1:
			agenda_show_internal_items_on_projector: false
			agenda_item_ids: [1, 2, 3 ,4 ,5, 6, 7, 8]
	agenda_item:
		1:
			item_number: Ino1
			content_object_id: topic/1
			meeting_id: 1
			level: 0
			weight: 2
			child_ids: [2, 3]
		2:
			item_number: Ino1.1
			content_object_id: topic/2
			meeting_id: 1
			level: 1
			weight: 3
			parent_id: 1
			child_ids: [4, 5]
		3:
			item_number: Ino1.2
			content_object_id: topic/3
			meeting_id: 1
			level: 1
			weight: 4
			parent_id: 1
			child_ids: []
		4:
			item_number: Ino1.1.1
			content_object_id: topic/4
			meeting_id: 1
			level: 3
			parent_id: 2
			child_ids: []
		5:
			item_number: Ino1.1.2
			content_object_id: topic/5
			meeting_id: 1
			level: 2
			parent_id: 2
			child_ids: []
		6:
			item_number: Ino2
			content_object_id: topic/6
			meeting_id: 1
			level: 0
			weight: 3
			parent_id: 0
			child_ids: [7]
		7:
			item_number: Ino2.1
			content_object_id: topic/7
			meeting_id: 1
			level: 1
			weight: 4
			parent_id: 6
			child_ids: [8]
		8:
			item_number: Ino2.1.1
			content_object_id: topic/8
			meeting_id: 1
			level: 2
			weight: 5
			parent_id: 7
			child_ids: []

	topic/1/title: topic title 1
	topic/2/title: topic title 1.1
	topic/3/title: topic title 1.2
	topic/4/title: topic title 1.1.1
	topic/5/title: topic title 1.1.2
	topic/6/title: topic title 2
	topic/7/title: topic title 2.1
	topic/8/title: topic title 2.1.1

    `)

	for _, tt := range []struct {
		name       string
		data       map[string]string
		expect     string
		expectKeys []string
	}{
		{
			"with_leveled_item",
			data,
			`{
				"items": [
					{
						"depth": 0,
						"title_information": {
							"collection": "topic",
							"agenda_item_number": "Ino1",
							"content_object_id": "topic/1",
							"title": "topic title 1"
					    }
					},
					{
						"depth": 1,
						"title_information": {
							"collection": "topic",
							"agenda_item_number": "Ino1.1",
							"content_object_id": "topic/2",
							"title": "topic title 1.1"
					    }
					},
					{
						"depth": 2,
						"title_information": {
							"collection": "topic",
							"agenda_item_number": "Ino1.1.1",
							"content_object_id": "topic/4",
							"title": "topic title 1.1.1"
					    }
					},
					{
						"depth": 2,
						"title_information": {
							"collection": "topic",
							"agenda_item_number": "Ino1.1.2",
							"content_object_id": "topic/5",
							"title": "topic title 1.1.2"
					    }
					},
					{
						"depth": 1,
						"title_information": {
							"collection": "topic",
							"agenda_item_number": "Ino1.2",
							"content_object_id": "topic/3",
							"title": "topic title 1.2"
					    }
					},
					{
						"depth": 0,
						"title_information": {
							"collection": "topic",
							"agenda_item_number": "Ino2",
							"content_object_id": "topic/6",
							"title": "topic title 2"
					    }
					},
					{
						"depth": 1,
						"title_information": {
							"collection": "topic",
							"agenda_item_number": "Ino2.1",
							"content_object_id": "topic/7",
							"title": "topic title 2.1"
					    }
					},
					{
						"depth": 2,
						"title_information": {
							"collection": "topic",
							"agenda_item_number": "Ino2.1.1",
							"content_object_id": "topic/8",
							"title": "topic title 2.1.1"
					    }
					}
				]
			}
			`,
			[]string{
				"meeting/1/agenda_item_ids",
				"meeting/1/agenda_show_internal_items_on_projector",
				"agenda_item/1/id",
				"agenda_item/1/item_number",
				"agenda_item/1/content_object_id",
				"agenda_item/1/meeting_id",
				"agenda_item/1/is_hidden",
				"agenda_item/1/is_internal",
				"agenda_item/1/level",
				"agenda_item/1/weight",
				"agenda_item/1/parent_id",
				"topic/1/id",
				"topic/1/title",
				"topic/1/agenda_item_id",
				"agenda_item/2/id",
				"agenda_item/2/item_number",
				"agenda_item/2/content_object_id",
				"agenda_item/2/meeting_id",
				"agenda_item/2/is_hidden",
				"agenda_item/2/is_internal",
				"agenda_item/2/level",
				"agenda_item/2/weight",
				"agenda_item/2/parent_id",
				"topic/2/id",
				"topic/2/title",
				"topic/2/agenda_item_id",
				"agenda_item/3/id",
				"agenda_item/3/item_number",
				"agenda_item/3/content_object_id",
				"agenda_item/3/meeting_id",
				"agenda_item/3/is_hidden",
				"agenda_item/3/is_internal",
				"agenda_item/3/level",
				"agenda_item/3/weight",
				"agenda_item/3/parent_id",
				"topic/3/id",
				"topic/3/title",
				"topic/3/agenda_item_id",
				"agenda_item/4/id",
				"agenda_item/4/item_number",
				"agenda_item/4/content_object_id",
				"agenda_item/4/meeting_id",
				"agenda_item/4/is_hidden",
				"agenda_item/4/is_internal",
				"agenda_item/4/level",
				"agenda_item/4/weight",
				"agenda_item/4/parent_id",
				"topic/4/id",
				"topic/4/title",
				"topic/4/agenda_item_id",
				"agenda_item/5/id",
				"agenda_item/5/item_number",
				"agenda_item/5/content_object_id",
				"agenda_item/5/meeting_id",
				"agenda_item/5/is_hidden",
				"agenda_item/5/is_internal",
				"agenda_item/5/level",
				"agenda_item/5/weight",
				"agenda_item/5/parent_id",
				"topic/5/id",
				"topic/5/title",
				"topic/5/agenda_item_id",
				"agenda_item/6/id",
				"agenda_item/6/item_number",
				"agenda_item/6/content_object_id",
				"agenda_item/6/meeting_id",
				"agenda_item/6/is_hidden",
				"agenda_item/6/is_internal",
				"agenda_item/6/level",
				"agenda_item/6/weight",
				"agenda_item/6/parent_id",
				"topic/6/id",
				"topic/6/title",
				"topic/6/agenda_item_id",
				"agenda_item/7/id",
				"agenda_item/7/item_number",
				"agenda_item/7/content_object_id",
				"agenda_item/7/meeting_id",
				"agenda_item/7/is_hidden",
				"agenda_item/7/is_internal",
				"agenda_item/7/level",
				"agenda_item/7/weight",
				"agenda_item/7/parent_id",
				"topic/7/id",
				"topic/7/title",
				"topic/7/agenda_item_id",
				"agenda_item/8/id",
				"agenda_item/8/item_number",
				"agenda_item/8/content_object_id",
				"agenda_item/8/meeting_id",
				"agenda_item/8/is_hidden",
				"agenda_item/8/is_internal",
				"agenda_item/8/level",
				"agenda_item/8/weight",
				"agenda_item/8/parent_id",
				"topic/8/id",
				"topic/8/title",
				"topic/8/agenda_item_id",
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			closed := make(chan struct{})
			defer close(closed)
			ds := dsmock.NewMockDatastore(closed, tt.data)
			fetch := datastore.NewFetcher(ds)

			p7on := &projector.Projection{
				ContentObjectID: "meeting/1",
				Type:            "agenda_item_list",
				MeetingID:       1,
			}

			bs, keys, err := ailSlide.Slide(context.Background(), fetch, p7on)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expect, string(bs))
			assert.ElementsMatch(t, tt.expectKeys, keys)
		})
	}
}
