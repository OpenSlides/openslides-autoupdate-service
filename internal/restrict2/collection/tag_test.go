package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
)

func TestTagModeA(t *testing.T) {
	var tg collection.Tag

	testCase(
		"see meeting",
		t,
		tg.Modes("A"),
		true,
		`---
		tag/1/meeting_id: 5
		meeting/5/enable_anonymous: true
		`,
	)

	testCase(
		"can not see meeting",
		t,
		tg.Modes("A"),
		false,
		`---
		tag/1/meeting_id: 5
		meeting/5/enable_anonymous: false
		`,
	)
}
