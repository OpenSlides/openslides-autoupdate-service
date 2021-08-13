package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

func TestPollModeA(t *testing.T) {
	f := collection.Poll{}.Modes("A")

	testCase(
		"no perms",
		t,
		f,
		false,
		`---
		poll/1/meeting_id: 1
		meeting/1/id: 1
		`,
	)

	testCase(
		"motion can see",
		t,
		f,
		true,
		`---
		poll/1:
			content_object_id: motion/1
		
		motion/1:
			meeting_id: 1
		`,
		withPerms(1, perm.MotionCanSee),
	)

	testCase(
		"motion can not see",
		t,
		f,
		false,
		`---
		poll/1:
			content_object_id: motion/1
		
		motion/1:
			meeting_id: 1
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
			meeting_id: 1
		`,
		withPerms(1, perm.AssignmentCanSee),
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
			meeting_id: 1
		`,
	)

	testCase(
		"other can see meeting",
		t,
		f,
		true,
		`---
		poll/1:
			meeting_id: 1
			content_object_id: topic/1
		
		meeting/1/enable_anonymous: true
		`,
	)

	testCase(
		"other can not see meeting",
		t,
		f,
		false,
		`---
		poll/1:
			meeting_id: 1
			content_object_id: topic/1
		
		meeting/1/enable_anonymous: false
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
			content_object_id: motion/1
			state: published
		
		motion/1:
			meeting_id: 1
		`,
		withPerms(1, perm.MotionCanSee),
	)

	testCase(
		"published can not see",
		t,
		f,
		false,
		`---
		poll/1:
			content_object_id: motion/1
			state: published
		
		motion/1:
			meeting_id: 1
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
			meeting_id: 1
		`,
		withPerms(1, perm.MotionCanManagePolls),
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
			meeting_id: 1
		`,
	)

	testCase(
		"finished can manage assignment",
		t,
		f,
		true,
		`---
		poll/1:
			content_object_id: assignment/1
			state: finished
		
		assignment/1:
			meeting_id: 1
		`,
		withPerms(1, perm.AssignmentCanManage),
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
			meeting_id: 1
		`,
	)

	testCase(
		"finished can manage other",
		t,
		f,
		true,
		`---
		poll/1:
			meeting_id: 1
			content_object_id: topic/1
			state: finished
		`,
		withPerms(1, perm.PollCanManage),
	)

	testCase(
		"finished can not manage other",
		t,
		f,
		false,
		`---
		poll/1:
			content_object_id: topic/1
			state: finished
		`,
	)

	testCase(
		"other",
		t,
		f,
		false,
		`---
		poll/1:
			content_object_id: motion/1
			state: other
			meeting_id: 1
		`,
		withPerms(1, perm.PollCanManage),
	)
}

func TestPollModeD(t *testing.T) {
	f := collection.Poll{}.Modes("D")

	testCase(
		"published can see",
		t,
		f,
		true,
		`---
		poll/1:
			content_object_id: motion/1
			state: published
		
		motion/1:
			meeting_id: 1
		`,
		withPerms(1, perm.MotionCanSee),
	)

	testCase(
		"published can not see",
		t,
		f,
		false,
		`---
		poll/1:
			content_object_id: motion/1
			state: published
		
		motion/1:
			meeting_id: 1
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
			meeting_id: 1
		`,
		withPerms(1, perm.MotionCanManagePolls),
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
			meeting_id: 1
		`,
	)

	testCase(
		"finished can manage list of speakers",
		t,
		f,
		true,
		`---
		poll/1:
			content_object_id: motion/1
			state: finished
			meeting_id: 1
		
		motion/1:
			meeting_id: 1
		`,
		withPerms(1, perm.ListOfSpeakersCanManage),
	)

}
