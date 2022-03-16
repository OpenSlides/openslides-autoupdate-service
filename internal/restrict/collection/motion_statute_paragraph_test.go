package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

func TestMotionStatuteParagraphModeA(t *testing.T) {
	f := collection.MotionStatuteParagraph{}.Modes("A")

	testCase(
		"no perms",
		t,
		f,
		false,
		"motion_statute_paragraph/1/meeting_id: 30",
	)

	testCase(
		"can see",
		t,
		f,
		true,
		"motion_statute_paragraph/1/meeting_id: 30",
		withPerms(30, perm.MotionCanSee),
	)
}
