package slide_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/slide"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/dsmock"
	"github.com/stretchr/testify/assert"
)

func TestPoll(t *testing.T) {
	s := new(projector.SlideStore)
	slide.Poll(s)
	slide.Motion(s)
	slide.Topic(s)

	pollSlide := s.GetSlider("poll")
	assert.NotNilf(t, pollSlide, "Slide with name `poll` not found.")

	data := dsmock.YAMLData(`
	poll:
	    1:
	        content_object_id: motion/1
	        title: Poll Title 1
	        description: Poll description 1
	        type: analog
	        state: published
	        global_yes: false
	        global_no: true
	        global_abstain: false
	        option_ids: [1, 2]
	        is_pseudoanonymized: false
	        pollmethod: YNA
	        onehundred_percent_base: YNA
	        votesvalid: "2.000000"
	        votesinvalid: "9.000000"
	        votescast: "2.000000"
	        global_option_id: 3
	        meeting_id: 111
	        entitled_users_at_stop: {"A": "bcd", "B":"def"}
	motion:
	    1:
	        title: Motion title 1
	        number: motion number 1234
	        agenda_item_id: 1
	option:
	    1:
	        text: Option text
	        content_object_id: topic/1
	        yes: "4.000000"
	        no: "5.000000"
	        abstain: "6.000000"
	        weight: 10
	    2:
	        text: Option text
	        content_object_id: topic/2
	        yes: "5.000000"
	        no: "4.000000"
	        abstain: "3.000000"
	        weight: 3
	    3:
	        yes: "14.000000"
	        no: "15.000000"
	        abstain: "16.000000"
	topic:
	    1:
	        title: Topic title 1
	        text: Topic text 1
	        agenda_item_id: 2
	    2:
	        title: Topic title 2
	        text: Topic text 2
	        agenda_item_id: 3
	agenda_item/1/item_number: itemNr. Motion1
	agenda_item/2/item_number: itemNr. Topic1
	agenda_item/3/item_number: itemNr. Topic2
	`)

	for _, tt := range []struct {
		name       string
		data       map[string]string
		expect     string
		expectKeys []string
	}{
		{
			"Poll state published",
			data,
			`{
                "id":1,
                "content_object_id":"motion/1",
                "title_information": {
                    "content_object_id":"motion/1",
                    "collection":"motion",
                    "title":"Motion title 1",
                    "number":"motion number 1234",
                    "agenda_item_number":"itemNr. Motion1"
                },
                "title":"Poll Title 1",
                "description":"Poll description 1",
                "type":"analog",
                "state":"published",
                "global_yes":false,
                "global_no":true,
                "global_abstain":false,
                "options": [
                    {
                        "content_object_id":"topic/2",
                        "text":"Option text",
                        "content_object":{
                            "content_object_id":"topic/2",
                            "collection":"topic",
                            "title":"Topic title 2",
                            "agenda_item_number":"itemNr. Topic2"
                        },
                        "yes":"5.000000",
                        "no":"4.000000",
                        "abstain":"3.000000"
                    },
                    {
                        "content_object_id":"topic/1",
                        "text":"Option text",
                        "content_object":{
                            "content_object_id":"topic/1",
                            "collection":"topic",
                            "title":"Topic title 1",
                            "agenda_item_number":"itemNr. Topic1"
                        },
                        "yes":"4.000000",
                        "no":"5.000000",
                        "abstain":"6.000000"
                    }
                ],
                "entitled_users_at_stop": {
                    "A":"bcd",
                    "B":"def"
                },
                "is_pseudoanonymized":false,
                "pollmethod":"YNA",
                "onehundred_percent_base":"YNA",
                "votesvalid": "2.000000",
                "votesinvalid": "9.000000",
                "votescast": "2.000000",
                "global_option":{
                    "yes":"14.000000",
                    "no":"15.000000",
                    "abstain":"16.000000"
                }
            }
            `,
			[]string{
				"poll/1/state",
				"poll/1/id",
				"poll/1/content_object_id",
				"poll/1/title",
				"poll/1/description",
				"poll/1/type",
				"poll/1/state",
				"poll/1/global_yes",
				"poll/1/global_no",
				"poll/1/global_abstain",
				"poll/1/option_ids",
				"poll/1/meeting_id",
				"poll/1/entitled_users_at_stop",
				"poll/1/is_pseudoanonymized",
				"poll/1/pollmethod",
				"poll/1/onehundred_percent_base",
				"poll/1/votesvalid",
				"poll/1/votesinvalid",
				"poll/1/votescast",
				"poll/1/global_option_id",
				"motion/1/id",
				"motion/1/number",
				"motion/1/title",
				"motion/1/agenda_item_id",
				"option/1/id",
				"option/1/text",
				"option/1/content_object_id",
				"option/1/yes",
				"option/1/no",
				"option/1/abstain",
				"option/1/weight",
				"option/2/id",
				"option/2/text",
				"option/2/content_object_id",
				"option/2/yes",
				"option/2/no",
				"option/2/abstain",
				"option/2/weight",
				"option/3/yes",
				"option/3/no",
				"option/3/abstain",
				"topic/1/id",
				"topic/1/title",
				"topic/1/agenda_item_id",
				"topic/2/id",
				"topic/2/title",
				"topic/2/agenda_item_id",
				"agenda_item/1/item_number",
				"agenda_item/2/item_number",
				"agenda_item/3/item_number",
			},
		},
		{
			"Poll state finished",
			changeData(data, map[string]string{
				"poll/1/state": `"finished"`,
			}),
			`{
                "id":1,
                "content_object_id":"motion/1",
                "title_information": {
                    "content_object_id":"motion/1",
                    "collection":"motion",
                    "title":"Motion title 1",
                    "number":"motion number 1234",
                    "agenda_item_number":"itemNr. Motion1"
                },
                "title":"Poll Title 1",
                "description":"Poll description 1",
                "type":"analog",
                "state":"finished",
                "global_yes":false,
                "global_no":true,
                "global_abstain":false,
                "options": [
                    {
                        "content_object_id":"topic/2",
                        "text":"Option text",
                        "content_object":{
                            "content_object_id":"topic/2",
                            "collection":"topic",
                            "title":"Topic title 2",
                            "agenda_item_number":"itemNr. Topic2"
                        }
                    },
                    {
                        "content_object_id":"topic/1",
                        "text":"Option text",
                        "content_object":{
                            "content_object_id":"topic/1",
                            "collection":"topic",
                            "title":"Topic title 1",
                            "agenda_item_number":"itemNr. Topic1"
                        }
                    }
                ]
            }
            `,
			[]string{
				"poll/1/id",
				"poll/1/content_object_id",
				"poll/1/title",
				"poll/1/description",
				"poll/1/type",
				"poll/1/state",
				"poll/1/state",
				"poll/1/global_yes",
				"poll/1/global_no",
				"poll/1/global_abstain",
				"poll/1/option_ids",
				"poll/1/meeting_id",
				"motion/1/id",
				"motion/1/number",
				"motion/1/title",
				"motion/1/agenda_item_id",
				"option/1/id",
				"option/1/weight",
				"option/1/text",
				"option/1/content_object_id",
				"option/2/id",
				"option/2/weight",
				"option/2/text",
				"option/2/content_object_id",
				"topic/1/id",
				"topic/1/title",
				"topic/1/agenda_item_id",
				"topic/2/id",
				"topic/2/title",
				"topic/2/agenda_item_id",
				"agenda_item/1/item_number",
				"agenda_item/2/item_number",
				"agenda_item/3/item_number",
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			closed := make(chan struct{})
			defer close(closed)
			ds := dsmock.NewMockDatastore(closed, tt.data)

			p7on := &projector.Projection{
				ContentObjectID: "poll/1",
			}

			bs, keys, err := pollSlide.Slide(context.Background(), ds, p7on)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expect, string(bs))
			assert.ElementsMatch(t, tt.expectKeys, keys)
		})
	}
}
