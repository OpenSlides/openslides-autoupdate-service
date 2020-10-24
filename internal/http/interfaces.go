package http

import (
	"context"
	"net/http"
)

// Authenticater gives an user id for an request. Returns 0 for anonymous.
type Authenticater interface {
	Authenticate(http.ResponseWriter, *http.Request) (context.Context, error)
	FromContext(context.Context) int
}

// ClientError is an expected error that are returned to the client.
type ClientError interface {
	Type() string
	Error() string
}
