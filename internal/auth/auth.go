package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go/v4"
)

// Auth authenticates a request against the auth service.
type Auth struct{}

// Authenticate uses the headers from the given request to get the user id. The
// returned context will be cancled, if the session is revoced.
func (a *Auth) Authenticate(r *http.Request) (ctx context.Context, cancel func(), err error) {
	p := new(payload)
	if err := loadToken(r, p); err != nil {
		return nil, nil, fmt.Errorf("reading token: %w", err)
	}

	if p.UserID == 0 {
		// Empty token or anonymous token. No need to save anything in the
		// context.
		return r.Context(), func() {}, nil
	}

	ctx = context.WithValue(r.Context(), userIDType, p.UserID)

	// TODO:
	// listen if p.SessionID is revoced and if so, cancel the ctx.
	// if the cancel func was called, stop listening.
	// Is a cancel function necessary or would it be possible to listen on r.Context()?
	return ctx, func() {}, nil
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

type authString string

const userIDType authString = "user_id"
const authHeader = "authentication"

func loadToken(r *http.Request, payload jwt.Claims) error {
	header := r.Header.Get(authHeader)

	encoded := strings.TrimPrefix(header, "bearer ")
	if header == encoded {
		// No valid token in request. Handle the request as anonymous requst.
		return nil
	}

	// TODO: SECURITY ISSUE!!! Use jwt.ParseWithClaims(encoded, KEY, payload)
	_, _, err := jwt.NewParser().ParseUnverified(encoded, payload)
	if err != nil {
		return err
	}
	return nil
}

type payload struct {
	UserID    int    `json:"userId"`
	SessionID string `json:"sessionId"`
}

func (p *payload) Valid(*jwt.ValidationHelper) error {
	return nil
}
