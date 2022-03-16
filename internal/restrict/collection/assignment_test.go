package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

func TestAssignmentModeA(t *testing.T) {
	var a collection.Assignment

	testCase(
		"Without perms",
		t,
		a.Modes("A"),
		false,
		`---
		assignment/1:
			meeting_id: 30
		`,
	)

	testCase(
		"Can see",
		t,
		a.Modes("A"),
		true,
		`---
		assignment/1:
			meeting_id: 30
		`,
		withPerms(30, perm.AssignmentCanSee),
	)

	testCase(
		"Can see agenda item",
		t,
		a.Modes("A"),
		true,
		`---
		assignment/1:
			agenda_item_id: 30
			meeting_id: 30
			list_of_speakers_id: 15

		agenda_item/30:
			meeting_id: 30

		list_of_speakers/15:
			id: 15
			meeting_id: 30
		`,
		withPerms(30, perm.AgendaItemCanSee),
	)

	testCase(
		"Can not see agenda item",
		t,
		a.Modes("A"),
		false,
		`---
		assignment/1:
			agenda_item_id: 30
			meeting_id: 30
			list_of_speakers_id: 15

		agenda_item/30:
			meeting_id: 30

		list_of_speakers/15:
			id: 15
			meeting_id: 30
		`,
	)

	testCase(
		"Can see agenda item but assignment has none",
		t,
		a.Modes("A"),
		false,
		`---
		assignment/1:
			meeting_id: 30
			
		agenda_item/30:
			meeting_id: 30
		`,
		withPerms(30, perm.AgendaItemCanSee),
	)
}

func TestAssignmentModeB(t *testing.T) {
	var a collection.Assignment

	testCase(
		"Without perms",
		t,
		a.Modes("B"),
		false,
		`---
		assignment/1:
			meeting_id: 30
			list_of_speakers_id: 404
		`,
	)

	testCase(
		"Can see",
		t,
		a.Modes("B"),
		true,
		`---
		assignment/1:
			meeting_id: 30
		`,
		withPerms(30, perm.AssignmentCanSee),
	)

	testCase(
		"Can see agenda item",
		t,
		a.Modes("B"),
		false,
		`---
		assignment/1:
			agenda_item_id: 30
			meeting_id: 30
			list_of_speakers_id: 404

		agenda_item/30:
			meeting_id: 30
			list_of_speakers_id: 404
		`,
		withPerms(30, perm.AgendaItemCanSee),
	)
}
