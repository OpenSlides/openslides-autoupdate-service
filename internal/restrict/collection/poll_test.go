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
		meeting/1:
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
			meeting_id: 1
			state_id: 3
		
		motion_state/3/id: 3
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
			content_object_id: motion/2
		
		motion/2:
			meeting_id: 1
			state_id: 3
		
		motion_state/3/id: 3
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
			list_of_speakers_id: 300
		
		list_of_speakers/300/meeting_id: 1
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
			content_object_id: null
		
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
			content_object_id: null
		
		meeting/1:
			enable_anonymous: false
			committee_id: 404
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
			state: published
		
		motion/2:
			meeting_id: 1
			state_id: 3
			
		motion_state/3/id: 3
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
			content_object_id: motion/2
			state: published
		
		motion/2:
			meeting_id: 1
			state_id: 3
			
		motion_state/3/id: 3
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
			content_object_id: null
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
			content_object_id: null
			state: finished
			meeting_id: 1
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

func TestPollModeC(t *testing.T) {
	f := collection.Poll{}.Modes("C")

	testCase(
		"No permission",
		t,
		f,
		false,
		`---
		poll/1:
			content_objct_id: null
			state: started
			meeting_id: 1
		`,
	)

	testCase(
		"Wrong state",
		t,
		f,
		false,
		`---
		poll/1:
			content_objct_id: null
			state: published
			meeting_id: 1
		`,
		withPerms(1, perm.PollCanManage),
	)

	testCase(
		"Correct",
		t,
		f,
		true,
		`---
		poll/1:
			content_objct_id: null
			state: started
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
			content_object_id: motion/2
			state: published
			meeting_id: 1
		
		motion/2:
			meeting_id: 1
			state_id: 3
		
		motion_state/3/id: 3
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
			meeting_id: 1
		
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
			meeting_id: 1
		
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
			meeting_id: 1
		
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
