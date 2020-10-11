package auth

import (
	"context"
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

// Auth authenticates a request against the auth service.
type Auth struct {
	logoutEventer    LogoutEventer
	logedoutSessions *topic.Topic
	closed           <-chan struct{}
}

// New initializes a Auth service.
func New(logoutEventer LogoutEventer, closed <-chan struct{}, errHandler func(error)) *Auth {
	a := &Auth{
		closed:           closed,
		logoutEventer:    logoutEventer,
		logedoutSessions: topic.New(topic.WithClosed(closed)),
	}

	go a.receiveLogoutEvent(errHandler)
	go a.pruneLogoutEvent(closed)

	return a
}

// Authenticate uses the headers from the given request to get the user id. The
// returned context will be cancled, if the session is revoced.
func (a *Auth) Authenticate(r *http.Request) (ctx context.Context, err error) {
	p := new(payload)
	if err := loadToken(r, p); err != nil {
		return nil, fmt.Errorf("reading token: %w", err)
	}

	if p.UserID == 0 {
		// Empty token or anonymous token. No need to save anything in the
		// context.
		return r.Context(), nil
	}

	ctx, cancelCtx := context.WithCancel(context.WithValue(r.Context(), userIDType, p.UserID))

	go func() {
		var cid uint64
		var sessionIDs []string
		var err error
		for {
			cid, sessionIDs, err = a.logedoutSessions.Receive(ctx, cid)
			if err != nil {
				// TODO: Do something with the error.
				return
			}

			for _, sid := range sessionIDs {
				if sid == p.SessionID {
					cancelCtx()
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
	for {
		select {
		case <-a.closed:
			return
		default:
		}

		data, err := a.logoutEventer.LogoutEvent()
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
