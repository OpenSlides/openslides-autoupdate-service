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
		"see with other group",
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
		"see with group",
		t,
		f,
		true,
		`---
		user/1/group_$30_ids: [2]
		group/2/id: 2

		motion_comment_section/1:
			meeting_id: 30
			read_group_ids: [2]
		`,
		withPerms(30, perm.MotionCanSee),
	)
}
