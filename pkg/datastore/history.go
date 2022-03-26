package datastore

import "context"

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
func (g *GetPosition) Get(ctx context.Context, keys ...string) (map[string][]byte, error) {
	return g.getter.GetPosition(ctx, g.position, keys...)
}
