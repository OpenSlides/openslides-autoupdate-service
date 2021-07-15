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
	        recommendation_extension: RecommendationExtension_motion1
	        recommendation_extension_reference_ids: ["motion/9", "motion/10"]
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
                        "merge_amendment_into_final": "do_merge",
                        "merge_amendment_into_diff": "do_merge"
                    },
                    {
                        "id": 4,
                        "title": "Amendment4 title",
                        "number": "Amendment4 4123",
                        "amendment_paragraphs": null,
                        "merge_amendment_into_final": "undefined",
                        "merge_amendment_into_diff": "undefined"
                    },
                    {
                        "id": 5,
                        "title": "Amendment5 title",
                        "number": "Amendment5 5123",
                        "amendment_paragraphs": null,
                        "merge_amendment_into_final": "undefined",
                        "merge_amendment_into_diff": "do_merge"
                    },
                    {
                        "id": 6,
                        "title": "Amendment6 title",
                        "number": "Amendment6 6123",
                        "amendment_paragraphs": null,
                        "merge_amendment_into_final": "undefined",
                        "merge_amendment_into_diff": "undefined"
                    }
                ]
            }
            `,
			[]string{
				"meeting/1/motions_enable_text_on_projector",
				"meeting/1/motions_enable_reason_on_projector",
				"meeting/1/motions_show_referring_motions",
				"meeting/1/motions_enable_recommendation_on_projector",
				"meeting/1/motions_statute_recommendations_by",
				"meeting/1/motions_recommendations_by",
				"meeting/1/motions_enable_sidebox_on_projector",
				"meeting/1/motions_line_length",
				"meeting/1/motions_preamble",
				"meeting/1/motions_default_line_numbering",
				"motion/1/id",
				"motion/1/title",
				"motion/1/number",
				"motion/1/meeting_id",
				"motion/1/lead_motion_id",
				"motion/1/statute_paragraph_id",
				"motion/1/amendment_paragraph_$",
				"motion/1/change_recommendation_ids",
				"motion/1/amendment_ids",
				"motion/1/submitter_ids",
				"motion/1/referenced_in_motion_recommendation_extension_ids",
				"motion/1/recommendation_id",
				"motion/1/recommendation_extension",
				"motion/1/recommendation_extension_reference_ids",
				"motion/1/amendment_paragraph_$1",
				"motion/1/amendment_paragraph_$2",
				"motion_submitter/1/user_id",
				"motion_submitter/1/weight",
				"motion_submitter/2/user_id",
				"motion_submitter/2/weight",
				"motion_submitter/3/user_id",
				"motion_submitter/3/weight",
				"user/11/username",
				"user/11/title",
				"user/11/first_name",
				"user/11/last_name",
				"user/11/default_structure_level",
				"user/11/structure_level_$1",
				"user/12/username",
				"user/12/title",
				"user/12/first_name",
				"user/12/last_name",
				"user/12/default_structure_level",
				"user/12/structure_level_$1",
				"user/13/username",
				"user/13/title",
				"user/13/first_name",
				"user/13/last_name",
				"user/13/default_structure_level",
				"user/13/structure_level_$1",
				"motion_change_recommendation/1/id",
				"motion_change_recommendation/1/rejected",
				"motion_change_recommendation/1/type",
				"motion_change_recommendation/1/other_description",
				"motion_change_recommendation/1/line_from",
				"motion_change_recommendation/1/line_to",
				"motion_change_recommendation/1/text",
				"motion_change_recommendation/1/creation_time",
				"motion_change_recommendation/1/internal",
				"motion_change_recommendation/2/id",
				"motion_change_recommendation/2/rejected",
				"motion_change_recommendation/2/type",
				"motion_change_recommendation/2/other_description",
				"motion_change_recommendation/2/line_from",
				"motion_change_recommendation/2/line_to",
				"motion_change_recommendation/2/text",
				"motion_change_recommendation/2/creation_time",
				"motion_change_recommendation/2/internal",
				"motion_change_recommendation/3/id",
				"motion_change_recommendation/3/rejected",
				"motion_change_recommendation/3/type",
				"motion_change_recommendation/3/other_description",
				"motion_change_recommendation/3/line_from",
				"motion_change_recommendation/3/line_to",
				"motion_change_recommendation/3/text",
				"motion_change_recommendation/3/creation_time",
				"motion_change_recommendation/3/internal",
				"motion/3/id",
				"motion/3/title",
				"motion/3/number",
				"motion/3/meeting_id",
				"motion/3/amendment_paragraph_$",
				"motion/3/state_id",
				"motion/3/recommendation_id",
				"motion/3/amendment_paragraph_$31",
				"motion/3/amendment_paragraph_$32",
				"motion_state/1/merge_amendment_into_final",
				"motion/4/id",
				"motion/4/title",
				"motion/4/number",
				"motion/4/meeting_id",
				"motion/4/amendment_paragraph_$",
				"motion/4/state_id",
				"motion/4/recommendation_id",
				"motion_state/2/merge_amendment_into_final",
				"motion/5/id",
				"motion/5/title",
				"motion/5/number",
				"motion/5/meeting_id",
				"motion/5/amendment_paragraph_$",
				"motion/5/state_id",
				"motion/5/recommendation_id",
				"motion_state/3/merge_amendment_into_final",
				"motion_state/1/merge_amendment_into_final",
				"motion/6/id",
				"motion/6/title",
				"motion/6/number",
				"motion/6/meeting_id",
				"motion/6/amendment_paragraph_$",
				"motion/6/state_id",
				"motion/6/recommendation_id",
				"motion_state/3/merge_amendment_into_final",
				"motion_state/2/merge_amendment_into_final",
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
                        "merge_amendment_into_final": "do_merge",
                        "merge_amendment_into_diff": "do_merge"
                    },
                    {
                        "id": 4,
                        "title": "Amendment4 title",
                        "number": "Amendment4 4123",
                        "amendment_paragraphs": null,
                        "merge_amendment_into_final": "undefined",
                        "merge_amendment_into_diff": "undefined"
                    },
                    {
                        "id": 5,
                        "title": "Amendment5 title",
                        "number": "Amendment5 5123",
                        "amendment_paragraphs": null,
                        "merge_amendment_into_final": "undefined",
                        "merge_amendment_into_diff": "do_merge"
                    },
                    {
                        "id": 6,
                        "title": "Amendment6 title",
                        "number": "Amendment6 6123",
                        "amendment_paragraphs": null,
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
                        "agenda_item_number":"",
                        "collection":"motion",
                        "content_object_id":"motion/7",
                        "title": "ReferencedInMotionRecommendationExtension7 title",
                        "number": "RIMRE7 number"
                    },
                    {
                        "agenda_item_number":"",
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
                        "agenda_item_number":"",
                        "collection":"motion",
                        "content_object_id":"motion/9",
                        "title": "RecommendationExtensionReferenceMotion9 title",
                        "number": "RecommendationExtensionReferenceMotion9 number"
                    },
                    "motion/10":{
                        "agenda_item_number":"",
                        "collection":"motion",
                        "content_object_id":"motion/10",
                        "title": "RecommendationExtensionReferenceMotion10 title",
                        "number": "RecommendationExtensionReferenceMotion10 number"
                    }
                },
                "recommender": "Meeting MotionsStatuteRecommendations"
            }
            `,
			[]string{
				"meeting/1/motions_enable_text_on_projector",
				"meeting/1/motions_enable_reason_on_projector",
				"meeting/1/motions_show_referring_motions",
				"meeting/1/motions_enable_recommendation_on_projector",
				"meeting/1/motions_statute_recommendations_by",
				"meeting/1/motions_recommendations_by",
				"meeting/1/motions_enable_sidebox_on_projector",
				"meeting/1/motions_line_length",
				"meeting/1/motions_preamble",
				"meeting/1/motions_default_line_numbering",
				"motion/1/id",
				"motion/1/title",
				"motion/1/number",
				"motion/1/meeting_id",
				"motion/1/lead_motion_id",
				"motion/1/statute_paragraph_id",
				"motion/1/amendment_paragraph_$",
				"motion/1/change_recommendation_ids",
				"motion/1/amendment_ids",
				"motion/1/submitter_ids",
				"motion/1/referenced_in_motion_recommendation_extension_ids",
				"motion/1/recommendation_id",
				"motion/1/recommendation_extension",
				"motion/1/recommendation_extension_reference_ids",
				"motion/1/text",
				"motion/1/reason",
				"motion/1/modified_final_version",
				"motion/1/amendment_paragraph_$1",
				"motion/1/amendment_paragraph_$2",
				"motion_submitter/1/user_id",
				"motion_submitter/1/weight",
				"motion_submitter/2/user_id",
				"motion_submitter/2/weight",
				"motion_submitter/3/user_id",
				"motion_submitter/3/weight",
				"user/11/username",
				"user/11/title",
				"user/11/first_name",
				"user/11/last_name",
				"user/11/default_structure_level",
				"user/11/structure_level_$1",
				"user/12/username",
				"user/12/title",
				"user/12/first_name",
				"user/12/last_name",
				"user/12/default_structure_level",
				"user/12/structure_level_$1",
				"user/13/username",
				"user/13/title",
				"user/13/first_name",
				"user/13/last_name",
				"user/13/default_structure_level",
				"user/13/structure_level_$1",
				"motion/2/title",
				"motion/2/number",
				"motion/2/text",
				"motion_statute_paragraph/1/title",
				"motion_statute_paragraph/1/text",
				"motion_change_recommendation/1/id",
				"motion_change_recommendation/1/rejected",
				"motion_change_recommendation/1/type",
				"motion_change_recommendation/1/other_description",
				"motion_change_recommendation/1/line_from",
				"motion_change_recommendation/1/line_to",
				"motion_change_recommendation/1/text",
				"motion_change_recommendation/1/creation_time",
				"motion_change_recommendation/1/internal",
				"motion_change_recommendation/2/id",
				"motion_change_recommendation/2/rejected",
				"motion_change_recommendation/2/type",
				"motion_change_recommendation/2/other_description",
				"motion_change_recommendation/2/line_from",
				"motion_change_recommendation/2/line_to",
				"motion_change_recommendation/2/text",
				"motion_change_recommendation/2/creation_time",
				"motion_change_recommendation/2/internal",
				"motion_change_recommendation/3/id",
				"motion_change_recommendation/3/rejected",
				"motion_change_recommendation/3/type",
				"motion_change_recommendation/3/other_description",
				"motion_change_recommendation/3/line_from",
				"motion_change_recommendation/3/line_to",
				"motion_change_recommendation/3/text",
				"motion_change_recommendation/3/creation_time",
				"motion_change_recommendation/3/internal",
				"motion/3/id",
				"motion/3/title",
				"motion/3/number",
				"motion/3/meeting_id",
				"motion/3/amendment_paragraph_$",
				"motion/3/state_id",
				"motion/3/recommendation_id",
				"motion/3/amendment_paragraph_$31",
				"motion/3/amendment_paragraph_$32",
				"motion_state/1/merge_amendment_into_final",
				"motion/4/id",
				"motion/4/title",
				"motion/4/number",
				"motion/4/meeting_id",
				"motion/4/amendment_paragraph_$",
				"motion/4/state_id",
				"motion/4/recommendation_id",
				"motion_state/2/merge_amendment_into_final",
				"motion/5/id",
				"motion/5/title",
				"motion/5/number",
				"motion/5/meeting_id",
				"motion/5/amendment_paragraph_$",
				"motion/5/state_id",
				"motion/5/recommendation_id",
				"motion_state/3/merge_amendment_into_final",
				"motion_state/1/merge_amendment_into_final",
				"motion/6/id",
				"motion/6/title",
				"motion/6/number",
				"motion/6/meeting_id",
				"motion/6/amendment_paragraph_$",
				"motion/6/state_id",
				"motion/6/recommendation_id",
				"motion_state/3/merge_amendment_into_final",
				"motion_state/2/merge_amendment_into_final",
				"motion/7/id",
				"motion/7/number",
				"motion/7/title",
				"motion/8/id",
				"motion/8/number",
				"motion/8/title",
				"motion_state/4/recommendation_label",
				"motion_state/4/show_recommendation_extension_field",
				"motion/9/id",
				"motion/9/number",
				"motion/9/title",
				"motion/10/id",
				"motion/10/number",
				"motion/10/title",
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
