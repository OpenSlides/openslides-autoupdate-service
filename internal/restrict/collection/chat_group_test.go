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
		chat_group/1/meeting_id: 30
		meeting/30/id: 30
		`,
	)

	testCase(
		"Chat Manager in meeting",
		t,
		c.Modes("A"),
		true,
		`---
		chat_group/1/meeting_id: 30

		meeting/30/id: 30
		`,
		withPerms(30, perm.ChatCanManage),
	)

	testCase(
		"In chat read group",
		t,
		c.Modes("A"),
		true,
		`---
		chat_group/1:
			meeting_id: 30
			read_group_ids: [4]

		meeting/30/id: 1
		group/4/id: 4

		user/1/meeting_user_ids: [10]
		meeting_user/10/group_ids: [4]
		meeting_user/10/meeting_id: 30
		`,
	)

	testCase(
		"In chat write group",
		t,
		c.Modes("A"),
		true,
		`---
		chat_group/1:
			meeting_id: 30
			write_group_ids: [4]

		meeting/30/id: 1
		group/4/id: 4

		user/1/meeting_user_ids: [10]
		meeting_user/10/group_ids: [4]
		meeting_user/10/meeting_id: 30
		`,
	)
}
