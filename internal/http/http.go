// Package http handles http requests for the autoupate service.
package http

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"net/http/pprof"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/keysbuilder"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/metric"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/oserror"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/redis"
	"github.com/klauspost/compress/zstd"
	"nhooyr.io/websocket"
)

const (
	prefixPublic   = "/system/autoupdate"
	prefixInternal = "/internal/autoupdate"
)

// Run starts the http server.
func Run(
	ctx context.Context,
	addr string,
	auth Authenticater,
	autoupdate *autoupdate.Autoupdate,
	history History,
	redisConnection *redis.Redis,
	saveIntercal time.Duration,
) error {
	var connectionCount *connectionCount
	if redisConnection != nil {
		connectionCount = newConnectionCount(ctx, redisConnection, saveIntercal)
		metric.Register(connectionCount.Metric)
	}

	mux := http.NewServeMux()
	HandleHealth(mux)
	HandleAutoupdate(mux, auth, autoupdate, history, connectionCount)
	HandleInternalAutoupdate(mux, auth, history, autoupdate)
	HandleShowConnectionCount(mux, autoupdate, auth, connectionCount)
	HandleHistoryInformation(mux, auth, history)
	HandleProfile(mux)

	srv := &http.Server{
		Addr:        addr,
		Handler:     mux,
		BaseContext: func(net.Listener) context.Context { return ctx },
	}

	// Shutdown logic in separate goroutine.
	wait := make(chan error)
	go func() {
		<-ctx.Done()
		if err := srv.Shutdown(context.WithoutCancel(ctx)); err != nil {
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
	Connect(ctx context.Context, userID int, kb autoupdate.KeysBuilder) (autoupdate.DataProvider, error)
	SingleData(ctx context.Context, userID int, kb autoupdate.KeysBuilder) (map[dskey.Key][]byte, error)
}

func autoupdateHandler(auth Authenticater, connecter Connecter, history History) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		defer r.Body.Close()
		uid := auth.FromContext(r.Context())

		queryBuilder, err := keysbuilder.FromKeys(strings.Split(r.URL.Query().Get("k"), ",")...)
		if err != nil {
			handleErrorWithStatus(w, fmt.Errorf("building keysbuilder from query: %w", err))
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintln(w, "Invalid body")
			return
		}

		compactedBody := new(bytes.Buffer)
		if err := json.Compact(compactedBody, body); err == nil {
			// Ignore error, it will be handled in the keysbuilder function.
			ctx = oserror.ContextWithBody(ctx, string(body))
		}

		bodyBuilder, err := keysbuilder.ManyFromJSON(bytes.NewReader(body))
		if err != nil {
			handleErrorWithStatus(w, fmt.Errorf("building keysbuilder from body: %w", err))
			return
		}

		builder := keysbuilder.FromBuilders(queryBuilder, bodyBuilder)

		rawPosition := r.URL.Query().Get("position")
		position := 0
		if rawPosition != "" {
			p, err := strconv.Atoi(rawPosition)
			if err != nil {
				handleErrorWithStatus(w, invalidRequestError{fmt.Errorf("position has to be a number, not %s", rawPosition)})
				return
			}
			position = p
		}

		if r.URL.Query().Has("profile_restrict") {
			ctx = oserror.ContextWithTag(ctx, "profile_restrict")
		}

		var compress bool
		if r.URL.Query().Has("compress") {
			compress = true
		}

		if r.URL.Query().Has("single") || position != 0 {
			handleSingle(ctx, uid, w, r, position, connecter, builder, history, compress)
			return
		}

		if r.URL.Query().Has("websocket") {
			if err := sendMessagesWebsocket(ctx, w, r, uid, builder, connecter, compress); err != nil {
				handleErrorWithoutStatus(w, err)
				return
			}
			return
		}

		if err := sendMessages(ctx, w, uid, builder, connecter, compress); err != nil {
			handleErrorWithoutStatus(w, err)
			return
		}
	})
}

func handleSingle(ctx context.Context, uid int, w http.ResponseWriter, r *http.Request, position int, connecter Connecter, builder *keysbuilder.Builder, history History, compress bool) {
	w.Header().Set("Cache-Control", "no-store, max-age=0")

	var data map[dskey.Key][]byte
	switch position {
	case 0:
		d, err := connecter.SingleData(ctx, uid, builder)
		if err != nil {
			handleErrorWithStatus(w, fmt.Errorf("getting single data: %w", err))
			return
		}

		data = d

	default:
		d, err := history.Data(ctx, uid, builder, position)
		if err != nil {
			handleErrorWithStatus(w, fmt.Errorf("getting history data: %w", err))
			return
		}
		data = d
	}

	if err := writeData(w, data, compress); err != nil {
		handleErrorWithoutStatus(w, err)
		return
	}
	return
}

