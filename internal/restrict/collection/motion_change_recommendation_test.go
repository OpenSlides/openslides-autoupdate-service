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
			meeting_id: 1
		`,
	)

	testCase(
		"can manage",
		t,
		f,
		true,
		`---
		motion_change_recommendation/1/meeting_id: 1
		`,
		withPerms(1, perm.MotionCanManage),
	)

	testCase(
		"can see not internal",
		t,
		f,
		true,
		`---
		motion_change_recommendation/1:
			meeting_id: 1
			internal: false
		`,
		withPerms(1, perm.MotionCanSee),
	)

	testCase(
		"can see internal",
		t,
		f,
		false,
		`---
		motion_change_recommendation/1:
			meeting_id: 1
			internal: true
		`,
		withPerms(1, perm.MotionCanSee),
	)
}
