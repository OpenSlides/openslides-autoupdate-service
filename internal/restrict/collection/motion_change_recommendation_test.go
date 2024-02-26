package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

func TestMotionChangeRecommendationModeA(t *testing.T) {
	f := collection.MotionChangeRecommendation{}.Modes("A")

	testCase(
		"no perms",
		t,
		f,
		false,
		`---
		motion_change_recommendation/1:
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
		motion_change_recommendation/1/meeting_id: 30
		`,
		withPerms(30, perm.MotionCanManage),
	)

	testCase(
		"can manage metadata",
		t,
		f,
		true,
		`---
		motion_change_recommendation/1/meeting_id: 30
		`,
		withPerms(30, perm.MotionCanManageMetadata),
	)

	testCase(
		"can see internal change_recommendation",
		t,
		f,
		false,
		`---
		motion_change_recommendation/1:
			meeting_id: 30
			internal: true
			motion_id: 15

		motion/15:
			state_id: 10

		motion_state/10:
			is_internal: false
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"can see internal motion state",
		t,
		f,
		false,
		`---
		motion_change_recommendation/1:
			meeting_id: 30
			internal: false
			motion_id: 15

		motion/15:
			state_id: 10
		
		motion_state/10:
			is_internal: true
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"can see not internal motion state and not internal change recommendation",
		t,
		f,
		true,
		`---
		motion_change_recommendation/1:
			meeting_id: 30
			internal: false
			motion_id: 15

		motion/15:
			state_id: 10
		
		motion_state/10:
			is_internal: false
		`,
		withPerms(30, perm.MotionCanSee),
	)
}
