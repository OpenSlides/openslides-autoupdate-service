package projector

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/models"
)

// SlideStore holds the slides by name.
type SlideStore struct {
	slides map[string]Slider
}

// Add registers a slide for a name.
func (s *SlideStore) Add(name string, slide Slider) {
	if s.slides == nil {
		s.slides = make(map[string]Slider)
	}

	if _, ok := s.slides[name]; ok {
		panic(fmt.Sprintf("Slide with name %s does already exist", name))
	}
	s.slides[name] = slide
}

// AddFunc is a helper to add a SliderFunc.
func (s *SlideStore) AddFunc(name string, f SliderFunc) {
	s.Add(name, f)
}

// Get returns a Slide for a name.
func (s *SlideStore) Get(name string) Slider {
	return s.slides[name]
}

// Slider knows how to create a slide.
type Slider interface {
	Slide(ctx context.Context, ds Datastore, p7on *models.Projection) (encoded []byte, keys []string, err error)
}

// SliderFunc is a function that implements the Slider interface.
type SliderFunc func(ctx context.Context, ds Datastore, p7on *models.Projection) (encoded []byte, keys []string, err error)

// Slide calls the func.
func (f SliderFunc) Slide(ctx context.Context, ds Datastore, p7on *models.Projection) (encoded []byte, keys []string, err error) {
	return f(ctx, ds, p7on)
}
