package projector

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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
			//
			// If an error happend, don't return the error but log it and show
			// the user a generic error message.
			hotKeys[fqfield] = keys
			if err != nil {
				log.Printf("Error parsing slide %s: %v", fqfield, err)
				bs = []byte(fmt.Sprintf(`{"error":"Error parsing slide %s!"}`, fqfield))
				err = nil
			}
		}()

		parts := strings.SplitN(fqfield, "/", 3)
		if len(parts) != 3 {
			err = fmt.Errorf("invalid key %s, expected two '/'", fqfield)
			return bs, err
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
			err = fmt.Errorf("fetching projection %s from datastore: %w", parts[1], err)
			return bs, err
		}

		p7on, err := p7onFromMap(data)
		if err != nil {
			err = fmt.Errorf("loading p7on: %w", err)
			return bs, err
		}

		if !p7on.exists() {
			return nil, nil
		}

		slideName, err := p7on.slideName()
		if err != nil {
			err = fmt.Errorf("getting slide name: %w", err)
			return bs, err
		}

		slider := slides.GetSlider(slideName)
		if slider == nil {
			err = fmt.Errorf("unknown slide %s", slideName)
			return bs, err
		}

		bs, slideKeys, err := slider.Slide(context.Background(), ds, p7on)
		if err != nil {
			err = fmt.Errorf("calculating slide: %w", err)
			return bs, err
		}
		keys = append(keys, slideKeys...)
		return bs, nil
	})
}

// Projections option holds the key/value pairs from the options field of Projection
type ProjectionOptions struct {
	OnlyMainItems bool `json:"only_main_items"`
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

func (p *Projection) slideName() (string, error) {
	if p.Type != "" {
		return p.Type, nil
	}
	i := strings.Index(p.ContentObjectID, "/")
	if i == -1 {
		return "", fmt.Errorf("invalid content_object_id `%s`, expected one '/'", p.ContentObjectID)
	}
	return p.ContentObjectID[:i], nil
}
