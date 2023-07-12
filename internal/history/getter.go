package history

import (
	"context"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

// getPositioner is like a Getter but also taks a position
type getPositioner interface {
	GetPosition(ctx context.Context, position int, keys ...dskey.Key) (map[dskey.Key][]byte, error)
}

// getPosition translates a GetPositioner to a Getter.
type getPosition struct {
	position int
	getter   getPositioner
}

// newGetPosition initializes a GetPosition.
func newGetPosition(g getPositioner, position int) *getPosition {
	return &getPosition{
		getter:   g,
		position: position,
	}
}

// Get fetches the keys at a position.
func (g *getPosition) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	return g.getter.GetPosition(ctx, g.position, keys...)
}
