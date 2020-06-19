// Package http handles http requests for the autoupate service.
package http

import (
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
	keepAlive int
}

// New create a new Handler with the correct urls.
func New(s *autoupdate.Autoupdate, auth Authenticator, keepAlive int) *Handler {
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

		kb, err := kbg(r, uid)
		if err != nil {
			return fmt.Errorf("build keysbuilder: %w", err)
		}

		connection := h.s.Connect(r.Context(), uid, kb)

		ticker := new(time.Ticker)
		if h.keepAlive > 0 {
			ticker = time.NewTicker(time.Duration(h.keepAlive) * time.Second)
			defer ticker.Stop()
		}

		// connection.Next() blocks, until there is new data or the client
		// context or the server is closed.
		var data map[string]json.RawMessage
		for {
			event := make(chan struct{})
			go func() {
				data, err = connection.Next()
				close(event)
			}()

			select {
			case <-ticker.C:
				if err := sendKeepAlive(w); err != nil {
					internalErr(w, err)
				}
				continue
			case <-event:
				// Received autoupdate event.
				// TODO: Reset ticker. This will be possible in go 1.15 that will be released in august:
				//       https://tip.golang.org/doc/go1.15#time
			}

			if err != nil {
				var derr DefinedError
				if errors.As(err, &derr) {
					writeErr(w, derr)
					return nil
				}
				internalErr(w, err)
				return nil
			}

			// data is nil when the topic or the connection are closed.
			if data == nil {
				return nil
			}

			if err := sendData(w, data); err != nil {
				internalErr(w, err)
				return nil
			}
		}
	}
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

// writeErr formats a DefinedError and sents it to the writer (normally the
// client).
func writeErr(w io.Writer, err DefinedError) {
	fmt.Fprintf(w, `{"error": {"type": "%s", "msg": "%s"}}`, err.Type(), quote(err.Error()))
	w.(http.Flusher).Flush()
}

// internalErr sends a nonsense error message to the client and logs the real
// message to stdout.
func internalErr(w io.Writer, err error) {
	log.Printf("Internal Error: %v", err)
	fmt.Fprintln(w, `{"error": {"type": "InternalError", "msg": "Ups, something went wrong!"}}`)
	w.(http.Flusher).Flush()
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
