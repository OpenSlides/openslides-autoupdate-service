package restrict

import (
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

func TestRelationList(t *testing.T) {
	perm := new(test.MockPermission)
	perm.Data = map[string]bool{
		"foo/1": true,
		"foo/2": false,
	}
	r := relationList{
		perm:  perm,
		model: "foo",
	}

	v, err := r.Check(1, "bar/1/foo_ids", []byte("[1,2]"))

	if err != nil {
		t.Errorf("Check returned an error: %v", err)
	}

	if got := string(v); got != "[1]" {
		t.Errorf("Check returned `%s`, expected `[1]`", got)
	}

}
