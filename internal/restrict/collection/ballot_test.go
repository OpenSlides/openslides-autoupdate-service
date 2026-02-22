package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-go/perm"
)

func TestBallotModeA(t *testing.T) {
	f := collection.Ballot{}.Modes("A")

	testCase(
		"no perms",
		t,
		f,
		false,
		`---
		ballot/1:
			poll_id: 3
			represented_meeting_user_id: 5
			acting_meeting_user_id: 6
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
		ballot/1:
			poll_id: 3
			represented_meeting_user_id: 5
			acting_meeting_user_id: 6
		poll/3:
			meeting_id: 30
			published: true
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
		ballot/1:
			poll_id: 3
			represented_meeting_user_id: 5
			acting_meeting_user_id: 6
		poll/3:
			meeting_id: 30
			published: true
			content_object_id: topic/5
		topic/5:
			meeting_id: 30
			agenda_item_id: 40
		agenda_item/40/meeting_id: 30
		`,
	)

	testCase(
		"can see progress",
		t,
		f,
		true,
		`---
		ballot/1:
			poll_id: 3
			represented_meeting_user_id: 5
			acting_meeting_user_id: 6
		poll/3:
			meeting_id: 30
			content_object_id: topic/5
		topic/5:
			meeting_id: 30
			agenda_item_id: 40
		agenda_item/40/meeting_id: 30
		`,
		withPerms(30, perm.PollCanSeeProgress, perm.AgendaItemCanSee),
	)

	testCase(
		"vote user",
		t,
		f,
		true,
		`---
		ballot/1:
			poll_id: 3
			represented_meeting_user_id: 5
			acting_meeting_user_id: 6
		meeting_user/5/user_id: 50
		meeting_user/6/user_id: 60
		poll/3:
			meeting_id: 30
			content_object_id: topic/5
		topic/5:
			meeting_id: 30
			agenda_item_id: 40
		agenda_item/40/meeting_id: 30
		`,
		withRequestUser(50),
		withPerms(30, perm.AgendaItemCanSee),
	)

	testCase(
		"vote user from delegated",
		t,
		f,
		true,
		`---
		ballot/1:
			poll_id: 3
			represented_meeting_user_id: 5
			acting_meeting_user_id: 6
		meeting_user/5/user_id: 50
		meeting_user/6/user_id: 60
		poll/3:
			meeting_id: 30
			content_object_id: topic/5
		topic/5:
			meeting_id: 30
			agenda_item_id: 40
		agenda_item/40/meeting_id: 30
		`,
		withRequestUser(60),
		withPerms(30, perm.AgendaItemCanSee),
	)
}

func TestBallotModeB(t *testing.T) {
	f := collection.Ballot{}.Modes("B")

	testCase(
		"poll is published",
		t,
		f,
		true,
		`---
		ballot/1:
			poll_id: 3
			represented_meeting_user_id: 5
			acting_meeting_user_id: 6
		poll/3:
			meeting_id: 30
			published: true
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
		ballot/1/poll_id: 3
		poll/3:
			meeting_id: 30
			state: finished
			content_object_id: topic/5
		topic/5/meeting_id: 30
		`,
		withPerms(30, perm.PollCanManage),
	)
}

func TestBallotModeC(t *testing.T) {
	f := collection.Ballot{}.Modes("C")

	testCase(
		"poll is published, but secret",
		t,
		f,
		false,
		`---
		ballot/1:
			poll_id: 3
		poll/3:
			meeting_id: 30
			published: true
			visibility: secret
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
		"poll is published, not secret",
		t,
		f,
		true,
		`---
		ballot/1:
			poll_id: 3
		poll/3:
			meeting_id: 30
			published: true
			visibility: open
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
		"state finished, not secret",
		t,
		f,
		true,
		`---
		ballot/1/poll_id: 3
		poll/3:
			meeting_id: 30
			state: finished
			visibility: open
			content_object_id: topic/5
		topic/5/meeting_id: 30
		`,
		withPerms(30, perm.PollCanManage),
	)

	testCase(
		"state finished, but secret",
		t,
		f,
		false,
		`---
		ballot/1/poll_id: 3
		poll/3:
			meeting_id: 30
			state: finished
			visibility: secret
			content_object_id: topic/5
		topic/5/meeting_id: 30
		`,
		withPerms(30, perm.PollCanManage),
	)

	testCase(
		"secret, own ballot",
		t,
		f,
		true,
		`---
		ballot/1:
			poll_id: 3
			represented_meeting_user_id: 10
		user/1/meeting_user_ids: [10]
		poll/3:
			meeting_id: 30
			visibility: secret
			content_object_id: topic/5
		topic/5/meeting_id: 30
		`,
		withPerms(30, perm.PollCanManage),
	)

	testCase(
		"secret, ballot of deletated",
		t,
		f,
		true,
		`---
		ballot/1:
			poll_id: 3
			represented_meeting_user_id: 20
		meeting_user/20:
			user_id: 7
		meeting_user/10:
			user_id: 1
			vote_delegations_from_ids: [20]
		poll/3:
			meeting_id: 30
			visibility: secret
			content_object_id: topic/5
		topic/5/meeting_id: 30
		`,
		withPerms(30, perm.PollCanManage),
	)
}
