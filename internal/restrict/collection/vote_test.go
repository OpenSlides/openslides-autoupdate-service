package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-go/perm"
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
		poll/3:
			meeting_id: 30
			content_object_id: topic/5
		topic/5:
			meeting_id: 30
			agenda_item_id: 40
		agenda_item/40/meeting_id: 30
		meeting/30/enable_anonymous: false
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
			meeting_id: 30
			state: published
			content_object_id: topic/5
		topic/5:
			meeting_id: 30
			agenda_item_id: 40
		agenda_item/40/meeting_id: 30
		`,
		withPerms(30, perm.AgendaItemCanSee),
	)

	testCase(
		"poll is published but can not see it",
		t,
		f,
		false,
		`---
		vote/1:
			option_id: 2
			user_id: 5
			delegated_user_id: 6
		option/2/poll_id: 3
		poll/3:
			meeting_id: 30
			state: published
			content_object_id: topic/5
		topic/5:
			meeting_id: 30
			agenda_item_id: 40
		agenda_item/40/meeting_id: 30
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
			meeting_id: 30
			content_object_id: topic/5
		topic/5:
			meeting_id: 30
			agenda_item_id: 40
		agenda_item/40/meeting_id: 30
		`,
		withPerms(30, perm.PollCanManage, perm.AgendaItemCanSee),
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
			meeting_id: 30
			content_object_id: topic/5
		topic/5:
			meeting_id: 30
			agenda_item_id: 40
		agenda_item/40/meeting_id: 30
		`,
		withRequestUser(5),
		withPerms(30, perm.AgendaItemCanSee),
	)

	testCase(
		"vote user from delegated",
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
			meeting_id: 30
			content_object_id: topic/5
		topic/5:
			meeting_id: 30
			agenda_item_id: 40
		agenda_item/40/meeting_id: 30
		`,
		withRequestUser(6),
		withPerms(30, perm.AgendaItemCanSee),
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
		poll/3/meeting_id: 30
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
			meeting_id: 30
			state: published
			content_object_id: topic/5
		topic/5:
			meeting_id: 30
			agenda_item_id: 40
		agenda_item/40/meeting_id: 30
		`,
		withRequestUser(5),
		withPerms(30, perm.AgendaItemCanSee),
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
			meeting_id: 30
			state: finished
			content_object_id: topic/5
		topic/5/meeting_id: 30
		`,
		withPerms(30, perm.PollCanManage),
	)
}
