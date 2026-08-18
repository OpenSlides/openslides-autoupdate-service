package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-go/perm"
)

func TestPollBallotUserModeA(t *testing.T) {
	f := collection.PollBallotUser{}.Modes("A")

	testCase(
		"no perms",
		t,
		f,
		false,
		`---
		poll_ballot_user/1:
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
		poll_ballot_user/1:
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
		poll_ballot_user/1:
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
		poll_ballot_user/1:
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
		"own ballot",
		t,
		f,
		true,
		`---
		poll_ballot_user/1:
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
		"request user is delegated, but did not send ballot",
		t,
		f,
		true,
		`---
		poll_ballot_user/1:
			poll_id: 3
			represented_meeting_user_id: 5
			acting_meeting_user_id: 5
		meeting_user/5:
			user_id: 50
			vote_delegated_to_id: 6
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

	testCase(
		"request user did sent ballot as delegated, but is not delegated anymore",
		t,
		f,
		false,
		`---
		poll_ballot_user/1:
			poll_id: 3
			represented_meeting_user_id: 5
			acting_meeting_user_id: 6
		meeting_user/5:
			user_id: 50
			vote_delegated_to_id: null
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

	testCase(
		"request user is delegated, but delegation is deactivated",
		t,
		f,
		true,
		`---
		poll_ballot_user/1:
			poll_id: 3
			represented_meeting_user_id: 5
			acting_meeting_user_id: 5
		meeting_user/5:
			user_id: 50
			vote_delegated_to_id: 6
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
