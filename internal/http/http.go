// Package http handles http requests for the autoupate service.
package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"syscall"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/keysbuilder"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/metric"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

const (
	prefixPublic   = "/system/autoupdate"
	prefixInternal = "/internal/autoupdate"
)

// Run starts the http server.
func Run(ctx context.Context, addr string, auth Authenticater, autoupdate *autoupdate.Autoupdate) error {
	requestCount := new(metric.CurrentCounter)
	metric.Register(requestCount.Metric)

	mux := http.NewServeMux()
	Health(mux)
	Autoupdate(mux, auth, autoupdate, requestCount)
	HistoryInformation(mux, auth, autoupdate)

	srv := &http.Server{
		Addr:        addr,
		Handler:     mux,
		BaseContext: func(net.Listener) context.Context { return ctx },
	}

	// Shutdown logic in separate goroutine.
	wait := make(chan error)
	go func() {
		<-ctx.Done()
		if err := srv.Shutdown(context.Background()); err != nil {
			// TODO EXTERNAL ERROR
			wait <- fmt.Errorf("HTTP server shutdown: %w", err)
			return
		}
		wait <- nil
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// TODO EXTERNAL ERROR
		return fmt.Errorf("HTTP Server failed: %v", err)
	}

	return <-wait
}

// Connecter returns an connect object.
type Connecter interface {
	Connect(userID int, kb autoupdate.KeysBuilder) autoupdate.DataProvider
	SingleData(ctx context.Context, userID int, kb autoupdate.KeysBuilder, position int) (map[datastore.Key][]byte, error)
}

// Autoupdate builds the requested keys from the body of a request. The
// body has to be in the format specified in the keysbuilder package.
func Autoupdate(mux *http.ServeMux, auth Authenticater, connecter Connecter, counter *metric.CurrentCounter) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Cache-Control", "no-store, max-age=0")

		defer r.Body.Close()
		uid := auth.FromContext(r.Context())

		queryBuilder, err := keysbuilder.FromKeys(strings.Split(r.URL.Query().Get("k"), ",")...)
		if err != nil {
			handleError(w, fmt.Errorf("building keysbuilder from query: %w", err), true)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			// TODO EXTERNAL ERROR
			handleError(w, fmt.Errorf("reading body: %w", err), true)
			return
		}

		bodyBuilder, err := keysbuilder.ManyFromJSON(bytes.NewReader(body))
		if err != nil {
			handleError(w, fmt.Errorf("building keysbuilder from body: %w", err), true)
			return
		}

		builder := keysbuilder.FromBuilders(queryBuilder, bodyBuilder)

		rawPosition := r.URL.Query().Get("position")
		position := 0
		if rawPosition != "" {
			p, err := strconv.Atoi(rawPosition)
			if err != nil {
				handleError(w, invalidRequestError{fmt.Errorf("position has to be a number, not %s", rawPosition)}, true)
				return
			}
			position = p
		}

		if r.URL.Query().Has("single") || position != 0 {
			data, err := connecter.SingleData(r.Context(), uid, builder, position)
			if err != nil {
				handleError(w, fmt.Errorf("getting single data: %w", err), true)
				return
			}

			converted := make(map[string]json.RawMessage, len(data))
			for k, v := range data {
				converted[k.String()] = v
			}

			if err := json.NewEncoder(w).Encode(converted); err != nil {
				// TODO EXTERNAL ERROR
				handleError(w, fmt.Errorf("encoding end sending next message: %w", err), true)
				return
			}
			return
		}

		if err := sendMessages(r.Context(), w, uid, builder, connecter); err != nil {
			handleError(w, err, false)
			return
		}
	})

	mux.Handle(
		prefixPublic,
		validRequest(
			authMiddleware(
				countMiddleware(
					handler,
					counter,
				),
				auth,
			),
		),
	)
}

// HistoryInformationer is an object, that can write the history information for
// an object.
type HistoryInformationer interface {
	HistoryInformation(ctx context.Context, uid int, fqid string, w io.Writer) error
}

// HistoryInformation registers the route to return the history information info
// for an fqid.
func HistoryInformation(mux *http.ServeMux, auth Authenticater, hi HistoryInformationer) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uid := auth.FromContext(r.Context())

		fqid := r.URL.Query().Get("fqid")
		if fqid == "" {
			handleError(w, invalidRequestError{fmt.Errorf("History Information needs an fqid")}, true)
			return
		}

		if err := hi.HistoryInformation(r.Context(), uid, fqid, w); err != nil {
			handleError(w, fmt.Errorf("getting history information: %w", err), true)
			return
		}
	})

	mux.Handle(prefixPublic+"/history_information", authMiddleware(handler, auth))
}

func sendMessages(ctx context.Context, w io.Writer, uid int, kb autoupdate.KeysBuilder, connecter Connecter) error {
	next := connecter.Connect(uid, kb)
	encoder := json.NewEncoder(w)

	for ctx.Err() == nil {
		// This blocks, until there is new data. It also unblocks, when the
		// client context is done.
		data, err := next(ctx)
		if err != nil {
			return fmt.Errorf("getting next message: %w", err)
		}

		converted := make(map[string]json.RawMessage, len(data))
		for k, v := range data {
			converted[k.String()] = v
		}

		if err := encoder.Encode(converted); err != nil {
			// TODO EXTERNAL ERROR
			return fmt.Errorf("encoding and sending next message: %w", err)
		}

		w.(http.Flusher).Flush()
	}
	return ctx.Err()
}

// Health tells, if the service is running.
func Health(mux *http.ServeMux) {
	url := prefixPublic + "/health"
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/octet-stream")
		fmt.Fprintln(w, `{"healthy": true}`)
	})

	mux.Handle(url, handler)
}

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

	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) || errors.Is(err, syscall.EPIPE) {
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

func countMiddleware(next http.Handler, counter *metric.CurrentCounter) http.Handler {
	if counter == nil {
		return next
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		counter.Add()
		defer counter.Done()

		next.ServeHTTP(w, r)
	})
}
