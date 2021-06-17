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

// AgendaTitler returns the needed information to parse the agenda item of an element.
type AgendaTitler interface {
	AgendaTitle(ctx context.Context, fetch *datastore.Fetcher, fqid string, itemNumber string) (json.RawMessage, error)
}

// SliderFunc is a function that implements the Slider interface.
type SliderFunc func(ctx context.Context, ds Datastore, p7on *Projection) (encoded []byte, keys []string, err error)

// Slide calls the func.
func (f SliderFunc) Slide(ctx context.Context, ds Datastore, p7on *Projection) (encoded []byte, keys []string, err error) {
	return f(ctx, ds, p7on)
}

// AgendaTitlerFunc is a function that implements the AgendaTitler interface.
type AgendaTitlerFunc func(ctx context.Context, fetch *datastore.Fetcher, fqid string, itemNumber string) (json.RawMessage, error)

// AgendaTitle calls the func.
func (f AgendaTitlerFunc) AgendaTitle(ctx context.Context, fetch *datastore.Fetcher, fqid string, itemNumber string) (json.RawMessage, error) {
	return f(ctx, fetch, fqid, itemNumber)
}

// SlideStore holds the Sliders and AgendaTitler by name.
type SlideStore struct {
	slides       map[string]Slider
	agendaTitler map[string]AgendaTitler
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

// RegisterAgendaTitlerFunc adds a AgendaTitlerFunc to the store.
func (s *SlideStore) RegisterAgendaTitlerFunc(collection string, f AgendaTitlerFunc) {
	if s.agendaTitler == nil {
		s.agendaTitler = make(map[string]AgendaTitler)
	}

	if _, ok := s.agendaTitler[collection]; ok {
		panic(fmt.Sprintf("GetTitle function for collection %s does already exist", collection))
	}
	s.agendaTitler[collection] = f
}

// GetAgendaTitler returns a AgendaTitler for the given name.
//
// Returns nil, if there if the name is unknown.
func (s *SlideStore) GetAgendaTitler(name string) AgendaTitler {
	return s.agendaTitler[name]
}
