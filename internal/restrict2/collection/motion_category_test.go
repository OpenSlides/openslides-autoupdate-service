package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
)

func TestMotionCategoryModeA(t *testing.T) {
	f := collection.MotionCategory{}.Modes("A")

	testCase(
		"no perms",
		t,
		f,
		false,
		"motion_category/1/id: 1",
	)
}
