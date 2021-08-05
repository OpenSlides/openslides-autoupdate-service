package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
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

	testData{
		"Can see all assignments",
		ds,
		permList(perm.AssignmentCanSee),
		true,
	}.test(t, a.Modes("A"))

	testData{
		"Can only see the linked assignment",
		ds,
		permList(perm.ListOfSpeakersCanSee),
		true,
	}.test(t, a.Modes("A"))
}
