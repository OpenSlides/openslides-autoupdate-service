package allowed

import (
	"errors"
	"testing"

	"github.com/OpenSlides/openslides-permission-service/internal/tests"
)

func AddBasicModel(collection string, dp *tests.TestDataProvider) {
	dp.Set(collection+"/1/id", "1")
	dp.Set(collection+"/1/meeting_id", "1")
	dp.Set("meeting/1/"+collection+"_ids", "[1]")
}

func AssertIsNotAllowed(t *testing.T, isAllowed IsAllowed, params *IsAllowedParams) {
	addition, err := isAllowed(params)
	if nil != addition {
		t.Errorf("Expected to fail without an addition: %s", addition)
	}

	if nil == err {
		t.Errorf("Expected to fail (reason must be set).")
	} else {
		var clientError interface {
			Type() string
		}
		if !errors.As(err, &clientError) || clientError.Type() != "ClientError" {
			t.Errorf("Expected to fail with a client error, not %v", err)
		}
	}
}

func AssertIsAllowed(t *testing.T, isAllowed IsAllowed, params *IsAllowedParams) {
	addition, err := isAllowed(params)
	if nil != addition {
		t.Errorf("Expected to fail without an addition: %s", addition)
	}
	if nil != err {
		t.Errorf("Expected to fail without an error (error: %s)", err)
	}
}
