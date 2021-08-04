package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
)

func TestAgendaSee(t *testing.T) {
	var a collection.AgendaItem

	for _, tt := range []restrictTestData{
		{
			"No permission",
			``,
			[]perm.TPermission{},
			false,
		},

		{
			"agenda does not exist",
			``,
			[]perm.TPermission{
				perm.AgendaItemCanManage,
			},
			false,
		},

		{
			"manager",
			`
			agenda_item/1/meeting_id: 1
			`,
			[]perm.TPermission{
				perm.AgendaItemCanManage,
			},
			true,
		},

		{
			"Can see internal with hidden",
			`
			agenda_item/1:
				meeting_id: 1
				is_hidden: true
			`,
			[]perm.TPermission{
				perm.AgendaItemCanSeeInternal,
			},
			false,
		},

		{
			"Can see internal not hidden",
			`
			agenda_item/1:
				meeting_id: 1
				is_hidden: false
			`,
			[]perm.TPermission{
				perm.AgendaItemCanSeeInternal,
			},
			true,
		},

		{
			"Can see with hidden and internal",
			`
			agenda_item/1:
				meeting_id: 1
				is_hidden: true
				is_internal: true
			`,
			[]perm.TPermission{
				perm.AgendaItemCanSee,
			},
			false,
		},

		{
			"Can see not hidden but internal",
			`
			agenda_item/1:
				meeting_id: 1
				is_hidden: false
				is_internal: true
			`,
			[]perm.TPermission{
				perm.AgendaItemCanSee,
			},
			false,
		},

		{
			"Can see with hidden but not internal",
			`
			agenda_item/1:
				meeting_id: 1
				is_hidden: true
				is_internal: false
			`,
			[]perm.TPermission{
				perm.AgendaItemCanSee,
			},
			false,
		},

		{
			"Can see not hidden and not internal",
			`
			agenda_item/1:
				meeting_id: 1
				is_hidden: false
				is_internal: false
			`,
			[]perm.TPermission{
				perm.AgendaItemCanSee,
			},
			true,
		},
	} {
		tt.test(t, a.See)
	}
}

func TestAgendaModeA(t *testing.T) {
	var a collection.AgendaItem

	restrictTestData{
		"simple",
		``,
		nil,
		true,
	}.test(t, a.Modes("A"))
}

func TestAgendaModeB(t *testing.T) {
	var a collection.AgendaItem
	r := a.Modes("B")
	ds := `---
	agenda_item/1/meeting_id: 1
	`

	restrictTestData{
		"Can see internal",
		ds,
		[]perm.TPermission{perm.AgendaItemCanSeeInternal},
		true,
	}.test(t, r)

	restrictTestData{
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

	restrictTestData{
		"Can see internal",
		ds,
		[]perm.TPermission{perm.AgendaItemCanSeeInternal},
		false,
	}.test(t, r)

	restrictTestData{
		"Can see",
		ds,
		[]perm.TPermission{perm.AgendaItemCanSee},
		false,
	}.test(t, r)

	restrictTestData{
		"Can manage",
		ds,
		[]perm.TPermission{perm.AgendaItemCanManage},
		true,
	}.test(t, r)

	restrictTestData{
		"No perm",
		ds,
		nil,
		false,
	}.test(t, r)
}
