package projector

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/datastore"
)

// Datastore gets values for keys and informs, if they change.
type Datastore interface {
	Get(ctx context.Context, keys ...datastore.Key) (map[datastore.Key][]byte, error)
	RegisterCalculatedField(field string, f func(ctx context.Context, key datastore.Key, changed map[datastore.Key][]byte) ([]byte, error))
}

// Register initializes a new projector.
func Register(ds Datastore, slides *SlideStore) {
	hotKeys := map[datastore.Key]map[datastore.Key]struct{}{}
	ds.RegisterCalculatedField("projection/content", func(ctx context.Context, fqfield datastore.Key, changed map[datastore.Key][]byte) (bs []byte, err error) {
		if changed != nil {
			var needUpdate bool
			for k := range changed {
				if _, ok := hotKeys[fqfield][k]; ok {
					needUpdate = true
					break
				}
			}
			if !needUpdate {
				old, err := ds.Get(ctx, fqfield)
				if err != nil {
					return nil, fmt.Errorf("getting old value: %w", err)
				}
				return old[fqfield], nil
			}
		}

		recorder := datastore.NewRecorder(ds)
		fetch := datastore.NewFetcher(recorder)

		defer func() {
			// At the end, save all requested keys to check later if one has
			// changed.
			hotKeys[fqfield] = recorder.Keys()
		}()

		data := fetch.Object(
			ctx,
			fqfield.FQID(),
			"id",
			"type",
			"content_object_id",
			"meeting_id",
			"options",
		)
		if err := fetch.Err(); err != nil {
			var errDoesNotExist datastore.DoesNotExistError
			if errors.As(err, &errDoesNotExist) {

				return nil, nil
			}
			return nil, fmt.Errorf("fetching projection %d from datastore: %w", fqfield.ID, err)
		}

		p7on, err := p7onFromMap(data)
		if err != nil {
			return nil, fmt.Errorf("loading p7on: %w", err)
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

		slider := slides.GetSlider(slideName)
		if slider == nil {
			return nil, fmt.Errorf("unknown slide %s", slideName)
		}

		bs, err = slider.Slide(ctx, fetch, p7on)
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
	})
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
	ID              int             `json:"id"`
	Type            string          `json:"type"`
	ContentObjectID string          `json:"content_object_id"`
	MeetingID       int             `json:"meeting_id"`
	Options         json.RawMessage `json:"options"`
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
