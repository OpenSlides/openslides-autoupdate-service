package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

func TestAgendaModeA(t *testing.T) {
	var a collection.AgendaItem

	testCase(
		"No permission",
		t,
		a.Modes("A"),
		false,
		`---
		agenda_item/1/meeting_id: 30
		`,
		withPerms(30),
	)

	testCase(
		"manager",
		t,
		a.Modes("A"),
		true,
		`---
		agenda_item/1/meeting_id: 30
		`,
		withPerms(30, perm.AgendaItemCanManage),
	)

	testCase(
		"Can see internal with hidden",
		t,
		a.Modes("A"),
		false,
		`---
		agenda_item/1:
			meeting_id: 30
			is_hidden: true
		`,
		withPerms(30, perm.AgendaItemCanSeeInternal),
	)

	testCase(
		"Can see internal not hidden",
		t,
		a.Modes("A"),
		true,
		`---
		agenda_item/1:
			meeting_id: 30
			is_hidden: false
		`,
		withPerms(30, perm.AgendaItemCanSeeInternal),
	)

	testCase(
		"Can see with hidden and internal",
		t,
		a.Modes("A"),
		false,
		`---
		agenda_item/1:
			meeting_id: 30
			is_hidden: true
			is_internal: true
		`,
		withPerms(30, perm.AgendaItemCanSee),
	)

	testCase(
		"Can see not hidden but internal",
		t,
		a.Modes("A"),
		false,
		`---
		agenda_item/1:
			meeting_id: 30
			is_hidden: false
			is_internal: true
		`,
		withPerms(30, perm.AgendaItemCanSee),
	)

	testCase(
		"Can see with hidden but not internal",
		t,
		a.Modes("A"),
		false,
		`---
		agenda_item/1:
			meeting_id: 30
			is_hidden: true
			is_internal: false
		`,
		withPerms(30, perm.AgendaItemCanSee),
	)

	testCase(
		"Can see not hidden and not internal",
		t,
		a.Modes("A"),
		true,
		`---
		agenda_item/1:
			meeting_id: 30
			is_hidden: false
			is_internal: false
		`,
		withPerms(30, perm.AgendaItemCanSee),
	)
}

func TestAgendaModeB(t *testing.T) {
	var a collection.AgendaItem
	ds := `---
	agenda_item/1/meeting_id: 30
	`

	testCase(
		"Can see internal",
		t,
		a.Modes("B"),
		true,
		ds,
		withPerms(30, perm.AgendaItemCanSeeInternal),
	)

	testCase(
		"Can not see internal",
		t,
		a.Modes("B"),
		false,
		ds,
		withPerms(30),
	)
}

func TestAgendaModeC(t *testing.T) {
	var a collection.AgendaItem
	ds := `---
	agenda_item/1/meeting_id: 30
	`

	testCase(
		"Can see internal",
		t,
		a.Modes("C"),
		false,
		ds,
		withPerms(30, perm.AgendaItemCanSeeInternal),
	)

	testCase(
		"Can see",
		t,
		a.Modes("C"),
		false,
		ds,
		withPerms(30, perm.AgendaItemCanSee),
	)

	testCase(
		"Can manage",
		t,
		a.Modes("C"),
		true,
		ds,
		withPerms(30, perm.AgendaItemCanManage),
	)

	testCase(
		"No perm",
		t,
		a.Modes("C"),
		false,
		ds,
		withPerms(30),
	)
}
