package slide_test

import (
	"context"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/projector"
	"github.com/openslides/openslides-autoupdate-service/internal/projector/slide"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestListOfSpeakers(t *testing.T) {
	s := new(projector.SlideStore)
	slide.ListOfSpeaker(s)

	losSlide := s.Get("list_of_speakers")
	assert.NotNilf(t, losSlide, "Slide with name `list_of_speakers` not found.")

	data := test.YAMLData(`
	list_of_speakers/1:
		content_object_id:	topic/1
		closed: 			true
		speaker_ids: 		[1,2,3]

	topic/1/title: topic title

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
		name   string
		data   map[string]string
		expect string
	}{
		{
			"Starter",
			data,
			`{
				"title": "topic title",
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
				"title_information": "title_information for topic/1"
			}
			`,
		},
		{
			"No Current spaker",
			changeData(data, map[string]string{
				"list_of_speakers/1/speaker_ids": "[1,3]",
			}),
			`{
				"title": "topic title",
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
				"title_information": "title_information for topic/1"
			}
			`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			closed := make(chan struct{})
			defer close(closed)
			ds := test.NewMockDatastore(closed, tt.data)

			p7on := &projector.Projection{
				ContentObjectID: "list_of_speakers/1",
			}

			bs, keys, err := losSlide.Slide(context.Background(), ds, p7on)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expect, string(bs))
			expectedKeys := []string{
				"list_of_speakers/1/speaker_ids",
				"list_of_speakers/1/content_object_id",
				"list_of_speakers/1/closed",
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
			}
			_, _ = keys, expectedKeys
			//assert.ElementsMatch(t, expectedKeys, keys)
		})
	}
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
