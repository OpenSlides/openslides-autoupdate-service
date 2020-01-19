package http_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	ahttp "github.com/openslides/openslides-autoupdate-service/internal/http"
)

func TestHandlerTestURLs(t *testing.T) {
	s := &autoupdate.Service{}
	srv := httptest.NewServer(ahttp.NewHandler(s, mockAuth{1}))
	defer srv.Close()

	tc := []struct {
		url    string
		status int
	}{
		{"", http.StatusNotFound},
		{"/autoupdate/", http.StatusBadRequest},
	}

	for _, tt := range tc {
		t.Run(tt.url, func(t *testing.T) {
			resp, err := http.Get(srv.URL + tt.url)
			if err != nil {
				t.Fatalf("Can not send request to %s: %v", tt.url, err)
			}

			if resp.StatusCode != tt.status {
				t.Errorf("Handler returned %s, expected %d, %s", resp.Status, tt.status, http.StatusText(tt.status))
			}
		})
	}
}

type mockAuth struct {
	uid int
}

func (a mockAuth) Authenticate(context.Context, *http.Request) (int, error) {
	return a.uid, nil
}
