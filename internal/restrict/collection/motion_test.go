package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-go/perm"
)

func TestMotionModeC(t *testing.T) {
	f := collection.Motion{}.Modes("C")

	testCase(
		"no permissions",
		t,
		f,
		false,
		`---
		motion/1:
			id: 1
			meeting_id: 30

		meeting/30/locked_from_inside: false
		`,
	)

	testCase(
		"motion.can_see",
		t,
		f,
		true,
		`---
		motion/1:
			meeting_id: 30
			state_id: 3

		motion_state/3/id: 3
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"motion/all_derived_motion_ids but not motion.can_see_origin",
		t,
		f,
		false,
		`---
		motion:
			1:
				meeting_id: 30
				state_id: 3
				all_derived_motion_ids: [10]

			10:
				meeting_id: 31
				state_id: 3

		motion_state/3/id: 3

		meeting/30/locked_from_inside: false
		`,
		withPerms(31, perm.MotionCanSee),
	)

	testCase(
		"motion/all_derived_motion_ids with motion.can_see_origin",
		t,
		f,
		true,
		`---
		motion:
			1:
				meeting_id: 30
				state_id: 3
				all_derived_motion_ids: [10]

			10:
				meeting_id: 31
				state_id: 3

		motion_state/3/id: 3

		meeting/30/locked_from_inside: false
		`,
		withPerms(31, perm.MotionCanSeeOrigin),
	)

	testCase(
		"motion/all_derived_motion_ids with motion.can_see_origin and locked meeting",
		t,
		f,
		false,
		`---
		motion:
			1:
				meeting_id: 30
				state_id: 3
				all_derived_motion_ids: [10]

			10:
				meeting_id: 31
				state_id: 3

		motion_state/3/id: 3

		meeting/30/locked_from_inside: true
		`,
		withPerms(31, perm.MotionCanSeeOrigin),
	)

	testCase(
		"motion.can_see with restrict is_submitter",
		t,
		f,
		true,
		`---
		motion/1:
			meeting_id: 30
			state_id: 3
			submitter_ids: [4]

		motion_state/3/restrictions:
		- is_submitter

		motion_submitter/4/meeting_user_id: 10
		meeting_user/10/user_id: 1
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"motion.can_see with restrict is_submitter not submitter",
		t,
		f,
		false,
		`---
		motion/1:
			meeting_id: 30
			state_id: 3
			submitter_ids: [4]

		motion_state/3/restrictions:
		- is_submitter

		motion_submitter/4/meeting_user_id: 20
		meeting_user/20/user_id: 2
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"admin with restrict is_submitter not submitter",
		t,
		f,
		true,
		`---
		motion/1:
			meeting_id: 30
			state_id: 3
			submitter_ids: [4]

		motion_state/3/restrictions:
		- is_submitter

		motion_submitter/4/meeting_user_id: 20
		meeting_user/20/user_id: 2

		meeting/30/admin_group_id: 13
		user/1/meeting_user_ids: [10]
		meeting_user/10:
			group_ids: [13]
			meeting_id: 30
		`,
	)

	testCase(
		"motion.can_see with restrict perm",
		t,
		f,
		true,
		`---
		motion/1:
			meeting_id: 30
			state_id: 3

		motion_state/3/restrictions:
		- motion.can_manage
		`,
		withPerms(30, perm.MotionCanSee, perm.MotionCanManage),
	)

	testCase(
		"motion.can_see with restrict perm without perm",
		t,
		f,
		false,
		`---
		motion/1:
			meeting_id: 30
			state_id: 3

		motion_state/3/restrictions:
		- motion.can_manage
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"motion.can_see with two restrictions, non fullfield",
		t,
		f,
		false,
		`---
		motion/1:
			meeting_id: 30
			state_id: 3

		motion_state/3/restrictions:
		- motion.can_manage
		- is_submitter
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"motion.can_see with two restrictions, is_submitter fullfield",
		t,
		f,
		true,
		`---
		motion/1:
			meeting_id: 30
			state_id: 3
			submitter_ids: [4]

		motion_state/3/restrictions:
		- motion.can_manage
		- is_submitter

		motion_submitter/4/meeting_user_id: 10
		meeting_user/10/user_id: 1
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"motion.can_see with two restrictions, motion.can_manage fullfield",
		t,
		f,
		true,
		`---
		motion/1:
			meeting_id: 30
			state_id: 3

		motion_state/3/restrictions:
		- motion.can_manage
		- is_submitter
		`,
		withPerms(30, perm.MotionCanManage),
	)

	testCase(
		"motion.can_see with two restrictions, both fullfield",
		t,
		f,
		true,
		`---
		motion/1:
			meeting_id: 30
			state_id: 3
			submitter_ids: [4]

		motion_state/3/restrictions:
		- motion.can_manage
		- is_submitter

		motion_submitter/4/meeting_user_id: 10
		meeting_user/10/user_id: 1
		`,
		withPerms(30, perm.MotionCanManage),
	)

	testCase(
		"Can see motion but not the lead motion",
		t,
		f,
		false,
		`---
		motion/1:
			meeting_id: 30
			lead_motion_id: 2
			state_id: 10

		motion/2:
			meeting_id: 30
			state_id: 20
			submitter_ids: [4]

		motion_state/10/id: 10

		motion_state/20/restrictions:
			- is_submitter

		motion_submitter/4/meeting_user_id: 400
		meeting_user/400/user_id: 40
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"Can see motion and the lead motion",
		t,
		f,
		true,
		`---
		motion/1:
			meeting_id: 30
			lead_motion_id: 2
			state_id: 10

		motion/2:
			meeting_id: 30
			state_id: 20
			submitter_ids: [4]

		motion_state/10/id: 10

		motion_state/20/restrictions:
			- is_submitter

		motion_submitter/4/meeting_user_id: 10

		meeting_user/10/user_id: 1
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"Motion is its own lead motion can see",
		t,
		f,
		true,
		`---
		motion/1:
			meeting_id: 30
			lead_motion_id: 1
			state_id: 10

		motion_state/10/id: 10
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"Motion is its own lead motion can not see",
		t,
		f,
		false,
		`---
		motion/1:
			meeting_id: 30
			lead_motion_id: 1
			state_id: 10

		motion_state/10/id: 10

		meeting/30/locked_from_inside: true
		`,
	)

	testCase(
		"motions lead_id circle allowed",
		t,
		f,
		true,
		`---
		motion/1:
			meeting_id: 30
			lead_motion_id: 2
			state_id: 10

		motion/2:
			meeting_id: 30
			lead_motion_id: 3
			state_id: 10

		motion/3:
			meeting_id: 30
			lead_motion_id: 1
			state_id: 10

		motion_state/10/id: 10

		motion_state/30/restrictions:
			- is_submitter

		motion_submitter/4/meeting_user_id: 10
		meeting_user/10/user_id: 1
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"motions lead_id circle now allowed",
		t,
		f,
		false,
		`---
		motion/1:
			meeting_id: 30
			lead_motion_id: 2
			state_id: 10

		motion/2:
			meeting_id: 30
			lead_motion_id: 3
			state_id: 10

		motion/3:
			meeting_id: 30
			lead_motion_id: 1
			state_id: 30

		motion_state/10/id: 10

		motion_state/30/restrictions:
			- is_submitter

		motion_submitter/4/meeting_user_id: 4040
		meeting_user/4040/user_id: 404
		`,
		withPerms(30, perm.MotionCanSee),
	)
}

