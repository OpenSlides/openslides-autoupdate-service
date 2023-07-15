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

func TestAssignment(t *testing.T) {
	s := new(projector.SlideStore)
	slide.Assignment(s)
	slide.User(s)

	assignmentSlide := s.GetSlider("assignment")
	assert.NotNilf(t, assignmentSlide, "Slide with name `assignment` not found.")

	for _, tt := range []struct {
		name   string
		data   map[dskey.Key][]byte
		expect string
	}{
		{
			"Assignment Complete",
			dsmock.YAMLData(`---
			assignment/1:
				id:                     1
				title:                  "title 1"
				description:            "description 1"
				number_poll_candidates: true
				candidate_ids:          [10,11]

			assignment_candidate:
				10:
					id:      10
					meeting_user_id: 1100
					weight:  10
				11:
					id:      11
					meeting_user_id: 1110
					weight:  3
			
			meeting_user:
				1100:
					user_id: 110
				1110:
					user_id: 111

			user:
				110:
					id:       110
					username: "user110"
				111:
					id:       111
					username: "user111"
			`),
			`{"title":"title 1", "description":"description 1","number_poll_candidates":true, "candidates":["user111", "user110"]}`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			ds := dsmock.NewFlow(tt.data)
			fetch := datastore.NewFetcher(ds)

			p7on := &projector.Projection{
				ContentObjectID: "assignment/1",
			}

			bs, err := assignmentSlide.Slide(ctx, fetch, p7on)
			assert.NoError(t, err)
			assert.NoError(t, fetch.Err())
			assert.JSONEq(t, tt.expect, string(bs))
		})
	}
}
