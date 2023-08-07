package auth_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/auth"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/auth/authtest"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
	"github.com/golang-jwt/jwt/v4"
)

func parseURL(raw string) (host, port, protocol string) {
	parsed, err := url.Parse(raw)
	if err != nil {
		panic(fmt.Sprintf("parsing url %s: %v", raw, err))
	}

	return parsed.Hostname(), parsed.Port(), parsed.Scheme
}

func TestAuth(t *testing.T) {
	const invalidSecret = "wrong-auth-dev-key"
	const cookieName = "refreshId"

	cookie, authHeader, validHeader, err := authtest.ValidTokens([]byte(auth.DebugCookieKey), []byte(auth.DebugTokenKey), 1)
	if err != nil {
		t.Fatalf("Create tokens: %v", err)
	}

	validCookie := cookie.String()

	oldHeader, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    1,
		"sessionId": "123",
		"exp":       123,
	}).SignedString([]byte(auth.DebugTokenKey))
	if err != nil {
		t.Fatalf("Can not sign token token: %v", err)
	}
	oldHeader = "bearer " + oldHeader

	invalidCookie, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sessionId": "123",
	}).SignedString([]byte(invalidSecret))
	if err != nil {
		t.Fatalf("Can not sign cookie token: %v", err)
	}
	invalidCookie = cookieName + "=bearer%20" + invalidCookie

	invalidHeader, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    1,
		"sessionId": "123",
	}).SignedString([]byte(invalidSecret))
	if err != nil {
		t.Fatalf("Can not sign token token: %v", err)
	}
	invalidHeader = "bearer " + invalidHeader

	authSRV := httptest.NewServer(&mockAuth{token: "NEWTOKEN"})
	defer authSRV.Close()

	host, port, schema := parseURL(authSRV.URL)
	env := environment.ForTests(map[string]string{
		"AUTH_HOST":     host,
		"AUTH_PORT":     port,
		"AUTH_PROTOCOL": schema,
	})
	a, _, _ := auth.New(env, nil)

	for _, tt := range []struct {
		name    string
		request *http.Request
		uid     int
		header  string
		errMSG  string
	}{
		{
			"No cookie no token",
			&http.Request{},
			0,
			"",
			"",
		},
		{
			"Valid cookie no token",
			&http.Request{
				Header: map[string][]string{
					"Cookie": {validCookie},
				},
			},
			0,
			"",
			"Can not find auth token",
		},
		{
			"No cookie Valid token",
			&http.Request{
				Header: map[string][]string{
					authHeader: {validHeader},
				},
			},
			0,
			"",
			"Can not find auth cookie",
		},
		{
			"Invalid cookie Valid token",
			&http.Request{
				Header: map[string][]string{
					"Cookie":   {invalidCookie},
					authHeader: {validHeader},
				},
			},
			0,
			"",
			"Invalid auth token",
		},
		{
			"Valid cookie Invalid token",
			&http.Request{
				Header: map[string][]string{
					"Cookie":   {validCookie},
					authHeader: {invalidHeader},
				},
			},
			0,
			"",
			"Invalid auth token",
		},
		{
			"Valid cookie Valid token",
			&http.Request{
				Header: map[string][]string{
					"Cookie":   {validCookie},
					authHeader: {validHeader},
				},
			},
			1,
			"",
			"",
		},
		{
			"Valid cookie Valid Old token",
			&http.Request{
				Header: map[string][]string{
					"Cookie":   {validCookie},
					authHeader: {oldHeader},
				},
			},
			1,
			"NEWTOKEN",
			"",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, err := a.Authenticate(w, tt.request)

			if tt.errMSG != "" {
				if err == nil {
					t.Fatalf("Got no error, expected `%s`", tt.errMSG)
				}

				var clientErr interface {
					Type() string
					Error() string
				}

				if !errors.As(err, &clientErr) {
					t.Fatalf("Expected a client error, got: %v", err)
				}

				if clientErr.Type() != "auth" {
					t.Fatalf("Got error of type %s, expected `auth`", clientErr.Type())
				}

				if got := clientErr.Error(); got != tt.errMSG {
					t.Errorf("Got err `%s`, expected `%s`", got, tt.errMSG)
				}
				return
			}

			if err != nil {
				t.Fatalf("Auth returned an unexpected error: %v", err)
			}

			if got := w.Result().Header.Get(authHeader); got != tt.header {
				t.Errorf("Got header `%s`, expected `%s`", got, tt.header)
			}

			if got := a.FromContext(ctx); got != tt.uid {
				t.Errorf("Got uid %d, expected %d", got, tt.uid)
			}
		})
	}
}

