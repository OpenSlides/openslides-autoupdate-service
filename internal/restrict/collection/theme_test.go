package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
)

func TestThemeModeA(t *testing.T) {
	f := collection.Theme{}.Modes("A")

	testCase(
		"no perm",
		t,
		f,
		true,
		"",
	)
}
