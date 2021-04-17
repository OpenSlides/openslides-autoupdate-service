package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/ostcar/topic"
)

// pruneTime defines how long a topic id will be valid. This should be higher
// then the max livetime of a token.
const pruneTime = 15 * time.Minute

const cookieName = "refreshId"
const authHeader = "Authentication"
const authPath = "/internal/auth/authenticate"

// Auth authenticates a request against the auth service.
type Auth struct {
	logoutEventer    LogoutEventer
	logedoutSessions *topic.Topic
	closed           <-chan struct{}
	errHandler       func(error)

	authServiceURL string

	tokenKey  []byte
	cookieKey []byte
}

// New initializes a Auth service.
func New(authServiceURL string, logoutEventer LogoutEventer, closed <-chan struct{}, errHandler func(error), tokenKey, cookieKey []byte) (*Auth, error) {
	a := &Auth{
		closed:           closed,
		errHandler:       errHandler,
		logoutEventer:    logoutEventer,
		logedoutSessions: topic.New(topic.WithClosed(closed)),
		authServiceURL:   authServiceURL,
		tokenKey:         tokenKey,
		cookieKey:        cookieKey,
	}

	// Make sure the topic is not empty
	a.logedoutSessions.Publish("")

	if logoutEventer != nil {
		go a.receiveLogoutEvent(errHandler)
	}

	go a.pruneLogoutEvent(closed)

	return a, nil
}

// Authenticate uses the headers from the given request to get the user id. The
// returned context will be cancled, if the session is revoced.
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

// FromContext returnes the user id from a context returned by Authenticate. If
// the context was not returned from Authenticate or the user is an anonymous
// user, then 0 is returned.
func (a *Auth) FromContext(ctx context.Context) int {
	v := ctx.Value(userIDType)
	if v == nil {
		return 0
	}

	return v.(int)
}

func (a *Auth) receiveLogoutEvent(errHandler func(error)) {
	if errHandler == nil {
		errHandler = func(error) {}
	}
	for {
		select {
		case <-a.closed:
			return
		default:
		}

		data, err := a.logoutEventer.LogoutEvent(a.closed)
		if err != nil {
			errHandler(fmt.Errorf("receiving logout event: %w", err))
			time.Sleep(time.Second)
			continue
		}

		a.logedoutSessions.Publish(data...)
	}
}

func (a *Auth) pruneLogoutEvent(closed <-chan struct{}) {
	tick := time.NewTicker(5 * time.Minute)
	defer tick.Stop()

	for {
		select {
		case <-closed:
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

	_, err = jwt.Parse(encodedCookie, jwt.KnownKeyfunc(jwt.SigningMethodHS256, a.cookieKey))
	if err != nil {
		var invalid *jwt.InvalidSignatureError
		if errors.As(err, &invalid) {
			return authError{"Invalid auth ticket", err}
		}
		return fmt.Errorf("validating auth cookie: %w", err)
	}

	_, err = jwt.ParseWithClaims(encodedToken, payload, jwt.KnownKeyfunc(jwt.SigningMethodHS256, a.tokenKey))
	if err != nil {
		var invalid *jwt.InvalidSignatureError
		if errors.As(err, &invalid) {
			return authError{"Invalid auth ticket", err}
		}

		var expired *jwt.TokenExpiredError
		if !errors.As(err, &expired) {
			return fmt.Errorf("validating auth token: %w", err)
		}

		token, err := a.refreshToken(r.Context(), encodedToken, encodedCookie)
		if err != nil {
			return fmt.Errorf("refreshing token: %w", err)
		}
		w.Header().Set(authHeader, "bearer "+token)
	}

	return nil
}

func (a *Auth) refreshToken(ctx context.Context, token, cookie string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", a.authServiceURL+authPath, nil)
	if err != nil {
		return "", fmt.Errorf("creating auth request: %w", err)
	}

	req.Header.Add(authHeader, "bearer "+token)
	req.AddCookie(&http.Cookie{Name: cookieName, Value: "bearer " + cookie})

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
