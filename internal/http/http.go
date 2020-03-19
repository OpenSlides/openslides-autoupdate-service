// Package http helps to handel http requests for the autoupate service.
package http

import (
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
	h.mux.Handle("/system/autoupdate", errHandleFunc(h.autoupdate(h.komplex())))
	h.mux.Handle("/system/autoupdate/keys", errHandleFunc(h.autoupdate(h.simple())))
	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func (h *Handler) autoupdate(kbg keysBuilderGetter) errHandleFunc {
	return errHandleFunc(func(w http.ResponseWriter, r *http.Request) error {
		uid, err := h.auth.Authenticate(r.Context(), r)
		if err != nil {
			return fmt.Errorf("can not authenticate request: %w", err)
		}

		kb, err := kbg(r, uid)
		if err != nil {
			return err
		}

		w.Header().Set("Content-Type", "application/octet-stream")

		connection := h.s.Connect(r.Context(), uid, kb)

		encoder := json.NewEncoder(w)

		// connection.Next() blocks, until there is new data or the client context or the server is
		// closed.
		for connection.Next() {
			if err := encoder.Encode(connection.Data()); err != nil {
				writeErr(w, err.Error())
				return nil
			}
			w.(http.Flusher).Flush()
		}
		if connection.Err() != nil {
			writeErr(w, err.Error())
			return nil
		}
		return nil
	})
}

type keysBuilderGetter func(r *http.Request, uid int) (autoupdate.KeysBuilder, error)

// komplex builds a keysbuilder from the body of a request. The body has to be
// in the format specified in the keysbuilder package. It returns an err400 if
// the input is wrong.
func (h *Handler) komplex() keysBuilderGetter {
	return keysBuilderGetter(func(r *http.Request, uid int) (autoupdate.KeysBuilder, error) {
		kb, err := keysbuilder.ManyFromJSON(r.Context(), r.Body, h.s.IDer(uid))
		defer r.Body.Close()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil, raiseErr400(fmt.Errorf("empty body, expected key request"))
			}
			var errInvalid keysbuilder.ErrInvalid
			if errors.As(err, &errInvalid) {
				return nil, raiseErr400(err)
			}

			var errJSON keysbuilder.ErrJSON
			if errors.As(err, &errJSON) {
				return nil, raiseErr400(fmt.Errorf("request body is not valid json: %w", err))
			}

			return nil, fmt.Errorf("can not parse request body: %w", err)
		}
		return kb, nil
	})
}

// simple builds a keysbuilder from the url query. It expects a
// comma seperated list of keysname.
func (h *Handler) simple() keysBuilderGetter {
	return keysBuilderGetter(func(r *http.Request, uid int) (autoupdate.KeysBuilder, error) {
		keys := strings.Split(r.URL.RawQuery, ",")
		return &keysbuilder.Simple{K: keys}, nil
	})
}

func pushData(w http.ResponseWriter, data map[string]string) {

	for key, value := range data {
		fmt.Fprintf(w, "%s: %s\n", key, value)
		w.(http.Flusher).Flush()
	}
}

type errHandleFunc func(w http.ResponseWriter, r *http.Request) error

func (f errHandleFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := f(w, r); err != nil {
		var inputErr err400
		if errors.As(err, &inputErr) {
			write400(w, fmt.Sprintf("Wrong input: %v", inputErr.Error()))
			return
		}
		log.Printf("Error: %v", err)
		http.Error(w, "Ups, something went wrong!", http.StatusInternalServerError)
	}
}

func write400(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, msg)
}

func writeErr(w io.Writer, msg string) {
	log.Printf("Error: %s", msg)
	fmt.Fprintf(w, `{"error": {"detail": "%s"}}`, msg)
}
