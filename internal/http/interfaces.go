package http

import (
	"context"
	"net/http"
)

// Authenticator gives an user id for an request. Returns 0 for anonymous.
type Authenticator interface {
	Authenticate(http.ResponseWriter, *http.Request) (context.Context, error)
	FromContext(context.Context) int
}

// ClientError is an expected error that are returned to the client.
type ClientError interface {
	Type() string
	Error() string
}
