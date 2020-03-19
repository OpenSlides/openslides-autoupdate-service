package restrict_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/restrict"
)

func TestRestrict(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(testSrv))

	s := &restrict.Service{
		Addr: srv.URL,
	}
	r, err := s.Restrict(context.Background(), 1, []string{"motion/1/name"})
	if err != nil {
		t.Errorf("did not expect an error, got: %v", err)
	}
	var data map[string]string
	json.NewDecoder(r).Decode(&data)

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

	out := []map[string]string{make(map[string]string)}
	for _, key := range reqData[0].Keys {
		out[0][key] = "restricted value for " + key
	}

	json.NewEncoder(w).Encode(out)
}
