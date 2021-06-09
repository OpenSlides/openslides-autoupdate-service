package slide_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/slide"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/dsmock"
	"github.com/stretchr/testify/assert"
)

func TestProjectorMessage(t *testing.T) {
	s := new(projector.SlideStore)
	slide.ProjectorMessage(s)

	pmSlide := s.Get("projector_message")
	assert.NotNilf(t, pmSlide, "Slide with name `projector_message` not found.")

	data := dsmock.YAMLData(`
	projector_message/1/message: Shine on you crazy diamond
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
			`{"message": "Shine on you crazy diamond"}`,
			[]string{
				"projector_message/1/id",
				"projector_message/1/message",
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			closed := make(chan struct{})
			defer close(closed)
			ds := dsmock.NewMockDatastore(closed, tt.data)

			p7on := &projector.Projection{
				ContentObjectID: "projector_message/1",
			}

			bs, keys, err := pmSlide.Slide(context.Background(), ds, p7on)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expect, string(bs))
			assert.ElementsMatch(t, tt.expectKeys, keys)
		})
	}
}
