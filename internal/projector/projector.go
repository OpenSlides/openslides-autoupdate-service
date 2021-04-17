package projector

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/openslides/openslides-autoupdate-service/pkg/datastore"
)

// Datastore gets values for keys and informs, if they change.
type Datastore interface {
	Get(ctx context.Context, keys ...string) ([]json.RawMessage, error)
	RegisterCalculatedField(field string, f func(ctx context.Context, key string, changed map[string]json.RawMessage) ([]byte, error))
}

// Register initializes a new projector.
func Register(ds Datastore, slides *SlideStore) {
	hotKeys := make(map[string][]string)
	ds.RegisterCalculatedField("projection/content", func(ctx context.Context, fqfield string, changed map[string]json.RawMessage) ([]byte, error) {
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

		parts := strings.SplitN(fqfield, "/", 3)
		if len(parts) != 3 {
			return nil, fmt.Errorf("invalid key %s, expected two '/'", fqfield)
		}

		var p7on Projection
		keys, err := datastore.GetObject(ctx, ds, parts[0]+"/"+parts[1], &p7on)
		if err != nil {
			return nil, fmt.Errorf("fetching projection %s from datastore: %w", parts[1], err)
		}

		if !p7on.exists() {
			return nil, nil
		}

		slideName, err := p7on.slideName()
		if err != nil {
			return nil, fmt.Errorf("getting slide name: %w", err)
		}

		slider := slides.Get(slideName)
		if slider == nil {
			return nil, fmt.Errorf("unknown slide %s", slideName)
		}

		bs, slideKeys, err := slider.Slide(context.Background(), ds, &p7on)
		if err != nil {
			return nil, fmt.Errorf("calculating slide: %w", err)
		}
		keys = append(keys, slideKeys...)
		hotKeys[fqfield] = keys
		return bs, nil
	})
}

// Projection holds the meta data to render a projection on a projecter.
type Projection struct {
	ID              int    `json:"id"`
	Type            string `json:"type"`
	ContentObjectID string `json:"content_object_id"`
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
