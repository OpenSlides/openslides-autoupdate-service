package slide_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/slide"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
	"github.com/stretchr/testify/assert"
)

func TestTopic(t *testing.T) {
	s := new(projector.SlideStore)
	slide.Topic(s)

	topicSlide := s.GetSlider("topic")
	assert.NotNilf(t, topicSlide, "Slide with name `topic` not found.")

	for _, tt := range []struct {
		name   string
		data   map[string]string
		expect string
	}{
		{
			"Topic Complete",
			map[string]string{
				"topic/1/id":                `1`,
				"topic/1/title":             `"topic title 1"`,
				"topic/1/text":              `"topic text 1"`,
				"topic/1/agenda_item_id":    `1`,
				"agenda_item/1/id":          "1",
				"agenda_item/1/item_number": `"AI-Item 1"`,
			},
			`{"title":"topic title 1","text":"topic text 1","item_number":"AI-Item 1"}`,
		},
		{
			"Without Agenda Item",
			map[string]string{
				"topic/1/id":    `1`,
				"topic/1/title": `"topic title 1"`,
				"topic/1/text":  `"topic text 1"`,
			},
			`{"item_number":"", "text":"topic text 1", "title":"topic title 1"}`,
		},
		{
			"Agenda Item without number",
			map[string]string{
				"topic/1/id":             `1`,
				"topic/1/title":          `"topic title 1"`,
				"topic/1/text":           `"topic text 1"`,
				"agenda_item/1/id":       "1",
				"topic/1/agenda_item_id": `1`,
			},
			`{"title":"topic title 1","text":"topic text 1","item_number":""}`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ds := dsmock.Stub(convertData(tt.data))
			fetch := datastore.NewFetcher(ds)

			p7on := &projector.Projection{
				ContentObjectID: "topic/1",
			}

			bs, err := topicSlide.Slide(context.Background(), fetch, p7on)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expect, string(bs))
		})
	}
}
