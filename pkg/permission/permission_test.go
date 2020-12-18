package permission_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/OpenSlides/openslides-permission-service/pkg/permission"
)

func TestDispatchNotFound(t *testing.T) {
	p := permission.New(nil, permission.WithCollections(fakeCollections()))
	_, err := p.IsAllowed(context.Background(), "", 0, nil)
	if err == nil {
		t.Errorf("Got no error, expected one")
	}
}

func TestDispatchAllowed(t *testing.T) {
	p := permission.New(nil, permission.WithCollections(fakeCollections()))
	additions, err := p.IsAllowed(context.Background(), "dummy_allowed", 0, []map[string]json.RawMessage{nil})
	if err != nil {
		t.Errorf("Got unexpected error: %v", err)
	}
	if additions == nil {
		t.Errorf("Got nil")
	}
}

func TestDispatchNotAllowed(t *testing.T) {
	p := permission.New(nil, permission.WithCollections(fakeCollections()))
	_, err := p.IsAllowed(context.Background(), "dummy_not_allowed", 0, []map[string]json.RawMessage{nil})
	var indexError interface {
		Index() int
	}
	if !errors.As(err, &indexError) {
		t.Fatalf("Got error `%v`, expected an index error", err)
	}
	if got := indexError.Index(); got != 0 {
		t.Errorf("Got index %d, expected 0", got)
	}
}

func TestDispatchEmptyDataAllowed(t *testing.T) {
	p := permission.New(nil, permission.WithCollections(fakeCollections()))
	additions, err := p.IsAllowed(context.Background(), "dummy_allowed", 0, []map[string]json.RawMessage{})
	if err != nil || len(additions) != 0 {
		t.Errorf("Fail")
	}
}

func TestDispatchEmptyDataNotAllowed(t *testing.T) {
	p := permission.New(nil, permission.WithCollections(fakeCollections()))
	additions, err := p.IsAllowed(context.Background(), "dummy_not_allowed", 0, []map[string]json.RawMessage{})
	if err != nil || len(additions) != 0 {
		t.Errorf("Fail")
	}
}
