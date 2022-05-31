package restrict

import (
	"strings"
	"testing"
)

func TestRestrictModeForAll(t *testing.T) {
	for field := range restrictionModes {
		parts := strings.Split(field, "/")

		fieldMode, err := buildFieldMode(parts[0], parts[1])
		if err != nil {
			t.Fatalf("building field mode: %v", err)
		}

		if _, err := restrictMode(parts[0], fieldMode, false); err != nil {
			t.Errorf("restrictMode(%s, %s) returned: %v", parts[0], parts[1], err)
		}
	}
}
