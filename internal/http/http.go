// Package http handles http requests for the autoupate service.
package http

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/openslides/openslides-autoupdate-service/internal/keysbuilder"
)

const prefix = "/system/autoupdate"

// Complex builds the requested keys from the body of a request. The
// body has to be in the format specified in the keysbuilder package.
func Complex(mux *http.ServeMux, auth Authenticater, db keysbuilder.DataProvider, liver Liver) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/octet-stream")

		defer r.Body.Close()
		uid := auth.FromContext(r.Context())

		kb, err := keysbuilder.ManyFromJSON(r.Body, db, uid)
		if err != nil {
			handleError(w, err, true)
			return
		}

		// TODO: This should not be run here. This is only for development
		kb.Update(r.Context())

		// This blocks until the request is done.
		if err := liver.Live(r.Context(), uid, w, kb); err != nil {
			handleError(w, err, false)
			return
		}
	})

	mux.Handle(prefix, validRequest(authMiddleware(handler, auth)))
}

// Simple builds a keysbuilder from the url query. It expects a comma
// separated list of keysname.
func Simple(mux *http.ServeMux, auth Authenticater, liver Liver) {
	url := prefix + "/keys"
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/octet-stream")

		keys := strings.Split(r.URL.RawQuery, ",")
		kb := &keysbuilder.Simple{K: keys}
		if err := kb.Validate(); err != nil {
			handleError(w, err, true)
			return
		}

		uid := auth.FromContext(r.Context())

		// This blocks until the request is done.
		if err := liver.Live(r.Context(), uid, w, kb); err != nil {
			handleError(w, err, false)
			return
		}
	})

	mux.Handle(url, validRequest(authMiddleware(handler, auth)))
}

// Health tells, if the service is running.
func Health(mux *http.ServeMux) {
	url := prefix + "/health"
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/octet-stream")
		fmt.Fprintln(w, `{"healthy": true}`)
	})

	mux.Handle(url, handler)
}

// // ProjectorLiver returns the projector data.
// type ProjectorLiver interface {
// 	Live(ctx context.Context, uid int, w io.Writer, pids []int) error
// }

// // Projector adds the projector route to the mutex.
// func Projector(mux *http.ServeMux, auth Authenticater, liver ProjectorLiver) {
// 	url := prefix + "/projector"
// 	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/octet-stream")

// 		uid := auth.FromContext(r.Context())

// 		projectorIDs, err := projectorIDs(r.URL.Query().Get("ids"))
// 		if err != nil {
// 			handleError(w, err, false)
// 			return
// 		}

// 		// This blocks until the request is done.
// 		if err := liver.Live(r.Context(), uid, w, projectorIDs); err != nil {
// 			handleError(w, err, false)
// 			return
// 		}
// 	})

// 	mux.Handle(url, handler)
// }

// func projectorIDs(raw string) ([]int, error) {
// 	parts := strings.Split(raw, ",")
// 	ids := make([]int, len(parts))
// 	for i, rpid := range parts {
// 		id, err := strconv.Atoi(rpid)
// 		if err != nil {
// 			return nil, invalidRequestError{fmt.Errorf("projector_ids has to be a list of ints not `%s`", raw)}
// 		}

// 		ids[i] = id
// 	}
// 	return ids, nil
// }

func authMiddleware(next http.Handler, auth Authenticater) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, err := auth.Authenticate(w, r)
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
	if writeStatusCode {
		w.Header().Set("Content-Type", "application/octet-stream")
	}

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

func validRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Only allow GET or POST requests.
		if !(r.Method == http.MethodPost || r.Method == http.MethodGet) {
			handleError(w, invalidRequestError{fmt.Errorf("Only GET or POST requests are supported")}, true)
			return
		}

		next.ServeHTTP(w, r)
	})
}
