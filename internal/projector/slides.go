package projector

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Slider knows how to create a slide.
type Slider interface {
	Slide(ctx context.Context, ds Datastore, p7on *Projection) (encoded []byte, keys []string, err error)
}

// Titler defines the interface for GetTitleInformation-function, used for individual objects.
type Titler interface {
	GetTitleInformation(ctx context.Context, fetch *datastore.Fetcher, fqid string, itemNumber string) (json.RawMessage, error)
}

// SliderFunc is a function that implements the Slider interface.
type SliderFunc func(ctx context.Context, ds Datastore, p7on *Projection) (encoded []byte, keys []string, err error)

// Slide calls the func.
func (f SliderFunc) Slide(ctx context.Context, ds Datastore, p7on *Projection) (encoded []byte, keys []string, err error) {
	return f(ctx, ds, p7on)
}

// TitlerFunc is a type that implements the Titler interface.
type TitlerFunc func(ctx context.Context, fetch *datastore.Fetcher, fqid string, itemNumber string) (json.RawMessage, error)

// GetTitleInformation calls the func.
func (f TitlerFunc) GetTitleInformation(ctx context.Context, fetch *datastore.Fetcher, fqid string, itemNumber string) (json.RawMessage, error) {
	return f(ctx, fetch, fqid, itemNumber)
}

// SlideStore holds the Slider- and Titler-functions by name.
type SlideStore struct {
	slides map[string]Slider
	titles map[string]Titler
}

// RegisterSliderFunc adds a SliderFunc to the store.
func (s *SlideStore) RegisterSliderFunc(name string, f SliderFunc) {
	if s.slides == nil {
		s.slides = make(map[string]Slider)
	}

	if _, ok := s.slides[name]; ok {
		panic(fmt.Sprintf("Slide with name %s does already exist", name))
	}
	s.slides[name] = f
}

// GetSlider returns the Slide for the given name.
//
// Returns nil, if there if the name is unknown.
func (s *SlideStore) GetSlider(name string) Slider {
	return s.slides[name]
}

// RegisterGetTitleInformationFunc adds a function of type TitlerFunc to the store.
func (s *SlideStore) RegisterGetTitleInformationFunc(collection string, f TitlerFunc) {
	if s.titles == nil {
		s.titles = make(map[string]Titler)
	}

	if _, ok := s.titles[collection]; ok {
		panic(fmt.Sprintf("GetTitleInformation function for collection %s does already exist", collection))
	}
	s.titles[collection] = f
}

// GetTitleInformationFunc returns a Titler-function for the given name.
//
// Returns nil, if the name is unknown.
func (s *SlideStore) GetTitleInformationFunc(name string) Titler {
	return s.titles[name]
}
