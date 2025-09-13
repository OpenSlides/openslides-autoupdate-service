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
	"iter"
	"mime"
	"mime/multipart"
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
	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/OpenSlides/openslides-go/oserror"
	"github.com/OpenSlides/openslides-go/redis"
	"github.com/klauspost/compress/zstd"
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
	redisConnection *redis.Redis,
	saveInterval time.Duration,
	heartbeat time.Duration,

) error {
	var connectionCount [2]*ConnectionCount
	connectionCount[0] = newConnectionCount(ctx, redisConnection, saveInterval, "connections_stream")
	connectionCount[1] = newConnectionCount(ctx, redisConnection, saveInterval, "connections_longpolling")
	metric.Register(connectionCount[0].Metric)
	metric.Register(connectionCount[1].Metric)

	mux := http.NewServeMux()
	HandleHealth(mux)
	HandleAutoupdate(mux, auth, autoupdate, connectionCount, heartbeat)
	HandleInternalAutoupdate(mux, auth, autoupdate, heartbeat)
	HandleShowConnectionCount(mux, autoupdate, auth, connectionCount)
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
	Connect(ctx context.Context, userID int, kb autoupdate.KeysBuilder) (autoupdate.Connection, error)
	SingleData(ctx context.Context, userID int, kb autoupdate.KeysBuilder) (map[dskey.Key][]byte, error)
}

func autoupdateHandler(auth Authenticater, connecter Connecter, heartbeat time.Duration) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Cache-Control", "no-store, max-age=0")
		ctx := r.Context()

		defer r.Body.Close()
		uid := auth.FromContext(r.Context())

		queryBuilder, err := keysbuilder.FromKeys(strings.Split(r.URL.Query().Get("k"), ",")...)
		if err != nil {
			handleErrorWithStatus(w, fmt.Errorf("building keysbuilder from query: %w", err))
			return
		}

		body, hashes, isLongPolling, err := parseBody(r)
		if err != nil {
			if !(errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF)) {
				// EOF errors happen, when clients close the conections. No need
				// to inform about it
				handleErrorWithStatus(w, fmt.Errorf("parse Body: %w", err))
			}
			return
		}

		bodyBuilder, err := keysbuilder.ManyFromJSON(bytes.NewReader(body))
		if err != nil {
			handleErrorWithStatus(w, fmt.Errorf("building keysbuilder from body: %w", err))
			return
		}

		builder := keysbuilder.FromBuilders(queryBuilder, bodyBuilder)

		var compress bool
		if r.URL.Query().Has("compress") {
			compress = true
		}

		if r.URL.Query().Has("single") {
			data, err := connecter.SingleData(ctx, uid, builder)
			if err != nil {
				handleErrorWithStatus(w, fmt.Errorf("getting single data: %w", err))
				return
			}

			if err := writeData(w, data, compress); err != nil {
				handleErrorWithoutStatus(w, err)
				return
			}
			return
		}

		if isLongPolling {
			if headersSent, err := handleLongpolling(ctx, w, uid, builder, connecter, compress, hashes); err != nil {
				if headersSent {
					handleErrorWithoutStatus(w, err)
				} else {
					handleErrorWithStatus(w, err)
				}
			}
			return
		}

		if err := sendMessages(ctx, w, uid, builder, connecter, compress, heartbeat); err != nil {
			handleErrorWithoutStatus(w, err)
			return
		}
	})
}

