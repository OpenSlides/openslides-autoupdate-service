package test

import (
	"context"
	"net/http"
)

// Auth implements the http.Authenticater interface. It allways returs the given
// user id.
type Auth int

// Authenticate does nothing.
func (a Auth) Authenticate(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	return r.Context(), nil
}

// FromContext returns the uid the object was initialiced with.
func (a Auth) FromContext(ctx context.Context) int {
	return int(a)
}
