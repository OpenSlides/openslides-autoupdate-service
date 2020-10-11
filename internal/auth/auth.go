package auth

import (
	"context"
	"net/http"
)

// Auth authenticates a request against the auth service.
type Auth struct{}

// Authenticate uses the headers from the given request to get the user id. The
// returned context will be cancled, if the session will be canceled.
func (a *Auth) Authenticate(r *http.Request) (ctx context.Context, cancel func(), err error) {
	// TODO:
	// 1. Extract the token and the sessionID from the request.
	// 2. Validate the token and get the uid from it.
	// 3. Save the uid in the the context.
	// 4. In a background task, listen if the session was invalidated by the auth service. IF so, cancel the context.
	// 5. Return the new context and a cancel func to stop the background task.
	cancel = func() {}
	ctx = context.WithValue(r.Context(), userIDType, 1)
	return ctx, cancel, nil
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
