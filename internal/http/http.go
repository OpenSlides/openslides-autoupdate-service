// Package http helps to handel http requests for the autoupate service.
package http

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/keysbuilder"
)

const uidHeader = "X-User-ID"

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
	h.mux.Handle("/system/autoupdate", h.authMiddleware(errHandleFunc(h.autoupdate(h.komplex()))))
	h.mux.Handle("/system/autoupdate/keys", h.authMiddleware(errHandleFunc(h.autoupdate(h.simple()))))
	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func (h *Handler) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uid, err := h.auth.Authenticate(r.Context(), r)
		if err != nil {
			// TODO: Differ between auth error and 500er error
			write400(w, fmt.Sprintf("Unauthorized access: %v", err))
			return
		}

		r2 := new(http.Request)
		*r2 = *r
		r2.Header.Set(uidHeader, strconv.Itoa(uid))
		next.ServeHTTP(w, r2)
	})
}

func (h *Handler) autoupdate(kbg keysBuilderGetter) errHandleFunc {
	return errHandleFunc(func(w http.ResponseWriter, r *http.Request) error {
		rawUID := r.Header.Get(uidHeader)
		uid, err := strconv.Atoi(rawUID)
		if err != nil {
			return fmt.Errorf("Invalid user id in header: %w", err)
		}

		kb, err := kbg(r, uid)
		if err != nil {
			return err
		}

		connection := h.s.Connect(r.Context(), uid, kb)

		// connection.Next() blocks, until there is new data or the client context or the server is
		// closed.
		for connection.Next() {
			pushData(w, connection.Data())
		}
		if connection.Err() != nil {
			// It is not possible to return the error after content was sent to the client
			fmt.Fprintf(w, "Error: Ups, something went wrong!")
			log.Printf("Error: %v", err)
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
