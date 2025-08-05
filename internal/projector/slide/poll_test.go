package slide_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/slide"
	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/OpenSlides/openslides-go/datastore/dsmock"
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
		name   string
		data   map[dskey.Key][]byte
		expect string
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
                "global_abstain":false,
                "live_voting_enabled":false,
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
		},
		{
			"Poll state finished",
			changeData(data, map[dskey.Key][]byte{
				dskey.MustKey("poll/1/state"): []byte(`"finished"`),
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
                "live_voting_enabled":false,
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
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ds := dsmock.Stub(tt.data)
			fetch := datastore.NewFetcher(ds)

			p7on := &projector.Projection{
				ContentObjectID: "poll/1",
			}

			bs, err := pollSlide.Slide(context.Background(), fetch, p7on)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expect, string(bs))
		})
	}
}

func TestPollSingleVotes(t *testing.T) {
	s := new(projector.SlideStore)
	slide.Poll(s)
	slide.Motion(s)
	slide.Topic(s)

	pollSlide := s.GetSlider("poll")
	assert.NotNilf(t, pollSlide, "Slide with name `poll` not found.")

	data := dsmock.YAMLData(`
	user:
		1:
			title: Billy
			first_name: Bob
			last_name: Buckingham
			username: BillyBobBuckingham
			member_number: 123456789abcdef
			meeting_user_ids: [1111]
			meeting_ids: [1]
		2:
			first_name: Johnny
			last_name: Johnson
			username: JohnnySonOfJohn
			meeting_user_ids: [1112]
			meeting_ids: [1]
		6:
			title: Sir
			first_name: Shawn Stanley Sheldon
			last_name: Sinclair
			username: SirSinclair
			meeting_user_ids: [1111]
			meeting_ids: [1]
	structure_level:
		1:
			name: Birmingham
			meeting_id: 111
	meeting:
		111:
			poll_ids: [1]
			meeting_user_ids: [1111, 1112, 1116]
			user_ids: [1,2,6]
			structure_level_ids: [1]
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
	        entitled_users_at_stop: [{"voted": false, "present": true, "user_id": 1, "vote_delegated_to_user_id": 2}, {"voted": false, "present": true, "user_id": 4, "user_merged_into_id": 6}]
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
		name   string
		data   map[dskey.Key][]byte
		expect string
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
                "live_voting_enabled":false,
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
                "entitled_users_at_stop": [
					{
						"voted": false,
						"present": true,
						"user_id": 1,
						"vote_delegated_to_user_id": 2,
						"user": {
							"id": 1,
							"title": "Billy",
							"first_name": "Bob",
							"last_name": "Buckingham"
						}
					},
					{
						"voted": false,
						"present": true,
						"user_id": 4,
						"user_merged_into_id": 6,
						"user": {
							"id": 6,
							"title": "Sir",
							"first_name": "Shawn Stanley Sheldon",
							"last_name": "Sinclair"
						}
					}
				],
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
		},
		{
			"Poll state finished",
			changeData(data, map[dskey.Key][]byte{
				dskey.MustKey("poll/1/state"): []byte(`"finished"`),
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
                "live_voting_enabled":false,
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
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ds := dsmock.Stub(tt.data)
			fetch := datastore.NewFetcher(ds)

			p7on := &projector.Projection{
				ContentObjectID: "poll/1",
				Options:         []byte(`{"single_votes":true}`),
			}

			bs, err := pollSlide.Slide(context.Background(), fetch, p7on)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expect, string(bs))
		})
	}
}

func TestPollLiveVote(t *testing.T) {
	s := new(projector.SlideStore)
	slide.Poll(s)
	slide.Motion(s)
	slide.Topic(s)

	pollSlide := s.GetSlider("poll")
	assert.NotNilf(t, pollSlide, "Slide with name `poll` not found.")

	data := dsmock.YAMLData(`
	user:
		1:
			title: Billy
			first_name: Bob
			last_name: Buckingham
			username: BillyBobBuckingham
			member_number: 123456789abcdef
			meeting_user_ids: [1113]
			meeting_ids: [1]
		2:
			first_name: Johnny
			last_name: Johnson
			username: JohnnySonOfJohn
			meeting_user_ids: [1112]
			meeting_ids: [1]
		6:
			title: Sir
			first_name: Shawn Stanley Sheldon
			last_name: Sinclair
			username: SirSinclair
			meeting_user_ids: [1111]
			meeting_ids: [1]
	meeting_user:
	    1111:
	        user_id: 6
	        structure_level_ids: [1]
	    1112:
	        user_id: 2
	        structure_level_ids: [2]
	    1113:
	        user_id: 1
	structure_level:
		1:
			name: Birmingham
			meeting_id: 111
		2:
			name: Fooo
			meeting_id: 111
	meeting:
		111:
			poll_ids: [1]
			meeting_user_ids: [1111, 1112, 1116]
			user_ids: [1,2,6]
			structure_level_ids: [1, 2]
	group:
	    1:
	        meeting_user_ids: [1111, 1112, 1113]
	poll:
	    1:
	        content_object_id: motion/1
	        title: Poll Title 1
	        description: Poll description 1
	        type: analog
	        state: running
	        global_yes: false
	        global_no: true
	        global_abstain: false
	        entitled_group_ids: [1]
	        option_ids: [1, 2]
	        is_pseudoanonymized: false
	        pollmethod: YNA
	        onehundred_percent_base: YNA
	        votesvalid: "2.000000"
	        votesinvalid: "9.000000"
	        votescast: "2.000000"
	        global_option_id: 3
	        meeting_id: 111
	        live_votes: {1: "{\"value\": {\"31\": \"Y\"}, \"weight\": \"1.000000\"}"}
	        live_voting_enabled: true
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
		name   string
		data   map[dskey.Key][]byte
		expect string
	}{
		{
			"Running poll",
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
                "state":"running",
                "global_yes":false,
                "global_no":true,
                "global_abstain":false,
                "live_voting_enabled":true,
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
                ],
                "entitled_structure_levels": {
                    "1": "Birmingham",
                    "2": "Fooo"
                },
                "entitled_users": {
                    "1": {
                        "present": false,
						"votes": {"31": "Y"},
						"weight": "1.000000",
						"user_data": {
							"id": 1,
							"title": "Billy",
							"first_name": "Bob",
							"last_name": "Buckingham"
						}
					},
                    "2": {
                        "present": false,
                        "structure_level_id": 2,
						"user_data": {
							"id": 2,
							"first_name": "Johnny",
							"last_name": "Johnson"
						}
					},
                    "6": {
                        "present": false,
                        "structure_level_id": 1,
						"user_data": {
							"id": 6,
							"title": "Sir",
							"first_name": "Shawn Stanley Sheldon",
							"last_name": "Sinclair"
						}
					}
                },
                "is_pseudoanonymized":false,
                "pollmethod":"YNA",
                "onehundred_percent_base":"YNA"
            }
            `,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ds := dsmock.Stub(tt.data)
			fetch := datastore.NewFetcher(ds)

			p7on := &projector.Projection{
				ContentObjectID: "poll/1",
				Options:         []byte(`{"single_votes":true}`),
			}

			bs, err := pollSlide.Slide(context.Background(), fetch, p7on)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expect, string(bs))
		})
	}
}
