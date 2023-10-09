package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
)

func TestMotionWorkflowModeA(t *testing.T) {
	t.Parallel()
	f := collection.MotionWorkflow{}.Modes("A")

	testCase(
		"no perms",
		t,
		f,
		false,
		"motion_workflow/1/meeting_id: 30",
	)

	testCase(
		"can see",
		t,
		f,
		true,
		"motion_workflow/1/meeting_id: 30",
		withPerms(30, perm.MotionCanSee),
	)
}
