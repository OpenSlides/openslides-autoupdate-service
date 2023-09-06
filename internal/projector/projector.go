package projector

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/oserror"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsrecorder"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/flow"
)

const longCalculation = time.Second

// NewProjector initializes a new Projector.
func NewProjector(ds flow.Flow, slides *SlideStore) *Projector {
	return &Projector{
		hotKeys: make(map[dskey.Key]map[dskey.Key]struct{}),
		cache:   make(map[dskey.Key][]byte),

		flow:   ds,
		slides: slides,
	}
}

// Projector is a Flow that adds the field projection/content
//
// When such a key is requested with Get, it gets calculated.
//
// When keys get updated via Update, that where needed to calculate a field, the
// field is updated.
//
// Only projections with a current_projector_id get calculated. Fields from
// other projections return nil. If current_projector_id get updated to nil/0,
// then the field is removed from the cache.
type Projector struct {
	mu      sync.RWMutex
	hotKeys map[dskey.Key]map[dskey.Key]struct{}
	cache   map[dskey.Key][]byte

	flow   flow.Flow
	slides *SlideStore
}

// Reset clears the projector object.
func (p *Projector) Reset() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.hotKeys = make(map[dskey.Key]map[dskey.Key]struct{})
}

// Get is a Getter middleware that passes all keys though but calculates
// projection/content keys.
func (p *Projector) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	normalKeys, contentKeys := splitKeys(keys)

	values, err := p.flow.Get(ctx, normalKeys...)
	if err != nil {
		return nil, fmt.Errorf("get from flow: %w", err)
	}

	if len(contentKeys) == 0 {
		return values, nil
	}

	p.mu.RLock()
	var needCalc []dskey.Key
	for _, k := range contentKeys {
		v, ok := p.cache[k]
		if !ok {
			needCalc = append(needCalc, k)
			continue
		}

		values[k] = v
	}
	p.mu.RUnlock()

	if len(needCalc) == 0 {
		return values, nil
	}

	p.mu.Lock()
	for _, k := range needCalc {
		v := p.calculate(ctx, k)
		p.cache[k] = v
		values[k] = v
	}
	p.mu.Unlock()

	return values, nil
}

// Update updates projection/content keys.
func (p *Projector) Update(ctx context.Context, updateFn func(map[dskey.Key][]byte, error)) {
	p.flow.Update(ctx, func(data map[dskey.Key][]byte, err error) {
		if err != nil {
			updateFn(nil, err)
			return
		}

		p.mu.Lock()
		defer p.mu.Unlock()

		needUpdate := p.needUpdate(data)

		if len(needUpdate) == 0 {
			updateFn(data, nil)
			return
		}

		for _, key := range needUpdate {
			value := p.calculate(ctx, key)
			data[key] = value
			p.cache[key] = value
		}

		updateFn(data, nil)
	})
}

func (p *Projector) needUpdate(data map[dskey.Key][]byte) []dskey.Key {
	var needUpdate []dskey.Key
	for calculated := range p.hotKeys {
		for key := range data {
			if _, ok := p.hotKeys[calculated][key]; ok {
				needUpdate = append(needUpdate, calculated)
				break
			}
		}
	}
	return needUpdate
}

func splitKeys(keys []dskey.Key) ([]dskey.Key, []dskey.Key) {
	var contentKeys []dskey.Key
	normalKeys := make([]dskey.Key, 0, len(keys))
	for _, k := range keys {
		if !(k.Collection() == "projection" && k.Field() == "content") {
			normalKeys = append(normalKeys, k)
			continue
		}

		contentKeys = append(contentKeys, k)
	}

	return normalKeys, contentKeys
}

func (p *Projector) calculate(ctx context.Context, fqfield dskey.Key) []byte {
	bs, err := p.calculateHelper(ctx, fqfield)
	if err != nil {
		oserror.Handle(fmt.Errorf("Error calculating key %s: %v", fqfield, err))
		msg := fmt.Sprintf("calculating key %s", fqfield)
		return []byte(fmt.Sprintf(`{"error": "%s"}`, msg))
	}

	return bs
}

