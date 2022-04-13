package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
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

		motion_submitter/4/user_id: 1
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

		motion_submitter/4/user_id: 2
		`,
		withPerms(30, perm.MotionCanSee),
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

		motion_submitter/4/user_id: 1
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

		motion_submitter/4/user_id: 1
		`,
		withPerms(30, perm.MotionCanManage),
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
		"See Agenda",
		t,
		f,
		true,
		`---
		motion/1:
			meeting_id: 30
			state_id: 3
			agenda_item_id: 8
		
		motion_state/3/id: 3
		agenda_item/8/meeting_id: 30
		`,
		withPerms(30, perm.AgendaItemCanSee),
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

		motion_submitter/4/user_id: 2
		motion_submitter/5/user_id: 1
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

		motion_submitter/4/user_id: 2
		motion_submitter/5/user_id: 2
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

		motion_submitter/4/user_id: 2
		motion_submitter/5/user_id: 1
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

		motion_submitter/4/user_id: 2
		motion_submitter/5/user_id: 2
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"see origin_id",
		t,
		f,
		true,
		`---
		motion:
			1:
				meeting_id: 30
				state_id: 3
				submitter_ids: [4]
				origin_id: 2
			2:
				meeting_id: 30
				state_id: 3
				submitter_ids: [5]

		motion_state/3/restrictions:
		- is_submitter

		motion_submitter/4/user_id: 2
		motion_submitter/5/user_id: 1
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"see derived_motion_ids",
		t,
		f,
		true,
		`---
		motion:
			1:
				meeting_id: 30
				state_id: 3
				submitter_ids: [4]
				derived_motion_ids: [2]
			2:
				meeting_id: 30
				state_id: 3
				submitter_ids: [5]

		motion_state/3/restrictions:
		- is_submitter

		motion_submitter/4/user_id: 2
		motion_submitter/5/user_id: 1
		`,
		withPerms(30, perm.MotionCanSee),
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
