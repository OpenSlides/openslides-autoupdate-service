package projector

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// SlideStore holds the slides by name.
type SlideStore struct {
	slides map[string]Slider
	titles map[string]TitlerFunc
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

// RegisterSliderFunc is a helper to add a SliderFunc.
func (s *SlideStore) RegisterSlideFunc(name string, f SliderFunc) {
	s.Add(name, f)
}

// Get returns a Slide for a name.
func (s *SlideStore) GetSlideFunc(name string) Slider {
	return s.slides[name]
}

// Slider knows how to create a slide.
type Slider interface {
	Slide(ctx context.Context, ds Datastore, p7on *Projection) (encoded []byte, keys []string, err error)
}

type TitlerFuncResult struct {
	Collection       string  `json:"collection"`
	ContentObjectId  string  `json:"content_object_id"`
	Title            *string `json:"title,omitempty"`
	AgendaItemNumber *string `json:"agenda_item_number,omitempty"`
	Number           *string `json:"number,omitempty"`
	Username         *string `json:"username,omitempty"`
}

// SliderFunc is a function that implements the Slider interface.
type SliderFunc func(ctx context.Context, ds Datastore, p7on *Projection) (encoded []byte, keys []string, err error)
type TitlerFunc func(ctx context.Context, fetch *datastore.Fetcher, fqid string, meeting_id int, value map[string]interface{}) (title *TitlerFuncResult, err error)

// Slide calls the func.
func (f SliderFunc) Slide(ctx context.Context, ds Datastore, p7on *Projection) (encoded []byte, keys []string, err error) {
	return f(ctx, ds, p7on)
}

// RegisterTitleFunc registers a function for a collection name.
func (s *SlideStore) RegisterTitleFunc(collection string, f TitlerFunc) {
	if s.titles == nil {
		s.titles = make(map[string]TitlerFunc)
	}

	if _, ok := s.titles[collection]; ok {
		panic(fmt.Sprintf("GetTitle function for collection %s does already exist", collection))
	}
	s.titles[collection] = f
}

// Get returns a TitleFunc for a name.
func (s *SlideStore) GetTitleFunc(collection string) TitlerFunc {
	f := s.titles[collection]
	if f == nil {
		panic(fmt.Sprintf("There is no TitlerFunc registered for collection %s", collection))
	}
	return f
}
