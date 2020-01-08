package autoupdate

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/openslides/openslides-autoupdate-service/internal/keysrequest"
)

// Handler is an http handler for the autoupdate service
type Handler struct {
	s    *Service
	mux  *http.ServeMux
	auth Authenticator
}

// NewHandler create a new Handler with the correct urls
func NewHandler(s *Service, auth Authenticator) *Handler {
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

	keysReq, err := keysrequest.FromJSON(r.Body)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return write400(w, "Empty body, expected key request.\n")
		}
		return fmt.Errorf("can not parse request body: %w", err)
	}
	defer r.Body.Close()

	tid, keys, data, err := h.s.prepare(r.Context(), uid, keysReq)
	if err != nil {
		return fmt.Errorf("can not get first data: %w", err)
	}
	pushData(w, data)

	for {
		tid, data, err = h.s.echo(r.Context(), uid, tid, keys)
		if err != nil {
			// It is not possible to return the error after content was sent to the client
			log.Printf("Error: %v", err)
			return nil
		}
		pushData(w, data)
		select {
		case <-r.Context().Done():
			return nil
		case <-h.s.closed:
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
		if errors.Is(err, err400{}) {
			write400(w, err.Error())
			return
		}
		log.Printf("Error: %v", err)
		http.Error(w, "Ups, something went wrong!", http.StatusInternalServerError)
	}
}

func write400(w http.ResponseWriter, msg string) error {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, msg)
	return nil
}

type err400 struct {
	err error
}

func new400(e error) err400 {
	return err400{err: e}
}

func (e err400) Error() string {
	return e.err.Error()
}

func (e err400) Unwrap() error {
	return e.err
}
