package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
)

func TestAssignmentCandidateModeA(t *testing.T) {
	t.Parallel()
	var a collection.AssignmentCandidate
	ds := `---
	assignment_candidate/1/assignment_id: 7
	
	assignment/7:
		meeting_id: 30
	`

	testCase(
		"Can see all assignments",
		t,
		a.Modes("A"),
		true,
		ds,
		withPerms(30, perm.AssignmentCanSee),
	)

	testCase(
		"No Perm",
		t,
		a.Modes("A"),
		false,
		ds,
	)
}