func (p *Projector) calculateHelper(ctx context.Context, fqfield dskey.Key) ([]byte, error) {
	recorder := dsrecorder.New(p.flow)
	fetch := datastore.NewFetcher(recorder)

	defer func() {
		// At the end, save all requested keys to check later if one has
		// changed.
		p.hotKeys[fqfield] = recorder.Keys()
	}()

	data := fetch.Object(
		ctx,
		fqfield.FQID(),
		"id",
		"type",
		"content_object_id",
		"meeting_id",
		"options",
		"current_projector_id",
	)
	if err := fetch.Err(); err != nil {
		var errDoesNotExist datastore.DoesNotExistError
		if errors.As(err, &errDoesNotExist) {
			return nil, nil
		}
		return nil, fmt.Errorf("fetching projection %d from datastore: %w", fqfield.ID(), err)
	}

	p7on, err := p7onFromMap(data)
	if err != nil {
		return nil, fmt.Errorf("loading p7on: %w", err)
	}

	if p7on.CurrentProjectorID == 0 {
		return nil, nil
	}

	if p7on.ContentObjectID == "" {
		// There are broken projections in the datastore. Ignore them.
		log.Printf("Bug in Backend: The projection %d has an empty content_object_id", p7on.ID)
		return nil, nil
	}

	slideName, err := p7on.slideName()
	if err != nil {
		return nil, fmt.Errorf("getting slide name: %w", err)
	}

	slider := p.slides.GetSlider(slideName)
	if slider == nil {
		return nil, fmt.Errorf("unknown slide %s", slideName)
	}

	bs, err := slider.Slide(ctx, fetch, p7on)
	if err != nil {
		return nil, fmt.Errorf("calculating slide %s for p7on %v: %w", slideName, p7on, err)
	}

	if err := fetch.Err(); err != nil {
		return nil, err
	}

	final, err := addCollection(bs, slideName)
	if err != nil {
		return nil, fmt.Errorf("adding name of collection %q to value %q: %w", slideName, bs, err)
	}
	return final, nil
}

// addCollection adds the collection addribute to the given encoded json.
//
// `bs` has to be a encoded json-object. `collection` has to be a valid json
// string.
func addCollection(bs []byte, collection string) ([]byte, error) {
	var decoded map[string]json.RawMessage
	if err := json.Unmarshal(bs, &decoded); err != nil {
		return nil, fmt.Errorf("decoding object: %w", err)
	}

	decoded["collection"] = []byte(`"` + collection + `"`)

	bs, err := json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("encoding object: %w", err)
	}
	return bs, nil
}

// Projection holds the meta data to render a projection on a projecter.
type Projection struct {
	ID                 int             `json:"id"`
	Type               string          `json:"type"`
	ContentObjectID    string          `json:"content_object_id"`
	MeetingID          int             `json:"meeting_id"`
	Options            json.RawMessage `json:"options"`
	CurrentProjectorID int             `json:"current_projector_id"`
}

func p7onFromMap(in map[string]json.RawMessage) (*Projection, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding projection data: %w", err)
	}

	var p Projection
	if err := json.Unmarshal(bs, &p); err != nil {
		return nil, fmt.Errorf("decoding projection: %w", err)
	}
	return &p, nil
}

// slideName extracts the name from Projection.
// Using Type as slideName is only possible together with collection meeting,
// otherwise use always collection.
func (p *Projection) slideName() (string, error) {
	parts := strings.Split(p.ContentObjectID, "/")
	if len(parts) != 2 {
		// TODO LAST ERROR
		return "", fmt.Errorf("invalid content_object_id `%s`, expected one '/'", p.ContentObjectID)
	}

	if p.Type != "" && parts[0] == "meeting" {
		return p.Type, nil
	}
	return parts[0], nil
}
