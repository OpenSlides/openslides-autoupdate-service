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

func TestMotionBlock(t *testing.T) {
	s := new(projector.SlideStore)
	slide.MotionBlock(s)
	slide.Motion(s)

	motionBlockSlide := s.GetSlider("motion_block")
	assert.NotNilf(t, motionBlockSlide, "Slide with name `motion_block` not found.")

	data := dsmock.YAMLData(`
	motion_block/1:
	    title: MotionBlock1 Title
	    motion_ids: [1,2]
	motion:
	    1:
	        title: Motion Title 1
	        number: MNr 123
	        recommendation_id: 1
	        recommendation_extension: RecommendationExtension_motion1
	        recommendation_extension_reference_ids: ["motion/3", "motion/4"]
	        meeting_id: 1
	        agenda_item_id: 1
	    2:
	        title: Motion Title 2
	        number: MNR 456
	        meeting_id: 1
	    3:
	        title: RecommendationExtensionReferenceMotion3 title
	        number: RecommendationExtensionReferenceMotion3 number
	        meeting_id: 1
	        agenda_item_id: 3
	    4:
	        title: RecommendationExtensionReferenceMotion4 title
	        number: RecommendationExtensionReferenceMotion4 number
	        meeting_id: 1
	        agenda_item_id: 4
	motion_state/1:
	    recommendation_label: RecommendationLabel_state1
	    css_class: Css-Class1
	    show_recommendation_extension_field: true
	agenda_item/1/item_number: ItemNr Motion1
	agenda_item/3/item_number: ItemNr Motion3
	agenda_item/4/item_number: ItemNr Motion4
	`)

	for _, tt := range []struct {
		name   string
		data   map[dskey.Key][]byte
		expect string
	}{
		{
			"MotionBlock Standard",
			data,
			`{
                "title":"MotionBlock1 Title",
                "motions":[
                    {
                        "title": "Motion Title 1",
                        "number": "MNr 123",
                        "agenda_item_number": "ItemNr Motion1",
                        "recommendation": {
                            "recommendation_label": "RecommendationLabel_state1",
                            "css_class": "Css-Class1"
                        },
                        "recommendation_extension": "RecommendationExtension_motion1"
                    },
                    {
                        "title": "Motion Title 2",
                        "number": "MNR 456"
                    }
                ],
                "referenced": {
                    "motion/3": {
                        "agenda_item_number": "ItemNr Motion3",
                        "title": "RecommendationExtensionReferenceMotion3 title",
                        "number": "RecommendationExtensionReferenceMotion3 number",
                        "collection": "motion",
                        "content_object_id": "motion/3"
                    },
                    "motion/4": {
                        "agenda_item_number": "ItemNr Motion4",
                        "title": "RecommendationExtensionReferenceMotion4 title",
                        "number": "RecommendationExtensionReferenceMotion4 number",
                        "collection": "motion",
                        "content_object_id": "motion/4"
                    }
                }
            }
            `,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			ds := dsmock.Stub(tt.data)
			fetch := datastore.NewFetcher(ds)

			p7on := &projector.Projection{
				ContentObjectID: "motion_block/1",
				MeetingID:       1,
			}

			bs, err := motionBlockSlide.Slide(ctx, fetch, p7on)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expect, string(bs))
		})
	}
}
