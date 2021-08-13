package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

func TestListOfSpeakersModeA(t *testing.T) {
	f := collection.ListOfSpeakers{}.Modes("A")

	testCase(
		"can see",
		t,
		f,
		true,
		"list_of_speakers/1/meeting_id: 1",
		withPerms(1, perm.ListOfSpeakersCanSee),
	)

	testCase(
		"no perm",
		t,
		f,
		false,
		"list_of_speakers/1/meeting_id: 1",
	)
}