// HandleAutoupdate builds the requested keys from the body of a request. The
// body has to be in the format specified in the keysbuilder package.
func HandleAutoupdate(mux *http.ServeMux, auth Authenticater, connecter Connecter, history History, connectionCount *connectionCount) {
	mux.Handle(
		prefixPublic,
		validRequest(
			authMiddleware(
				connectionCountMiddleware(
					autoupdateHandler(auth, connecter, history),
					auth,
					connectionCount,
				),
				auth,
			),
		),
	)
}

// HandleInternalAutoupdate is the same as the normal autoupdate route, but it
// uses the user_id from an argument.
//
// /internal/autoupdate?user_id=23&single=1&k=user/1/username
func HandleInternalAutoupdate(mux *http.ServeMux, auth Authenticater, history History, connecter Connecter) {
	mux.Handle(
		prefixInternal,
		validRequest(
			internalAuthMiddleware(
				autoupdateHandler(auth, connecter, history),
				auth,
			),
		),
	)
}

func writeData(w io.Writer, data map[dskey.Key][]byte, compress bool) error {
	converted := make(map[string]json.RawMessage, len(data))
	for k, v := range data {
		converted[k.String()] = v
	}

	var wr io.Writer = w
	if compress {
		defer fmt.Fprintln(wr)
		base64Encoder := base64.NewEncoder(base64.RawStdEncoding, wr)
		defer base64Encoder.Close()

		zstdEncoder, err := zstd.NewWriter(base64Encoder)
		if err != nil {
			return fmt.Errorf("create encoder: %w", err)
		}
		defer zstdEncoder.Close()
		wr = zstdEncoder
	}

	if err := json.NewEncoder(wr).Encode(converted); err != nil {
		return fmt.Errorf("encode data: %w", err)
	}

	return nil
}

// HandleShowConnectionCount adds a handler to show the result of the connection counter.
func HandleShowConnectionCount(mux *http.ServeMux, autoupdate *autoupdate.Autoupdate, auth Authenticater, connectionCount *connectionCount) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if connectionCount == nil {
			oserror.Handle(fmt.Errorf("Error connection count is not initialized"))
			http.Error(w, "Counting not possible", 500)
			return
		}

		ctx := r.Context()
		uid := auth.FromContext(ctx)

		allowed, meetingIDs, err := autoupdate.CanSeeConnectionCount(ctx, uid)
		if err != nil {
			oserror.Handle(fmt.Errorf("Error checking count permission %w", err))
			http.Error(w, "Counting not possible", 500)
			return
		}

		if !allowed {
			http.Error(w, "Connection counting not allowed", 400)
			return
		}

		val, err := connectionCount.Show(ctx)
		if err != nil {
			oserror.Handle(fmt.Errorf("Error counting connection: %w", err))
			http.Error(w, "Counting not possible", 500)
			return
		}

		if err := autoupdate.FilterConnectionCount(ctx, meetingIDs, val); err != nil {
			oserror.Handle(fmt.Errorf("Error filtering connection count: %w", err))
			http.Error(w, "Counting not possible", 500)
			return
		}

		if err := json.NewEncoder(w).Encode(val); err != nil {
			oserror.Handle(fmt.Errorf("Error decoding counter %w", err))
			http.Error(w, "Counting not possible", 500)
			return
		}
	})

	mux.Handle(prefixPublic+"/connection_count", authMiddleware(handler, auth))
}

// HistoryInformationer is an object, that can write the history information for
// an object.
type HistoryInformationer interface {
	HistoryInformation(ctx context.Context, uid int, fqid string, w io.Writer) error
}

// HandleHistoryInformation registers the route to return the history information info
// for an fqid.
func HandleHistoryInformation(mux *http.ServeMux, auth Authenticater, hi HistoryInformationer) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uid := auth.FromContext(r.Context())

		fqid := r.URL.Query().Get("fqid")
		if fqid == "" {
			handleErrorWithStatus(w, invalidRequestError{fmt.Errorf("History Information needs an fqid")})
			return
		}

		if err := hi.HistoryInformation(r.Context(), uid, fqid, w); err != nil {
			handleErrorWithStatus(w, fmt.Errorf("getting history information: %w", err))
			return
		}
	})

	mux.Handle(prefixPublic+"/history_information", authMiddleware(handler, auth))
}

func sendMessages(ctx context.Context, w http.ResponseWriter, uid int, kb autoupdate.KeysBuilder, connecter Connecter, compress bool) error {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Cache-Control", "no-store, max-age=0")

	next, err := connecter.Connect(ctx, uid, kb)
	if err != nil {
		return fmt.Errorf("getting connection: %w", err)
	}

	for f, ok := next(); ok; f, ok = next() {
		// This blocks, until there is new data. It also unblocks, when the
		// client context is done.
		data, err := f(ctx)
		if err != nil {
			return fmt.Errorf("getting next message: %w", err)
		}

		if err := writeData(w, data, compress); err != nil {
			return fmt.Errorf("write data: %w", err)
		}
		w.(http.Flusher).Flush()
	}
	return ctx.Err()
}

