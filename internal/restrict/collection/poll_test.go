package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-go/perm"
)

func TestPollModeA(t *testing.T) {
	f := collection.Poll{}.Modes("A")

	testCase(
		"no perms",
		t,
		f,
		false,
		`---
		poll/1:
			meeting_id: 30
			content_object_id: topic/5

		topic/5:
			meeting_id: 30
			agenda_item_id: 7
		agenda_item/7/meeting_id: 30

		meeting/30:
			id: 1
			committee_id: 300
		`,
	)

	testCase(
		"motion can see",
		t,
		f,
		true,
		`---
		poll/1:
			content_object_id: motion/2

		motion/2:
			meeting_id: 30
			state_id: 3

		motion_state/3/id: 3
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"motion can not see",
		t,
		f,
		false,
		`---
		poll/1:
			content_object_id: motion/2

		motion/2:
			meeting_id: 30
			state_id: 3

		motion_state/3/id: 3

		meeting/30/locked_from_inside: false
		`,
	)

	testCase(
		"assignment can see",
		t,
		f,
		true,
		`---
		poll/1:
			content_object_id: assignment/1

		assignment/1:
			meeting_id: 30
		`,
		withPerms(30, perm.AssignmentCanSee),
	)

	testCase(
		"assignment can not see",
		t,
		f,
		false,
		`---
		poll/1:
			content_object_id: assignment/1

		assignment/1:
			meeting_id: 30
			list_of_speakers_id: 300

		list_of_speakers/300/meeting_id: 30
		`,
	)

	testCase(
		"topic can see",
		t,
		f,
		true,
		`---
		poll/1:
			meeting_id: 30
			content_object_id: topic/5

		topic/5:
			meeting_id: 30
			agenda_item_id: 3

		agenda_item/3/meeting_id: 30
		`,
		withPerms(30, perm.AgendaItemCanSee),
	)

	testCase(
		"other can not see agenda",
		t,
		f,
		false,
		`---
		poll/1:
			meeting_id: 30
			content_object_id: topic/5

		topic/5:
			meeting_id: 30
			agenda_item_id: 3

		agenda_item/3/meeting_id: 30
		`,
	)
}

func TestPollModeB(t *testing.T) {
	f := collection.Poll{}.Modes("B")

	testCase(
		"published can see",
		t,
		f,
		true,
		`---
		poll/1:
			content_object_id: motion/2
			published: true

		motion/2:
			meeting_id: 30
			state_id: 3

		motion_state/3/id: 3
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"published can not see",
		t,
		f,
		false,
		`---
		poll/1:
			content_object_id: motion/2
			published: true

		motion/2:
			meeting_id: 30
			state_id: 3

		motion_state/3/id: 3

		meeting/30/locked_from_inside: false
		`,
	)

	testCase(
		"finished can manage motion",
		t,
		f,
		true,
		`---
		poll/1:
			content_object_id: motion/1
			state: finished

		motion/1:
			meeting_id: 30
		`,
		withPerms(30, perm.MotionCanManagePolls),
	)

	testCase(
		"finished can not manage motion",
		t,
		f,
		false,
		`---
		poll/1:
			content_object_id: motion/1
			state: finished

		motion/1:
			meeting_id: 30
		`,
	)

	testCase(
		"finished can manage assignment polls",
		t,
		f,
		true,
		`---
		poll/1:
			content_object_id: assignment/1
			state: finished

		assignment/1:
			meeting_id: 30
		`,
		withPerms(30, perm.AssignmentCanManagePolls),
	)

	testCase(
		"finished can manage assignment",
		t,
		f,
		false,
		`---
		poll/1:
			content_object_id: assignment/1
			state: finished

		assignment/1:
			meeting_id: 30
		`,
		withPerms(30, perm.AssignmentCanManage),
	)

	testCase(
		"finished can not manage assignment",
		t,
		f,
		false,
		`---
		poll/1:
			content_object_id: assignment/1
			state: finished

		assignment/1:
			meeting_id: 30
		`,
	)

	testCase(
		"finished can manage poll",
		t,
		f,
		true,
		`---
		poll/1:
			meeting_id: 30
			content_object_id: topic/5
			state: finished
		topic/5/meeting_id: 30
		`,
		withPerms(30, perm.PollCanManage),
	)

	testCase(
		"finished can not manage poll",
		t,
		f,
		false,
		`---
		poll/1:
			meeting_id: 30
			content_object_id: topic/5
			state: finished
		topic/5/meeting_id: 30
		`,
	)
}
