package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-go/perm"
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
			motion_id: 50

		motion/50:
			meeting_id: 30
			state_id: 40

		meeting/30/locked_from_inside: false
		`,
	)

	testCase(
		"can manage",
		t,
		f,
		true,
		`---
		motion_change_recommendation/1:
			meeting_id: 30
			motion_id: 50

		motion/50:
			meeting_id: 30
			state_id: 40

		motion_state/40/restrictions: []
		`,
		withPerms(30, perm.MotionCanManage),
	)

	testCase(
		"can manage metadata",
		t,
		f,
		true,
		`---
		motion_change_recommendation/1:
			meeting_id: 30
			motion_id: 50

		motion/50:
			meeting_id: 30
			state_id: 40

		motion_state/40/restrictions: []
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
			meeting_id: 30

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
			meeting_id: 30

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
			meeting_id: 30

		motion_state/10:
			is_internal: false
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"can see change recommendation of forwarded motion",
		t,
		f,
		true,
		`---
        motion:
            3:
                meeting_id: 30
                state_id: 10
                all_derived_motion_ids: [10]
            10:
                meeting_id: 31
                state_id: 10

        motion_change_recommendation/1:
            meeting_id: 30
            motion_id: 3

        motion_state/10/id: 10

        meeting/30/locked_from_inside: false
        `,
		withPerms(31, perm.MotionCanSeeOrigin),
		withElementID(1),
	)
}
