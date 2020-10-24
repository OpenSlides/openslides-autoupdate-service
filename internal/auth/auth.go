package auth

import (
	"bytes"
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
const authPath = "/service/auth/api/refresh"

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
			// TODO: handle closing error
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

	_, err = jwt.Parse(cookie.Value, jwt.KnownKeyfunc(jwt.SigningMethodHS256, a.cookieKey))
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

		token, err := a.refreshToken(encodedToken, cookie.Value)
		if err != nil {
			return fmt.Errorf("refreshing token: %w", err)
		}
		w.Header().Set(authHeader, "bearer "+token)
	}

	return nil
}

func (a *Auth) refreshToken(token, cookie string) (string, error) {
	payload := struct {
		Auth    string `json:"authentication"`
		Cookies string `json:"cookies"`
	}{
		token,
		cookie,
	}
	enc, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("encoding payload to auth service: %w", err)
	}
	resp, err := http.Post(a.authServiceURL, "application/json", bytes.NewReader(enc))
	if err != nil {
		return "", fmt.Errorf("asking auth-service for new token: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("auth-service returned status %s", resp.Status)
	}

	var rPayload struct {
		Token   string `json:"token"`
		Message string `json:"message"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&rPayload); err != nil {
		return "", fmt.Errorf("decoding auth response: %w", err)
	}
	if rPayload.Token == "" {
		if rPayload.Message == "" {
			rPayload.Message = "Can not refresh token"
		}
		return "", authError{rPayload.Message, nil}
	}
	return rPayload.Token, nil
}

type authString string

const userIDType authString = "user_id"

type payload struct {
	jwt.StandardClaims
	UserID    int    `json:"userId"`
	SessionID string `json:"sessionId"`
}
