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
			meeting_id: 1
		`,
	)

	testCase(
		"motion.can_see",
		t,
		f,
		true,
		`---
		motion/1:
			meeting_id: 1
			state_id: 3
		
		motion_state/3/id: 3
		`,
		withPerms(1, perm.MotionCanSee),
	)

	testCase(
		"motion.can_see with restrict is_submitter",
		t,
		f,
		true,
		`---
		motion/1:
			meeting_id: 1
			state_id: 3
			submitter_ids: [4]
		
		motion_state/3/restrictions:
		- is_submitter

		motion_submitter/4/user_id: 1
		`,
		withPerms(1, perm.MotionCanSee),
	)

	testCase(
		"motion.can_see with restrict is_submitter not submitter",
		t,
		f,
		false,
		`---
		motion/1:
			meeting_id: 1
			state_id: 3
			submitter_ids: [4]
		
		motion_state/3/restrictions:
		- is_submitter

		motion_submitter/4/user_id: 2
		`,
		withPerms(1, perm.MotionCanSee),
	)

	testCase(
		"motion.can_see with restrict perm",
		t,
		f,
		true,
		`---
		motion/1:
			meeting_id: 1
			state_id: 3
		
		motion_state/3/restrictions:
		- motion.can_manage
		`,
		withPerms(1, perm.MotionCanSee, perm.MotionCanManage),
	)

	testCase(
		"motion.can_see with restrict perm without perm",
		t,
		f,
		false,
		`---
		motion/1:
			meeting_id: 1
			state_id: 3
		
		motion_state/3/restrictions:
		- motion.can_manage
		`,
		withPerms(1, perm.MotionCanSee),
	)

	testCase(
		"motion.can_see with two restrictions, non fullfield",
		t,
		f,
		false,
		`---
		motion/1:
			meeting_id: 1
			state_id: 3
		
		motion_state/3/restrictions:
		- motion.can_manage
		- is_submitter
		`,
		withPerms(1, perm.MotionCanSee),
	)

	testCase(
		"motion.can_see with two restrictions, is_submitter fullfield",
		t,
		f,
		true,
		`---
		motion/1:
			meeting_id: 1
			state_id: 3
			submitter_ids: [4]
		
		motion_state/3/restrictions:
		- motion.can_manage
		- is_submitter

		motion_submitter/4/user_id: 1
		`,
		withPerms(1, perm.MotionCanSee),
	)

	testCase(
		"motion.can_see with two restrictions, motion.can_manage fullfield",
		t,
		f,
		true,
		`---
		motion/1:
			meeting_id: 1
			state_id: 3
		
		motion_state/3/restrictions:
		- motion.can_manage
		- is_submitter
		`,
		withPerms(1, perm.MotionCanManage),
	)

	testCase(
		"motion.can_see with two restrictions, both fullfield",
		t,
		f,
		true,
		`---
		motion/1:
			meeting_id: 1
			state_id: 3
			submitter_ids: [4]
		
		motion_state/3/restrictions:
		- motion.can_manage
		- is_submitter

		motion_submitter/4/user_id: 1
		`,
		withPerms(1, perm.MotionCanManage),
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
			meeting_id: 1
			list_of_speakers_id: 300
		
		list_of_speakers/300/meeting_id: 1
		`,
	)

	testCase(
		"motion.can_see",
		t,
		f,
		true,
		`---
		motion/1:
			meeting_id: 1
			state_id: 3
		
		motion_state/3/id: 3
		`,
		withPerms(1, perm.MotionCanSee),
	)

	testCase(
		"See List of speakers",
		t,
		f,
		true,
		`---
		motion/1:
			meeting_id: 1
			state_id: 3
			list_of_speakers_id: 7
		
		motion_state/3/id: 3
		list_of_speakers/7/meeting_id: 1
		`,
		withPerms(1, perm.ListOfSpeakersCanSee),
	)

	testCase(
		"See Agenda",
		t,
		f,
		true,
		`---
		motion/1:
			meeting_id: 1
			state_id: 3
			agenda_item_id: 8
		
		motion_state/3/id: 3
		agenda_item/8/meeting_id: 1
		`,
		withPerms(1, perm.AgendaItemCanSee),
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
			meeting_id: 1
		
		meeting/1/committee_id: 300
		`,
	)

	testCase(
		"motion.can_see",
		t,
		f,
		true,
		`---
		motion/1:
			meeting_id: 1
			state_id: 3
		
		motion_state/3/id: 3
		`,
		withPerms(1, perm.MotionCanSee),
	)

	testCase(
		"see all_origin_ids",
		t,
		f,
		true,
		`---
		motion:
			1:
				meeting_id: 1
				state_id: 3
				submitter_ids: [4]
				all_origin_ids: [2]
			2:
				meeting_id: 1
				state_id: 3
				submitter_ids: [5]

		motion_state/3/restrictions:
		- is_submitter

		motion_submitter/4/user_id: 2
		motion_submitter/5/user_id: 1
		`,
		withPerms(1, perm.MotionCanSee),
	)

	testCase(
		"not see all_origin_ids",
		t,
		f,
		false,
		`---
		motion:
			1:
				meeting_id: 1
				state_id: 3
				submitter_ids: [4]
				all_origin_ids: [2]
			2:
				meeting_id: 1
				state_id: 3
				submitter_ids: [5]

		motion_state/3/restrictions:
		- is_submitter

		motion_submitter/4/user_id: 2
		motion_submitter/5/user_id: 2
		`,
		withPerms(1, perm.MotionCanSee),
	)

	testCase(
		"see all_derived_motion_ids",
		t,
		f,
		true,
		`---
		motion:
			1:
				meeting_id: 1
				state_id: 3
				submitter_ids: [4]
				all_derived_motion_ids: [2]
			2:
				meeting_id: 1
				state_id: 3
				submitter_ids: [5]

		motion_state/3/restrictions:
		- is_submitter

		motion_submitter/4/user_id: 2
		motion_submitter/5/user_id: 1
		`,
		withPerms(1, perm.MotionCanSee),
	)

	testCase(
		"not see all_derived_motion_ids",
		t,
		f,
		false,
		`---
		motion:
			1:
				meeting_id: 1
				state_id: 3
				submitter_ids: [4]
				all_derived_motion_ids: [2]
			2:
				meeting_id: 1
				state_id: 3
				submitter_ids: [5]

		motion_state/3/restrictions:
		- is_submitter

		motion_submitter/4/user_id: 2
		motion_submitter/5/user_id: 2
		`,
		withPerms(1, perm.MotionCanSee),
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
