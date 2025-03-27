package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

func TestMotionStateModeA(t *testing.T) {
	f := collection.MotionState{}.Modes("A")

	testCase(
		"no perms",
		t,
		f,
		false,
		"motion_state/1/meeting_id: 30",
	)

	testCase(
		"can see",
		t,
		f,
		true,
		"motion_state/1/meeting_id: 30",
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"can see motion",
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
				state_id: 7

		motion_state/3:
			meeting_id: 30
			motion_ids: [1]
			
		motion_state/7/id: 7

		meeting/30/locked_from_inside: false
		`,
		withPerms(31, perm.MotionCanSeeOrigin),
		withElementID(3),
	)
}
