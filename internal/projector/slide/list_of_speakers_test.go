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
			user_id:        10
			speech_state:   contribution
			note:           Seq2Waiting
			point_of_order: false
			weight:         10
		2:
			# Waiting
			user_id:        11
			speech_state:   contribution
			note:           Seq1Waiting
			point_of_order: true
			weight:         5
		
		3:
			# Current
			user_id:        20
			speech_state:   pro
			note:           SeqCurrent
			point_of_order: false
			weight:         20
			begin_time:     100
			
		
		4:
			# Finished
			user_id:        30
			speech_state:   contra
			note:           Seq3Finished
			point_of_order: true
			weight:         30
			begin_time:     20
			end_time:       23
			
		5:
			# Finished
			user_id:        31
			speech_state:   contra
			note:           Seq1Finished
			point_of_order: true
			weight:         30
			begin_time:     29
			end_time:       32
		6:
			# Finished
			user_id:        32
			speech_state:   contra
			note:           Seq2Finished
			point_of_order: true
			weight:         30
			begin_time:     24
			end_time:       28

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
		name       string
		data       map[string]string
		expect     string
		expectKeys []string
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
			[]string{
				"meeting/1/list_of_speakers_amount_next_on_projector",
				"meeting/1/list_of_speakers_amount_last_on_projector",
				"meeting/1/list_of_speakers_show_amount_of_speakers_on_slide",
				"list_of_speakers/1/speaker_ids",
				"list_of_speakers/1/content_object_id",
				"list_of_speakers/1/closed",
				"assignment/1/id",
				"assignment/1/title",
				"assignment/1/agenda_item_id",
				"agenda_item/1/item_number",
				"speaker/1/user_id",
				"speaker/1/speech_state",
				"speaker/1/note",
				"speaker/1/point_of_order",
				"speaker/1/weight",
				"speaker/1/begin_time",
				"speaker/1/end_time",
				"speaker/2/user_id",
				"speaker/2/speech_state",
				"speaker/2/note",
				"speaker/2/point_of_order",
				"speaker/2/weight",
				"speaker/2/begin_time",
				"speaker/2/end_time",
				"speaker/3/user_id",
				"speaker/3/speech_state",
				"speaker/3/note",
				"speaker/3/point_of_order",
				"speaker/3/weight",
				"speaker/3/begin_time",
				"speaker/3/end_time",
				"speaker/4/user_id",
				"speaker/4/speech_state",
				"speaker/4/note",
				"speaker/4/point_of_order",
				"speaker/4/weight",
				"speaker/4/begin_time",
				"speaker/4/end_time",
				"speaker/5/user_id",
				"speaker/5/speech_state",
				"speaker/5/note",
				"speaker/5/point_of_order",
				"speaker/5/weight",
				"speaker/5/begin_time",
				"speaker/5/end_time",
				"speaker/6/user_id",
				"speaker/6/speech_state",
				"speaker/6/note",
				"speaker/6/point_of_order",
				"speaker/6/weight",
				"speaker/6/begin_time",
				"speaker/6/end_time",
				"user/10/username",
				"user/10/title",
				"user/10/first_name",
				"user/10/last_name",
				"user/10/default_structure_level",
				"user/10/structure_level_$1",
				"user/11/username",
				"user/11/title",
				"user/11/first_name",
				"user/11/last_name",
				"user/11/default_structure_level",
				"user/11/structure_level_$1",
				"user/20/username",
				"user/20/title",
				"user/20/first_name",
				"user/20/last_name",
				"user/20/default_structure_level",
				"user/20/structure_level_$1",
				"user/30/username",
				"user/30/title",
				"user/30/first_name",
				"user/30/last_name",
				"user/30/default_structure_level",
				"user/30/structure_level_$1",
				"user/31/username",
				"user/31/title",
				"user/31/first_name",
				"user/31/last_name",
				"user/31/default_structure_level",
				"user/31/structure_level_$1",
				"user/32/username",
				"user/32/title",
				"user/32/first_name",
				"user/32/last_name",
				"user/32/default_structure_level",
				"user/32/structure_level_$1",
			},
		},
		{
			"No Current speaker",
			changeData(data, map[string]string{
				"list_of_speakers/1/speaker_ids":                              "[1,4]",
				"meeting/1/list_of_speakers_show_amount_of_speakers_on_slide": "false",
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
			[]string{
				"meeting/1/list_of_speakers_amount_next_on_projector",
				"meeting/1/list_of_speakers_amount_last_on_projector",
				"meeting/1/list_of_speakers_show_amount_of_speakers_on_slide",
				"list_of_speakers/1/speaker_ids",
				"list_of_speakers/1/content_object_id",
				"list_of_speakers/1/closed",
				"assignment/1/id",
				"assignment/1/title",
				"assignment/1/agenda_item_id",
				"agenda_item/1/item_number",
				"speaker/1/user_id",
				"speaker/1/speech_state",
				"speaker/1/note",
				"speaker/1/point_of_order",
				"speaker/1/weight",
				"speaker/1/begin_time",
				"speaker/1/end_time",
				"speaker/4/user_id",
				"speaker/4/speech_state",
				"speaker/4/note",
				"speaker/4/point_of_order",
				"speaker/4/weight",
				"speaker/4/begin_time",
				"speaker/4/end_time",
				"user/10/username",
				"user/10/title",
				"user/10/first_name",
				"user/10/last_name",
				"user/10/default_structure_level",
				"user/10/structure_level_$1",
				"user/30/username",
				"user/30/title",
				"user/30/first_name",
				"user/30/last_name",
				"user/30/default_structure_level",
				"user/30/structure_level_$1",
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			closed := make(chan struct{})
			defer close(closed)
			ds := dsmock.NewMockDatastore(closed, tt.data)

			p7on := &projector.Projection{
				ContentObjectID: "list_of_speakers/1",
				MeetingID:       1,
			}

			bs, keys, err := losSlide.Slide(context.Background(), ds, p7on)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expect, string(bs))
			assert.ElementsMatch(t, tt.expectKeys, keys)
		})
	}
}