func parseBody(r *http.Request) ([]byte, string, bool, error) {
	contentType := r.Header.Get("Content-Type")

	if contentType == "" {
		return parseBodyNormal(r)
	}

	mediaType, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		return nil, "", false, fmt.Errorf("parsing multipart: %w", err)
	}

	if !strings.HasPrefix(mediaType, "multipart/") {
		return parseBodyNormal(r)
	}

	mr := multipart.NewReader(r.Body, params["boundary"])

	kbPart, err := mr.NextPart()
	if err != nil {
		return nil, "", false, fmt.Errorf("no keysbuilder part: %w", err)
	}

	body, err := io.ReadAll(kbPart)
	if err != nil {
		return nil, "", false, fmt.Errorf("invalid multipart body, key-builder-part: %w", err)
	}

	lpPart, err := mr.NextPart()
	if err != nil {
		return nil, "", false, fmt.Errorf("no longpolling part: %w", err)
	}

	longPollinHashes, err := io.ReadAll(lpPart)
	if err != nil {
		return nil, "", false, fmt.Errorf("invalid multipart body, long-poll-part: %w", err)
	}

	return body, string(longPollinHashes), true, nil
}

func parseBodyNormal(r *http.Request) ([]byte, string, bool, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, "", false, fmt.Errorf("invalid body: %w", err)
	}

	return body, "", r.URL.Query().Has("longpolling"), nil
}

