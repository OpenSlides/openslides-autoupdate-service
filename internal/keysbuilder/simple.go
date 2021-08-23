package keysbuilder

import (
	"context"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Simple implements the autoupdate.Keysbuilder interface. It returns the keys
// it was initialized with.
type Simple struct {
	K []string
}

// Update does nothing. The keys of a simple keysbuilder can not change.
func (s *Simple) Update(context.Context, datastore.Getter) error {
	return nil
}

// Keys returns the keys the keysbuilder.Simple was initialized.
func (s *Simple) Keys() []string {
	return s.K
}
