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
)

func TestMotion(t *testing.T) {
	s := new(projector.SlideStore)
	slide.Motion(s)

	motionSlide := s.GetSlider("motion")
	assert.NotNilf(t, motionSlide, "Slide with name `motion` not found.")

	data := dsmock.YAMLData(`
    projection:
        1:
            content_object_id: motion/1
            meeting_id: 1
            options: {
                mode: final
            }
    meeting:
        1:
            motions_enable_text_on_projector: false
            motions_enable_reason_on_projector: false
            motions_show_referring_motions: false
            motions_enable_recommendation_on_projector: false
            motions_statute_recommendations_by: Meeting MotionsStatuteRecommendations
            motions_recommendations_by: Meeting not used variant
            motions_enable_sidebox_on_projector: true
            motions_line_length: 85
            motions_preamble: The assembly may decide
            motions_default_line_numbering: outside
    motion:
        1:
            title: Motion Title 1
            number: MNr1234
            text: <p>Motion1 Text HTML</p>
            reason: <p>Motion1 reason HTML</p>
            modified_final_version: <p>Motion1 modifiedFinalVersion HTML</p>
            submitter_ids: [1,2,3]
            amendment_paragraphs: {"1": "amendmentParagraph1", "2": "amendmentParagraph2"}
            change_recommendation_ids: [1,2,3]
            amendment_ids: [3,4,5,6]
            referenced_in_motion_recommendation_extension_ids: [7,8]
            recommendation_id: 4
            recommendation_extension: RecommendationExtension_motion1
            recommendation_extension_reference_ids: ["motion/9", "motion/10"]
            meeting_id: 1
            agenda_item_id: 1
        2:
            title: Lead Motion Title
            number: Lead Motion 111
            text: <p>Lead Motion Text HTML</p>
            agenda_item_id: 2
        3:
            title: Amendment3 title
            number: Amendment3 123
            amendment_paragraphs: {"31": "amendmentParagraph31", "32": "amendmentParagraph32"}
            change_recommendation_ids: [4, 5]
            state_id: 1
            agenda_item_id: 3
        4:
            title: Amendment4 title
            number: Amendment4 4123
            amendment_paragraphs: null
            state_id: 2
            agenda_item_id: 4
        5:
            title: Amendment5 title
            number: Amendment5 5123
            amendment_paragraphs: null
            state_id: 3
            recommendation_id: 1
            agenda_item_id: 5
        6:
            title: Amendment6 title
            number: Amendment6 6123
            amendment_paragraphs: null
            state_id: 3
            recommendation_id: 2
            agenda_item_id: 6
        7:
            title: ReferencedInMotionRecommendationExtension7 title
            number: RIMRE7 number
            agenda_item_id: 7
        8:
            title: ReferencedInMotionRecommendationExtension8 title
            number: RIMRE8 number
            agenda_item_id: 8
        9:
            title: RecommendationExtensionReferenceMotion9 title
            number: RecommendationExtensionReferenceMotion9 number
            agenda_item_id: 9
        10:
            title: RecommendationExtensionReferenceMotion10 title
            number: RecommendationExtensionReferenceMotion10 number
            agenda_item_id: 10
    motion_state:
        1:
            merge_amendment_into_final: do_merge
        2:
            merge_amendment_into_final: do_not_merge
        3:
            merge_amendment_into_final: undefined
        4:
            recommendation_label: RecommendationLabel_state4
            show_recommendation_extension_field: true
    
    motion_submitter:
        1:
            weight: 100
            meeting_user_id: 130
            motion_id: 1
        2:
            weight: 2
            meeting_user_id: 110
            motion_id: 1
        3:
            weight: 30
            meeting_user_id: 120
            motion_id: 1
    
    meeting_user:
        130:
            user_id: 13
        110:
            user_id: 11
        120:
            user_id: 12

    user:
        11:
            username: user11
        12:
            username: user12
        13:
            username: user13
    motion_statute_paragraph:
        1:
            title: MotionStatuteParagraph1 title
            text: <p>MotionStatuteParagraph1 text html</p>
    motion_change_recommendation:
        1:
            internal: false
            rejected: true
            type: replacement
            other_description: Other Description1
            line_from: 1
            line_to: 3
            text: <p>text1 HTML</p>
            creation_time: 12345
        2:
            internal: true
        3:
            internal: false
            rejected: false
            type: insertion
            other_description: Other Description3
            line_from: 5
            line_to: 5
            text: <p>text3 HTML</p>
            creation_time: 32345
        4:
            internal: false
            rejected: true
            type: replacement
            other_description: ChangeRecommendation4 for amendment3
            line_from: 4
            line_to: 5
            text: <p>text4 HTML</p>
            creation_time: 42345
        5:
            internal: true
    agenda_item/1/item_number: ItemNr Motion1
    agenda_item/2/item_number: ItemNr Motion2
    agenda_item/3/item_number: ItemNr Motion3
    agenda_item/4/item_number: ItemNr Motion4
    agenda_item/5/item_number: ItemNr Motion5
    agenda_item/6/item_number: ItemNr Motion6
    agenda_item/7/item_number: ItemNr Motion7
    agenda_item/8/item_number: ItemNr Motion8
    agenda_item/9/item_number: ItemNr Motion9
    agenda_item/10/item_number: ItemNr Motion10
    `)

	for _, tt := range []struct {
		name    string
		options []byte
		data    map[dskey.Key][]byte
		expect  string
	}{
		{
			"Motion only non conditional",
			nil,
			data,
			`{
                "id":1,
                "title":"Motion Title 1",
                "number":"MNr1234",
                "submitters":[
                    "user11",
                    "user12",
                    "user13"
                ],
                "show_sidebox": true,
                "line_length": 85,
                "preamble": "The assembly may decide",
                "line_numbering": "outside",
                "amendment_paragraphs":{
                    "1":"amendmentParagraph1",
                    "2":"amendmentParagraph2"
                },
                "change_recommendations":[
                    {
                        "id": 1,
                        "rejected": true,
                        "type": "replacement",
                        "other_description": "Other Description1",
                        "line_from": 1,
                        "line_to": 3,
                        "text": "<p>text1 HTML</p>",
                        "creation_time": 12345
                    },
                    {
                        "id": 3,
                        "rejected": false,
                        "type": "insertion",
                        "other_description": "Other Description3",
                        "line_from": 5,
                        "line_to": 5,
                        "text": "<p>text3 HTML</p>",
                        "creation_time": 32345
                    }
                ],
                "amendments":[
                    {
                        "id": 3,
                        "title": "Amendment3 title",
                        "number": "Amendment3 123",
                        "amendment_paragraphs":{
                            "31":"amendmentParagraph31",
                            "32":"amendmentParagraph32"
                        },
                        "change_recommendations":[
                            {
                                "id": 4,
                                "rejected": true,
                                "type": "replacement",
                                "other_description": "ChangeRecommendation4 for amendment3",
                                "line_from": 4,
                                "line_to": 5,
                                "text": "<p>text4 HTML</p>",
                                "creation_time": 42345
                            }
                        ],
                        "merge_amendment_into_final": "do_merge",
                        "merge_amendment_into_diff": "do_merge"
                    },
                    {
                        "id": 4,
                        "title": "Amendment4 title",
                        "number": "Amendment4 4123",
                        "amendment_paragraphs": null,
                        "change_recommendations": null,
                        "merge_amendment_into_final": "undefined",
                        "merge_amendment_into_diff": "undefined"
                    },
                    {
                        "id": 5,
                        "title": "Amendment5 title",
                        "number": "Amendment5 5123",
                        "amendment_paragraphs": null,
                        "change_recommendations": null,
                        "merge_amendment_into_final": "undefined",
                        "merge_amendment_into_diff": "do_merge"
                    },
                    {
                        "id": 6,
                        "title": "Amendment6 title",
                        "number": "Amendment6 6123",
                        "amendment_paragraphs": null,
                        "change_recommendations": null,
                        "merge_amendment_into_final": "undefined",
                        "merge_amendment_into_diff": "undefined"
                    }
                ]
            }
            `,
		},
		{
			"motion including conditional fields",
			[]byte(`{"mode":"final"}`),
			changeData(data, map[dskey.Key][]byte{
				dskey.MustKey("meeting/1/motions_enable_text_on_projector"):           []byte(`true`),
				dskey.MustKey("meeting/1/motions_enable_reason_on_projector"):         []byte(`true`),
				dskey.MustKey("meeting/1/motions_show_referring_motions"):             []byte(`true`),
				dskey.MustKey("meeting/1/motions_enable_recommendation_on_projector"): []byte(`true`),
				dskey.MustKey("motion/1/lead_motion_id"):                              []byte(`2`),
				dskey.MustKey("motion/1/statute_paragraph_id"):                        []byte(`1`),
			}),
			`{
                "id":1,
                "title":"Motion Title 1",
                "number":"MNr1234",
                "submitters":[
                    "user11",
                    "user12",
                    "user13"
                ],
                "show_sidebox": true,
                "line_length": 85,
                "preamble": "The assembly may decide",
                "line_numbering": "outside",
                "amendment_paragraphs":{
                    "1":"amendmentParagraph1",
                    "2":"amendmentParagraph2"
                },
                "change_recommendations":[
                    {
                        "id": 1,
                        "rejected": true,
                        "type": "replacement",
                        "other_description": "Other Description1",
                        "line_from": 1,
                        "line_to": 3,
                        "text": "<p>text1 HTML</p>",
                        "creation_time": 12345
                    },
                    {
                        "id": 3,
                        "rejected": false,
                        "type": "insertion",
                        "other_description": "Other Description3",
                        "line_from": 5,
                        "line_to": 5,
                        "text": "<p>text3 HTML</p>",
                        "creation_time": 32345
                    }
                ],
                "amendments":[
                    {
                        "id": 3,
                        "title": "Amendment3 title",
                        "number": "Amendment3 123",
                        "amendment_paragraphs":{
                            "31":"amendmentParagraph31",
                            "32":"amendmentParagraph32"
                        },
                        "change_recommendations":[
                            {
                                "id": 4,
                                "rejected": true,
                                "type": "replacement",
                                "other_description": "ChangeRecommendation4 for amendment3",
                                "line_from": 4,
                                "line_to": 5,
                                "text": "<p>text4 HTML</p>",
                                "creation_time": 42345
                            }
                        ],
                        "merge_amendment_into_final": "do_merge",
                        "merge_amendment_into_diff": "do_merge"
                    },
                    {
                        "id": 4,
                        "title": "Amendment4 title",
                        "number": "Amendment4 4123",
                        "amendment_paragraphs": null,
                        "change_recommendations": null,
                        "merge_amendment_into_final": "undefined",
                        "merge_amendment_into_diff": "undefined"
                    },
                    {
                        "id": 5,
                        "title": "Amendment5 title",
                        "number": "Amendment5 5123",
                        "amendment_paragraphs": null,
                        "change_recommendations": null,
                        "merge_amendment_into_final": "undefined",
                        "merge_amendment_into_diff": "do_merge"
                    },
                    {
                        "id": 6,
                        "title": "Amendment6 title",
                        "number": "Amendment6 6123",
                        "amendment_paragraphs": null,
                        "change_recommendations": null,
                        "merge_amendment_into_final": "undefined",
                        "merge_amendment_into_diff": "undefined"
                    }
                ],
                "text":"<p>Motion1 Text HTML</p>",
                "reason": "<p>Motion1 reason HTML</p>",
                "modified_final_version":"<p>Motion1 modifiedFinalVersion HTML</p>",
                "lead_motion":{
                    "title":"Lead Motion Title",
                    "number":"Lead Motion 111",
                    "text":"<p>Lead Motion Text HTML</p>"
                },
                "base_statute":{
                    "title":"MotionStatuteParagraph1 title",
                    "text":"<p>MotionStatuteParagraph1 text html</p>"
                },
                "recommendation_referencing_motions":[
                    {
                        "agenda_item_number":"ItemNr Motion7",
                        "collection":"motion",
                        "content_object_id":"motion/7",
                        "title": "ReferencedInMotionRecommendationExtension7 title",
                        "number": "RIMRE7 number"
                    },
                    {
                        "agenda_item_number":"ItemNr Motion8",
                        "collection":"motion",
                        "content_object_id":"motion/8",
                        "title": "ReferencedInMotionRecommendationExtension8 title",
                        "number": "RIMRE8 number"
                    }
                ],
                "recommendation_label":"RecommendationLabel_state4",
                "recommendation_extension":"RecommendationExtension_motion1",
                "recommendation_referenced_motions":{
                    "motion/9":{
                        "agenda_item_number":"ItemNr Motion9",
                        "collection":"motion",
                        "content_object_id":"motion/9",
                        "title": "RecommendationExtensionReferenceMotion9 title",
                        "number": "RecommendationExtensionReferenceMotion9 number"
                    },
                    "motion/10":{
                        "agenda_item_number":"ItemNr Motion10",
                        "collection":"motion",
                        "content_object_id":"motion/10",
                        "title": "RecommendationExtensionReferenceMotion10 title",
                        "number": "RecommendationExtensionReferenceMotion10 number"
                    }
                },
                "recommender": "Meeting MotionsStatuteRecommendations"
            }
            `,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			ds := dsmock.Stub(tt.data)
			fetch := datastore.NewFetcher(ds)

			p7on := &projector.Projection{
				ContentObjectID: "motion/1",
				MeetingID:       1,
				Options:         tt.options,
			}

			bs, err := motionSlide.Slide(ctx, fetch, p7on)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expect, string(bs))
		})
	}
}
