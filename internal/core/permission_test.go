package core_test

import (
	"context"
	"errors"
	"testing"

	"github.com/OpenSlides/openslides-permission-service/internal/definitions"

	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
	"github.com/OpenSlides/openslides-permission-service/internal/core"
)

func DummyAllowed(params *allowed.IsAllowedParams) (map[string]interface{}, error) {
	return nil, nil
}

func DummyNotAllowed(params *allowed.IsAllowedParams) (map[string]interface{}, error) {
	return nil, allowed.NotAllowed("Some reason here")
}

var Queries = map[string]allowed.IsAllowed{
	"dummy_allowed":     DummyAllowed,
	"dummy_not_allowed": DummyNotAllowed,
}

func TestDispatchNotFound(t *testing.T) {
	core.Queries = Queries
	p := core.NewPermissionService(nil)
	_, err := p.IsAllowed(context.Background(), "", 0, nil)
	if err == nil {
		t.Errorf("Got no error, expected one")
	}
}

func TestDispatchAllowed(t *testing.T) {
	core.Queries = Queries
	p := core.NewPermissionService(nil)
	additions, err := p.IsAllowed(context.Background(), "dummy_allowed", 0, []definitions.FqfieldData{nil})
	if err != nil {
		t.Errorf("Got unexpected error: %v", err)
	}
	if additions == nil {
		t.Errorf("Got nil")
	}
}

func TestDispatchNotAllowed(t *testing.T) {
	core.Queries = Queries
	p := core.NewPermissionService(nil)
	_, err := p.IsAllowed(context.Background(), "dummy_not_allowed", 0, []definitions.FqfieldData{nil})
	var indexError interface {
		Index() int
	}
	if !errors.As(err, &indexError) {
		t.Errorf("Got error `%v`, expected an index error", err)
	}
	if got := indexError.Index(); got != 0 {
		t.Errorf("Got index %d, expected 0", got)
	}
}

func TestDispatchEmptyDataAllowed(t *testing.T) {
	core.Queries = Queries
	p := core.NewPermissionService(nil)
	additions, err := p.IsAllowed(context.Background(), "dummy_allowed", 0, []definitions.FqfieldData{})
	if err != nil || len(additions) != 0 {
		t.Errorf("Fail")
	}
}

func TestDispatchEmptyDataNotAllowed(t *testing.T) {
	core.Queries = Queries
	p := core.NewPermissionService(nil)
	additions, err := p.IsAllowed(context.Background(), "dummy_not_allowed", 0, []definitions.FqfieldData{})
	if err != nil || len(additions) != 0 {
		t.Errorf("Fail")
	}
}
