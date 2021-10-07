package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
)

func TestTagModeA(t *testing.T) {
	f := collection.Tag{}.Modes("A")

	testCase(
		"see meeting",
		t,
		f,
		true,
		`---
		tag/1/meeting_id: 5
		meeting/5:
			enable_anonymous: true
			committee_id: 300
		`,
	)

	testCase(
		"can not see meeting",
		t,
		f,
		false,
		`---
		tag/1/meeting_id: 5
		meeting/5:
			enable_anonymous: false
			committee_id: 300
		`,
	)
}
