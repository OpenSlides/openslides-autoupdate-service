package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

func TestVoteModeA(t *testing.T) {
	f := collection.Vote{}.Modes("A")

	testCase(
		"no perms",
		t,
		f,
		false,
		`---
		vote/1:
			option_id: 2
			user_id: 5
			delegated_user_id: 6
		option/2/poll_id: 3
		poll/3/meeting_id: 1
		meeting/1/enable_anonymous: false
		`,
	)

	testCase(
		"poll is published",
		t,
		f,
		true,
		`---
		vote/1:
			option_id: 2
			user_id: 5
			delegated_user_id: 6
		option/2/poll_id: 3
		poll/3:
			meeting_id: 1
			state: published
		`,
	)

	testCase(
		"can manage poll",
		t,
		f,
		true,
		`---
		vote/1:
			option_id: 2
			user_id: 5
			delegated_user_id: 6
		option/2/poll_id: 3
		poll/3:
			meeting_id: 1
			content_object_id: topic/1
		`,
		withPerms(1, perm.PollCanManage),
	)

	testCase(
		"vote user",
		t,
		f,
		true,
		`---
		vote/1:
			option_id: 2
			user_id: 5
			delegated_user_id: 6
		option/2/poll_id: 3
		poll/3:
			meeting_id: 1
			content_object_id: topic/1
		`,
		withRequestUser(5),
	)

	testCase(
		"vote user",
		t,
		f,
		true,
		`---
		vote/1:
			option_id: 2
			user_id: 5
			delegated_user_id: 6
		option/2/poll_id: 3
		poll/3:
			meeting_id: 1
			content_object_id: topic/1
		`,
		withRequestUser(6),
	)
}

func TestVoteModeB(t *testing.T) {
	f := collection.Vote{}.Modes("B")

	testCase(
		"other state",
		t,
		f,
		false,
		`---
		vote/1/option_id: 2
		option/2/poll_id: 3
		poll/3/meeting_id: 1
		`,
	)

	testCase(
		"state published",
		t,
		f,
		true,
		`---
		vote/1:
			option_id: 2
			user_id: 5
			delegated_user_id: 6
		option/2/poll_id: 3
		poll/3:
			meeting_id: 1
			state: published
			content_object_id: topic/1
		`,
		withRequestUser(5),
	)

	testCase(
		"state finished",
		t,
		f,
		true,
		`---
		vote/1/option_id: 2
		option/2/poll_id: 3
		poll/3:
			meeting_id: 1
			state: finished
			content_object_id: topic/1
		`,
		withPerms(1, perm.PollCanManage),
	)
}
