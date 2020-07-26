package keysbuilder

import (
	"context"
	"strings"
)

// Simple implements the autoupdate.Keysbuilder interface. It returns the keys
// it was initialized with.
type Simple struct {
	K []string
}

// Update does nothing. The keys of a simple keysbuilder can not change.
func (s *Simple) Update(context.Context) error {
	return nil
}

// Keys returns the keys the keysbuilder.Simple was initialized.
func (s *Simple) Keys() []string {
	return s.K
}

// Validate checks, if the given keys are valid.
func (s *Simple) Validate() error {
	for _, key := range s.K {
		keyParts := strings.SplitN(key, "/", 3)
		if len(keyParts) != 3 {
			return InvalidError{msg: "Invalid keys"}
		}
	}
	return nil
}
