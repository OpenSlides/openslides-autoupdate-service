package http_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	ahttp "github.com/openslides/openslides-autoupdate-service/internal/http"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

func TestHandlerTestURLs(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	s := autoupdate.New(new(test.MockDatastore), new(test.MockRestricter), closed)
	srv := httptest.NewUnstartedServer(ahttp.New(s, test.Auth(1)))
	srv.EnableHTTP2 = true
	srv.StartTLS()
	defer srv.Close()

	tc := []struct {
		url    string
		status int
	}{
		{"", http.StatusNotFound},
		{"/system/autoupdate", http.StatusBadRequest},
		{"/system/autoupdate/keys?user/1/name", http.StatusOK},
		{"/system/autoupdate/health", http.StatusOK},
	}

	for _, tt := range tc {
		t.Run(tt.url, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, srv.URL+tt.url, nil)
			if err != nil {
				t.Fatalf("Can not create request: %v", err)
			}
			resp, err := srv.Client().Do(req)
			if err != nil {
				t.Fatalf("Can not send request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.status {
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					body = nil
				}
				body = bytes.TrimSpace(body)

				t.Errorf("Handler returned %s: %s, expected %d, %s", resp.Status, body, tt.status, http.StatusText(tt.status))
			}
		})
	}
}

func TestSimple(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	s := autoupdate.New(new(test.MockDatastore), new(test.MockRestricter), closed)
	srv := httptest.NewUnstartedServer(ahttp.New(s, test.Auth(1)))
	srv.EnableHTTP2 = true
	srv.StartTLS()
	defer srv.Close()

	tc := []struct {
		query  string
		keys   []string
		status int
		errMsg string
	}{
		{"user/1/name", keys("user/1/name"), http.StatusOK, ""},
		{"user/1/name,user/2/name", keys("user/1/name", "user/2/name"), http.StatusOK, ""},
		{"key1,key2", keys("key1", "key2"), http.StatusBadRequest, "Invalid keys"},
	}

	for _, tt := range tc {
		t.Run("?"+tt.query, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, srv.URL+"/system/autoupdate/keys?"+tt.query, nil)
			if err != nil {
				t.Fatalf("Can not create request: %v", err)
			}
			resp, err := srv.Client().Do(req)
			if err != nil {
				t.Fatalf("Can not send request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.status {
				t.Errorf("Expected status %s, got %s", http.StatusText(tt.status), resp.Status)
			}

			expected := "application/octet-stream"
			if got := resp.Header.Get("Content-Type"); got != expected {
				t.Errorf("Got content-type %s, expected: %s", got, expected)
			}

			if tt.errMsg != "" {
				var body map[string]map[string]string
				if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
					t.Errorf("Got invalid json: %v", err)
				}

				if v := body["error"]["msg"]; v != tt.errMsg {
					t.Errorf("Got error message `%s`, expected `%s`", v, tt.errMsg)
				}
				return
			}

			var body map[string]json.RawMessage
			if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
				t.Errorf("Got invalid json: %v", err)
			}

			if v, ok := body["error"]; ok {
				t.Errorf("Error: %v", v)
			}

			got := make([]string, 0, len(body))
			for key := range body {
				got = append(got, key)
			}

			if !cmpSlice(got, tt.keys) {
				t.Errorf("Got keys %v, expected %v", got, tt.keys)
			}
		})
	}
}

func TestErrors(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	s := autoupdate.New(new(test.MockDatastore), new(test.MockRestricter), closed)
	srv := httptest.NewUnstartedServer(ahttp.New(s, test.Auth(1)))
	srv.EnableHTTP2 = true
	srv.StartTLS()
	defer srv.Close()

	for _, tt := range []struct {
		name    string
		request *http.Request
		status  int
		errType string
		errMsg  string
	}{
		{
			"No Body",
			mustRequest(http.NewRequest(
				"GET",
				srv.URL+"/system/autoupdate",
				nil,
			)),
			400,
			`SyntaxError`,
			`No data`,
		},
		{
			"Empty List",
			mustRequest(http.NewRequest(
				"GET",
				srv.URL+"/system/autoupdate",
				strings.NewReader("[]"),
			)),
			400,
			`SyntaxError`,
			`No data`,
		},
		{
			"Invalid json",
			mustRequest(http.NewRequest(
				"GET",
				srv.URL+"/system/autoupdate",
				strings.NewReader("{5"),
			)),
			400,
			`JsonError`,
			`invalid character '5' looking for beginning of object key string`,
		},
		{
			"Invalid KeyRequest",
			mustRequest(http.NewRequest(
				"GET",
				srv.URL+"/system/autoupdate",
				strings.NewReader(`[{"ids":[123]}]`),
			)),
			400,
			`SyntaxError`,
			`no collection`,
		},
		{
			"No list",
			mustRequest(http.NewRequest(
				"GET",
				srv.URL+"/system/autoupdate",
				strings.NewReader(`{"ids":[1],"collection":"foo","fields":{}}`),
			)),
			400,
			`SyntaxError`,
			"wrong type at field ``. Got object, expected list",
		},
		{
			"String ID",
			mustRequest(http.NewRequest(
				"GET",
				srv.URL+"/system/autoupdate",
				strings.NewReader(`[{"ids":["1"],"collection":"foo","fields":{}}]`),
			)),
			400,
			`SyntaxError`,
			"wrong type at field `ids`. Got string, expected number",
		},
		{
			"Wrong field value",
			mustRequest(http.NewRequest(
				"GET",
				srv.URL+"/system/autoupdate",
				strings.NewReader(`
				[{
					"ids": [1],
					"collection": "foo",
					"fields": {
						"name": {
							"type": "relation",
							"collection": "bar",
							"fields": {}
						}
					}
				}]`),
			)),
			400,
			`ValueError`,
			`invalid value in key foo/1/name. Got string, expected int`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			resp, err := srv.Client().Do(tt.request.WithContext(ctx))
			if err != nil {
				t.Fatalf("Can not send request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.status {
				t.Errorf("Expected status %d %s, got %s", tt.status, http.StatusText(tt.status), resp.Status)
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Can not read body: %v", err)
			}

			var data struct {
				Error struct {
					Type string `json:"type"`
					Msg  string `json:"msg"`
				} `json:"error"`
			}
			if err := json.Unmarshal(body, &data); err != nil {
				t.Fatalf("Can not decode body `%s`: %v", body, err)
			}

			if data.Error.Type != tt.errType {
				t.Errorf("Got error type %s, expected %s", data.Error.Type, tt.errType)
			}

			if data.Error.Msg != tt.errMsg {
				t.Errorf("Got error message `%s`, expected %s", data.Error.Msg, tt.errMsg)
			}
		})
	}
}
