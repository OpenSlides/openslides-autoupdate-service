package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
)

func TestListOfSpeakersModeA(t *testing.T) {
	var los collection.ListOfSpeakers
	ds := `---
	list_of_speakers/1/meeting_id: 1
	`

	testCase(
		"Can see internal",
		true,
		ds,
		withPerms(1, perm.ListOfSpeakersCanSee),
	).test(t, los.Modes("A"))

	testCase(
		"Can not see internal",
		false,
		ds,
	).test(t, los.Modes("A"))
}
