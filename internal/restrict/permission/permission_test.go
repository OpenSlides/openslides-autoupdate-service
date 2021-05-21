package permission

import (
	"context"
	"encoding/json"
	"strings"
	"testing"
)

// TODO Activate after developer-mode is removed.
// func TestDispatchNotFound(t *testing.T) {
// 	p := permission.New(nil, permission.WithCollections(fakeCollections()))
// 	_, err := p.IsAllowed(context.Background(), "", 0, nil)
// 	if err == nil {
// 		t.Errorf("Got no error, expected one")
// 	}
// }

func TestDispatchAllowed(t *testing.T) {
	p := NewTestPermission()
	allowed, err := p.IsAllowed(context.Background(), "dummy_allowed", 0, []map[string]json.RawMessage{nil})
	if err != nil {
		t.Fatalf("Got unexpected error: %v", err)
	}
	if allowed == false {
		t.Errorf("Got false, expected true")
	}
}

func TestDispatchNotAllowed(t *testing.T) {
	p := NewTestPermission()
	allowed, err := p.IsAllowed(context.Background(), "dummy_not_allowed", 0, []map[string]json.RawMessage{nil})
	if err != nil {
		t.Fatalf("Got unexpected error: %v", err)
	}
	if allowed == true {
		t.Errorf("Got true, expected false")
	}
}

func TestErrorMessage(t *testing.T) {
	p := NewTestPermission()
	_, err := p.IsAllowed(context.Background(), "dummy_error", 0, []map[string]json.RawMessage{{"id": []byte("1")}})
	if err == nil {
		t.Fatalf("Got no error.")
	}

	// Check action name.
	if !strings.Contains(err.Error(), "dummy_error") {
		t.Errorf("Error does not contain action name: %v", err)
	}

	// Check payload.
	if !strings.Contains(err.Error(), `{"id":1}`) {
		t.Errorf("Error does not contain payload: %v", err)
	}

	// Check original error.
	if !strings.Contains(err.Error(), `original error message`) {
		t.Errorf("Error does not contain original error: %v", err)
	}
}
