package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
)

func TestResourceModeA(t *testing.T) {
	f := collection.Resource{}.Modes("A")

	testCase(
		"no perms",
		t,
		f,
		true,
		``,
	)
}