func TestFromContext(t *testing.T) {
	a, _, _ := auth.New(environment.ForTests{}, nil)

	t.Run("Empty Context", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("FromContext() did not panic")
			}
		}()

		a.FromContext(context.Background())
	})

	t.Run("Context from Authenticate", func(t *testing.T) {
		ctx, err := a.Authenticate(validSession(t))
		if err != nil {
			t.Fatalf("Can not create context from Authenticate: %v", err)
		}

		got := a.FromContext(ctx)
		if got != 1 {
			t.Errorf("Got uid %d from auth-context. Expected 1", got)
		}
	})

	t.Run("Context from AuthenticatedContext", func(t *testing.T) {
		ctx := context.Background()
		ctx = a.AuthenticatedContext(ctx, 7)

		got := a.FromContext(ctx)
		if got != 7 {
			t.Errorf("Got uid %d from auth-context. Expected 7", got)
		}
	})
}

func TestLogout(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var lastErr error
	errHandler := func(err error) {
		lastErr = err
	}

	logouter := NewLockoutEventMock()
	defer logouter.Close()

	a, bg, _ := auth.New(environment.ForTests{}, logouter)
	go bg(shutdownCtx, errHandler)

	t.Run("Closing session", func(t *testing.T) {
		ctx, err := a.Authenticate(validSession(t, withSessionID("session1")))
		if err != nil {
			t.Fatalf("Can not authenticat: %v", err)
		}

		logouter.Send([]string{"session1"})

		timer := time.NewTimer(time.Millisecond)
		defer timer.Stop()
		select {
		case <-ctx.Done():
		case <-timer.C:
			t.Errorf("context is not closed after logout")
		}

		if lastErr != nil {
			t.Errorf("Got error on logout: %v", err)
		}
	})

	t.Run("Already closed session", func(t *testing.T) {
		_, err := a.Authenticate(validSession(t, withSessionID("session1")))
		if err == nil {
			t.Fatalf("Got no error. Expected an auth error")
		}

		var clientErr interface {
			Type() string
			Error() string
		}

		if !errors.As(err, &clientErr) {
			t.Fatalf("Expected a client error, got: %v", err)
		}

		if clientErr.Type() != "auth" {
			t.Fatalf("Got error of type %s, expected `auth`", clientErr.Type())
		}
	})

	t.Run("Closing other session", func(t *testing.T) {
		ctx, err := a.Authenticate(validSession(t, withSessionID("session2")))
		if err != nil {
			t.Fatalf("Can not authenticat: %v", err)
		}

		logouter.Send([]string{"session3"})

		timer := time.NewTimer(time.Millisecond)
		defer timer.Stop()
		select {
		case <-ctx.Done():
			t.Errorf("context is closed after logout of other session")
		case <-timer.C:
		}

		if lastErr != nil {
			t.Errorf("Got error on logout: %v", err)
		}
	})
}

func validSession(t *testing.T, opts ...validOption) (http.ResponseWriter, *http.Request) {
	config := &validConfig{
		sessionID: "123",
	}
	for _, o := range opts {
		o(config)
	}

	validCookie, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sessionId": config.sessionID,
	}).SignedString([]byte(auth.DebugCookieKey))
	if err != nil {
		t.Fatalf("Can not sign cookie token: %v", err)
	}
	validCookie = "refreshId=" + validCookie

	validHeader, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    1,
		"sessionId": config.sessionID,
	}).SignedString([]byte(auth.DebugTokenKey))
	if err != nil {
		t.Fatalf("Can not sign token token: %v", err)
	}
	validHeader = "bearer " + validHeader

	w := httptest.NewRecorder()
	r := &http.Request{
		Header: map[string][]string{
			"Cookie":         {validCookie},
			"Authentication": {validHeader},
		},
	}
	return w, r
}

type validConfig struct {
	sessionID string
}

type validOption func(*validConfig)

func withSessionID(sid string) validOption {
	return func(v *validConfig) {
		v.sessionID = sid
	}
}
