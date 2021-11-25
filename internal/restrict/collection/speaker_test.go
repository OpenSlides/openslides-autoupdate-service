package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

func TestSpeakerModeA(t *testing.T) {
	f := collection.Speaker{}.Modes("A")

	testCase(
		"no perms",
		t,
		f,
		false,
		`---
		speaker/1:
			list_of_speakers_id: 15

		list_of_speakers/15:
			id: 15
			meeting_id: 1
		`,
	)

	testCase(
		"Can see list of speakers",
		t,
		f,
		true,
		`---
		speaker/1:
			list_of_speakers_id: 15

		list_of_speakers/15:
			id: 15
			meeting_id: 1
			content_object_id: topic/5
		
		topic/5/meeting_id: 1
		`,
		withPerms(1, perm.ListOfSpeakersCanSee, perm.AgendaItemCanSee),
	)
}