func TestMotionModeA(t *testing.T) {
	f := collection.Motion{}.Modes("A")

	testCase(
		"no permissions",
		t,
		f,
		false,
		`---
		motion/1:
			id: 1
			meeting_id: 30

		meeting/30/committee_id: 300
		`,
	)

	testCase(
		"motion.can_see",
		t,
		f,
		true,
		`---
		motion/1:
			meeting_id: 30
			state_id: 3

		motion_state/3/id: 3
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"see all_origin_ids",
		t,
		f,
		true,
		`---
		motion:
			1:
				meeting_id: 30
				state_id: 3
				submitter_ids: [4]
				all_origin_ids: [2]
			2:
				meeting_id: 30
				state_id: 3
				submitter_ids: [5]

		motion_state/3/restrictions:
		- is_submitter

		motion_submitter/4/meeting_user_id: 20
		motion_submitter/5/meeting_user_id: 10
		meeting_user/20/user_id: 2
		meeting_user/10/user_id: 1
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"not see all_origin_ids",
		t,
		f,
		false,
		`---
		motion:
			1:
				meeting_id: 30
				state_id: 3
				submitter_ids: [4]
				all_origin_ids: [2]
			2:
				meeting_id: 30
				state_id: 3
				submitter_ids: [5]

		motion_state/3/restrictions:
		- is_submitter

		motion_submitter/4/meeting_user_id: 20
		motion_submitter/5/meeting_user_id: 20
		meeting_user/20/user_id: 2
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"see all_derived_motion_ids",
		t,
		f,
		true,
		`---
		motion:
			1:
				meeting_id: 30
				state_id: 3
				submitter_ids: [4]
				all_derived_motion_ids: [2]
			2:
				meeting_id: 30
				state_id: 3
				submitter_ids: [5]

		motion_state/3/restrictions:
		- is_submitter

		motion_submitter/4/meeting_user_id: 20
		motion_submitter/5/meeting_user_id: 10
		meeting_user/20/user_id: 2
		meeting_user/10/user_id: 1
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"not see all_derived_motion_ids",
		t,
		f,
		false,
		`---
		motion:
			1:
				meeting_id: 30
				state_id: 3
				submitter_ids: [4]
				all_derived_motion_ids: [2]
			2:
				meeting_id: 30
				state_id: 3
				submitter_ids: [5]

		motion_state/3/restrictions:
		- is_submitter

		motion_submitter/4/meeting_user_id: 20
		motion_submitter/5/meeting_user_id: 20
		meeting_user/20/user_id: 2
		`,
		withPerms(30, perm.MotionCanSee),
	)
}

func TestMotionModeB(t *testing.T) {
	f := collection.Motion{}.Modes("B")

	testCase(
		"no permissions",
		t,
		f,
		false,
		`---
		motion/1:
			id: 1
			meeting_id: 30
		`,
	)

	testCase(
		"motion.can_manage_metadata",
		t,
		f,
		true,
		`---
		motion/1:
			meeting_id: 30
			editor_ids: [3]
		`,
		withPerms(30, perm.MotionCanManageMetadata),
	)
}

func TestMotionModeD(t *testing.T) {
	f := collection.Motion{}.Modes("D")

	testCase(
		"no permissions",
		t,
		f,
		false,
		`motion/1/id: 1`,
	)
}

func TestMotionModeE(t *testing.T) {
	f := collection.Motion{}.Modes("E")

	testCase(
		"no permissions",
		t,
		f,
		false,
		`---
		motion/1:
			meeting_id: 30
			state_id: 3

		motion_state/3/is_internal: true
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"is_internal false",
		t,
		f,
		true,
		`---
		motion/1:
			meeting_id: 30
			state_id: 3

		motion_state/3/is_internal: false
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"motion.can_manage_metadata",
		t,
		f,
		true,
		`---
		motion/1:
			meeting_id: 30
			state_id: 3

		motion_state/3/is_internal: false
		`,
		withPerms(30, perm.MotionCanManageMetadata),
	)
}
