package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
)

func TestAgendaModeA(t *testing.T) {
	var a collection.AgendaItem

	for _, tt := range []testData{
		testCase(
			"No permission",
			false,
			`---
			agenda_item/1/meeting_id: 1
			`,
		),
		testCase(
			"manager",
			true,
			`---
			agenda_item/1/meeting_id: 1
			`,
			withPerms(1, perm.AgendaItemCanManage),
		),

		testCase(
			"Can see internal with hidden",
			false,
			`---
			agenda_item/1:
				meeting_id: 1
				is_hidden: true
			`,
			withPerms(1, perm.AgendaItemCanSeeInternal),
		),

		testCase(
			"Can see internal not hidden",
			true,
			`---
			agenda_item/1:
				meeting_id: 1
				is_hidden: false
			`,
			withPerms(1, perm.AgendaItemCanSeeInternal),
		),

		testCase(
			"Can see with hidden and internal",
			false,
			`---
			agenda_item/1:
				meeting_id: 1
				is_hidden: true
				is_internal: true
			`,
			withPerms(1, perm.AgendaItemCanSee),
		),

		testCase(
			"Can see not hidden but internal",
			false,
			`---
			agenda_item/1:
				meeting_id: 1
				is_hidden: false
				is_internal: true
			`,
			withPerms(1, perm.AgendaItemCanSee),
		),

		testCase(
			"Can see with hidden but not internal",
			false,
			`---
			agenda_item/1:
				meeting_id: 1
				is_hidden: true
				is_internal: false
			`,
			withPerms(1, perm.AgendaItemCanSee),
		),

		testCase(
			"Can see not hidden and not internal",
			true,
			`---
			agenda_item/1:
				meeting_id: 1
				is_hidden: false
				is_internal: false
			`,
			withPerms(1, perm.AgendaItemCanSee),
		),
	} {
		tt.test(t, a.Modes("A"))
	}
}

func TestAgendaModeB(t *testing.T) {
	var a collection.AgendaItem
	r := a.Modes("B")
	ds := `---
	agenda_item/1/meeting_id: 1
	`

	testCase(
		"Can see internal",
		true,
		ds,
		withPerms(1, perm.AgendaItemCanSeeInternal),
	).test(t, r)

	testCase(
		"Can not see internal",
		false,
		ds,
	).test(t, r)
}

func TestAgendaModeC(t *testing.T) {
	var a collection.AgendaItem
	r := a.Modes("C")
	ds := `---
	agenda_item/1/meeting_id: 1
	`

	testCase(
		"Can see internal",
		false,
		ds,
		withPerms(1, perm.AgendaItemCanSeeInternal),
	).test(t, r)

	testCase(
		"Can see",
		false,
		ds,
		withPerms(1, perm.AgendaItemCanSee),
	).test(t, r)

	testCase(
		"Can manage",
		true,
		ds,
		withPerms(1, perm.AgendaItemCanManage),
	).test(t, r)

	testCase(
		"No perm",
		false,
		ds,
	).test(t, r)
}
