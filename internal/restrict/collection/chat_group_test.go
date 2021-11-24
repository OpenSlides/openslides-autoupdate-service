package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
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
		"Chat Manager in meeting",
		t,
		c.Modes("A"),
		true,
		`---
		chat_group/1/meeting_id: 1

		meeting/1/id: 1
		`,
		withPerms(1, perm.ChatCanManage),
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
