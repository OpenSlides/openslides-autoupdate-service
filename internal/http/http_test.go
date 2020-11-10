package http_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	permHTTP "github.com/OpenSlides/openslides-permission-service/internal/http"
)

func TestHttpIsAllowed(t *testing.T) {
	mux := http.NewServeMux()
	allowed := new(IsAllowedMock)
	permHTTP.IsAllowed(mux, allowed)

	for _, tt := range []struct {
		name string

		reqBody  string
		addition [](map[string]interface{})
		err      error

		expectResponse    string
		expectStatuseCode int
	}{
		{
			name:    "Allowed",
			reqBody: `{"name": "everything", "user_id": 1}`,

			expectResponse:    `{"allowed":true,"additions":null}`,
			expectStatuseCode: 200,
		},
		{
			name:    "Not Allowed with reason",
			reqBody: `{"name": "everything", "user_id": 1}`,

			err: clientError{errType: "ClientError", msg: "This explains why"},

			expectResponse:    `{"allowed":false,"reason":"This explains why","error_index":0}`,
			expectStatuseCode: 200,
		},
		{
			name:    "With addition",
			reqBody: `{"name": "everything", "user_id": 1}`,

			addition: [](map[string]interface{}){map[string]interface{}{"with_addition": 5}},

			expectResponse:    `{"allowed":true,"additions":[{"with_addition":5}]}`,
			expectStatuseCode: 200,
		},
		{
			name:    "Internal Error",
			reqBody: `{"name": "everything", "user_id": 1}`,

			err: fmt.Errorf("something happend :("),

			expectResponse:    `{"error":{"type":"InternalError","msg":"Ups, something went wrong!"}}`,
			expectStatuseCode: 500,
		},
		{
			name:    "Custom Error",
			reqBody: `{"name": "everything", "user_id": 1}`,

			err: clientError{errType: "SomethingError", msg: "This explains why"},

			expectResponse:    `{"error":{"type":"SomethingError","msg":"calling IsAllowed: This explains why"}}`,
			expectStatuseCode: 400,
		},
		{
			name:    "Invalid JSON",
			reqBody: `{"name": "ever`,

			err: clientError{errType: "JSONError", msg: "Can not decode request body"},

			expectResponse:    `{"error":{"type":"JSONError","msg":"Can not decode request body '{\"name\": \"ever': unexpected end of JSON input"}}`,
			expectStatuseCode: 400,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
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

			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Cannot read response: %v", err)
			}
			body := strings.TrimSpace(string(bodyBytes))
			if body != tt.expectResponse {
				t.Errorf("Got '%s', expected '%s'", body, tt.expectResponse)
			}
		})
	}
}

type IsAllowedMock struct {
	addition [](map[string]interface{})
	err      error
	index    int
}

func (a *IsAllowedMock) IsAllowed(ctx context.Context, name string, userID int, data [](map[string]json.RawMessage)) ([](map[string]interface{}), error, int) {
	return a.addition, a.err, a.index
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
