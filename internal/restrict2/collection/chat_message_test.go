package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
)

func TestChatMessageModeA(t *testing.T) {
	t.Parallel()
	var c collection.ChatMessage

	testCase(
		"No permission",
		t,
		c.Modes("A"),
		false,
		`---
		chat_message/1:
			chat_group_id: 5
			meeting_user_id: 20
		meeting_user/20/user_id: 25
		chat_group/5:
			meeting_id: 30
			read_group_ids: []
			write_group_ids: []
		meeting/10/id: 10
		`,
	)

	testCase(
		"chat manager",
		t,
		c.Modes("A"),
		true,
		`---
		chat_message/1:
			chat_group_id: 5
			meeting_user_id: 404
		meeting_user/404/user_id: 44
		chat_group/5:
			meeting_id: 30
			read_group_ids: []
			write_group_ids: []
		
		meeting/30/id: 30
		`,
		withPerms(30, perm.ChatCanManage),
	)

	testCase(
		"read group",
		t,
		c.Modes("A"),
		true,
		`---
		chat_message/1:
			chat_group_id: 5
			meeting_user_id: 404
		meeting_user/404/user_id: 44
		chat_group/5:
			meeting_id: 30
			read_group_ids: [4]
		
		meeting/30/id: 30
		group/4/id: 4

		user/1/meeting_user_ids: [10]
		meeting_user/10/group_ids: [4]
		meeting_user/10/meeting_id: 30
		`,
	)

	testCase(
		"author",
		t,
		c.Modes("A"),
		true,
		`---
		chat_message/1:
			chat_group_id: 5
			meeting_user_id: 20
		meeting_user/20/user_id: 1

		chat_group/5:
			meeting_id: 30
		
		meeting/10/id: 10		
		`,
	)
}
