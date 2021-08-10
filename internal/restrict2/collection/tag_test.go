package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
)

func TestTagModeA(t *testing.T) {
	var tg collection.Tag

	testCase(
		"see meeting",
		true,
		`---
		tag/1/meeting_id: 5
		meeting/5/enable_anonymous: true
		`,
	).test(t, tg.Modes("A"))

	testCase(
		"can not see meeting",
		false,
		`---
		tag/1/meeting_id: 5
		meeting/5/enable_anonymous: false
		`,
	).test(t, tg.Modes("A"))
}
