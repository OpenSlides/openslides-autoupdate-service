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

func TestProjectorCountdown(t *testing.T) {
	s := new(projector.SlideStore)
	slide.ProjectorCountdown(s)

	pcSlide := s.GetSlider("projector_countdown")
	assert.NotNilf(t, pcSlide, "Slide with name `projector_countdown` not found.")

	data := dsmock.YAMLData(`
	projector_countdown/1:
		description: description text
		running: true
		countdown_time: 200.3445678
		meeting_id: 1
	meeting/1/projector_countdown_warning_time: 100
    `)

	for _, tt := range []struct {
		name   string
		data   map[dskey.Key][]byte
		expect string
	}{
		{
			"Starter",
			data,
			`{
				"countdown_time":200.3445678,
				"description":"description text",
				"running":true,
				"warning_time":100}`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ds := dsmock.Stub(tt.data)
			fetch := datastore.NewFetcher(ds)

			p7on := &projector.Projection{
				ContentObjectID: "projector_countdown/1",
			}

			bs, err := pcSlide.Slide(context.Background(), fetch, p7on)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expect, string(bs))
		})
	}
}

func TestProjectorMessage(t *testing.T) {
	s := new(projector.SlideStore)
	slide.ProjectorMessage(s)

	pmSlide := s.GetSlider("projector_message")
	assert.NotNilf(t, pmSlide, "Slide with name `projector_message` not found.")

	data := dsmock.YAMLData(`
	projector_message/1/message: Shine on you crazy diamond
    `)

	for _, tt := range []struct {
		name   string
		data   map[dskey.Key][]byte
		expect string
	}{
		{
			"Starter",
			data,
			`{"message": "Shine on you crazy diamond"}`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ds := dsmock.Stub(tt.data)
			fetch := datastore.NewFetcher(ds)

			p7on := &projector.Projection{
				ContentObjectID: "projector_message/1",
			}

			bs, err := pmSlide.Slide(context.Background(), fetch, p7on)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expect, string(bs))
		})
	}
}
