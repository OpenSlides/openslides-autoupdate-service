package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-go/perm"
)

func TestMotionCategoryModeA(t *testing.T) {
	f := collection.MotionCategory{}.Modes("A")

	testCase(
		"no perms",
		t,
		f,
		false,
		"motion_category/1/meeting_id: 30",
	)

	testCase(
		"can see",
		t,
		f,
		true,
		"motion_category/1/meeting_id: 30",
		withPerms(30, perm.MotionCanSee),
	)
}
