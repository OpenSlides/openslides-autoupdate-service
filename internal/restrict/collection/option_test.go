package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

func TestOptionModeA(t *testing.T) {
	f := collection.Option{}.Modes("A")

	testCase(
		"no perms",
		t,
		f,
		false,
		`---
		option/1/poll_id: 3
		poll/3:
			meeting_id: 1
			content_object_id: topic/5
		topic/5/meeting_id: 1
		meeting/1:
			enable_anonymous: false
			committee_id: 300
		`,
	)

	testCase(
		"can see poll",
		t,
		f,
		true,
		`---
		option/1/poll_id: 3
		poll/3:
			meeting_id: 1
			content_object_id: topic/5
		topic/5/meeting_id: 1
		`,
		withPerms(1, perm.AgendaItemCanSee),
	)
}

func TestOptionModeB(t *testing.T) {
	f := collection.Option{}.Modes("B")

	testCase(
		"no perms",
		t,
		f,
		false,
		`---
		option/1/poll_id: 3
		poll/3:
			meeting_id: 1
			content_object_id: topic/5
		topic/5/meeting_id: 1
		meeting/1/committee_id: 1
		`,
	)

	testCase(
		"can see poll",
		t,
		f,
		false,
		`---
		option/1/poll_id: 3
		poll/3:
			meeting_id: 1
			content_object_id: topic/5
		topic/5/meeting_id: 1
		meeting/1/user_ids: [1]
		`,
		withPerms(1, perm.AgendaItemCanSee),
	)

	testCase(
		"can manage poll",
		t,
		f,
		true,
		`---
		option/1/poll_id: 3
		poll/3:
			meeting_id: 1
			content_object_id: topic/5
		topic/5/meeting_id: 1
		meeting/1/user_ids: [1]
		`,
		withPerms(1, perm.PollCanManage, perm.AgendaItemCanSee),
	)

	testCase(
		"can manage poll no see",
		t,
		f,
		false,
		`---
		option/1/poll_id: 3
		poll/3:
			meeting_id: 1
			content_object_id: topic/5
		topic/5/meeting_id: 1
		meeting/1/user_ids: [1]
		`,
		withPerms(1, perm.PollCanManage),
	)

	testCase(
		"poll is published",
		t,
		f,
		true,
		`---
		option/1/poll_id: 3
		poll/3:
			meeting_id: 1
			content_object_id: topic/5
			state: published
		topic/5/meeting_id: 1
		meeting/1/user_ids: [1]
		`,
		withPerms(1, perm.AgendaItemCanSee),
	)

	testCase(
		"poll is published not see",
		t,
		f,
		false,
		`---
		option/1/poll_id: 3
		poll/3:
			meeting_id: 1
			content_object_id: topic/5
			state: published
		topic/5/meeting_id: 1
		meeting/1/user_ids: [1]
		`,
	)
}
