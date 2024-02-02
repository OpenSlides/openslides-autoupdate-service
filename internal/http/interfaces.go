package http

import (
	"context"
	"io"
	"net/http"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/history"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

// Authenticater gives an user id for an request. Returns 0 for anonymous.
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

// History gets old data.
type History interface {
	HistoryInformation(ctx context.Context, uid int, fqid string, w io.Writer) error
	Data(ctx context.Context, userID int, kb history.KeysBuilder, position int) (map[dskey.Key][]byte, error)
}
