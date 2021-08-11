package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
)

func TestChatGroupModeA(t *testing.T) {
	var c collection.ChatGroup

	testCase(
		"No permission",
		t,
		c.Modes("A"),
		false,
		`---
		chat_group/1/meeting_id: 1
		meeting/1/id: 1
		`,
	)

	testCase(
		"Admin in meeting",
		t,
		c.Modes("A"),
		true,
		`---
		chat_group/1/meeting_id: 1

		meeting/1/admin_group_id: 3

		user/1/group_$1_ids: [3]
		`,
	)

	testCase(
		"In chat read group",
		t,
		c.Modes("A"),
		true,
		`---
		chat_group/1:
			meeting_id: 1
			read_group_ids: [4]

		meeting/1/id: 1
		group/4/id: 4

		user/1/group_$1_ids: [4]
		`,
	)
}
