package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
)

func TestThemeModeA(t *testing.T) {
	t.Parallel()
	f := collection.Theme{}.Modes("A")

	testCase(
		"no perm",
		t,
		f,
		true,
		"",
	)
}
