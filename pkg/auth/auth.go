// Package auth implement the auth system from the openslides-auth-service:
// https://github.com/OpenSlides/openslides-auth-service
package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/oserror"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ostcar/topic"
)

// DebugTokenKey and DebugCookieKey are non random auth keys for development.
const (
	DebugTokenKey  = "auth-dev-token-key"
	DebugCookieKey = "auth-dev-cookie-key"
)

var (
	envAuthHost     = environment.NewVariable("AUTH_HOST", "localhost", "Host of the auth service.")
	envAuthPort     = environment.NewVariable("AUTH_PORT", "9004", "Port of the auth service.")
	envAuthProtocol = environment.NewVariable("AUTH_PROTOCOL", "http", "Protocol of the auth service.")
	envAuthFake     = environment.NewVariable("AUTH_FAKE", "false", "Use user id 1 for every request. Ignores all other auth environment variables.")

	envAuthTokenFile  = environment.NewVariable("AUTH_TOKEN_KEY_FILE", "/run/secrets/auth_token_key", "Key to sign the JWT auth tocken.")
	envAuthCookieFile = environment.NewVariable("AUTH_COOKIE_KEY_FILE", "/run/secrets/auth_cookie_key", "Key to sign the JWT auth cookie.")
)

// pruneTime defines how long a topic id will be valid. This should be higher
// then the max livetime of a token.
const pruneTime = 15 * time.Minute

const (
	cookieName = "refreshId"
	authHeader = "Authentication"
	authPath   = "/internal/auth/authenticate"
)

// LogoutEventer tells, when a sessionID gets revoked.
//
// The method LogoutEvent has to block until there are new data. The returned
// data is a list of sessionIDs that are revoked.
type LogoutEventer interface {
	LogoutEvent(context.Context) ([]string, error)
}

// Auth authenticates a request against the openslides-auth-service.
//
// Has to be initialized with auth.New().
type Auth struct {
	fake bool

	logedoutSessions *topic.Topic[string]

	authServiceURL string

	tokenKey  string
	cookieKey string
}

// New initializes the Auth object.
//
// Returns the initialized Auth objectand a function to be called in the
// background.
func New(lookup environment.Environmenter, messageBus LogoutEventer) (*Auth, func(context.Context, func(error)), error) {
	url := fmt.Sprintf(
		"%s://%s:%s",
		envAuthProtocol.Value(lookup),
		envAuthHost.Value(lookup),
		envAuthPort.Value(lookup),
	)

	fake, _ := strconv.ParseBool(envAuthFake.Value(lookup))

	authToken, err := environment.ReadSecretWithDefault(lookup, envAuthTokenFile, DebugTokenKey)
	if err != nil {
		return nil, nil, fmt.Errorf("reading auth token: %w", err)
	}

	cookieToken, err := environment.ReadSecretWithDefault(lookup, envAuthCookieFile, DebugCookieKey)
	if err != nil {
		return nil, nil, fmt.Errorf("reading cookie token: %w", err)
	}

	a := &Auth{
		fake:             fake,
		logedoutSessions: topic.New[string](),
		authServiceURL:   url,
		tokenKey:         authToken,
		cookieKey:        cookieToken,
	}

	// Make sure the topic is not empty
	a.logedoutSessions.Publish("")

	background := func(ctx context.Context, errorHandler func(error)) {
		if fake {
			return
		}

		go a.listenOnLogouts(ctx, messageBus, errorHandler)
		go a.pruneOldData(ctx)
	}

	return a, background, nil
}

// Authenticate uses the headers from the given request to get the user id. The
// returned context will be cancled, if the session is revoked.
func (a *Auth) Authenticate(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	if a.fake {
		return r.Context(), nil
	}

	ctx := r.Context()

	p := new(payload)
	if err := a.loadToken(w, r, p); err != nil {
		return nil, fmt.Errorf("reading token: %w", err)
	}

	if p.UserID == 0 {
		return a.AuthenticatedContext(ctx, 0), nil
	}

	_, sessionIDs, err := a.logedoutSessions.Receive(ctx, 0)
	if err != nil {
		return nil, fmt.Errorf("getting already logged out sessions: %w", err)
	}
	for _, sid := range sessionIDs {
		if sid == p.SessionID {
			return nil, &authError{"invalid session", nil}
		}
	}

	ctx, cancelCtx := context.WithCancel(a.AuthenticatedContext(ctx, p.UserID))

	go func() {
		defer cancelCtx()

		var cid uint64
		var sessionIDs []string
		var err error
		for {
			cid, sessionIDs, err = a.logedoutSessions.Receive(ctx, cid)
			if err != nil {
				return
			}

			for _, sid := range sessionIDs {
				if sid == p.SessionID {
					return
				}
			}
		}
	}()

	return ctx, nil
}

