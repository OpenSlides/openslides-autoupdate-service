package restrict_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/restrict"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

func TestRestrict(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(testSrv))

	s := &restrict.Service{
		Addr: srv.URL,
	}
	r, err := s.Restrict(context.Background(), 1, []string{"motion/1/name"})
	if err != nil {
		t.Fatalf("did not expect an error, got: %v", err)
	}

	var data map[string]json.RawMessage
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		t.Fatalf("can not decode restricted data: %v", err)
	}

	if len(data) != 1 {
		t.Errorf("expect data to have one value, got: %d", len(data))
	}
}

func testSrv(w http.ResponseWriter, r *http.Request) {
	var reqData []struct {
		UID  int      `json:"user_id"`
		Keys []string `json:"fqfields"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Invalid json: %s", err)
		return
	}

	if len(reqData) == 0 {
		w.WriteHeader(400)
		fmt.Fprintf(w, "no data in request")
		return
	}
	mr := new(test.MockRestricter)
	reader, err := mr.Restrict(r.Context(), reqData[0].UID, reqData[0].Keys)
	if err != nil {
		http.Error(w, "can not get restricted data", http.StatusInternalServerError)
		return
	}
	readers := io.MultiReader(strings.NewReader("["), reader, strings.NewReader("]"))
	if _, err := io.Copy(w, readers); err != nil {
		http.Error(w, "copy data from restricter to response", http.StatusInternalServerError)
	}
}
