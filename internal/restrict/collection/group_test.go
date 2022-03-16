package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
)

func TestGroupModeA(t *testing.T) {
	var g collection.Group

	testCase(
		"no perms",
		t,
		g.Modes("A"),
		false,
		`---
		group/1/meeting_id: 30
		meeting/30/id: 30
		meeting/30/committee_id: 404
		`,
	)

	testCase(
		"anonymous enabled",
		t,
		g.Modes("A"),
		true,
		`---
		group/1/meeting_id: 30
		meeting/30/enable_anonymous: true
		`,
	)

	testCase(
		"In meeting",
		t,
		g.Modes("A"),
		true,
		`---
		group/1/meeting_id: 30
		meeting/30/user_ids: [1]
		`,
	)
}
