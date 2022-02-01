// Package auth implement the auth system from the openslides-auth-service:
// https://github.com/OpenSlides/openslides-auth-service
package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/ostcar/topic"
)

// DebugTokenKey and DebugCookieKey are non random auth keys for development.
const (
	DebugTokenKey  = "auth-dev-token-key"
	DebugCookieKey = "auth-dev-cookie-key"
)

// pruneTime defines how long a topic id will be valid. This should be higher
// then the max livetime of a token.
const pruneTime = 15 * time.Minute

const cookieName = "refreshId"
const authHeader = "Authentication"
const authPath = "/internal/auth/authenticate"

// Auth authenticates a request against the openslides-auth-service.
//
// Has to be initialized with auth.New().
type Auth struct {
	logedoutSessions *topic.Topic

	authServiceURL string

	tokenKey  []byte
	cookieKey []byte
}

// New initializes an Auth service.
func New(
	authServiceURL string,
	closed <-chan struct{},
	tokenKey,
	cookieKey []byte,
) (*Auth, error) {
	a := &Auth{
		logedoutSessions: topic.New(topic.WithClosed(closed)),
		authServiceURL:   authServiceURL,
		tokenKey:         tokenKey,
		cookieKey:        cookieKey,
	}

	// Make sure the topic is not empty
	a.logedoutSessions.Publish("")

	return a, nil
}

// Authenticate uses the headers from the given request to get the user id. The
// returned context will be cancled, if the session is revoked.
func (a *Auth) Authenticate(w http.ResponseWriter, r *http.Request) (ctx context.Context, err error) {
	p := new(payload)
	if err := a.loadToken(w, r, p); err != nil {
		return nil, fmt.Errorf("reading token: %w", err)
	}

	if p.UserID == 0 {
		// Empty token or anonymous token. No need to save anything in the
		// context.
		return r.Context(), nil
	}

	_, sessionIDs, err := a.logedoutSessions.Receive(context.Background(), 0)
	if err != nil {
		return nil, fmt.Errorf("getting already logged out sessions: %w", err)
	}
	for _, sid := range sessionIDs {
		if sid == p.SessionID {
			return nil, &authError{"invalid session", nil}
		}
	}

	ctx, cancelCtx := context.WithCancel(context.WithValue(r.Context(), userIDType, p.UserID))

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

// FromContext returnes the user id from a context returned by Authenticate().
// If the context was not returned from Authenticate or the user is an anonymous
// user, then 0 is returned.
func (a *Auth) FromContext(ctx context.Context) int {
	v := ctx.Value(userIDType)
	if v == nil {
		return 0
	}

	return v.(int)
}

// LogoutEventer tells, when a sessionID gets revoked.
//
// The method LogoutEvent has to block until there are new data. The returned
// data is a list of sessionIDs that are revoked.
type LogoutEventer interface {
	LogoutEvent(context.Context) ([]string, error)
}

// ListenOnLogouts listen on logout events and closes the connections.
func (a *Auth) ListenOnLogouts(ctx context.Context, logoutEventer LogoutEventer, errHandler func(error)) {
	if errHandler == nil {
		errHandler = func(error) {}
	}

	for {
		data, err := logoutEventer.LogoutEvent(ctx)
		if err != nil {
			errHandler(fmt.Errorf("receiving logout event: %w", err))
			time.Sleep(time.Second)
			continue
		}

		a.logedoutSessions.Publish(data...)
	}
}

// PruneOldData removes old logout events.
func (a *Auth) PruneOldData(ctx context.Context) {
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

// loadToken loads and validates the ticket. If the token is expires, it tries
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
		return a.cookieKey, nil
	})
	if err != nil {
		var invalid *jwt.ValidationError
		if errors.As(err, &invalid) {
			return authError{"Invalid auth ticket", err}
		}
		return fmt.Errorf("validating auth cookie: %w", err)
	}

	_, err = jwt.ParseWithClaims(encodedToken, payload, func(token *jwt.Token) (interface{}, error) {
		return a.tokenKey, nil
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
		return authError{"Invalid auth ticket:", invalid}
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
		return "", fmt.Errorf("send request to auth service: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
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

const userIDType authString = "user_id"

type payload struct {
	jwt.StandardClaims
	UserID    int    `json:"userId"`
	SessionID string `json:"sessionId"`
}
