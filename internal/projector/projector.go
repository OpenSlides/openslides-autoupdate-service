package projector

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/openslides/openslides-autoupdate-service/internal/datastore"
)

// Datastore gets values for keys and informs, if they change.
type Datastore interface {
	Get(ctx context.Context, keys ...string) ([]json.RawMessage, error)
	RegisterChangeListener(f func(map[string]json.RawMessage) error)
}

// Service builds and shows the projector.
type Service struct {
	ds     Datastore
	slides *SlideStore
	closed <-chan struct{}

	projectors map[int]*projector
}

// New initializes a new projector.
func New(ds Datastore, slides *SlideStore, closed <-chan struct{}) *Service {
	return &Service{
		ds:         ds,
		slides:     slides,
		closed:     closed,
		projectors: make(map[int]*projector),
	}
}

// Live writes all projector changes for the given projector ids on w.
func (p *Service) Live(ctx context.Context, uid int, w io.Writer, pids []int) error {
	out := make(map[int]json.RawMessage)
	for _, pid := range pids {
		// TODO: Check permission

		pr, ok := p.projectors[pid]
		if !ok {
			var err error
			pr, err = newProjector(ctx, p.ds, p.slides, pid)
			if err != nil {
				return fmt.Errorf("create projector %d: %w", pid, err)
			}
		}

		bs, err := pr.Bytes()
		if err != nil {
			return fmt.Errorf("get data for projector %d: %w", pid, err)
		}
		out[pid] = bs
	}

	if err := json.NewEncoder(w).Encode(out); err != nil {
		return fmt.Errorf("encode and write projectors: %w", err)
	}
	return nil
}

type projector struct {
	buf    []byte
	ds     Datastore
	slides *SlideStore
	id     int
}

func newProjector(ctx context.Context, ds Datastore, slides *SlideStore, id int) (*projector, error) {
	var projectionIDs []int
	if err := datastore.GetIfExist(ctx, ds, fmt.Sprintf("projector/%d/current_projection_ids", id), &projectionIDs); err != nil {
		return nil, fmt.Errorf("get projections: %w", err)
	}

	pr := &projector{
		id:     id,
		ds:     ds,
		slides: slides,
	}

	if err := pr.calc(ctx); err != nil {
		return nil, fmt.Errorf("calculate projector for first time: %w", err)
	}
	return pr, nil
}

func (pr *projector) calc(ctx context.Context) error {
	var p7onIDs []int
	if err := datastore.Get(ctx, pr.ds, fmt.Sprintf("projector/%d/current_projection_ids", pr.id), &p7onIDs); err != nil {
		var errDoesNotExist datastore.DoesNotExistError
		if errors.As(err, &errDoesNotExist) {
			// Projector does not exist.
			pr.buf = nil
			return nil
		}
		return fmt.Errorf("get projections: %w", err)
	}

	projectionsData := make(map[int]json.RawMessage)
	for _, id := range p7onIDs {
		var p7on Projection
		if err := datastore.GetObject(ctx, pr.ds, fmt.Sprintf("projection/%d", id), &p7on); err != nil {
			return fmt.Errorf("fetch projection: %w", err)
		}

		slideName, err := p7on.slideName()
		if err != nil {
			return fmt.Errorf("getting slide name: %w", err)
		}
		slide := pr.slides.Get(slideName)
		if slide == nil {
			return fmt.Errorf("unknown slide %s", slideName)
		}

		bs, err := p7on.calc(ctx, pr.ds, slide)
		if err != nil {
			return fmt.Errorf("calculate projection %d: %w", id, err)
		}
		projectionsData[id] = bs
	}

	bs, err := json.Marshal(projectionsData)
	if err != nil {
		return fmt.Errorf("encode projector %d: %w", pr.id, err)
	}
	pr.buf = bs

	return nil
}

func (pr *projector) Bytes() ([]byte, error) {
	return pr.buf, nil
}

// Projection holds the meta data to render a projection on a projecter.
type Projection struct {
	Option          json.RawMessage `json:"option,omitempty"`
	Stable          bool            `json:"stable"`
	Type            string          `json:"type,omitempty"`
	ContentObjectID string          `json:"content_object_id"`
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

func (p *Projection) calc(ctx context.Context, ds Datastore, slide Slider) ([]byte, error) {
	var outProjection struct {
		Projection
		Data json.RawMessage `json:"data"`
	}
	outProjection.Projection = *p

	slideData, _, err := slide.Slide(ctx, ds, p)
	if err != nil {
		return nil, fmt.Errorf("calculating slide: %w", err)
	}
	outProjection.Data = slideData

	bs, err := json.Marshal(outProjection)
	if err != nil {
		return nil, fmt.Errorf("decoding calculated projection: %w", err)
	}
	return bs, nil
}
