package permission_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/OpenSlides/openslides-permission-service/pkg/permission"
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
	p := permission.New(nil, permission.WithCollections(fakeCollections()))
	allowed, err := p.IsAllowed(context.Background(), "dummy_allowed", 0, []map[string]json.RawMessage{nil})
	if err != nil {
		t.Fatalf("Got unexpected error: %v", err)
	}
	if allowed == false {
		t.Errorf("Got false, expected true")
	}
}

func TestDispatchNotAllowed(t *testing.T) {
	p := permission.New(nil, permission.WithCollections(fakeCollections()))
	allowed, err := p.IsAllowed(context.Background(), "dummy_not_allowed", 0, []map[string]json.RawMessage{nil})
	if err != nil {
		t.Fatalf("Got unexpected error: %v", err)
	}
	if allowed == true {
		t.Errorf("Got true, expected false")
	}
}
