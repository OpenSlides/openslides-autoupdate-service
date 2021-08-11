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
		t,
		los.Modes("A"),
		true,
		ds,
		withPerms(1, perm.ListOfSpeakersCanSee),
	)

	testCase(
		"Can not see internal",
		t,
		los.Modes("A"),
		false,
		ds,
	)
}
