package slide_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/slide"
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

	ailSlide := s.Get("agenda_item_list")
	assert.NotNilf(t, ailSlide, "Slide with name `agenda_item_list` not found.")

	data := dsmock.YAMLData(`
	meeting:
		1:
			agenda_show_internal_items_on_projector: false
			agenda_item_ids: [1,2,3,4,5,6,7]

	agenda_item:
		1:
			item_number: Ino1
			content_object_id: topic/1
			meeting_id: 1
			is_hidden: false
			is_internal: false
			level: 0

		2:
			item_number: Ino2
			content_object_id: motion/1
			meeting_id: 1
			is_hidden: false
			is_internal: false
			level: 0

		3:
			item_number: Ino3
			content_object_id: motion_block/1
			meeting_id: 1
			is_hidden: false
			is_internal: false
			level: 0

		4:
			item_number: Ino4
			content_object_id: assignment/1
			meeting_id: 1
			is_hidden: false
			is_internal: false
			level: 0

		5:
			item_number: Ino5 misses because of level
			content_object_id: topic/2
			meeting_id: 1
			is_hidden: false
			is_internal: false
			level: 1

		6:
			item_number: Ino6 misses because of hidden
			content_object_id: topic/3
			meeting_id: 1
			is_hidden: true
			is_internal: false
			level: 0

		7:
			item_number: Ino7 misses because of internal
			content_object_id: topic/4
			meeting_id: 1
			is_hidden: false
			is_internal: True
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
			"Starter",
			data,
			`{
				"items": [
					{
						"depth": 0,
						"title_information": {
							"agenda_item_number": "Ino1",
							"content_object_id": "topic/1",
							"title": "topic title 1"
					    }
					},
					{
						"depth": 0,
						"title_information": {
							"agenda_item_number": "Ino2",
							"content_object_id": "motion/1",
							"number": "motion number 1",
							"title": "motion title 1"
						}
					},
					{
						"depth": 0,
						"title_information": {
							"agenda_item_number": "Ino3",
							"content_object_id": "motion_block/1",
							"title": "motion_block title 1"
						}
					},
					{
						"depth": 0,
						"title_information": {
							"agenda_item_number": "Ino4",
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
				"topic/1/id",
				"topic/1/title",
				"agenda_item/2/id",
				"agenda_item/2/item_number",
				"agenda_item/2/content_object_id",
				"agenda_item/2/meeting_id",
				"agenda_item/2/is_hidden",
				"agenda_item/2/is_internal",
				"agenda_item/2/level",
				"motion/1/id",
				"motion/1/number",
				"motion/1/title",
				"agenda_item/3/id",
				"agenda_item/3/item_number",
				"agenda_item/3/content_object_id",
				"agenda_item/3/meeting_id",
				"agenda_item/3/is_hidden",
				"agenda_item/3/is_internal",
				"agenda_item/3/level",
				"motion_block/1/id",
				"motion_block/1/title",
				"agenda_item/4/id",
				"agenda_item/4/item_number",
				"agenda_item/4/content_object_id",
				"agenda_item/4/meeting_id",
				"agenda_item/4/is_hidden",
				"agenda_item/4/is_internal",
				"agenda_item/4/level",
				"assignment/1/id",
				"assignment/1/title",
				"agenda_item/5/id",
				"agenda_item/5/item_number",
				"agenda_item/5/content_object_id",
				"agenda_item/5/meeting_id",
				"agenda_item/5/is_hidden",
				"agenda_item/5/is_internal",
				"agenda_item/5/level",
				"agenda_item/6/id",
				"agenda_item/6/item_number",
				"agenda_item/6/content_object_id",
				"agenda_item/6/meeting_id",
				"agenda_item/6/is_hidden",
				"agenda_item/6/is_internal",
				"agenda_item/6/level",
				"agenda_item/7/id",
				"agenda_item/7/item_number",
				"agenda_item/7/content_object_id",
				"agenda_item/7/meeting_id",
				"agenda_item/7/is_hidden",
				"agenda_item/7/is_internal",
				"agenda_item/7/level",
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			closed := make(chan struct{})
			defer close(closed)
			ds := dsmock.NewMockDatastore(closed, tt.data)

			p7on := &projector.Projection{
				ContentObjectID: "meeting/1",
				Type:            "agenda_item_list",
				MeetingID:       1,
				Options:         map[string]interface{}{"only_main_items": true},
			}

			bs, keys, err := ailSlide.Slide(context.Background(), ds, p7on)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expect, string(bs))
			assert.ElementsMatch(t, tt.expectKeys, keys)
		})
	}
}

func TestAgendaItem(t *testing.T) {
	s := new(projector.SlideStore)
	slide.AgendaItem(s)
	slide.Topic(s)

	aiSlide := s.Get("agenda_item")
	assert.NotNilf(t, aiSlide, "Slide with name `agenda_item` not found.")

	data := dsmock.YAMLData(`
	agenda_item/1:
		item_number: Ino1
		content_object_id: topic/1
		meeting_id: 1
		is_hidden: false
		is_internal: false
		level: 0

	topic/1/title: topic title 1
    `)

	for _, tt := range []struct {
		name       string
		data       map[string]string
		expect     string
		expectKeys []string
	}{
		{
			"Starter",
			data,
			`{
				"depth": 0,
				"title_information": {
					"agenda_item_number": "Ino1",
					"content_object_id": "topic/1",
					"title": "topic title 1"
				}
			}
			`,
			[]string{
				"agenda_item/1/id",
				"agenda_item/1/item_number",
				"agenda_item/1/content_object_id",
				"agenda_item/1/meeting_id",
				"agenda_item/1/is_hidden",
				"agenda_item/1/is_internal",
				"agenda_item/1/level",
				"topic/1/id",
				"topic/1/title",
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			closed := make(chan struct{})
			defer close(closed)
			ds := dsmock.NewMockDatastore(closed, tt.data)

			p7on := &projector.Projection{
				ContentObjectID: "agenda_item/1",
				Type:            "agenda_item_list",
				MeetingID:       1,
			}

			bs, keys, err := aiSlide.Slide(context.Background(), ds, p7on)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expect, string(bs))
			assert.ElementsMatch(t, tt.expectKeys, keys)
		})
	}
}
