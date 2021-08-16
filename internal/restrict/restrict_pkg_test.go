package restrict

import (
	"strings"
	"testing"
)

func TestRestrictModeForAll(t *testing.T) {
	// TODO: unskip
	t.Skip()
	for field := range restrictionModes {
		parts := strings.Split(field, "/")

		_, err := restrictMode(parts[0], parts[1], false)

		if err != nil {
			t.Errorf("restrictMode(%s, %s) returned: %v", parts[0], parts[1], err)
		}
	}
}
