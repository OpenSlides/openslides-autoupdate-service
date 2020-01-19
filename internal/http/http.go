package http

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate/keysrequest"
)

// Handler is an http handler for the autoupdate service
type Handler struct {
	s    *autoupdate.Service
	mux  *http.ServeMux
	auth Authenticator
}

// New create a new Handler with the correct urls
func New(s *autoupdate.Service, auth Authenticator) *Handler {
	h := &Handler{
		s:    s,
		mux:  http.NewServeMux(),
		auth: auth,
	}
	h.mux.Handle("/autoupdate/", errHandleFunc(h.autoupdate))
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
			return write400(w, "Empty body, expected key request.\n")
		}
		var errInvalid keysrequest.ErrInvalid
		if errors.As(err, &errInvalid) {
			return write400(w, fmt.Sprintf("Can not parse key request: %v\n", errInvalid.Error()))
		}
		var errJSON keysrequest.ErrJSON
		if errors.As(err, &errJSON) {
			return write400(w, "Request body has to be valid json.\n")
		}

		return fmt.Errorf("can not parse request body: %w", err)
	}
	defer r.Body.Close()

	c, data, err := h.s.Prepare(r.Context(), uid, keysReqs)
	if err != nil {
		return fmt.Errorf("can not get first data: %w", err)
	}
	pushData(w, data)

	for {
		data, err = h.s.Echo(r.Context(), c)
		if err != nil {
			// It is not possible to return the error after content was sent to the client
			fmt.Fprintf(w, "Error: Ups, something went wrong!")
			return nil
		}
		pushData(w, data)
		select {
		case <-r.Context().Done():
			return nil
		case <-h.s.IsClosed():
			return nil
		default:
		}
	}
}

func pushData(w http.ResponseWriter, data map[string][]byte) {
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

func write400(w http.ResponseWriter, msg string) error {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, msg)
	return nil
}
