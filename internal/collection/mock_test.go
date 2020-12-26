package collection_test

import (
	"errors"
	"testing"

	"github.com/OpenSlides/openslides-permission-service/internal/perm"
)

func mustFQfields(fqfields ...string) []perm.FQField {
	out := make([]perm.FQField, len(fqfields))
	var err error
	for i, fqfield := range fqfields {
		out[i], err = perm.ParseFQField(fqfield)
		if err != nil {
			panic(err)
		}
	}
	return out
}

func checkRead(t *testing.T, r map[string]bool, allowed ...string) {
	for _, a := range allowed {
		if !r[a] {
			t.Errorf("fqfield %s not in allowed", a)
		}
	}

	if len(allowed) != len(r) {
		keys := make([]string, 0, len(r))
		for k := range r {
			keys = append(keys, k)
		}
		t.Errorf("got %v, expected %v", keys, allowed)
	}
}

func assertNotAllowed(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Got no error, expected not allowed")
		return
	}

	var errNotAllowed perm.NotAllowedError
	if !errors.As(err, &errNotAllowed) {
		t.Errorf("Got undexpected error: %v", err)
	}
}
