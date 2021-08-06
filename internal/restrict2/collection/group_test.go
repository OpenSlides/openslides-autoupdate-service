package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
)

func TestGroupModeA(t *testing.T) {
	var g collection.Group

	testCase(
		"no perms",
		false,
		`---
		group/1/meeting_id: 1
		meeting/1/id: 1
		`,
	).test(t, g.Modes("A"))

	testCase(
		"anonymous enabled",
		true,
		`---
		group/1/meeting_id: 1
		meeting/1/enable_anonymous: true
		`,
	).test(t, g.Modes("A"))

	testCase(
		"In meeting",
		true,
		`---
		group/1/meeting_id: 1
		meeting/1/user_ids: [1]
		`,
	).test(t, g.Modes("A"))
}
