package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

func TestMotionCommentSectionModeA(t *testing.T) {
	f := collection.MotionCommentSection{}.Modes("A")

	testCase(
		"no perms",
		t,
		f,
		false,
		`---
		motion_comment_section/1:
			id: 1
			meeting_id: 30
		`,
	)

	testCase(
		"can manage",
		t,
		f,
		true,
		`---
		motion_comment_section/1/meeting_id: 30
		`,
		withPerms(30, perm.MotionCanManage),
	)

	testCase(
		"see without any group",
		t,
		f,
		false,
		`---
		motion_comment_section/1/meeting_id: 30
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"see with not related other group",
		t,
		f,
		false,
		`---
		motion_comment_section/1:
			meeting_id: 30
			read_group_ids: [2]
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"see with read group",
		t,
		f,
		true,
		`---
		user/1/meeting_user_ids: [10]
		meeting_user/10/group_ids: [2]

		group/2/id: 2

		motion_comment_section/1:
			meeting_id: 30
			read_group_ids: [2]
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"see with write group",
		t,
		f,
		true,
		`---
		user/1/meeting_user_ids: [10]
		meeting_user/10:
			group_ids: [2]
			meeting_id: 30
		group/2/id: 2

		motion_comment_section/1:
			meeting_id: 30
			write_group_ids: [2]
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"see with submitter_can_write",
		t,
		f,
		true,
		`---
		motion_comment_section/1:
			meeting_id: 30
			submitter_can_write: true
		`,
		withPerms(30, perm.MotionCanSee),
	)
}
