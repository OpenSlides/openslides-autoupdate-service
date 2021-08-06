package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
)

func TestMeetingModeA(t *testing.T) {
	var m collection.Meeting

	testCase("without perms", true, "meeting/1/id: 1").test(t, m.Modes("A"))
}

func TestMeetingModeB(t *testing.T) {
	var m collection.Meeting

	for _, tt := range []testData{
		testCase(
			"No perms",
			false,
			`meeting/1/id: 1`,
		),

		testCase(
			"anonymous enabled",
			true,
			`meeting/1/enable_anonymous: true`,
		),

		testCase(
			"user in meeting",
			true,
			"meeting/1/user_ids: [1]",
		),

		testCase(
			"CML can manage",
			true,
			`---
			meeting/1/committee_id: 4
			user/1/committee_$4_management_level: can_manage
			`,
		),
	} {
		tt.test(t, m.Modes("B"))
	}
}

func TestMeetingModeC(t *testing.T) {
	var m collection.Meeting

	testCase(
		"No perms",
		false,
		`meeting/1/id: 1`,
	).test(t, m.Modes("C"))

	testCase(
		"See frontpage",
		true,
		`meeting/1/id: 1`,
		withPerms(1, perm.MeetingCanSeeFrontpage),
	).test(t, m.Modes("C"))
}

func TestMeetingModeD(t *testing.T) {
	var m collection.Meeting

	testCase(
		"No perms",
		false,
		`meeting/1/id: 1`,
	).test(t, m.Modes("D"))

	testCase(
		"See frontpage",
		true,
		`meeting/1/id: 1`,
		withPerms(1, perm.MeetingCanSeeLivestream),
	).test(t, m.Modes("D"))
}