// HandleAutoupdate builds the requested keys from the body of a request. The
// body has to be in the format specified in the keysbuilder package.
func HandleAutoupdate(
	mux *http.ServeMux,
	auth Authenticater,
	connecter Connecter,
	connectionCount [2]*ConnectionCount,
	heartbeat time.Duration,
) {
	mux.Handle(
		prefixPublic,
		validRequest(
			authMiddleware(
				connectionCountMiddleware(
					autoupdateHandler(auth, connecter, heartbeat),
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
func HandleInternalAutoupdate(
	mux *http.ServeMux,
	auth Authenticater,
	connecter Connecter,
	heartbeat time.Duration,
) {
	mux.Handle(
		prefixInternal,
		validRequest(
			internalAuthMiddleware(
				autoupdateHandler(auth, connecter, heartbeat),
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

	if compress {
		defer fmt.Fprintln(w)
		base64Encoder := base64.NewEncoder(base64.RawStdEncoding, w)
		defer base64Encoder.Close()

		zstdEncoder, err := zstd.NewWriter(base64Encoder)
		if err != nil {
			return fmt.Errorf("create encoder: %w", err)
		}
		defer zstdEncoder.Close()
		w = zstdEncoder
	}

	if err := json.NewEncoder(w).Encode(converted); err != nil {
		return fmt.Errorf("encode data: %w", err)
	}

	return nil
}

// HandleShowConnectionCount adds a handler to show the result of the connection counter.
func HandleShowConnectionCount(mux *http.ServeMux, autoupdate *autoupdate.Autoupdate, auth Authenticater, connectionCount [2]*ConnectionCount) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if connectionCount[0] == nil {
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

		filter := func(ctx context.Context, count map[int]int) error {
			return autoupdate.FilterConnectionCount(ctx, meetingIDs, count)
		}

		val1, err := connectionCount[0].Show(ctx, filter)
		if err != nil {
			oserror.Handle(fmt.Errorf("Error counting normal connection: %w", err))
			http.Error(w, "Counting not possible", 500)
			return
		}

		val2, err := connectionCount[1].Show(ctx, filter)
		if err != nil {
			oserror.Handle(fmt.Errorf("Error counting longpolling connection: %w", err))
			http.Error(w, "Counting not possible", 500)
			return
		}

		if err := autoupdate.FilterConnectionCount(ctx, meetingIDs, val2); err != nil {
			oserror.Handle(fmt.Errorf("Error filtering connection count: %w", err))
			http.Error(w, "Counting not possible", 500)
			return
		}

		if err := json.NewEncoder(w).Encode([2]map[int]int{val1, val2}); err != nil {
			oserror.Handle(fmt.Errorf("Error decoding counter %w", err))
			http.Error(w, "Counting not possible", 500)
			return
		}
	})

	mux.Handle(prefixPublic+"/connection_count", authMiddleware(handler, auth))
}

func handleLongpolling(ctx context.Context, w http.ResponseWriter, uid int, kb autoupdate.KeysBuilder, connecter Connecter, compress bool, hashes string) (bool, error) {
	conn, err := connecter.Connect(ctx, uid, kb)
	if err != nil {
		return false, fmt.Errorf("getting connection: %w", err)
	}

	data, newHashes, err := conn.NextWithFilter(ctx, hashes)
	if err != nil {
		return false, fmt.Errorf("getting data: %w", err)
	}

	mp := multipart.NewWriter(w)
	w.Header().Set("Content-Type", mp.FormDataContentType())
	dataWriter, err := mp.CreateFormField("data")
	if err != nil {
		return true, fmt.Errorf("creating data part: %w", err)
	}

	if err := writeData(dataWriter, data, compress); err != nil {
		return true, fmt.Errorf("write data: %w", err)
	}

	hashWriter, err := mp.CreateFormField("hash")
	if err != nil {
		return true, fmt.Errorf("creating hashes part: %w", err)
	}

	if _, err := hashWriter.Write([]byte(newHashes)); err != nil {
		return true, fmt.Errorf("writing hashes: %w", err)
	}

	if err := mp.Close(); err != nil {
		return true, fmt.Errorf("close multipart: %w", err)
	}

	return true, nil
}

func sendMessages(
	ctx context.Context,
	w io.Writer,
	uid int,
	kb autoupdate.KeysBuilder,
	connecter Connecter,
	compress bool,
	heartbeat time.Duration,
) error {
	conn, err := connecter.Connect(ctx, uid, kb)
	if err != nil {
		return fmt.Errorf("getting connection: %w", err)
	}

	for data, err := range insertKeepAlive(conn.Messages(ctx), heartbeat) {
		// This blocks, until there is new data. It also unblocks, when the
		// client context is done.
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

// insertKeepAlive sends a empty message, if there is no message after
// heartbeat.
//
// This is a workaround for a firefox bug. Can be removed when we find a better
// method to close old connections.
func insertKeepAlive(in iter.Seq2[map[dskey.Key][]byte, error], heartbeat time.Duration) iter.Seq2[map[dskey.Key][]byte, error] {
	return func(yield func(key map[dskey.Key][]byte, value error) bool) {
		dataCh := make(chan struct {
			data map[dskey.Key][]byte
			err  error
		})

		go func() {
			defer close(dataCh)
			for data, err := range in {
				dataCh <- struct {
					data map[dskey.Key][]byte
					err  error
				}{data, err}
			}
		}()

		timer := time.NewTimer(heartbeat)
		defer timer.Stop()

		for {
			select {
			case item, ok := <-dataCh:
				if !ok {
					return
				}
				timer.Reset(heartbeat)

				if !yield(item.data, item.err) {
					return
				}

			case <-timer.C:
				if !yield(map[dskey.Key][]byte{}, nil) {
					return
				}
				timer.Reset(heartbeat)
			}
		}
	}
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
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
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
			handleErrorWithStatus(w, invalidRequestError{fmt.Errorf("only GET or POST requests are supported")})
			return
		}

		next.ServeHTTP(w, r)
	})
}

// isLongPollingRequest returns, if the request is a longpolling fallback
// request.
//
// This is the case, if it has the argument "longpolling" or if the body is
// multipart.
func isLongPollingRequest(r *http.Request) bool {
	return r.URL.Query().Has("longpolling") || strings.HasPrefix(strings.ToLower(r.Header.Get("Content-Type")), "multipart/")
}

func connectionCountMiddleware(next http.Handler, auth Authenticater, counter [2]*ConnectionCount) http.Handler {
	if counter[0] == nil {
		return next
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		uid := auth.FromContext(ctx)
		count := counter[0]
		if isLongPollingRequest(r) {
			count = counter[1]
		}

		count.Add(uid)

		defer func() {
			count.Done(uid)
		}()

		next.ServeHTTP(w, r)
	})
}
