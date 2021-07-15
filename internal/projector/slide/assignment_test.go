package slide_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/slide"
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
		name         string
		data         map[string]string
		expect       string
		expectedKeys []string
	}{
		{
			"Assignment Complete",
			map[string]string{
				"assignment/1/id":                     `1`,
				"assignment/1/title":                  `"title 1"`,
				"assignment/1/open_posts":             `10`,
				"assignment/1/description":            `"description 1"`,
				"assignment/1/number_poll_candidates": `true`,
				"assignment/1/candidate_ids":          `[10,11]`,

				"assignment_candidate/10/id":      `10`,
				"assignment_candidate/10/user_id": `110`,
				"assignment_candidate/10/weight":  `10`,
				"assignment_candidate/11/id":      `11`,
				"assignment_candidate/11/user_id": `111`,
				"assignment_candidate/11/weight":  `3`,

				"user/110/username": `"user110"`,
				"user/111/username": `"user111"`,
			},
			`{"title":"title 1","phase":"","open_posts":10, "description":"description 1","number_poll_candidates":true, "candidates":["user111", "user110"]}`,
			[]string{
				"assignment/1/id",
				"assignment/1/title",
				"assignment/1/phase",
				"assignment/1/open_posts",
				"assignment/1/description",
				"assignment/1/number_poll_candidates",
				"assignment/1/candidate_ids",
				"assignment_candidate/10/user_id",
				"assignment_candidate/10/weight",
				"assignment_candidate/11/user_id",
				"assignment_candidate/11/weight",
				"user/111/username",
				"user/111/title",
				"user/111/first_name",
				"user/111/last_name",
				"user/111/default_structure_level",
				"user/110/username",
				"user/110/title",
				"user/110/first_name",
				"user/110/last_name",
				"user/110/default_structure_level",
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			closed := make(chan struct{})
			defer close(closed)
			ds := dsmock.NewMockDatastore(closed, tt.data)

			p7on := &projector.Projection{
				ContentObjectID: "assignment/1",
			}

			bs, keys, err := assignmentSlide.Slide(context.Background(), ds, p7on)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expect, string(bs))
			assert.ElementsMatch(t, keys, tt.expectedKeys)
		})
	}
}
