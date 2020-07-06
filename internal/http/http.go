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
	"time"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/keysbuilder"
)

// Handler is an http handler for the autoupdate service.
type Handler struct {
	s         *autoupdate.Autoupdate
	mux       *http.ServeMux
	auth      Authenticator
	keepAlive time.Duration
}

// New create a new Handler with the correct urls.
func New(s *autoupdate.Autoupdate, auth Authenticator, keepAlive time.Duration) *Handler {
	h := &Handler{
		s:         s,
		mux:       http.NewServeMux(),
		auth:      auth,
		keepAlive: keepAlive,
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
			return fmt.Errorf("authenticate request: %w", err)
		}

		// Save tid before the keybuilder is generated. If the datastore gets an
		// update, the update can be handeled.
		tid := h.s.LastID()

		kb, err := kbg(r, uid)
		if err != nil {
			return fmt.Errorf("build keysbuilder: %w", err)
		}

		defer func() {
			// After this line, it is not allowed for the handler to set a
			// status error.
			if err != nil {
				err = noStatusCodeError{err}
			}
		}()

		connection := h.s.Connect(uid, kb, tid)

		for {
			if err := autoupdateLoop(r.Context(), h.keepAlive, w, connection); err != nil {
				return err
			}
		}
	}
}

func autoupdateLoop(ctx context.Context, timeout time.Duration, w io.Writer, connection *autoupdate.Connection) error {
	if timeout > 0 {
		var cancel func()
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}

	// connection.Next() blocks, until there is new data or the client
	// context or the server is closed.
	data, err := connection.Next(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			if err := sendKeepAlive(w); err != nil {
				return err
			}
			return nil
		}
		return err
	}

	if err := sendData(w, data); err != nil {
		return err
	}
	return nil
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

// errHandleFunc is like a http.Handler, but has a error as return value.
//
// If the returned error implements the DefinedError interface, then the error
// message is sent to the client. In other cases the error is interpredet as an
// internal error and logged to stdout.
type errHandleFunc func(w http.ResponseWriter, r *http.Request) error

func (f errHandleFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := f(w, r); err != nil {
		var noStatusErr noStatusCodeError
		status := true
		if errors.As(err, &noStatusErr) {
			status = false
		}

		var closing interface {
			Closing()
		}
		if errors.As(err, &closing) || errors.Is(err, context.Canceled) {
			// Shutdown or connection closed.
			return
		}

		var derr DefinedError
		if errors.As(err, &derr) {
			if status {
				w.WriteHeader(http.StatusBadRequest)
			}
			fmt.Fprintf(w, `{"error": {"type": "%s", "msg": "%s"}}`, derr.Type(), quote(derr.Error()))
			return
		}

		if status {
			w.WriteHeader(http.StatusInternalServerError)
		}
		log.Printf("Internal Error: %v", err)
		fmt.Fprintln(w, `{"error": {"type": "InternalError", "msg": "Ups, something went wrong!"}}`)
	}
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
		w.Write(value)
	}
	w.Write([]byte("}\n"))
	w.(http.Flusher).Flush()
	return nil
}

// sendKeepAlive sends an empty message to the client.
func sendKeepAlive(w io.Writer) error {
	_, err := fmt.Fprintln(w, `{}`)
	w.(http.Flusher).Flush()
	return err
}
