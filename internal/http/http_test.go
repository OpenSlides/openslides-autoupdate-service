package http_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"iter"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/autoupdate"
	ahttp "github.com/OpenSlides/openslides-autoupdate-service/internal/http"
	"github.com/OpenSlides/openslides-go/datastore/dskey"
)

var (
	myKey1, _ = dskey.FromParts("user", 1, "username")
)

type connectionMock struct {
	f func(ctx context.Context) (map[dskey.Key][]byte, error)
}

func (n *connectionMock) Messages(ctx context.Context) iter.Seq2[map[dskey.Key][]byte, error] {
	return func(yield func(map[dskey.Key][]byte, error) bool) {
		data, err := n.f(ctx)
		if !yield(data, err) {
			return
		}
	}
}

func (n *connectionMock) NextWithFilter(context.Context, string) (map[dskey.Key][]byte, string, error) {
	return nil, "", fmt.Errorf("Not Implemented")
}

type connecterMock struct {
	f func(ctx context.Context) (map[dskey.Key][]byte, error)
}

func (c *connecterMock) Connect(ctx context.Context, userID int, kb autoupdate.KeysBuilder) (autoupdate.Connection, error) {
	return &connectionMock{f: c.f}, nil
}

func (c *connecterMock) SingleData(ctx context.Context, userID int, kb autoupdate.KeysBuilder) (map[dskey.Key][]byte, error) {
	return c.f(ctx)
}

func TestKeysHandler(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()

	mux := http.NewServeMux()

	connecter := &connecterMock{
		f: func(ctx context.Context) (map[dskey.Key][]byte, error) {
			if ctx.Err() != nil {
				return nil, ctx.Err()
			}
			cancel()
			return map[dskey.Key][]byte{myKey1: []byte(`"bar"`)}, nil
		},
	}

	ahttp.HandleAutoupdate(mux, fakeAuth(1), connecter, [2]*ahttp.ConnectionCount{}, time.Hour)

	req := httptest.NewRequest("GET", "/system/autoupdate?k=user/1/username,user/2/username", nil).WithContext(ctx)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	res := rec.Result()

	if res.StatusCode != 200 {
		t.Errorf("Got status %q, expected %q", res.Status, http.StatusText(200))
	}

	expect := `{"user/1/username":"bar"}` + "\n"
	got, _ := io.ReadAll(res.Body)
	if string(got) != expect {
		t.Errorf("Got content `%s`, expected `%s`", got, expect)
	}
}

func TestComplexHandler(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()

	mux := http.NewServeMux()

	connecter := &connecterMock{
		f: func(ctx context.Context) (map[dskey.Key][]byte, error) {
			if ctx.Err() != nil {
				return nil, ctx.Err()
			}
			cancel()
			return map[dskey.Key][]byte{myKey1: []byte(`"bar"`)}, nil
		},
	}

	ahttp.HandleAutoupdate(mux, fakeAuth(1), connecter, [2]*ahttp.ConnectionCount{}, time.Hour)

	req := httptest.NewRequest(
		"GET",
		"/system/autoupdate",
		strings.NewReader(`[{"ids":[1],"collection":"user","fields":{"username":null}}]`),
	).WithContext(ctx)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	res := rec.Result()

	if res.StatusCode != 200 {
		t.Errorf("Got status %s, expected %s", res.Status, http.StatusText(200))
	}

	expect := `{"user/1/username":"bar"}` + "\n"
	got, _ := io.ReadAll(res.Body)
	if string(got) != expect {
		t.Errorf("Got %s, expected %s", got, expect)
	}
}

func TestHealth(t *testing.T) {
	mux := http.NewServeMux()
	ahttp.HandleHealth(mux)

	req := httptest.NewRequest("", "/system/autoupdate/health", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Result().StatusCode != 200 {
		t.Errorf("Got status %s, expected %s", rec.Result().Status, http.StatusText(200))
	}

	got, _ := io.ReadAll(rec.Body)
	expect := `{"healthy": true}` + "\n"
	if string(got) != expect {
		t.Errorf("Got %q, expected %q", got, expect)
	}
}

func TestErrors(t *testing.T) {
	mux := http.NewServeMux()

	connecter := &connecterMock{
		f: func(ctx context.Context) (map[dskey.Key][]byte, error) {
			if ctx.Err() != nil {
				return nil, ctx.Err()
			}
			return map[dskey.Key][]byte{myKey1: []byte(`"bar"`)}, nil
		},
	}

	ahttp.HandleAutoupdate(mux, fakeAuth(1), connecter, [2]*ahttp.ConnectionCount{}, time.Hour)

	for _, tt := range []struct {
		name    string
		request *http.Request
		status  int
		errType string
		errMsg  string
	}{
		{
			"Empty List",
			httptest.NewRequest(
				"",
				"/system/autoupdate",
				strings.NewReader("[]"),
			),
			400,
			`SyntaxError`,
			`No data`,
		},
		{
			"Invalid json",
			httptest.NewRequest(
				"GET",
				"/system/autoupdate",
				strings.NewReader("{5"),
			),
			400,
			`JsonError`,
			`invalid character '5' looking for beginning of object key string`,
		},
		{
			"Invalid KeyRequest",
			httptest.NewRequest(
				"GET",
				"/system/autoupdate",
				strings.NewReader(`[{"ids":[123]}]`),
			),
			400,
			`SyntaxError`,
			`attribute collection is missing`,
		},
		{
			"No list",
			httptest.NewRequest(
				"GET",
				"/system/autoupdate",
				strings.NewReader(`{"ids":[1],"collection":"user","fields":{}}`),
			),
			400,
			`SyntaxError`,
			"wrong type at field ``. Got object, expected list",
		},
		{
			"String ID",
			httptest.NewRequest(
				"GET",
				"/system/autoupdate",
				strings.NewReader(`[{"ids":["1"],"collection":"user","fields":{}}]`),
			),
			400,
			`SyntaxError`,
			"wrong type at field `ids`. Got string, expected number",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			resp := httptest.NewRecorder()

			mux.ServeHTTP(resp, tt.request)

			if resp.Result().StatusCode != tt.status {
				t.Errorf("Got status %s, expected %s", resp.Result().Status, http.StatusText(tt.status))
			}

			body, _ := io.ReadAll(resp.Body)

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

// fakeAuth implements the http.Authenticater interface. It allways returs the given
// user id.
type fakeAuth int

// Authenticate does nothing.
func (a fakeAuth) Authenticate(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	return r.Context(), nil
}

// FromContext returns the uid the object was initialiced with.
func (a fakeAuth) FromContext(ctx context.Context) int {
	return int(a)
}

func (a fakeAuth) AuthenticatedContext(ctx context.Context, _ int) context.Context {
	return ctx
}
