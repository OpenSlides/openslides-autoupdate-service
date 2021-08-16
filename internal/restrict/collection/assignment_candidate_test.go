package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

func TestAssignmentCandidateModeA(t *testing.T) {
	var a collection.AssignmentCandidate
	ds := `---
	assignment_candidate/1/assignment_id: 7
	
	assignment/7:
		list_of_speakers_id: 9
		meeting_id: 1

	list_of_speakers/9/meeting_id: 1
	`

	testCase(
		"Can see all assignments",
		t,
		a.Modes("A"),
		true,
		ds,
		withPerms(1, perm.AssignmentCanSee),
	)

	testCase(
		"Can only see the linked assignment",
		t,
		a.Modes("A"),
		true,
		ds,
		withPerms(1, perm.ListOfSpeakersCanSee),
	)

	testCase(
		"No Perm",
		t,
		a.Modes("A"),
		false,
		ds,
	)
}
