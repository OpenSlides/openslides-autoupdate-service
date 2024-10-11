package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
)

func TestGenderModeA(t *testing.T) {
	var a collection.Gender

	testCase(
		"No permission",
		t,
		a.Modes("A"),
		true,
		`---
		gender/1/id: 30
		`,
	)
}
