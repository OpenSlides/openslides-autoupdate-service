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

// definedError is an expected error that are returned to the client.
type definedError interface {
	Type() string
	Error() string
}