func sendMessagesWebsocket(ctx context.Context, w http.ResponseWriter, r *http.Request, uid int, kb autoupdate.KeysBuilder, connecter Connecter, compress bool) error {
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		return fmt.Errorf("accept websocket: %w", err)
	}
	defer c.CloseNow()

	next, err := connecter.Connect(ctx, uid, kb)
	if err != nil {
		return fmt.Errorf("getting connection: %w", err)
	}

	for f, ok := next(); ok; f, ok = next() {
		// This blocks, until there is new data. It also unblocks, when the
		// client context is done.
		data, err := f(ctx)
		if err != nil {
			return fmt.Errorf("getting next message: %w", err)
		}

		if err := c.Write(r.Context(), websocket.MessageText, []byte(fmt.Sprintf(`<div id="websocket">Websocket works: %d.</div>`, rand.Intn(1000)))); err != nil {
			return fmt.Errorf("write data: %w", err)
		}

		wr, err := c.Writer(ctx, websocket.MessageText)
		if err != nil {
			return fmt.Errorf("create writer: %w", err)
		}

		if err := writeData(wr, data, compress); err != nil {
			wr.Close()
			return fmt.Errorf("write data: %w", err)
		}
		wr.Close()
	}
	return ctx.Err()
}

// HandleHealth tells, if the service is running.
func HandleHealth(mux *http.ServeMux) {
	url := prefixPublic + "/health"
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/octet-stream")
		fmt.Fprintln(w, `{"healthy": true}`)
	})

	mux.Handle(url, handler)
}

// HandleProfile adds routes for profiling.
func HandleProfile(mux *http.ServeMux) {
	mux.Handle(prefixInternal+"/debug/pprof/", http.HandlerFunc(pprof.Index))
	mux.Handle(prefixInternal+"/debug/pprof/heap", pprof.Handler("heap"))
	mux.Handle(prefixInternal+"/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	mux.Handle(prefixInternal+"/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	mux.Handle(prefixInternal+"/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	mux.Handle(prefixInternal+"/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
}

func authMiddleware(next http.Handler, auth Authenticater) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, err := auth.Authenticate(w, r)
		if err != nil {
			handleErrorWithStatus(w, fmt.Errorf("authenticate request: %w", err))
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func internalAuthMiddleware(next http.Handler, auth Authenticater) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawUserID := r.URL.Query().Get("user_id")
		userID, err := strconv.Atoi(rawUserID)
		if err != nil {
			handleErrorInternal(w, fmt.Errorf("user_id has to be an int, not %s", rawUserID))
			return
		}

		ctx := auth.AuthenticatedContext(r.Context(), userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func handleErrorWithStatus(w http.ResponseWriter, err error) {
	handleError(w, err, true, false)
}

func handleErrorWithoutStatus(w http.ResponseWriter, err error) {
	handleError(w, err, false, false)
}

// handleErrorInternal is only for internal request routes. It returns the full
// error message to the client.
func handleErrorInternal(w http.ResponseWriter, err error) {
	handleError(w, err, true, true)
}

// handleError interprets the given error and writes a corresponding message to
// the client and/or stdout.
//
// Do not use this function directly but use handleErrorWithStatus,
// handleErrorWithoutStatus or handleErrorInternal.
//
// If the handler already started to write the body then it is not allowed to
// set the http-status-code. In this case, writeStatusCode has to be fales.
func handleError(w http.ResponseWriter, err error, writeStatusCode bool, internal bool) {
	if writeStatusCode {
		w.Header().Set("Content-Type", "application/octet-stream")
	}

	if oserror.ContextDone(err) || errors.Is(err, syscall.EPIPE) || errors.Is(err, syscall.ECONNRESET) {
		// Client closed connection.
		return
	}

	status := http.StatusBadRequest
	var StatusCoder interface{ StatusCode() int }
	if errors.As(err, &StatusCoder) {
		status = StatusCoder.StatusCode()
	}

	var errClient ClientError
	if errors.As(err, &errClient) {
		if writeStatusCode {
			w.WriteHeader(status)
		}

		fmt.Fprintf(w, `{"error": {"type": "%s", "msg": "%s"}}`, errClient.Type(), quote(errClient.Error()))
		return
	}

	if writeStatusCode {
		w.WriteHeader(http.StatusInternalServerError)
	}

	clientOutput := `{"error": {"type": "InternalError", "msg": "Something went wrong on the server. The admin is already informed."}}`
	if internal {
		clientOutput = err.Error()
	}

	oserror.Handle(err)
	fmt.Fprintln(w, clientOutput)
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
			handleErrorWithStatus(w, invalidRequestError{fmt.Errorf("Only GET or POST requests are supported")})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func connectionCountMiddleware(next http.Handler, auth Authenticater, counter *connectionCount) http.Handler {
	if counter == nil {
		return next
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		uid := auth.FromContext(ctx)
		counter.Add(uid)

		defer func() {
			counter.Done(uid)
		}()

		next.ServeHTTP(w, r)
	})
}
