package projector

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Datastore gets values for keys and informs, if they change.
type Datastore interface {
	Get(ctx context.Context, keys ...string) ([]json.RawMessage, error)
	RegisterCalculatedField(field string, f func(ctx context.Context, key string, changed map[string]json.RawMessage) ([]byte, error))
}

// Register initializes a new projector.
func Register(ds Datastore, slides *SlideStore) {
	hotKeys := make(map[string][]string)
	ds.RegisterCalculatedField("projection/content", func(ctx context.Context, fqfield string, changed map[string]json.RawMessage) (bs []byte, err error) {
		if changed != nil {
			var needUpdate bool
			for _, k := range hotKeys[fqfield] {
				if _, ok := changed[k]; ok {
					needUpdate = true
					break
				}
			}
			if !needUpdate {
				old, err := ds.Get(ctx, fqfield)
				if err != nil {
					return nil, fmt.Errorf("getting old value: %w", err)
				}
				return old[0], nil
			}
		}

		var keys []string
		defer func() {
			// At the end, save all requested keys to check later if one has
			// changed.
			hotKeys[fqfield] = keys
		}()

		parts := strings.SplitN(fqfield, "/", 3)
		if len(parts) != 3 {
			return nil, fmt.Errorf("invalid key %s, expected two '/'", fqfield)
		}

		data, keys, err := datastore.Object(
			ctx,
			ds,
			parts[0]+"/"+parts[1],
			[]string{
				"id",
				"type",
				"content_object_id",
				"meeting_id",
				"options",
			},
		)
		if err != nil {
			return nil, fmt.Errorf("fetching projection %s from datastore: %w", parts[1], err)
		}

		p7on, err := p7onFromMap(data)
		if err != nil {
			return nil, fmt.Errorf("loading p7on: %w", err)
		}

		if !p7on.exists() {
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

		bs, slideKeys, err := slider.Slide(ctx, ds, p7on)
		if err != nil {
			return nil, fmt.Errorf("calculating slide %s for p7on %v: %w", slideName, p7on, err)
		}
		keys = append(keys, slideKeys...)
		return bs, nil
	})
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
		return nil, fmt.Errorf("encoding projection data")
	}

	var p Projection
	if err := json.Unmarshal(bs, &p); err != nil {
		return nil, fmt.Errorf("decoding projection: %w", err)
	}
	return &p, nil
}

func (p *Projection) exists() bool {
	return p.Type != "" || p.ContentObjectID != ""
}

// slideName extracts the name from Projection.
// Using Type as slideName is only possible together with collection meeting,
// otherwise use always collection.
func (p *Projection) slideName() (string, error) {
	parts := strings.Split(p.ContentObjectID, "/")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid content_object_id `%s`, expected one '/'", p.ContentObjectID)
	}

	if p.Type != "" && parts[0] == "meeting" {
		return p.Type, nil
	}
	return parts[0], nil
}

// GetProjection as convenience function reading projections
func GetProjection(ctx context.Context, fetch *datastore.Fetcher, projectionID int) (projection *Projection, err error) {
	defer func() {
		if err == nil {
			err = fetch.Error()
		}
	}()

	data := fetch.Object(ctx, []string{"id", "type", "content_object_id", "meeting_id", "options"}, "projection/%d", projectionID)

	projection, err = p7onFromMap(data)
	if err != nil {
		return nil, fmt.Errorf("get projection: %w", err)
	}
	return projection, nil
}
