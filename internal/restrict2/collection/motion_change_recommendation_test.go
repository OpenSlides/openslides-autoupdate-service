package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
)

func TestMotionChangeRecommendationModeA(t *testing.T) {
	t.Parallel()
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
		"can see not internal",
		t,
		f,
		true,
		`---
		motion_change_recommendation/1:
			meeting_id: 30
			internal: false
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"can see internal",
		t,
		f,
		false,
		`---
		motion_change_recommendation/1:
			meeting_id: 30
			internal: true
		`,
		withPerms(30, perm.MotionCanSee),
	)
}
