package auth_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/openslides/openslides-autoupdate-service/internal/auth"
)

func TestAuth(t *testing.T) {
	const secret = "auth-dev-key"
	const invalidSecret = "wrong-auth-dev-key"
	const cookieName = "refreshId"
	const authHeader = "Authentication"

	validCookie, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sessionId": "123",
	}).SignedString([]byte(secret))
	if err != nil {
		t.Fatalf("Can not sign cookie token: %v", err)
	}
	validCookie = cookieName + "=" + validCookie

	validHeader, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    1,
		"sessionId": "123",
	}).SignedString([]byte(secret))
	if err != nil {
		t.Fatalf("Can not sign token token: %v", err)
	}
	validHeader = "bearer " + validHeader

	oldHeader, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    1,
		"sessionId": "123",
		"exp":       123,
	}).SignedString([]byte(secret))
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
	invalidCookie = cookieName + "=" + invalidCookie

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
	a, err := auth.New(authSRV.URL, nil, nil, nil, []byte(secret), []byte(secret))
	if err != nil {
		t.Fatalf("Can not create auth service: %v", err)
	}

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
			"Invalid auth ticket",
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
			"Invalid auth ticket",
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
			"bearer NEWTOKEN",
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
	a, err := auth.New("", nil, nil, nil, []byte(""), []byte(""))
	if err != nil {
		t.Fatalf("Can not create auth serivce: %v", err)
	}

	t.Run("Empty Context", func(t *testing.T) {
		got := a.FromContext(context.Background())
		if got != 0 {
			t.Errorf("Got uid %d from empty context. Expected 0", got)
		}
	})

	t.Run("Context from Authenticate", func(t *testing.T) {
		validCookie, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sessionId": "123",
		}).SignedString([]byte(""))
		if err != nil {
			t.Fatalf("Can not sign cookie token: %v", err)
		}
		validCookie = "refreshId=" + validCookie

		validHeader, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId":    1,
			"sessionId": "123",
		}).SignedString([]byte(""))
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
		ctx, err := a.Authenticate(w, r)
		if err != nil {
			t.Fatalf("Can not create context from Authenticate: %v", err)
		}

		got := a.FromContext(ctx)
		if got != 1 {
			t.Errorf("Got uid %d from auth-context. Expected 1", got)
		}
	})
}

type mockAuth struct {
	token   string
	message string
}

func (m *mockAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Auth    string `json:"authentication"`
		Cookies string `json:"cookies"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, fmt.Sprintf("Can not decode body: %v", err), 500)
		return
	}

	p := struct {
		Token   string `json:"token"`
		Message string `json:"message"`
	}{
		m.token,
		m.message,
	}
	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, fmt.Sprintf("Can not encode data: %v", err), 500)
	}
}