// AuthenticatedContext returns a new context that contains an userID.
//
// Should only used for internal URLs. All other URLs should use auth.Authenticate.
func (a *Auth) AuthenticatedContext(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, userIDType, userID)
}

// FromContext returnes the user id from a context returned by Authenticate().
//
// If the user is an anonymous user 0 is returned.
//
// Panics, if the context was not returned from Authenticate
func (a *Auth) FromContext(ctx context.Context) int {
	if a.fake {
		return 1
	}

	v := ctx.Value(userIDType)
	if v == nil {
		panic("call to auth.FromContext() without auth.Authenticate()")
	}

	return v.(int)
}

// listenOnLogouts listen on logout events and closes the connections.
func (a *Auth) listenOnLogouts(ctx context.Context, logoutEventer LogoutEventer, errHandler func(error)) {
	if errHandler == nil {
		errHandler = func(error) {}
	}

	for {
		data, err := logoutEventer.LogoutEvent(ctx)
		if err != nil {
			if oserror.ContextDone(err) {
				return
			}

			errHandler(fmt.Errorf("receiving logout event: %w", err))
			time.Sleep(time.Second)
			continue
		}

		a.logedoutSessions.Publish(data...)
	}
}

// pruneOldData removes old logout events.
func (a *Auth) pruneOldData(ctx context.Context) {
	tick := time.NewTicker(5 * time.Minute)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			a.logedoutSessions.Prune(time.Now().Add(-pruneTime))
		}
	}
}

// loadToken loads and validates the token. If the token is expires, it tries
// to renews it and writes the new token to the responsewriter.
func (a *Auth) loadToken(w http.ResponseWriter, r *http.Request, payload jwt.Claims) error {
	header := r.Header.Get(authHeader)
	cookie, err := r.Cookie(cookieName)
	if err != nil && err != http.ErrNoCookie {
		return fmt.Errorf("reading cookie: %w", err)
	}

	encodedToken := strings.TrimPrefix(header, "bearer ")

	if cookie == nil && header == encodedToken {
		// No token and no auth cookie. Handle the request as anonymous requst.
		return nil
	}

	if cookie == nil && header != encodedToken {
		return authError{"Can not find auth cookie", nil}
	}

	if cookie != nil && header == encodedToken {
		return authError{"Can not find auth token", nil}
	}

	encodedCookie := strings.TrimPrefix(cookie.Value, "bearer%20")

	_, err = jwt.Parse(encodedCookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.cookieKey), nil
	})
	if err != nil {
		var invalid *jwt.ValidationError
		if errors.As(err, &invalid) {
			return authError{"Invalid auth token", err}
		}
		return fmt.Errorf("validating auth cookie: %w", err)
	}

	_, err = jwt.ParseWithClaims(encodedToken, payload, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.tokenKey), nil
	})
	if err != nil {
		var invalid *jwt.ValidationError
		if errors.As(err, &invalid) {
			return a.handleInvalidToken(r.Context(), invalid, w, encodedToken, encodedCookie)
		}
	}

	return nil
}

func (a *Auth) handleInvalidToken(ctx context.Context, invalid *jwt.ValidationError, w http.ResponseWriter, encodedToken, encodedCookie string) error {
	if !tokenExpired(invalid.Errors) {
		return authError{"Invalid auth token", invalid}
	}

	token, err := a.refreshToken(ctx, encodedToken, encodedCookie)
	if err != nil {
		return fmt.Errorf("refreshing token: %w", err)
	}

	w.Header().Set(authHeader, token)
	return nil
}

func tokenExpired(errNo uint32) bool {
	return errNo&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0
}

func (a *Auth) refreshToken(ctx context.Context, token, cookie string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", a.authServiceURL+authPath, nil)
	if err != nil {
		return "", fmt.Errorf("creating auth request: %w", err)
	}

	req.Header.Add(authHeader, "bearer "+token)
	req.AddCookie(&http.Cookie{Name: cookieName, Value: "bearer " + cookie, HttpOnly: true, Secure: true})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// TODO External ERROR
		return "", fmt.Errorf("send request to auth service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		if resp.StatusCode == 403 {
			return "", authError{msg: "Invalid Session", wrapped: err}
		}
		// TODO LAST ERROR
		return "", fmt.Errorf("auth-service returned status %s", resp.Status)
	}

	newToken := resp.Header.Get(authHeader)
	if newToken == "" {
		var rPayload struct {
			Message string `json:"message"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&rPayload); err != nil {
			return "", fmt.Errorf("decoding auth response: %w", err)
		}
		if rPayload.Message == "" {
			rPayload.Message = "Can not refresh token"
		}
		return "", authError{rPayload.Message, nil}

	}

	return newToken, nil
}

type authString string

const (
	userIDType authString = "user_id"
)

type payload struct {
	jwt.StandardClaims
	UserID    int    `json:"userId"`
	SessionID string `json:"sessionId"`
}
