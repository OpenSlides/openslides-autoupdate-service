package http

import (
	"context"
	"net/http"
)

// Authenticater gives an user id for an request. Returns 0 for public access.
type Authenticater interface {
	Authenticate(http.ResponseWriter, *http.Request) (context.Context, error)
	FromContext(context.Context) int
	AuthenticatedContext(context.Context, int) context.Context
}

// ClientError is an expected error that are returned to the client.
type ClientError interface {
	Type() string
	Error() string
}
