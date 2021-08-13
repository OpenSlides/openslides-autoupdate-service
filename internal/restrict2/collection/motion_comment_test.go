package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
)

func TestMotionCommentModeA(t *testing.T) {
	f := collection.MotionComment{}.Modes("A")

	testCase(
		"no perms",
		t,
		f,
		false,
		`---
		motion_comment/1:
			meeting_id: 1
			motion_id: 5
			section_id: 7
		
		motion/5:
			meeting_id: 1
			state_id: 3
		
		motion_state/3/id: 3

		motion_comment_section/7:
			read_group_ids: [2]
		
		user/1/group_$1_ids: [2]
		group/2/id: 2
		meeting/1/id: 1
		`,
	)

	testCase(
		"can see motion but not comment section",
		t,
		f,
		false,
		`---
		motion_comment/1:
			meeting_id: 1
			motion_id: 5
			section_id: 7
		
		motion/5:
			meeting_id: 1
			state_id: 3
		
		motion_state/3/id: 3

		motion_comment_section/7:
			read_group_ids: []
		
		user/1/group_$1_ids: [2]
		group/2/id: 2
		`,
		withPerms(1, perm.MotionCanSee),
	)

	testCase(
		"can not see motion but see comment section",
		t,
		f,
		false,
		`---
		motion_comment/1:
			meeting_id: 1
			motion_id: 5
			section_id: 7
		
		motion/5:
			meeting_id: 1
			state_id: 3
			submitter_ids: [4]
		
		motion_state/3/restrictions:
		- is_submitter

		motion_comment_section/7:
			read_group_ids: [2]

		motion_submitter/4/user_id: 2
		
		user/1/group_$1_ids: [2]
		group/2/id: 2
		`,
		withPerms(1, perm.MotionCanSee),
	)

	testCase(
		"can see motion and comment section",
		t,
		f,
		true,
		`---
		motion_comment/1:
			meeting_id: 1
			motion_id: 5
			section_id: 7
		
		motion/5:
			meeting_id: 1
			state_id: 3
		
		motion_state/3/id: 3

		motion_comment_section/7:
			meeting_id: 1
			read_group_ids: [2]
		
		user/1/group_$1_ids: [2]
		group/2/id: 2
		`,
		withPerms(1, perm.MotionCanSee),
	)

}
