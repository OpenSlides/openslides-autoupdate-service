package datastore

import (
	"context"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

// GetPositioner is like a Getter but also taks a position
type GetPositioner interface {
	GetPosition(ctx context.Context, position int, keys ...dskey.Key) (map[dskey.Key][]byte, error)
}

// GetPosition translates a GetPositioner to a Getter.
type GetPosition struct {
	position int
	getter   GetPositioner
}

// NewGetPosition initializes a GetPosition.
func NewGetPosition(g GetPositioner, position int) *GetPosition {
	return &GetPosition{
		getter:   g,
		position: position,
	}
}

// Get fetches the keys at a position.
func (g *GetPosition) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	return g.getter.GetPosition(ctx, g.position, keys...)
}
