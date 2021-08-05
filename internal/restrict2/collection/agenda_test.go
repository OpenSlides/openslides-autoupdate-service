package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
)

func TestAgendaModeA(t *testing.T) {
	var a collection.AgendaItem

	for _, tt := range []testData{
		{
			"No permission",
			`---
			agenda_item/1/meeting_id: 1
			`,
			nil,
			false,
		},

		{
			"manager",
			`---
			agenda_item/1/meeting_id: 1
			`,
			permList(perm.AgendaItemCanManage),
			true,
		},

		{
			"Can see internal with hidden",
			`---
			agenda_item/1:
				meeting_id: 1
				is_hidden: true
			`,
			permList(perm.AgendaItemCanSeeInternal),
			false,
		},

		{
			"Can see internal not hidden",
			`---
			agenda_item/1:
				meeting_id: 1
				is_hidden: false
			`,
			permList(perm.AgendaItemCanSeeInternal),
			true,
		},

		{
			"Can see with hidden and internal",
			`---
			agenda_item/1:
				meeting_id: 1
				is_hidden: true
				is_internal: true
			`,
			permList(perm.AgendaItemCanSee),
			false,
		},

		{
			"Can see not hidden but internal",
			`---
			agenda_item/1:
				meeting_id: 1
				is_hidden: false
				is_internal: true
			`,
			permList(perm.AgendaItemCanSee),
			false,
		},

		{
			"Can see with hidden but not internal",
			`---
			agenda_item/1:
				meeting_id: 1
				is_hidden: true
				is_internal: false
			`,
			permList(perm.AgendaItemCanSee),
			false,
		},

		{
			"Can see not hidden and not internal",
			`---
			agenda_item/1:
				meeting_id: 1
				is_hidden: false
				is_internal: false
			`,
			permList(perm.AgendaItemCanSee),
			true,
		},
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

	testData{
		"Can see internal",
		ds,
		permList(perm.AgendaItemCanSeeInternal),
		true,
	}.test(t, r)

	testData{
		"Can not see internal",
		ds,
		nil,
		false,
	}.test(t, r)
}

func TestAgendaModeC(t *testing.T) {
	var a collection.AgendaItem
	r := a.Modes("C")
	ds := `---
	agenda_item/1/meeting_id: 1
	`

	testData{
		"Can see internal",
		ds,
		permList(perm.AgendaItemCanSeeInternal),
		false,
	}.test(t, r)

	testData{
		"Can see",
		ds,
		permList(perm.AgendaItemCanSee),
		false,
	}.test(t, r)

	testData{
		"Can manage",
		ds,
		permList(perm.AgendaItemCanManage),
		true,
	}.test(t, r)

	testData{
		"No perm",
		ds,
		nil,
		false,
	}.test(t, r)
}
