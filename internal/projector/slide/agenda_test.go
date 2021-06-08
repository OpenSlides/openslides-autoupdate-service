package slide_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/slide"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/dsmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestAgendaItemList(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	s := new(projector.SlideStore)
	slide.CurrentListOfSpeakers(s)

	slide := s.Get("current_list_of_speakers")
	require.NotNilf(t, slide, "Slide with name `curent_list_of_speakers` not found.")

	// This one is a bit compicated:
	//
	// The slide gets a projection object with id 1
	// projection/1 points to projector/50
	// projector/50 points to meeting/6
	// meeting/6 has reference_projector 60
	// projector/60 has projection/2
	// projection/2 has	content_object_id topic/5
	// topic/5 points list_of_speakers/7
	// list_of_speakers/7 points to speaker/8
	// speaker/8 points to user/10
	// user/10 has username jonny123
	//
	// lets find out if this username is on the slide-data...
	data := dsmock.YAMLData(`
	projection/1/current_projector_id: 50
	projector/50/meeting_id: 6
	meeting/6/reference_projector_id: 60
	projector/60/current_projection_ids: [2]
	projection/2/content_object_id: topic/5

	topic/5:
		list_of_speakers_id: 7
		title: topic title

	list_of_speakers/7:
		content_object_id:	topic/5
		closed: 			true
		speaker_ids: 		[8]

	speaker/8:
			user_id:        10
			marked:         false
			point_of_order: false
			weight:         10

	user/10/username: jonny123
	`)

	t.Run("Find list of speakers", func(t *testing.T) {
		ds := dsmock.NewMockDatastore(closed, data)

		p7on := &projector.Projection{
			ID:              1,
			ContentObjectID: "list_of_speakers/1",
			Type:            "current_list_of_speakers",
		}

		bs, keys, err := slide.Slide(context.Background(), ds, p7on)

		assert.NoError(t, err)
		expect := `{
			"title": "topic title",
			"waiting": [{
				"user": "jonny123",
				"marked": false,
				"point_of_order": false,
				"weight": 10
			}],
			"current": null,
			"finished": null,
			"closed": true,
			"content_object_collection": "topic",
			"title_information": "title_information for topic/5"
		}
		`
		assert.JSONEq(t, expect, string(bs))
		expectKeys := []string{
			"projection/1/current_projector_id",
			"projector/50/meeting_id",
			"meeting/6/reference_projector_id",
			"projector/60/current_projection_ids",
			"projection/2/content_object_id",
			"topic/5/title",
			"topic/5/list_of_speakers_id",
			"list_of_speakers/7/speaker_ids",
			"list_of_speakers/7/content_object_id",
			"list_of_speakers/7/closed",
			"speaker/8/user_id",
			"speaker/8/marked",
			"speaker/8/point_of_order",
			"speaker/8/weight",
			"speaker/8/begin_time",
			"speaker/8/end_time",
			"user/10/username",
			"user/10/title",
			"user/10/first_name",
			"user/10/last_name",
			"user/10/default_structure_level",
		}
		assert.ElementsMatch(t, expectKeys, keys)
	})
}
