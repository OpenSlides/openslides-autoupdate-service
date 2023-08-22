package slide_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/slide"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListOfSpeakers(t *testing.T) {
	s := new(projector.SlideStore)
	slide.ListOfSpeaker(s)
	slide.Assignment(s)

	losSlide := s.GetSlider("list_of_speakers")
	assert.NotNilf(t, losSlide, "Slide with name `list_of_speakers` not found.")

	data := dsmock.YAMLData(`
	meeting/1:
		list_of_speakers_amount_next_on_projector: 4
		list_of_speakers_amount_last_on_projector: 2
		list_of_speakers_show_amount_of_speakers_on_slide: true
	list_of_speakers/1:
		content_object_id:	assignment/1
		closed: 			true
		speaker_ids: 		[1,2,3,4,5,6]

	assignment/1:
		title: assignment1 title
		agenda_item_id: 1
	agenda_item/1/item_number: ItemNr Assignment1

	speaker:
		1:
			# Waiting
			meeting_user_id:        100
			speech_state:   contribution
			note:           Seq2Waiting
			point_of_order: false
			weight:         10
		2:
			# Waiting
			meeting_user_id:        110
			speech_state:   contribution
			note:           Seq1Waiting
			point_of_order: true
			weight:         5
		
		3:
			# Current
			meeting_user_id:        200
			speech_state:   pro
			note:           SeqCurrent
			point_of_order: false
			weight:         20
			begin_time:     100
			
		
		4:
			# Finished
			meeting_user_id:        300
			speech_state:   contra
			note:           Seq3Finished
			point_of_order: true
			weight:         30
			begin_time:     20
			end_time:       23
			
		5:
			# Finished
			meeting_user_id:        310
			speech_state:   contra
			note:           Seq1Finished
			point_of_order: true
			weight:         30
			begin_time:     29
			end_time:       32
		6:
			# Finished
			meeting_user_id:        320
			speech_state:   contra
			note:           Seq2Finished
			point_of_order: true
			weight:         30
			begin_time:     24
			end_time:       28

	meeting_user:
		100:
			user_id: 10
		
		110:
			user_id: 11

		200:
			user_id: 20
		
		300:
			user_id: 30

		310:
			user_id: 31
		
		320:
			user_id: 32
		
	user:
		10:
			username: jonny123
		11:
			username: elenor
		20:
			first_name: Jonny
		30:
			last_name: Bo
		31:
			username: Ernest
		32:
			username: Calli
	`)

	for _, tt := range []struct {
		name   string
		data   map[dskey.Key][]byte
		expect string
	}{
		{
			"Starter",
			data,
			`{
				"waiting": [
					{
						"user": "elenor",
						"speech_state": "contribution",
						"note": "Seq1Waiting",
						"point_of_order": true
					},
					{
						"user": "jonny123",
						"speech_state": "contribution",
						"note": "Seq2Waiting",
						"point_of_order": false
					}
				],
				"current": {
					"user": "Jonny",
					"speech_state": "pro",
					"note": "SeqCurrent",
					"point_of_order": false
				},
				"finished": [
					{
						"user": "Ernest",
						"speech_state": "contra",
						"note": "Seq1Finished",
						"point_of_order": true
					},
					{
						"user": "Calli",
						"speech_state": "contra",
						"note": "Seq2Finished",
						"point_of_order": true
					}
				],
				"closed": true,
				"title_information": {
					"agenda_item_number": "ItemNr Assignment1",
					"collection": "assignment",
					"content_object_id": "assignment/1",
					"title": "assignment1 title"
				},
				"number_of_waiting_speakers": 2
			}
			`,
		},
		{
			"No Current speaker",
			changeData(data, map[dskey.Key][]byte{
				dskey.MustKey("list_of_speakers/1/speaker_ids"):                              []byte("[1,4]"),
				dskey.MustKey("meeting/1/list_of_speakers_show_amount_of_speakers_on_slide"): []byte("false"),
			}),
			`{
				"waiting": [{
					"user": "jonny123",
					"speech_state": "contribution",
					"note": "Seq2Waiting",
					"point_of_order": false
				}],
				"current": null,
				"finished": [{
					"user": "Bo",
					"speech_state": "contra",
					"note": "Seq3Finished",
					"point_of_order": true
				}],
				"closed": true,
				"title_information": {
					"agenda_item_number": "ItemNr Assignment1",
					"collection": "assignment",
					"content_object_id": "assignment/1",
					"title": "assignment1 title"
				}
			}
			`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ds := dsmock.NewFlow(tt.data)
			fetch := datastore.NewFetcher(ds)

			p7on := &projector.Projection{
				ContentObjectID: "list_of_speakers/1",
				MeetingID:       1,
			}

			bs, err := losSlide.Slide(context.Background(), fetch, p7on)
			assert.NoError(t, err)
			assert.NoError(t, fetch.Err())
			assert.JSONEq(t, tt.expect, string(bs))
		})
	}
}

