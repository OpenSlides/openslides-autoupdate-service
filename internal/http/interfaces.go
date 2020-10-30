package http

import (
	"context"
	"io"
	"net/http"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
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

// Liver provides a Live method, that writes continues data to the given writer.
type Liver interface {
	Live(ctx context.Context, uid int, w io.Writer, kb autoupdate.KeysBuilder) error
}
