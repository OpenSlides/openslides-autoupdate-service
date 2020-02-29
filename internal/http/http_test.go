package http_test

import (
	"bufio"
	"context"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	ahttp "github.com/openslides/openslides-autoupdate-service/internal/http"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

func TestHandlerTestURLs(t *testing.T) {
	var r test.MockRestricter
	keyschanges := test.NewMockKeysChanged()
	defer keyschanges.Close()
	s := autoupdate.New(r, keyschanges)
	srv := httptest.NewServer(ahttp.New(s, mockAuth{1}))
	defer srv.Close()

	tc := []struct {
		url    string
		status int
	}{
		{"", http.StatusNotFound},
		{"/system/autoupdate", http.StatusBadRequest},
		{"/system/autoupdate/keys", http.StatusOK},
	}

	for _, tt := range tc {
		t.Run(tt.url, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, srv.URL+tt.url, nil)
			if err != nil {
				t.Fatalf("Can not create request: %v", err)
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("Can not send request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.status {
				t.Errorf("Handler returned %s, expected %d, %s", resp.Status, tt.status, http.StatusText(tt.status))
			}
		})
	}
}

func TestSimple(t *testing.T) {
	var r test.MockRestricter
	keyschanges := test.NewMockKeysChanged()
	defer keyschanges.Close()
	s := autoupdate.New(r, keyschanges)
	srv := httptest.NewServer(ahttp.New(s, mockAuth{1}))
	defer srv.Close()

	tc := []struct {
		query  string
		keys   []string
		status int
	}{
		{"user/1/name", keys("user/1/name"), http.StatusOK},
		{"user/1/name,user/2/name", keys("user/1/name", "user/2/name"), http.StatusOK},
		{"key1,key2", keys("key1", "key2"), http.StatusOK},
	}

	for _, tt := range tc {
		t.Run("?"+tt.query, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, srv.URL+"/system/autoupdate/keys?"+tt.query, nil)
			if err != nil {
				t.Fatalf("Can not create request: %v", err)
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("Can not send request: %v", err)
			}
			defer resp.Body.Close()

			// Close connection
			cancel()

			if resp.StatusCode != tt.status {
				t.Errorf("Expected status %s, got %s", http.StatusText(tt.status), resp.Status)
			}

			s := bufio.NewScanner(resp.Body)
			keys := make([]string, 0, len(tt.keys))
			for s.Scan() {
				line := s.Text()
				keys = append(keys, line[:strings.Index(line, ": ")])
			}
			if !cmpSlice(tt.keys, keys) {
				t.Errorf("Returned keys are not equal, expected %v, got %v", tt.keys, keys)
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

func keys(ks ...string) []string {
	return ks
}

func cmpSlice(one, two []string) bool {
	if len(one) != len(two) {
		return false
	}
	sort.Strings(one)
	sort.Strings(two)
	for i := range one {
		if one[i] != two[i] {
			return false
		}
	}
	return true
}
