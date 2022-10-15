package restrict

import (
	"strings"
	"testing"
)

func TestRestrictModeForAll(t *testing.T) {
	for collectionField := range restrictionModes {
		collection, field, found := strings.Cut(collectionField, "/")

		if !found {
			t.Fatalf("invalid field %s, expected one /", collectionField)
		}

		fieldMode, err := restrictModeName(collection, field)
		if err != nil {
			t.Fatalf("building field mode: %v", err)
		}

		if _, err := restrictModefunc(collection, fieldMode); err != nil {
			t.Errorf("restrictMode(%s, %s) returned: %v", collection, field, err)
		}
	}
}
