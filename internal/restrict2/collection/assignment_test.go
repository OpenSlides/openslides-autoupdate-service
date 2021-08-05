package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
)

func TestAssignmentModeA(t *testing.T) {
	var a collection.Assignment

	for _, tt := range []testData{
		{
			"Without perms",
			`---
			assignment/1:
				meeting_id: 1
			`,
			nil,
			false,
		},

		{
			"Can see",
			`---
			assignment/1:
				meeting_id: 1
			`,
			permList(perm.AssignmentCanSee),
			true,
		},

		{
			"Can see list of speakers",
			`---
			assignment/1:
				list_of_speakers_id: 15
				meeting_id: 1

			list_of_speakers/15:
				id: 15
				meeting_id: 1
			`,
			permList(perm.ListOfSpeakersCanSee),
			true,
		},

		{
			"Can not see list of speakers",
			`---
			assignment/1:
				list_of_speakers_id: 15
				meeting_id: 1

			list_of_speakers/15:
				id: 15
				meeting_id: 1
			`,
			nil,
			false,
		},

		{
			"Can see list of speakers but assignment has no list",
			`---
			assignment/1:
				meeting_id: 1

			list_of_speakers/15: 
				id: 15
				meeting_id: 1
			`,
			permList(perm.ListOfSpeakersCanSee),
			false,
		},

		{
			"Can see agenda item",
			`---
			assignment/1:
				agenda_item_id: 30
				meeting_id: 1

			agenda_item/30:
				meeting_id: 1
			`,
			permList(perm.AgendaItemCanSee),
			true,
		},

		{
			"Can not see agenda item",
			`---
			assignment/1:
				agenda_item_id: 30
				meeting_id: 1

			agenda_item/30:
				meeting_id: 1
			`,
			nil,
			false,
		},

		{
			"Can see agenda item but assignment has none",
			`---
			assignment/1:
				meeting_id: 1

			agenda_item/30:
				meeting_id: 1
			`,
			permList(perm.AgendaItemCanSee),
			false,
		},
	} {
		tt.test(t, a.Modes("A"))
	}
}

func TestAssignmentModeB(t *testing.T) {
	var a collection.Assignment

	testData{
		"Without perms",
		`---
		assignment/1:
			meeting_id: 1
		`,
		nil,
		false,
	}.test(t, a.Modes("B"))

	testData{
		"Can see",
		`---
		assignment/1:
			meeting_id: 1
		`,
		permList(perm.AssignmentCanSee),
		true,
	}.test(t, a.Modes("B"))

	testData{
		"Can see list of speakers",
		`---
		assignment/1:
			list_of_speakers_id: 15
			meeting_id: 1

		list_of_speakers/15:
			id: 15
			meeting_id: 1
		`,
		permList(perm.ListOfSpeakersCanSee),
		false,
	}.test(t, a.Modes("B"))

	testData{
		"Can see agenda item",
		`---
		assignment/1:
			agenda_item_id: 30
			meeting_id: 1

		agenda_item/30:
			meeting_id: 1
		`,
		permList(perm.AgendaItemCanSee),
		false,
	}.test(t, a.Modes("B"))
}
