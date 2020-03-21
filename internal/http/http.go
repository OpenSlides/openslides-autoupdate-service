// Package http helps to handel http requests for the autoupate service.
package http

import (
	"bytes"
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
	h.mux.Handle("/system/autoupdate", errHandleFunc(h.autoupdate(h.komplex)))
	h.mux.Handle("/system/autoupdate/keys", errHandleFunc(h.autoupdate(h.simple)))
	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func (h *Handler) autoupdate(kbg func(*http.Request, int) (autoupdate.KeysBuilder, error)) errHandleFunc {
	return errHandleFunc(func(w http.ResponseWriter, r *http.Request) error {
		uid, err := h.auth.Authenticate(r.Context(), r)
		if err != nil {
			return fmt.Errorf("can not authenticate request: %w", err)
		}

		kb, err := kbg(r, uid)
		if err != nil {
			return fmt.Errorf("can not build keysbuilder: %w", err)
		}

		w.Header().Set("Content-Type", "application/octet-stream")

		connection := h.s.Connect(r.Context(), uid, kb)

		// connection.Next() blocks, until there is new data or the client context or the server is
		// closed.
		for connection.Next() {
			if err := decode(w, connection.Data()); err != nil {
				internalErr(w, err)
				return nil
			}
			w.(http.Flusher).Flush()
		}
		if connection.Err() != nil {
			var derr definedError
			if errors.As(err, &derr) {
				writeErr(w, derr)
			} else {
				internalErr(w, err)
			}
			return nil
		}
		return nil
	})
}

// komplex builds a keysbuilder from the body of a request. The body has to be
// in the format specified in the keysbuilder package.
func (h *Handler) komplex(r *http.Request, uid int) (autoupdate.KeysBuilder, error) {
	defer r.Body.Close()
	return keysbuilder.ManyFromJSON(r.Context(), r.Body, h.s.IDer(uid))
}

// simple builds a keysbuilder from the url query. It expects a
// comma seperated list of keysname.
func (h *Handler) simple(r *http.Request, uid int) (autoupdate.KeysBuilder, error) {
	keys := strings.Split(r.URL.RawQuery, ",")
	return &keysbuilder.Simple{K: keys}, nil
}

type errHandleFunc func(w http.ResponseWriter, r *http.Request) error

func (f errHandleFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := f(w, r); err != nil {
		var derr definedError
		if errors.As(err, &derr) {
			w.WriteHeader(http.StatusBadRequest)
			writeErr(w, derr)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		internalErr(w, err)
	}
}

func writeErr(w io.Writer, err definedError) {
	fmt.Fprintf(w, `{"error": {"type": "%s", "msg": "%s"}}`, err.Type(), quote(err.Error()))
}

func internalErr(w io.Writer, err error) {
	log.Printf("Error: %v", err)
	fmt.Fprintln(w, `{"error": {"type": "InternalError", "msg": "Ups, something went wrong!"}}`)
}

func decode(w io.Writer, m map[string]string) error {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "{")
	for k, v := range m {
		fmt.Fprintf(buf, `"%s":%s,`, k, v)
	}
	buf.Truncate(buf.Len() - 1)
	fmt.Fprintf(buf, "}\n")
	_, err := io.Copy(w, buf)
	return err
}

func quote(s string) string {
	return strings.ReplaceAll(s, `"`, `\"`)
}
