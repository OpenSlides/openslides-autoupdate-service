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
	slide.Topic(s)

	losSlide := s.GetSlider("list_of_speakers")
	assert.NotNilf(t, losSlide, "Slide with name `list_of_speakers` not found.")

	data := dsmock.YAMLData(`
	list_of_speakers/1:
		content_object_id:	topic/1
		closed: 			true
		speaker_ids: 		[1,2,3]

	topic/1/title: topic1 title

	speaker:
		1:
			# Waiting
			user_id:        10
			marked:         false
			point_of_order: false
			weight:         10
		
		2:
			# Current
			user_id:        20
			begin_time:     100
			marked:         true
			point_of_order: false
			weight:         20
		
		3:
			# Finished
			user_id:        30
			begin_time:     10
			end_time:       20
			marked:         true
			point_of_order: true
			weight:         30

	user:
		10:
			username: jonny123
		20:
			first_name: Jonny
		30:
			last_name: Bo
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
				"waiting": [{
					"user": "jonny123",
					"marked": false,
					"point_of_order": false,
					"weight": 10
				}],
				"current": {
					"user": "Jonny",
					"marked": true,
					"point_of_order": false,
					"weight": 20
				},
				"finished": [{
					"user": "Bo",
					"marked": true,
					"point_of_order": true,
					"weight": 30,
					"end_time": 20
				}],
				"closed": true,
				"content_object_collection": "topic",
				"title_information": {
					"agenda_item_number": "",
					"collection": "topic",
					"content_object_id": "topic/1",
					"title": "topic1 title"
				}
			}
			`,
			[]string{
				"list_of_speakers/1/speaker_ids",
				"list_of_speakers/1/content_object_id",
				"list_of_speakers/1/closed",
				"topic/1/id",
				"topic/1/title",
				"speaker/1/user_id",
				"speaker/1/marked",
				"speaker/1/point_of_order",
				"speaker/1/weight",
				"speaker/1/begin_time",
				"speaker/1/end_time",
				"speaker/2/user_id",
				"speaker/2/marked",
				"speaker/2/point_of_order",
				"speaker/2/weight",
				"speaker/2/begin_time",
				"speaker/2/end_time",
				"speaker/3/user_id",
				"speaker/3/marked",
				"speaker/3/point_of_order",
				"speaker/3/weight",
				"speaker/3/begin_time",
				"speaker/3/end_time",
				"user/10/username",
				"user/10/title",
				"user/10/first_name",
				"user/10/last_name",
				"user/10/default_structure_level",
				"user/20/username",
				"user/20/title",
				"user/20/first_name",
				"user/20/last_name",
				"user/20/default_structure_level",
				"user/30/username",
				"user/30/title",
				"user/30/first_name",
				"user/30/last_name",
				"user/30/default_structure_level",
			},
		},
		{
			"No Current speaker",
			changeData(data, map[string]string{
				"list_of_speakers/1/speaker_ids": "[1,3]",
			}),
			`{
				"waiting": [{
					"user": "jonny123",
					"marked": false,
					"point_of_order": false,
					"weight": 10
				}],
				"current": null,
				"finished": [{
					"user": "Bo",
					"marked": true,
					"point_of_order": true,
					"weight": 30,
					"end_time": 20
				}],
				"closed": true,
				"content_object_collection": "topic",
				"title_information": {
					"agenda_item_number": "",
					"collection": "topic",
					"content_object_id": "topic/1",
					"title": "topic1 title"
				}
			}
			`,
			[]string{
				"list_of_speakers/1/speaker_ids",
				"list_of_speakers/1/content_object_id",
				"list_of_speakers/1/closed",
				"topic/1/id",
				"topic/1/title",
				"speaker/1/user_id",
				"speaker/1/marked",
				"speaker/1/point_of_order",
				"speaker/1/weight",
				"speaker/1/begin_time",
				"speaker/1/end_time",
				"speaker/3/user_id",
				"speaker/3/marked",
				"speaker/3/point_of_order",
				"speaker/3/weight",
				"speaker/3/begin_time",
				"speaker/3/end_time",
				"user/10/username",
				"user/10/title",
				"user/10/first_name",
				"user/10/last_name",
				"user/10/default_structure_level",
				"user/30/username",
				"user/30/title",
				"user/30/first_name",
				"user/30/last_name",
				"user/30/default_structure_level",
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			closed := make(chan struct{})
			defer close(closed)
			ds := dsmock.NewMockDatastore(closed, tt.data)

			p7on := &projector.Projection{
				ContentObjectID: "list_of_speakers/1",
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
	// topic/5 points list_of_speakers/7
	// list_of_speakers/7 points to speaker/8
	// speaker/8 points to user/10
	// user/10 has username jonny123
	//
	// lets find out if this username is on the slide-data...
	return dsmock.YAMLData(`
		meeting/6/reference_projector_id: 60
		projector/60/current_projection_ids: [2]
		projection/2/content_object_id: topic/5

		topic/5:
			list_of_speakers_id: 7
			title: topic5 title

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

}
func TestCurrentListOfSpeakers(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	s := new(projector.SlideStore)
	slide.CurrentListOfSpeakers(s)
	slide.Topic(s)

	slide := s.GetSlider("current_list_of_speakers")
	require.NotNilf(t, slide, "Slide with name `current_list_of_speakers` not found.")

	data := getDataForCurrentList()
	t.Run("Find list of speakers", func(t *testing.T) {
		ds := dsmock.NewMockDatastore(closed, data)

		p7on := &projector.Projection{
			ID:              1,
			ContentObjectID: "meeting/6",
			Type:            "current_list_of_speakers",
		}

		bs, keys, err := slide.Slide(context.Background(), ds, p7on)

		assert.NoError(t, err)
		expect := `{
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
			"title_information": {
				"agenda_item_number": "",
				"collection": "topic",
				"content_object_id": "topic/5",
				"title": "topic5 title"
			}
		}
		`
		assert.JSONEq(t, expect, string(bs))
		expectKeys := []string{
			"meeting/6/reference_projector_id",
			"projector/60/current_projection_ids",
			"projection/2/content_object_id",
			"topic/5/id",
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
			"topic/5/list_of_speakers_id",
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
