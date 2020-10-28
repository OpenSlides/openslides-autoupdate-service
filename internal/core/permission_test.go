package core_test

import (
	"testing"

	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
	"github.com/OpenSlides/openslides-permission-service/internal/core"
)

func DummyAllowed(params *allowed.IsAllowedParams) (bool, map[string]interface{}, error) {
	return true, nil, nil
}

var Queries = map[string]allowed.IsAllowed{
	"dummy": DummyAllowed,
}

func TestDispatchNotFound(t *testing.T) {
	core.Queries = Queries
	p := core.NewPermissionService(nil)
	result, addition, err := p.IsAllowed("", 0, nil)
	if err == nil || addition != nil || result == true {
		t.Errorf("Fail")
	}
}

func TestDispatch(t *testing.T) {
	core.Queries = Queries
	p := core.NewPermissionService(nil)
	result, addition, err := p.IsAllowed("dummy", 0, nil)
	if err != nil || addition != nil || result != true {
		t.Errorf("Fail")
	}
}
