package slide_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/slide"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/dsmock"
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
	        amendment_paragraph_$: ["1", "2"]
	        amendment_paragraph_$1: amendmentParagraph1
	        amendment_paragraph_$2: amendmentParagraph2
	        change_recommendation_ids: [1,2,3]
	        amendment_ids: [3,4,5,6]
	        referenced_in_motion_recommendation_extension_ids: [7,8]
	        recommendation_id: 4
	        global_yes: false
	        global_no: true
	        global_abstain: false
	        option_ids: [1, 2]
	        is_pseudoanonymized: false
	        pollmethod: YNA
	        onehundred_percent_base: YNA
	        majority_method: simple
	        votesvalid: 2.000000
	        votesinvalid: 9.000000
	        votescast: 2.000000
	        global_option_id: 3
	        meeting_id: 1
	    2:
	        title: Lead Motion Title
	        number: Lead Motion 111
	        text: <p>Lead Motion Text HTML</p>
	    3:
	        title: Amendment3 title
	        number: Amendment3 123
	        amendment_paragraph_$: ["31", "32"]
	        amendment_paragraph_$31: amendmentParagraph31
	        amendment_paragraph_$32: amendmentParagraph32
	        state_id: 1
	    4:
	        title: Amendment4 title
	        number: Amendment4 4123
	        amendment_paragraph_$: []
	        state_id: 2
	    5:
	        title: Amendment5 title
	        number: Amendment5 5123
	        amendment_paragraph_$: []
	        state_id: 3
	        recommendation_id: 1
	    6:
	        title: Amendment6 title
	        number: Amendment6 6123
	        amendment_paragraph_$: []
	        state_id: 3
	        recommendation_id: 2
	    7:
	        title: ReferencedInMotionRecommendationExtension7 title
	        number: RIMRE7 number
	    8:
	        title: ReferencedInMotionRecommendationExtension8 title
	        number: RIMRE8 number
	    9:
	        title: RecommendationExtensionReferenceMotion9 title
	        number: RecommendationExtensionReferenceMotion9 number
	    10:
	        title: RecommendationExtensionReferenceMotion10 title
	        number: RecommendationExtensionReferenceMotion10 number
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
	        recommendation_extension: RecommendationExtension_state4
	        recommendation_extension_reference_ids: ["motion/9", "motion/10"]
	motion_submitter:
	    1:
	        weight: 100
	        user_id: 13
	        motion_id: 1
	    2:
	        weight: 2
	        user_id: 11
	        motion_id: 1
	    3:
	        weight: 30
	        user_id: 12
	        motion_id: 1
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
	`)

	for _, tt := range []struct {
		name       string
		options    []byte
		data       map[string]string
		expect     string
		expectKeys []string
	}{
		{
			"Motion only non conditional",
			nil,
			data,
			`{
                "id":1,
                "title":"Motion Title 1",
                "number":"MNr1234",
                "submitters":{
                    "2":"user11",
                    "3":"user12",
                    "1":"user13",
                },
                "amendment_paragraphs":{
                    "1":"amendmentParagraph1",
                    "2":"amendmentParagraph2",
                },
                "change_recommendations":{
                    "1":{
                        "rejected": true,
                        "type": "replacement",
                        "other_description": "Other Description1",
                        "line_from": 1,
                        "line_to": 3,
                        "text": "<p>text1 HTML</p>",
                        "creation_time": 12345,
                    },
                    "3":{
                        "rejected": false,
                        "type": "insertion",
                        "other_description": "Other Description3",
                        "line_from": 5,
                        "line_to": 5,
                        "text": "<p>text3 HTML</p>",
                        "creation_time": 32345,
                    },
                },
                "amendments":{
                    "3":{
                        "id": 3,
                        "title": "Amendment3 title",
                        "number": "Amendment3 123",
                        "amendment_paragraphs":{
                            "31":"amendmentParagraph31",
                            "32":"amendmentParagraph32",
                        },
                        "merge_amendment_into_final": "do_merge",
                        "merge_amendment_into_diff": "do_merge",
                    },
                    "4":{
                        "id": 4,
                        "title": "Amendment4 title",
                        "number": "Amendment4 4123",
                        "amendment_paragraphs":{},
                        "merge_amendment_into_final": "undefined",
                        "merge_amendment_into_diff": "undefined",
                    },
                    "5":{
                        "id": 5,
                        "title": "Amendment5 title",
                        "number": "Amendment5 5123",
                        "amendment_paragraphs":{},
                        "merge_amendment_into_final": "undefined",
                        "merge_amendment_into_diff": "do_merge",
                    },
                    "6":{
                        "id": 6,
                        "title": "Amendment6 title",
                        "number": "Amendment6 6123",
                        "amendment_paragraphs":{},
                        "merge_amendment_into_final": "undefined",
                        "merge_amendment_into_diff": "undefined"
                    }
                }
            }
            `,
			[]string{
				"poll/1/state",
				"poll/1/id",
			},
		},
		{
			"motion including conditional fields",
			[]byte(`{"mode":"final"}`),
			changeData(data, map[string]string{
				"meeting/1/motions_enable_text_on_projector":           `true`,
				"meeting/1/motions_enable_reason_on_projector":         `true`,
				"meeting/1/motions_show_referring_motions":             `true`,
				"meeting/1/motions_enable_recommendation_on_projector": `true`,
				"motion/1/lead_motion_id":                              `2`,
				"motion/1/statute_paragraph_id":                        `1`,
			}),
			`{
                "id":1,
                "abHierConditional": true,
                "text":"<p>Motion1 Text HTML</p>",
                "reason": "<p>Motion1 reason HTML</p>",
                "modified_final_version":"<p>Motion1 modifiedFinalVersion HTML</p>",
                "lead_motion":{
                    "title":"Lead Motion Title",
                    "number":"Lead Motion 111",
                    "text":"<p>Lead Motion Text HTML</p>",
                },
                "base_statute":{
                    "title":"MotionStatuteParagraph1 title",
                    "text":"<p>MotionStatuteParagraph1 text html</p>",
                },
                "recommendation_referencing_motions":{
                    "7":{
                        "title": "ReferencedInMotionRecommendationExtension7 title",
                        "number": "RIMRE7 number",
                    },
                    "8":{
                        "title": "ReferencedInMotionRecommendationExtension8 title",
                        "number": "RIMRE8 number",
                    },
                },
                "recommendation_label":"RecommendationLabel_state4",
                "recommendation_extension":"RecommendationExtension_state4",
                "recommendation_referenced_motions":{
                    "motion/9":{
                        "title": "RecommendationExtensionReferenceMotion9 title",
                        "number": "RecommendationExtensionReferenceMotion9 number",
                    },
                    "motion/10":{
                        "title": "RecommendationExtensionReferenceMotion10 title",
                        "number": "RecommendationExtensionReferenceMotion10 number",
                    },
                },
                "recommender": "Meeting MotionsStatuteRecommendations",
            }
            `,
			[]string{
				"poll/1/state",
				"poll/1/id",
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			closed := make(chan struct{})
			defer close(closed)
			ds := dsmock.NewMockDatastore(closed, tt.data)

			p7on := &projector.Projection{
				ContentObjectID: "motion/1",
				MeetingID:       1,
				Options:         tt.options,
			}

			bs, keys, err := motionSlide.Slide(context.Background(), ds, p7on)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expect, string(bs))
			assert.ElementsMatch(t, tt.expectKeys, keys)
		})
	}
}
