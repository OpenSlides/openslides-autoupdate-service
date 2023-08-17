package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

func TestMotionSubmitterModeA(t *testing.T) {
	f := collection.MotionSubmitter{}.Modes("A")

	testCase(
		"no perms",
		t,
		f,
		false,
		`---
		motion_submitter/1:
			motion_id: 5
		
		motion/5:
			meeting_id: 30
			state_id: 3
		
		motion_state/3/id: 3
		`,
	)

	testCase(
		"see motion",
		t,
		f,
		true,
		`---
		motion_submitter/1:
			motion_id: 5
		
		motion/5:
			meeting_id: 30
			state_id: 3
		
		motion_state/3/id: 3
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCaseMulti(
		"see many same motion",
		t,
		f,
		[]int{1, 2},
		[]int{1, 2},
		`---
		motion_submitter:
			1:
				motion_id: 5
			2:
				motion_id: 5
		
		motion/5:
			meeting_id: 30
			state_id: 3
			
		motion_state/3/id: 3
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCaseMulti(
		"see many different motions",
		t,
		f,
		[]int{1, 2},
		[]int{1, 2},
		`---
		motion_submitter:
			1:
				motion_id: 5
			2:
				motion_id: 6
		
		motion:
			5:
				meeting_id: 30
				state_id: 3
			
			6: 
				meeting_id: 30
				state_id: 3
			
		motion_state/3/id: 3
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCaseMulti(
		"see many different motions only one",
		t,
		f,
		[]int{1, 2},
		[]int{1},
		`---
		motion_submitter:
			1:
				motion_id: 5
			2:
				motion_id: 6
		
		motion:
			5:
				meeting_id: 30
				state_id: 3
				submitter_ids: [4]
			
			6: 
				meeting_id: 30
				state_id: 3
			
		motion_state/3/restrictions:
			- is_submitter

		motion_submitter/4/meeting_user_id: 10
		meeting_user/10/user_id: 1
		`,
		withPerms(30, perm.MotionCanSee),
	)
}
