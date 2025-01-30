package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

func TestStructureLevelListOfSpeakersModeA(t *testing.T) {
	f := collection.StructureLevelListOfSpeakers{}.Modes("A")

	testCase(
		"no perms",
		t,
		f,
		false,
		`---
		structure_level_list_of_speakers/1:
			list_of_speakers_id: 5
		list_of_speakers/5:
			meeting_id: 30
		`,
	)

	testCase(
		"Can see linked list_of_speakers",
		t,
		f,
		true,
		`---
		structure_level_list_of_speakers/1:
			list_of_speakers_id: 5
		list_of_speakers/5:
			meeting_id: 30
		`,
		withPerms(30, perm.ListOfSpeakersCanSee),
	)

	testCase(
		"Can not see linked list_of_speakers",
		t,
		f,
		false,
		`---
		structure_level_list_of_speakers/1:
			list_of_speakers_id: 5
		list_of_speakers/5:
			meeting_id: 30
		`,
		withPerms(30, perm.UserCanSee),
	)
}
