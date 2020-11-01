package http_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	permHTTP "github.com/OpenSlides/openslides-permission-service/internal/http"
)

func TestIsAllowed(t *testing.T) {
	mux := http.NewServeMux()
	allowed := new(IsAllowedMock)
	permHTTP.IsAllowed(mux, allowed)

	for _, tt := range []struct {
		name string

		reqBody  string
		allowed  bool
		addition map[string]interface{}
		err      error

		expectResponse    string
		expectStatuseCode int
	}{
		{
			name:    "Allowed",
			reqBody: `{"name": "everything", "user_id": 1}`,

			allowed: true,

			expectResponse:    `{"allowed":true}`,
			expectStatuseCode: 200,
		},

		{
			name:    "Not Allowed",
			reqBody: `{"name": "everything", "user_id": 1}`,

			allowed: false,

			expectResponse:    `{"allowed":false}`,
			expectStatuseCode: 200,
		},

		{
			name:    "Without addition",
			reqBody: `{"name": "everything", "user_id": 1}`,

			allowed:  true,
			addition: map[string]interface{}{"with_addition": 5},

			expectResponse:    `{"allowed":true,"addition":{"with_addition":5}}`,
			expectStatuseCode: 200,
		},

		{
			name:    "Internal Error",
			reqBody: `{"name": "everything", "user_id": 1}`,

			err: fmt.Errorf("something happend :("),

			expectResponse:    `{"error": {"type": "InternalError", "msg": "Ups, something went wrong!"}}`,
			expectStatuseCode: 500,
		},

		{
			name:    "Client Error",
			reqBody: `{"name": "everything", "user_id": 1}`,

			err: clientError{errType: "SomethingError", msg: "This explains why"},

			expectResponse:    `{"error": {"type": "SomethingError", "msg": "This explains why"}}`,
			expectStatuseCode: 400,
		},

		{
			name:    "Invalid JSON",
			reqBody: `{"name": "ever`,

			err: clientError{errType: "JSONError", msg: "Can not decode request body"},

			expectResponse:    `{"error": {"type": "SomethingError", "msg": "This explains why"}}`,
			expectStatuseCode: 400,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			allowed.allowed = tt.allowed
			allowed.addition = tt.addition
			allowed.err = tt.err

			req, err := http.NewRequest("POST", "/internal/permission/is_allowed", strings.NewReader(tt.reqBody))
			if err != nil {
				t.Fatalf("Creating request: %v", err)
			}

			resp := httptest.NewRecorder()
			mux.ServeHTTP(resp, req)

			if resp.Result().StatusCode != tt.expectStatuseCode {
				t.Errorf("Got status %s, expected %s", resp.Result().Status, http.StatusText(tt.expectStatuseCode))
			}
		})
	}
}

type IsAllowedMock struct {
	allowed  bool
	addition map[string]interface{}
	err      error
}

func (a *IsAllowedMock) IsAllowed(name string, userID int, data map[string]string) (bool, map[string]interface{}, error) {
	return a.allowed, a.addition, a.err
}

type clientError struct {
	errType string
	msg     string
}

func (e clientError) Error() string {
	return e.msg
}

func (e clientError) Type() string {
	return e.errType
}
