package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

func TestTopicModeA(t *testing.T) {
	f := collection.Topic{}.Modes("A")

	testCase(
		"no perm",
		t,
		f,
		false,
		`---
		topic/1/meeting_id: 1
		`,
	)

	testCase(
		"see perm",
		t,
		f,
		true,
		`---
		topic/1/meeting_id: 1
		`,
		withPerms(1, perm.AgendaItemCanSee),
	)

	testCase(
		"see list of speakers",
		t,
		f,
		true,
		`---
		topic/1:
			meeting_id: 1
			list_of_speakers_id: 3

		list_of_speakers/3/meeting_id: 1
		`,
		withPerms(1, perm.ListOfSpeakersCanSee),
	)

	testCase(
		"can not see list of speakers",
		t,
		f,
		false,
		`---
		topic/1:
			meeting_id: 1
			list_of_speakers_id: 3

		list_of_speakers/3/meeting_id: 1
		`,
	)

	testCase(
		"see agenda item",
		t,
		f,
		true,
		`---
		topic/1:
			meeting_id: 1
			agenda_item_id: 3

		agenda_item_id/3/meeting_id: 1
		`,
		withPerms(1, perm.AgendaItemCanSee),
	)

	testCase(
		"can not see agenda_item",
		t,
		f,
		false,
		`---
		topic/1:
			meeting_id: 1
			agenda_item_id: 3

		agenda_item_id/3/meeting_id: 1
		`,
	)
}

func TestTopicModeB(t *testing.T) {
	f := collection.Topic{}.Modes("B")

	testCase(
		"see perm",
		t,
		f,
		true,
		`---
		topic/1/meeting_id: 1
		`,
		withPerms(1, perm.AgendaItemCanSee),
	)

	testCase(
		"no perm",
		t,
		f,
		false,
		`---
		topic/1/meeting_id: 1
		`,
	)
}
