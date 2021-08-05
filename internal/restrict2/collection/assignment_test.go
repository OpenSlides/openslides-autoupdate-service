package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
)

func TestAssignmentModeA(t *testing.T) {
	var a collection.Assignment

	for _, tt := range []testData{
		testCase(
			"Without perms",
			false,
			`---
			assignment/1:
				meeting_id: 1
			`,
		),

		testCase(
			"Can see",
			true,
			`---
			assignment/1:
				meeting_id: 1
			`,
			withPerms(1, perm.AssignmentCanSee),
		),

		testCase(
			"Can see list of speakers",
			true,
			`---
			assignment/1:
				list_of_speakers_id: 15
				meeting_id: 1

			list_of_speakers/15:
				id: 15
				meeting_id: 1
			`,
			withPerms(1, perm.ListOfSpeakersCanSee),
		),

		testCase(
			"Can not see list of speakers",
			false,
			`---
			assignment/1:
				list_of_speakers_id: 15
				meeting_id: 1

			list_of_speakers/15:
				id: 15
				meeting_id: 1
			`,
		),

		testCase(
			"Can see list of speakers but assignment has no list",
			false,
			`---
			assignment/1:
				meeting_id: 1

			list_of_speakers/15: 
				id: 15
				meeting_id: 1
			`,
			withPerms(1, perm.ListOfSpeakersCanSee),
		),

		testCase(
			"Can see agenda item",
			true,
			`---
			assignment/1:
				agenda_item_id: 30
				meeting_id: 1

			agenda_item/30:
				meeting_id: 1
			`,
			withPerms(1, perm.AgendaItemCanSee),
		),

		testCase(
			"Can not see agenda item",
			false,
			`---
			assignment/1:
				agenda_item_id: 30
				meeting_id: 1

			agenda_item/30:
				meeting_id: 1
			`,
		),

		testCase(
			"Can see agenda item but assignment has none",
			false,
			`---
			assignment/1:
				meeting_id: 1

			agenda_item/30:
				meeting_id: 1
			`,
			withPerms(1, perm.AgendaItemCanSee),
		),
	} {
		tt.test(t, a.Modes("A"))
	}
}

func TestAssignmentModeB(t *testing.T) {
	var a collection.Assignment

	testCase(
		"Without perms",
		false,
		`---
		assignment/1:
			meeting_id: 1
		`,
	).test(t, a.Modes("B"))

	testCase(
		"Can see",
		true,
		`---
		assignment/1:
			meeting_id: 1
		`,
		withPerms(1, perm.AssignmentCanSee),
	).test(t, a.Modes("B"))

	testCase(
		"Can see list of speakers",
		false,
		`---
		assignment/1:
			list_of_speakers_id: 15
			meeting_id: 1

		list_of_speakers/15:
			id: 15
			meeting_id: 1
		`,
		withPerms(1, perm.ListOfSpeakersCanSee),
	).test(t, a.Modes("B"))

	testCase(
		"Can see agenda item",
		false,
		`---
		assignment/1:
			agenda_item_id: 30
			meeting_id: 1

		agenda_item/30:
			meeting_id: 1
		`,
		withPerms(1, perm.AgendaItemCanSee),
	).test(t, a.Modes("B"))
}
