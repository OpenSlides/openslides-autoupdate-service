package core_test

import (
	"context"
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
	additions, err, index := p.IsAllowed(context.Background(), "", 0, nil)
	if additions != nil || err == nil || index != -1 {
		t.Errorf("Fail")
	}
}

func TestDispatchAllowed(t *testing.T) {
	core.Queries = Queries
	p := core.NewPermissionService(nil)
	additions, err, index := p.IsAllowed(context.Background(), "dummy_allowed", 0, []definitions.FqfieldData{nil})
	if err != nil || additions == nil || index != -1 {
		t.Errorf("Fail")
	}
}

func TestDispatchNotAllowed(t *testing.T) {
	core.Queries = Queries
	p := core.NewPermissionService(nil)
	additions, err, index := p.IsAllowed(context.Background(), "dummy_not_allowed", 0, []definitions.FqfieldData{nil})
	if err == nil || additions != nil || index != 0 {
		t.Errorf("Fail")
	}
}

func TestDispatchEmptyDataAllowed(t *testing.T) {
	core.Queries = Queries
	p := core.NewPermissionService(nil)
	additions, err, index := p.IsAllowed(context.Background(), "dummy_allowed", 0, []definitions.FqfieldData{})
	if err != nil || len(additions) != 0 || index != -1 {
		t.Errorf("Fail")
	}
}

func TestDispatchEmptyDataNotAllowed(t *testing.T) {
	core.Queries = Queries
	p := core.NewPermissionService(nil)
	additions, err, index := p.IsAllowed(context.Background(), "dummy_not_allowed", 0, []definitions.FqfieldData{})
	if err != nil || len(additions) != 0 || index != -1 {
		t.Errorf("Fail")
	}
}
