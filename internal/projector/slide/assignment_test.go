package slide_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/slide"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/dsmock"
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
		data   map[string]string
		expect string
	}{
		{
			"Assignment Complete",
			map[string]string{
				"assignment/1/id":                     `1`,
				"assignment/1/title":                  `"title 1"`,
				"assignment/1/description":            `"description 1"`,
				"assignment/1/number_poll_candidates": `true`,
				"assignment/1/candidate_ids":          `[10,11]`,

				"assignment_candidate/10/id":      `10`,
				"assignment_candidate/10/user_id": `110`,
				"assignment_candidate/10/weight":  `10`,
				"assignment_candidate/11/id":      `11`,
				"assignment_candidate/11/user_id": `111`,
				"assignment_candidate/11/weight":  `3`,

				"user/110/id":       "110",
				"user/110/username": `"user110"`,
				"user/111/id":       "111",
				"user/111/username": `"user111"`,
			},
			`{"title":"title 1", "description":"description 1","number_poll_candidates":true, "candidates":["user111", "user110"]}`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			shutdownCtx, cancel := context.WithCancel(context.Background())
			defer cancel()

			fetch := datastore.NewFetcher(dsmock.NewMockDatastore(shutdownCtx.Done(), convertData(tt.data)))

			p7on := &projector.Projection{
				ContentObjectID: "assignment/1",
			}

			bs, err := assignmentSlide.Slide(context.Background(), fetch, p7on)
			assert.NoError(t, err)
			assert.NoError(t, fetch.Err())
			assert.JSONEq(t, tt.expect, string(bs))
		})
	}
}
