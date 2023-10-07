package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
)

func TestListOfSpeakersModeA(t *testing.T) {
	f := collection.ListOfSpeakers{}.Modes("A")

	testCase(
		"no perm",
		t,
		f,
		false,
		`---
		list_of_speakers/1: 
			meeting_id: 30
			content_object_id: topic/5

		topic/5/meeting_id: 30
		`,
	)

	testCase(
		"can see",
		t,
		f,
		true,
		`---
		list_of_speakers/1: 
			meeting_id: 30
			content_object_id: topic/5

		topic/5/meeting_id: 30
		`,
		withPerms(30, perm.ListOfSpeakersCanSee),
	)

	testCase(
		"see content_object",
		t,
		f,
		false,
		`---
		list_of_speakers/1:
			meeting_id: 30
			content_object_id: topic/5
		
		topic/5/meeting_id: 30
		`,
		withPerms(30, perm.AgendaItemCanSee),
	)

	testCase(
		"see content_object and can see",
		t,
		f,
		true,
		`---
		list_of_speakers/1:
			meeting_id: 30
			content_object_id: topic/5
		
		topic/5/meeting_id: 30
		`,
		withPerms(30, perm.ListOfSpeakersCanSee, perm.AgendaItemCanSee),
	)

	testCase(
		"can see",
		t,
		f,
		true,
		`---
		list_of_speakers/1: 
			meeting_id: 30
			content_object_id: topic/5

		topic/5/meeting_id: 30
		`,
		withPerms(30, perm.ListOfSpeakersCanBeSpeaker),
	)
}
