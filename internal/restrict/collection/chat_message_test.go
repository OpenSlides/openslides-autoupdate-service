package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

func TestChatMessageModeA(t *testing.T) {
	var c collection.ChatMessage

	testCase(
		"No permission",
		t,
		c.Modes("A"),
		false,
		`---
		chat_message/1/chat_group_id: 5
		chat_group/5/meeting_id: 10
		meeting/10/id: 10
		`,
	)

	testCase(
		"chat manager",
		t,
		c.Modes("A"),
		true,
		`---
		chat_message/1/chat_group_id: 5
		chat_group/5:
			meeting_id: 10
		
		meeting/10/id: 10
		`,
		withPerms(10, perm.ChatCanManage),
	)

	testCase(
		"read group",
		t,
		c.Modes("A"),
		true,
		`---
		chat_message/1/chat_group_id: 5
		chat_group/5:
			meeting_id: 10
			read_group_ids: [4]
		
		meeting/10/id: 10
		user/1/group_$10_ids: [4]
		group/4/id: 4
		`,
	)

	testCase(
		"author",
		t,
		c.Modes("A"),
		true,
		`---
		chat_message/1:
			user_id: 1
			chat_group_id: 5

		chat_group/5:
			meeting_id: 10
		
		meeting/10/id: 10		
		`,
	)
}
