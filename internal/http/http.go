// Package http helps to handel http requests for the autoupate service.
package http

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/keysbuilder"
)

// Handler is an http handler for the autoupdate service.
type Handler struct {
	s    *autoupdate.Service
	mux  *http.ServeMux
	auth Authenticator
}

// New create a new Handler with the correct urls.
func New(s *autoupdate.Service, auth Authenticator) *Handler {
	h := &Handler{
		s:    s,
		mux:  http.NewServeMux(),
		auth: auth,
	}
	h.mux.Handle("/system/autoupdate", h.autoupdate(h.complex))
	h.mux.Handle("/system/autoupdate/keys", h.autoupdate(h.simple))
	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

// autoupdate creates a Handler for a specific Keysbuilder.
func (h *Handler) autoupdate(kbg func(*http.Request, int) (autoupdate.KeysBuilder, error)) errHandleFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Content-Type", "application/octet-stream")

		uid, err := h.auth.Authenticate(r.Context(), r)
		if err != nil {
			return fmt.Errorf("can not authenticate request: %w", err)
		}

		kb, err := kbg(r, uid)
		if err != nil {
			return fmt.Errorf("can not build keysbuilder: %w", err)
		}

		connection := h.s.Connect(r.Context(), uid, kb)

		// connection.Next() blocks, until there is new data or the client
		// context or the server is closed.
		for {
			reader, err := connection.Next()
			if err != nil {
				var derr DefinedError
				if errors.As(err, &derr) {
					writeErr(w, derr)
					return nil
				}
				internalErr(w, err)
				return nil
			}

			// reader is nil when the topic or the connection are closed.
			if reader == nil {
				return nil
			}

			if _, err := io.Copy(w, reader); err != nil {
				internalErr(w, err)
				return nil
			}
			w.(http.Flusher).Flush()
		}
	}
}

// complex builds a keysbuilder from the body of a request. The body has to be
// in the format specified in the keysbuilder package.
func (h *Handler) complex(r *http.Request, uid int) (autoupdate.KeysBuilder, error) {
	defer r.Body.Close()
	return keysbuilder.ManyFromJSON(r.Context(), r.Body, h.s.IDer(uid))
}

// simple builds a keysbuilder from the url query. It expects a comma seperated
// list of keysname.
func (h *Handler) simple(r *http.Request, uid int) (autoupdate.KeysBuilder, error) {
	keys := strings.Split(r.URL.RawQuery, ",")
	kb := &keysbuilder.Simple{K: keys}
	if err := kb.Validate(); err != nil {
		return nil, err
	}
	return kb, nil
}

// errHandleFunc is like a http.Handler, but has a error as return value.
//
// If the returned error implements the DefinedError interface, then the error
// message is sent to the client. In other cases the error is interpredet as an
// internal error and logged to stdout.
type errHandleFunc func(w http.ResponseWriter, r *http.Request) error

func (f errHandleFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := f(w, r); err != nil {
		var derr DefinedError
		if errors.As(err, &derr) {
			w.WriteHeader(http.StatusBadRequest)
			writeErr(w, derr)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		internalErr(w, err)
	}
}

// writeErr formats a DefinedError and sents it to the writer (normaly the
// client).
func writeErr(w io.Writer, err DefinedError) {
	fmt.Fprintf(w, `{"error": {"type": "%s", "msg": "%s"}}`, err.Type(), quote(err.Error()))
}

// internalErr sends a nonsense error message to the client and logs the real
// message to stdout.
func internalErr(w io.Writer, err error) {
	log.Printf("Error: %v", err)
	fmt.Fprintln(w, `{"error": {"type": "InternalError", "msg": "Ups, something went wrong!"}}`)
}

// quote decodes changes quotation marks with a backslash to make sure, they are
// valid json.
func quote(s string) string {
	return strings.ReplaceAll(s, `"`, `\"`)
}
