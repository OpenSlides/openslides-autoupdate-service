package permission

import "github.com/OpenSlides/openslides-permission-service/internal/perm"

// Option is an optional argument for permission.New()
type Option func(*Permission)

// WithCollections initializes a Permission Service with specific connecters. Per
// default, the OpenSlides collections are used.
func WithCollections(cons []perm.Connecter) Option {
	return func(p *Permission) {
		p.connecters = cons
	}
}
