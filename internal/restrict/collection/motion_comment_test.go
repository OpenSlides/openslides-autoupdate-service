package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
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
			meeting_id: 30
			motion_id: 5
			section_id: 7
		
		motion/5:
			meeting_id: 30
			state_id: 3
		
		motion_state/3/id: 3

		motion_comment_section/7:
			read_group_ids: [2]
			meeting_id: 30
		
		group/2/id: 2
		meeting/30/id: 30
		`,
	)

	testCase(
		"can see motion but not comment section",
		t,
		f,
		false,
		`---
		motion_comment/1:
			meeting_id: 30
			motion_id: 5
			section_id: 7
		
		motion/5:
			meeting_id: 30
			state_id: 3
		
		motion_state/3/id: 3

		motion_comment_section/7:
			read_group_ids: []
			meeting_id: 30
		
		group/2/id: 2
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"can not see motion but see comment section",
		t,
		f,
		false,
		`---
		motion_comment/1:
			meeting_id: 30
			motion_id: 5
			section_id: 7
		
		motion/5:
			meeting_id: 30
			state_id: 3
			submitter_ids: [4]
		
		motion_state/3/restrictions:
		- is_submitter

		motion_comment_section/7:
			read_group_ids: [2]
			meeting_id: 30

		motion_submitter/4/meeting_user_id: 20
		meeting_user/20/user_id: 2
		
		group/2/id: 2
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"can see motion and comment section with read_group",
		t,
		f,
		true,
		`---
		motion_comment/1:
			meeting_id: 30
			motion_id: 5
			section_id: 7
		
		motion/5:
			meeting_id: 30
			state_id: 3
		
		motion_state/3/id: 3

		motion_comment_section/7:
			meeting_id: 30
			read_group_ids: [2]
		
		meeting_user/10/group_ids: [2]
		group/2/id: 2
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"can see motion and comment section with write_group",
		t,
		f,
		true,
		`---
		motion_comment/1:
			meeting_id: 30
			motion_id: 5
			section_id: 7
		
		motion/5:
			meeting_id: 30
			state_id: 3
		
		motion_state/3/id: 3

		motion_comment_section/7:
			meeting_id: 30
			write_group_ids: [2]
		
		meeting_user/10/group_ids: [2]
		group/2/id: 2
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"can see motion and comment section with submitter_can_write but no submitter",
		t,
		f,
		false,
		`---
		motion_comment/1:
			meeting_id: 30
			motion_id: 5
			section_id: 7
		
		motion/5:
			meeting_id: 30
			state_id: 3
		
		motion_state/3/id: 3

		motion_comment_section/7:
			meeting_id: 30
			submitter_can_write: true
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"can see motion and comment section with submitter_can_write as submitter",
		t,
		f,
		true,
		`---
		motion_comment/1:
			meeting_id: 30
			motion_id: 5
			section_id: 7
		
		motion/5:
			meeting_id: 30
			state_id: 3
			submitter_ids: [13]

		motion_submitter/13:
			meeting_user_id: 10
		
		meeting_user/10/user_id: 1
		
		motion_state/3/id: 3

		motion_comment_section/7:
			meeting_id: 30
			submitter_can_write: true
		`,
		withPerms(30, perm.MotionCanSee),
	)
}
