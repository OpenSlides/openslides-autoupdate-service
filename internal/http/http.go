// Package http handles http requests for the autoupate service.
package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/keysbuilder"
)

// handler holds the state for the handlers.
type handler struct {
	au   *autoupdate.Autoupdate
	mux  *http.ServeMux
	auth Authenticater

	prefix string
}

// New create a new Handler with the correct urls.
func New(au *autoupdate.Autoupdate, auth Authenticater) http.Handler {
	h := &handler{
		au:     au,
		mux:    http.NewServeMux(),
		auth:   auth,
		prefix: "/system/autoupdate",
	}

	mux := http.NewServeMux()

	h.complexAutoupdate(mux)
	h.simpleAutoupdate(mux)
	h.health(mux)

	return mux
}

func (h *handler) middlewares(next http.Handler) http.Handler {
	return validRequest(h.authMiddleware(next))
}

// complex builds the requested keys from the body of a request. The body has to
// be in the format specified in the keysbuilder package.
func (h *handler) complexAutoupdate(mux *http.ServeMux) {
	url := h.prefix
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		uid := h.auth.FromContext(r.Context())

		kb, err := keysbuilder.ManyFromJSON(r.Context(), r.Body, h.au, uid)
		if err != nil {
			handleError(w, err, true)
			return
		}

		h.autoupdate(w, r, kb)
	})

	mux.Handle(url, h.middlewares(handler))
}

// simpleAutoupdate builds a keysbuilder from the url query. It expects a comma
// separated list of keysname.
func (h *handler) simpleAutoupdate(mux *http.ServeMux) {
	url := h.prefix + "/keys"
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		keys := strings.Split(r.URL.RawQuery, ",")
		kb := &keysbuilder.Simple{K: keys}
		if err := kb.Validate(); err != nil {
			handleError(w, err, true)
			return
		}

		h.autoupdate(w, r, kb)
	})

	mux.Handle(url, h.middlewares(handler))
}

// health tells, if the service is running.
func (h *handler) health(mux *http.ServeMux) {
	url := h.prefix + "/health"
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/octet-stream")
		fmt.Fprintln(w, `{"healthy": true}`)
	})

	mux.Handle(url, handler)
}

// autoupdate returns the values for the keys specified by the given
// KeysBuilder.
func (h *handler) autoupdate(w http.ResponseWriter, r *http.Request, kb autoupdate.KeysBuilder) {
	w.Header().Set("Content-Type", "application/octet-stream")

	uid := h.auth.FromContext(r.Context())
	connection := h.au.Connect(uid, kb)

	for {
		// connection.Next() blocks, until there is new data or the client context
		// or the server is closed.
		data, err := connection.Next(r.Context())
		if err != nil {
			handleError(w, err, false)
			return
		}

		if err := sendData(w, data); err != nil {
			handleError(w, err, false)
			return
		}
	}
}

func (h *handler) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, err := h.auth.Authenticate(w, r)
		if err != nil {
			handleError(w, fmt.Errorf("authenticate request: %w", err), true)
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// handleError interprets the given error and writes a corresponding message to
// the client and/or stdout.
//
// If the handler already started to write the body then it is not allowed to
// set the http-status-code. In this case, writeStatusCode has to be fales.
func handleError(w http.ResponseWriter, err error, writeStatusCode bool) {
	if writeStatusCode {
		w.Header().Set("Content-Type", "application/octet-stream")
	}

	var closing interface {
		Closing()
	}
	if errors.As(err, &closing) {
		// Server is closing.
		return
	}

	if errors.Is(err, context.Canceled) {
		// Client closed connection.
		return
	}

	var errClient ClientError
	if errors.As(err, &errClient) {
		if writeStatusCode {
			w.WriteHeader(http.StatusBadRequest)
		}

		fmt.Fprintf(w, `{"error": {"type": "%s", "msg": "%s"}}`, errClient.Type(), quote(errClient.Error()))
		return
	}

	if writeStatusCode {
		w.WriteHeader(http.StatusInternalServerError)
	}
	log.Printf("Internal Error: %v", err)
	fmt.Fprintln(w, `{"error": {"type": "InternalError", "msg": "Ups, something went wrong!"}}`)
}

// quote decodes changes quotation marks with a backslash to make sure, they are
// valid json.
func quote(s string) string {
	return strings.ReplaceAll(s, `"`, `\"`)
}

func sendData(w io.Writer, data map[string]json.RawMessage) error {
	// TODO: Handle errors
	first := true
	w.Write([]byte("{"))
	for key, value := range data {
		if !first {
			w.Write([]byte{','})
		}
		first = false
		w.Write([]byte{'"'})
		w.Write([]byte(key))
		w.Write([]byte{'"', ':'})
		if value == nil {
			value = []byte("null")
		}
		w.Write(value)
	}
	w.Write([]byte("}\n"))
	w.(http.Flusher).Flush()
	return nil
}

func validRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Only allow http2 requests.
		if !r.ProtoAtLeast(2, 0) {
			handleError(w, invalidRequestError{fmt.Errorf("Only http2 is supported")}, true)
			return
		}

		// Only allow GET or POST requests.
		if !(r.Method == http.MethodPost || r.Method == http.MethodGet) {
			handleError(w, invalidRequestError{fmt.Errorf("Only GET or POST requests are supported")}, true)
			return
		}

		next.ServeHTTP(w, r)
	})
}
