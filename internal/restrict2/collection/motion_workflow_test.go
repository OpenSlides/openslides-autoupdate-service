package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
)

func TestMotionWorkflowModeA(t *testing.T) {
	f := collection.MotionWorkflow{}.Modes("A")

	testCase(
		"no perms",
		t,
		f,
		false,
		"motion_workflow/1/meeting_id: 1",
	)

	testCase(
		"can see",
		t,
		f,
		true,
		"motion_workflow/1/meeting_id: 1",
		withPerms(1, perm.MotionCanSee),
	)
}
