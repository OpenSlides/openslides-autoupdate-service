// Package http helps to handel http requests for the autoupate service.
package http

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate/keysbuilder"
	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate/keysrequest"
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
	h.mux.Handle("/system/autoupdate", errHandleFunc(h.autoupdate))
	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func (h *Handler) autoupdate(w http.ResponseWriter, r *http.Request) error {
	uid, err := h.auth.Authenticate(r.Context(), r)
	if err != nil {
		return fmt.Errorf("can not authenticate request: %w", err)
	}

	keysReqs, err := keysrequest.ManyFromJSON(r.Body)
	if err != nil {
		if errors.Is(err, io.EOF) {
			write400(w, "Empty body, expected key request.\n")
			return nil
		}
		var errInvalid keysrequest.ErrInvalid
		if errors.As(err, &errInvalid) {
			write400(w, fmt.Sprintf("Can not parse key request: %v\n", errInvalid.Error()))
			return nil
		}

		var errJSON keysrequest.ErrJSON
		if errors.As(err, &errJSON) {
			write400(w, "Request body has to be valid json.\n")
			return nil
		}

		return fmt.Errorf("can not parse request body: %w", err)
	}
	defer r.Body.Close()

	kb, err := keysbuilder.New(r.Context(), h.s.IDer(uid), keysReqs...)
	if err != nil {
		if errors.Is(err, keysrequest.ErrInvalid{}) {
			write400(w, err.Error())
			return nil
		}
		return fmt.Errorf("can not build keys: %w", err)
	}

	connection := h.s.Connect(r.Context(), uid, kb)

	for connection.Next() {
		pushData(w, connection.Data())
		select {
		case <-r.Context().Done():
			return nil
		case <-h.s.Done():
			return nil
		default:
		}
	}
	if connection.Err() != nil {
		// It is not possible to return the error after content was sent to the client
		fmt.Fprintf(w, "Error: Ups, something went wrong!")
		log.Printf("Error: %v", err)
	}
	return nil
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
		var inputErr autoupdate.ErrInput
		if errors.As(err, &inputErr) {
			write400(w, err.Error())
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
