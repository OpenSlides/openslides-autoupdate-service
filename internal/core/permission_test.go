package core_test

import (
	"testing"

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
	result, addition, err := p.IsAllowed(nil, "", 0, nil)
	if err == nil || addition != nil || result == true {
		t.Errorf("Fail")
	}
}

func TestDispatchAllowed(t *testing.T) {
	core.Queries = Queries
	p := core.NewPermissionService(nil)
	isAllowed, addition, err := p.IsAllowed(nil, "dummy_allowed", 0, nil)
	if err != nil || addition != nil || !isAllowed {
		t.Errorf("Fail")
	}
}

func TestDispatchNotAllowed(t *testing.T) {
	core.Queries = Queries
	p := core.NewPermissionService(nil)
	isAllowed, addition, err := p.IsAllowed(nil, "dummy_not_allowed", 0, nil)
	if err == nil || addition != nil || isAllowed {
		t.Errorf("Fail")
	}
}