func getDataForCurrentList() map[dskey.Key][]byte {
	// This one is a bit complicated and will be used
	// for tests current_list_of_speakers and, slightly modified,
	// for current_speaker_chyron
	//
	// The slide gets a contentObjectID of meeting/6
	// meeting/6 has reference_projector 60
	// projector/60 has projection/2
	// projection/2 has	content_object_id topic/5
	// motion_block/1 points list_of_speakers/7
	// list_of_speakers/7 points to speaker/8
	// speaker/8 points to user/10
	// user/10 has username jonny123
	//
	// lets find out if this username is on the slide-data...
	return dsmock.YAMLData(`
		projector/60/current_projection_ids: [1, 2]
		projection/1/content_object_id: user/10
		projection/2/content_object_id: motion_block/1

		meeting/6:
			list_of_speakers_show_amount_of_speakers_on_slide: false
			reference_projector_id: 60
		motion_block/1:
			list_of_speakers_id: 7
			title: motion_block1 title
			agenda_item_id: 1

		list_of_speakers/7:
			content_object_id:	motion_block/1
			closed: 			true
			speaker_ids: 		[8]

		speaker/8:
				meeting_user_id:        100
				speech_state:   pro
				note:           Lonesome speaker
				point_of_order: false
				weight:         10
		
		meeting_user/100/user_id: 10
		user/10/username: jonny123
		agenda_item/1/item_number: ItemNr. MotionBlock1
	`)
}

func TestCurrentListOfSpeakers(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	s := new(projector.SlideStore)
	slide.CurrentListOfSpeakers(s)
	slide.MotionBlock(s)

	slide := s.GetSlider("current_list_of_speakers")
	require.NotNilf(t, slide, "Slide with name `current_list_of_speakers` not found.")

	data := getDataForCurrentList()
	for _, tt := range []struct {
		name   string
		data   map[dskey.Key][]byte
		expect string
	}{
		{
			"find second current projection with speaker list",
			data,
			`{
				"waiting": [{
					"user": "jonny123",
					"speech_state": "pro",
					"note": "Lonesome speaker",
					"point_of_order": false
				}],
				"current": null,
				"finished": null,
				"closed": true,
				"title_information": {
					"agenda_item_number": "ItemNr. MotionBlock1",
					"collection": "motion_block",
					"content_object_id": "motion_block/1",
					"title": "motion_block1 title"
				}
			}
			`,
		},
		{
			"don't find speaker list in current projections",
			changeData(data, map[dskey.Key][]byte{
				dskey.MustKey("motion_block/1/list_of_speakers_id"): []byte("0"),
			}),
			`{}`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ds := dsmock.NewFlow(tt.data)
			fetch := datastore.NewFetcher(ds)

			p7on := &projector.Projection{
				ID:              1,
				ContentObjectID: "meeting/6",
				Type:            "current_list_of_speakers",
				MeetingID:       6,
			}

			bs, err := slide.Slide(context.Background(), fetch, p7on)
			assert.NoError(t, err)
			assert.NoError(t, fetch.Err())
			assert.JSONEq(t, tt.expect, string(bs))
		})
	}
}

func TestCurrentSpeakerChyron(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	s := new(projector.SlideStore)
	slide.CurrentSpeakerChyron(s)

	slide := s.GetSlider("current_speaker_chyron")
	require.NotNilf(t, slide, "Slide with name `current_speaker_chyron` not found.")

	data := getDataForCurrentList()
	for k, v := range dsmock.YAMLData(`
		speaker/8/begin_time: 100
		speaker/8/end_time: 0

		user/10:
			title: Admiral
			first_name: Don
			last_name: Snyder
			default_structure_level: GB
			meeting_user_ids: [100]
		
		meeting_user/100:
			meeting_id: 6
			structure_level: Dinner

		projector/60:
			chyron_background_color: green
			chyron_font_color: red
	`) {
		data[k] = v
	}

	for _, tt := range []struct {
		name   string
		data   map[dskey.Key][]byte
		expect string
	}{
		{
			"current speaker chyron test find second",
			data,
			`{
				"background_color": "green",
				"font_color": "red",
				"current_speaker_name": "Admiral Don Snyder",
				"current_speaker_level": "Dinner"
			}
			`,
		},
		{
			"current speaker chyron test no current projection",
			changeData(data, map[dskey.Key][]byte{
				dskey.MustKey("motion_block/1/list_of_speakers_id"): []byte("0"),
			}),
			`{
				"background_color": "green",
				"font_color": "red",
				"current_speaker_name": "",
				"current_speaker_level": ""
			}
			`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ds := dsmock.NewFlow(tt.data)
			fetch := datastore.NewFetcher(ds)

			p7on := &projector.Projection{
				ID:              1,
				ContentObjectID: "meeting/6",
				Type:            "current_speaker_chyron",
				MeetingID:       6,
			}

			bs, err := slide.Slide(context.Background(), fetch, p7on)
			assert.NoError(t, err)
			assert.NoError(t, fetch.Err())
			assert.JSONEq(t, tt.expect, string(bs))
		})
	}
}

func changeData(orig, change map[dskey.Key][]byte) map[dskey.Key][]byte {
	out := make(map[dskey.Key][]byte)
	for k, v := range orig {
		out[k] = v
	}
	for k, v := range change {
		out[k] = v
	}
	return out
}
