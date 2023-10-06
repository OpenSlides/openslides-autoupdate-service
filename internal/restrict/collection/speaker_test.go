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
			meeting_id: 30

		list_of_speakers/15:
			id: 15
			meeting_id: 30
		`,
	)

	testCase(
		"Has can see",
		t,
		f,
		true,
		`---
		speaker/1:
			list_of_speakers_id: 15
			meeting_id: 30

		list_of_speakers/15:
			id: 15
			meeting_id: 30
			content_object_id: topic/5
		
		topic/5/meeting_id: 30
		`,
		withPerms(30, perm.ListOfSpeakersCanSee),
	)

	testCase(
		"Has can be speaker other user",
		t,
		f,
		false,
		`---
		speaker/1:
			list_of_speakers_id: 15
			meeting_id: 30
			meeting_user_id: 4040
		
		meeting_user/4040/user_id: 404

		list_of_speakers/15:
			id: 15
			meeting_id: 30
			content_object_id: topic/5
		
		topic/5/meeting_id: 30
		`,
		withPerms(30, perm.ListOfSpeakersCanBeSpeaker),
	)

	testCase(
		"Has can be speaker see him self",
		t,
		f,
		true,
		`---
		speaker/1:
			list_of_speakers_id: 15
			meeting_user_id: 10
			meeting_id: 30

		meeting_user/10/user_id: 1

		list_of_speakers/15:
			id: 15
			meeting_id: 30
			content_object_id: topic/5
		
		topic/5/meeting_id: 30
		`,
		withPerms(30, perm.ListOfSpeakersCanBeSpeaker),
	)
}
