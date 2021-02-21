package http_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
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

		reqBody string
		allowed bool
		err     error

		expectResponse    string
		expectStatuseCode int
	}{
		{
			name:    "Allowed",
			allowed: true,
			reqBody: `{"name": "everything", "user_id": 1}`,

			expectResponse:    `true`,
			expectStatuseCode: 200,
		},
		{
			name:    "Not Allowed",
			reqBody: `{"name": "everything", "user_id": 1}`,

			expectResponse:    `false`,
			expectStatuseCode: 200,
		},
		{
			name:    "Internal Error",
			reqBody: `{"name": "everything", "user_id": 1}`,

			err: fmt.Errorf("something happend :("),

			expectResponse:    `"Internal Error. Norman, Do not sent it to client: something happend :("`,
			expectStatuseCode: 500,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			allowed.allowed = tt.allowed
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

			bodyBytes, err := io.ReadAll(resp.Body)
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
	allowed bool
	err     error
}

func (a *IsAllowedMock) IsAllowed(ctx context.Context, name string, userID int, data [](map[string]json.RawMessage)) (bool, error) {
	return a.allowed, a.err
}
