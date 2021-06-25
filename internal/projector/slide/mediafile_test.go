package slide_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/slide"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/dsmock"
	"github.com/stretchr/testify/assert"
)

func TestMediafile(t *testing.T) {
	s := new(projector.SlideStore)
	slide.Mediafile(s)

	mfSlide := s.GetSlider("mediafile")
	assert.NotNilf(t, mfSlide, "Slide with name `mediafile` not found.")

	data := dsmock.YAMLData(`
	mediafile/1/mimetype: application/pdf
    `)

	for _, tt := range []struct {
		name       string
		data       map[string]string
		expect     string
		expectKeys []string
	}{
		{
			"Starter",
			data,
			`{
				"id": 1,
				"mimetype": "application/pdf"
			}`,
			[]string{
				"mediafile/1/id",
				"mediafile/1/mimetype",
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			closed := make(chan struct{})
			defer close(closed)
			ds := dsmock.NewMockDatastore(closed, tt.data)

			p7on := &projector.Projection{
				ContentObjectID: "mediafile/1",
			}

			bs, keys, err := mfSlide.Slide(context.Background(), ds, p7on)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expect, string(bs))
			assert.ElementsMatch(t, tt.expectKeys, keys)
		})
	}
}
