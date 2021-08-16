package http_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/autoupdate"
	ahttp "github.com/OpenSlides/openslides-autoupdate-service/internal/http"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/test"
)

type liverMock struct {
	content io.Reader
}

func (m *liverMock) Live(ctx context.Context, uid int, w io.Writer, kb autoupdate.KeysBuilder) error {
	io.Copy(w, m.content)
	return nil
}

func TestSimpleHandler(t *testing.T) {
	mux := http.NewServeMux()
	liver := &liverMock{
		content: strings.NewReader("content"),
	}
	ahttp.Simple(mux, test.Auth(1), liver)

	req, _ := http.NewRequest("GET", "/system/autoupdate/keys?user/1/name,user/2/name", nil)
	req.ProtoMajor = 2
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	res := rec.Result()

	if res.StatusCode != 200 {
		t.Errorf("Got status %s, expected %s", res.Status, http.StatusText(200))
	}

	expect := "content"
	got, _ := io.ReadAll(res.Body)
	if string(got) != expect {
		t.Errorf("Got content `%s`, expected `%s`", got, expect)
	}
}

func TestComplexHandler(t *testing.T) {
	mux := http.NewServeMux()
	liver := &liverMock{
		content: strings.NewReader("content"),
	}
	ahttp.Complex(mux, test.Auth(1), new(test.DataProvider), liver)

	req, _ := http.NewRequest("GET", "/system/autoupdate", strings.NewReader(`[{"ids":[1],"collection":"user","fields":{"name":null}}]`))
	req.ProtoMajor = 2
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	res := rec.Result()

	if res.StatusCode != 200 {
		t.Errorf("Got status %s, expected %s", res.Status, http.StatusText(200))
	}

	expect := "content"
	got, _ := io.ReadAll(res.Body)
	if string(got) != expect {
		t.Errorf("Got `%s`, expected `%s`", got, expect)
	}
}

func TestHealth(t *testing.T) {
	mux := http.NewServeMux()
	ahttp.Health(mux)

	req := httptest.NewRequest("", "/system/autoupdate/health", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Result().StatusCode != 200 {
		t.Errorf("Got status %s, expected %s", rec.Result().Status, http.StatusText(200))
	}

	got, _ := io.ReadAll(rec.Body)
	expect := `{"healthy": true}` + "\n"
	if string(got) != expect {
		t.Errorf("Got %s, expected %s", got, expect)
	}
}

func TestErrors(t *testing.T) {
	mux := http.NewServeMux()
	liver := &liverMock{
		content: strings.NewReader(`"content"`),
	}
	db := &test.DataProvider{
		Data: map[string][]byte{
			"foo/1/name": []byte(`"hugo"`),
		},
	}
	ahttp.Complex(mux, test.Auth(1), db, liver)

	for _, tt := range []struct {
		name    string
		request *http.Request
		status  int
		errType string
		errMsg  string
	}{
		{
			"No Body",
			httptest.NewRequest(
				"",
				"/system/autoupdate",
				nil,
			),
			400,
			`SyntaxError`,
			`No data`,
		},
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
			`no collection`,
		},
		{
			"No list",
			httptest.NewRequest(
				"GET",
				"/system/autoupdate",
				strings.NewReader(`{"ids":[1],"collection":"foo","fields":{}}`),
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
				strings.NewReader(`[{"ids":["1"],"collection":"foo","fields":{}}]`),
			),
			400,
			`SyntaxError`,
			"wrong type at field `ids`. Got string, expected number",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			tt.request.ProtoMajor = 2
			req := httptest.NewRecorder()
			mux.ServeHTTP(req, tt.request)

			if req.Result().StatusCode != tt.status {
				t.Errorf("Got status %s, expected %s", req.Result().Status, http.StatusText(tt.status))
			}

			body, _ := io.ReadAll(req.Body)

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
