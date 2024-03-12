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
		structure_level_list_of_speakers/1/meeting_id: 30
		`,
	)

	testCase(
		"list_of_speakers.can_see",
		t,
		f,
		true,
		`---
		structure_level_list_of_speakers/1/meeting_id: 30
		`,
		withPerms(30, perm.ListOfSpeakersCanSee),
	)

	testCase(
		"user.can_see",
		t,
		f,
		false,
		`---
		structure_level_list_of_speakers/1/meeting_id: 30
		`,
		withPerms(30, perm.UserCanSee),
	)
}
