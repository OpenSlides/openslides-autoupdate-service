package http

import (
	"context"
	"net/http"
)

// Authenticator gives an user id for an request.
// returns 0 for anonymous.
type Authenticator interface {
	Authenticate(context.Context, *http.Request) (int, error)
}

// DefinedError is an expected error that are returned to the client.
type DefinedError interface {
	Type() string
	Error() string
}
