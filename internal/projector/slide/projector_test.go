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

func TestProjectorCountdown(t *testing.T) {
	s := new(projector.SlideStore)
	slide.ProjectorCountdown(s)

	pcSlide := s.GetSlider("projector_countdown")
	assert.NotNilf(t, pcSlide, "Slide with name `projector_countdown` not found.")

	data := dsmock.YAMLData(`
	projector_countdown/1:
		description: description text
		running: true
		countdown_time: 200.34
		meeting_id: 1
	meeting/1/projector_countdown_warning_time: 100
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
				"countdown_time":200.34,
			    "description":"description text",
				"running":true,
				"warning_time":100}`,
			[]string{
				"projector_countdown/1/id",
				"projector_countdown/1/description",
				"projector_countdown/1/running",
				"projector_countdown/1/countdown_time",
				"projector_countdown/1/meeting_id",
				"meeting/1/projector_countdown_warning_time",
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			closed := make(chan struct{})
			defer close(closed)
			ds := dsmock.NewMockDatastore(closed, tt.data)
			fetch := datastore.NewFetcher(ds)

			p7on := &projector.Projection{
				ContentObjectID: "projector_countdown/1",
			}

			bs, keys, err := pcSlide.Slide(context.Background(), fetch, p7on)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expect, string(bs))
			assert.ElementsMatch(t, tt.expectKeys, keys)
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
			fetch := datastore.NewFetcher(ds)

			p7on := &projector.Projection{
				ContentObjectID: "projector_message/1",
			}

			bs, keys, err := pmSlide.Slide(context.Background(), fetch, p7on)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expect, string(bs))
			assert.ElementsMatch(t, tt.expectKeys, keys)
		})
	}
}
