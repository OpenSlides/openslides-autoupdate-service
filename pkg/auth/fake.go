package auth

import (
	"context"
	"net/http"
)

// Fake implements the http.Authenticater interface. It allways returs the given
// user id.
type Fake int

// Authenticate does nothing.
func (a Fake) Authenticate(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	return r.Context(), nil
}

// FromContext returns the uid the object was initialiced with.
func (a Fake) FromContext(ctx context.Context) int {
	return int(a)
}
