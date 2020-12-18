package collection

import (
	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
)

// Generic is a helper object to create a collection with usual functions.
type Generic struct {
	dp         dataprovider.DataProvider
	collection string
	managePerm string
	readPerm   string

	manageRoutes []string
}

// NewGeneric creates a generic collection.
func NewGeneric(dp dataprovider.DataProvider, collection string, readPerm, managePerm string, ops ...GenericOption) *Generic {
	g := &Generic{
		dp:         dp,
		collection: collection,
		managePerm: managePerm,
		readPerm:   readPerm,
	}

	for _, o := range ops {
		o(g)
	}

	return g
}

// Connect sets the generic routs to the given reader and writer.
func (g *Generic) Connect(s HandlerStore) {
	s.RegisterWriteHandler(g.collection+".create", Create(g.dp, g.managePerm, g.collection))
	s.RegisterWriteHandler(g.collection+".update", Modify(g.dp, g.managePerm, g.collection))
	s.RegisterWriteHandler(g.collection+".delete", Modify(g.dp, g.managePerm, g.collection))

	for _, r := range g.manageRoutes {
		s.RegisterWriteHandler(g.collection+"."+r, Modify(g.dp, g.managePerm, g.collection))
	}

	s.RegisterReadHandler(g.collection, Restrict(g.dp, g.readPerm, g.collection))
}

// GenericOption are options for NewGeneric.
type GenericOption func(*Generic)

// WithManageRoute adds additional routes that should be checked for manage
// permissions.
func WithManageRoute(routes ...string) GenericOption {
	return func(g *Generic) {
		g.manageRoutes = append(g.manageRoutes, routes...)
	}

}