func getDataForCurrentList() map[string]string {
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
		projector/60/current_projection_ids: [2]
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
				user_id:        10
				speech_state:   pro
				note:           Lonesome speaker
				point_of_order: false
				weight:         10

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
	t.Run("Find list of speakers", func(t *testing.T) {
		ds := dsmock.NewMockDatastore(closed, data)

		p7on := &projector.Projection{
			ID:              1,
			ContentObjectID: "meeting/6",
			Type:            "current_list_of_speakers",
			MeetingID:       6,
		}

		bs, keys, err := slide.Slide(context.Background(), ds, p7on)

		assert.NoError(t, err)
		expect := `{
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
		`
		assert.JSONEq(t, expect, string(bs))
		expectKeys := []string{
			"meeting/6/list_of_speakers_amount_next_on_projector",
			"meeting/6/list_of_speakers_amount_last_on_projector",
			"meeting/6/list_of_speakers_show_amount_of_speakers_on_slide",
			"meeting/6/reference_projector_id",
			"projector/60/current_projection_ids",
			"projection/2/content_object_id",
			"motion_block/1/id",
			"motion_block/1/title",
			"motion_block/1/agenda_item_id",
			"motion_block/1/list_of_speakers_id",
			"agenda_item/1/item_number",
			"list_of_speakers/7/speaker_ids",
			"list_of_speakers/7/content_object_id",
			"list_of_speakers/7/closed",
			"speaker/8/user_id",
			"speaker/8/speech_state",
			"speaker/8/note",
			"speaker/8/point_of_order",
			"speaker/8/weight",
			"speaker/8/begin_time",
			"speaker/8/end_time",
			"user/10/username",
			"user/10/title",
			"user/10/first_name",
			"user/10/last_name",
			"user/10/default_structure_level",
			"user/10/structure_level_$6",
		}
		assert.ElementsMatch(t, expectKeys, keys)
	})
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
			structure_level_$6: Dinner

		projector/60:
			chyron_background_color: green
			chyron_font_color: red
	`) {
		data[k] = v
	}

	t.Run("current speaker chyron test", func(t *testing.T) {
		ds := dsmock.NewMockDatastore(closed, data)

		p7on := &projector.Projection{
			ID:              1,
			ContentObjectID: "meeting/6",
			Type:            "current_speaker_chyron",
		}

		bs, keys, err := slide.Slide(context.Background(), ds, p7on)

		assert.NoError(t, err)
		expect := `{
			"background_color": "green",
			"font_color": "red",
			"current_speaker_name": "Admiral Don Snyder",
			"current_speaker_level": "Dinner"
		}
		`
		assert.JSONEq(t, expect, string(bs))
		expectKeys := []string{
			"meeting/6/reference_projector_id",
			"projector/60/current_projection_ids",
			"projector/60/chyron_background_color",
			"projector/60/chyron_font_color",
			"projection/2/content_object_id",
			"motion_block/1/list_of_speakers_id",
			"list_of_speakers/7/speaker_ids",
			"list_of_speakers/7/content_object_id",
			"list_of_speakers/7/closed",
			"speaker/8/user_id",
			"speaker/8/begin_time",
			"speaker/8/end_time",
			"user/10/username",
			"user/10/title",
			"user/10/first_name",
			"user/10/last_name",
			"user/10/default_structure_level",
			"user/10/structure_level_$6",
		}
		assert.ElementsMatch(t, expectKeys, keys)
	})
}

func changeData(orig, change map[string]string) map[string]string {
	out := make(map[string]string)
	for k, v := range orig {
		out[k] = v
	}
	for k, v := range change {
		out[k] = v
	}
	return out
}
