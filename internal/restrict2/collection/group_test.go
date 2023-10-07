package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
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
		meeting/30/committee_id: 40
		committee/40/id: 40
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
		group/1:
			meeting_id: 30
			meeting_user_ids: [50]
		meeting_user/50:
			user_id: 5
			group_ids: [1]
		user/5/meeting_user_ids: [50]
		meeting/30:
			group_ids: [1]
			committee_id: 2
		committee/2/id: 2
		`,
		withRequestUser(5),
	)
}
