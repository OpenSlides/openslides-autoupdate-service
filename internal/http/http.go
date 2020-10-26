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

// Handler is an http handler for the autoupdate service.
type Handler struct {
	s    *autoupdate.Autoupdate
	mux  *http.ServeMux
	auth Authenticater
}

// New create a new Handler with the correct urls.
func New(s *autoupdate.Autoupdate, auth Authenticater) *Handler {
	h := &Handler{
		s:    s,
		mux:  http.NewServeMux(),
		auth: auth,
	}

	middlewares := func(next http.Handler) http.Handler {
		return validRequest(h.authMiddleware(next))
	}

	h.mux.Handle("/system/autoupdate", middlewares(h.autoupdate(h.complex)))
	h.mux.Handle("/system/autoupdate/keys", middlewares(h.autoupdate(h.simple)))
	h.mux.Handle("/system/autoupdate/health", middlewares(h.health()))
	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

// autoupdate creates a Handler for a specific Keysbuilder.
func (h *Handler) autoupdate(kbg func(*http.Request, int) (autoupdate.KeysBuilder, error)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/octet-stream")

		uid := h.auth.FromContext(r.Context())

		// Save tid before the keybuilder is generated. If the datastore gets an
		// update, the update can be handeled.
		tid := h.s.LastID()

		kb, err := kbg(r, uid)
		if err != nil {
			handleError(w, fmt.Errorf("build keysbuilder: %w", err), true)
			return
		}

		defer func() {
			// After this line, it is not allowed for the handler to set a
			// status error.
			if err != nil {
				handleError(w, err, false)
				return
			}
		}()

		connection := h.s.Connect(uid, kb, tid)

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
	})
}

// complex builds a keysbuilder from the body of a request. The body has to be
// in the format specified in the keysbuilder package.
func (h *Handler) complex(r *http.Request, uid int) (autoupdate.KeysBuilder, error) {
	defer r.Body.Close()
	return keysbuilder.ManyFromJSON(r.Context(), r.Body, h.s, uid)
}

// simple builds a keysbuilder from the url query. It expects a comma separated
// list of keysname.
func (h *Handler) simple(r *http.Request, uid int) (autoupdate.KeysBuilder, error) {
	keys := strings.Split(r.URL.RawQuery, ",")
	kb := &keysbuilder.Simple{K: keys}
	if err := kb.Validate(); err != nil {
		return nil, err
	}
	return kb, nil
}

func (h *Handler) health() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"healthy": true}`)
	})
}

func (h *Handler) authMiddleware(next http.Handler) http.Handler {
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
