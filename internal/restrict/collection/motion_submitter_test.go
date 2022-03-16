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
}
