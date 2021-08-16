package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

func TestMeetingModeA(t *testing.T) {
	var m collection.Meeting

	testCase(
		"without perms",
		t,
		m.Modes("A"),
		true,
		"meeting/1/id: 1",
	)
}

func TestMeetingModeB(t *testing.T) {
	var m collection.Meeting

	testCase(
		"No perms",
		t,
		m.Modes("B"),
		false,
		`meeting/1/id: 1`,
	)

	testCase(
		"anonymous enabled",
		t,
		m.Modes("B"),
		true,
		`meeting/1/enable_anonymous: true`,
	)

	testCase(
		"user in meeting",
		t,
		m.Modes("B"),
		true,
		"meeting/1/user_ids: [1]",
	)

	testCase(
		"CML can manage",
		t,
		m.Modes("B"),
		true,
		`---
			meeting/1/committee_id: 4
			user/1/committee_$4_management_level: can_manage
			`,
	)
}

func TestMeetingModeC(t *testing.T) {
	var m collection.Meeting

	testCase(
		"No perms",
		t,
		m.Modes("C"),
		false,
		`meeting/1/id: 1`,
	)

	testCase(
		"See frontpage",
		t,
		m.Modes("C"),
		true,
		`meeting/1/id: 1`,
		withPerms(1, perm.MeetingCanSeeFrontpage),
	)
}

func TestMeetingModeD(t *testing.T) {
	var m collection.Meeting

	testCase(
		"No perms",
		t,
		m.Modes("D"),
		false,
		`meeting/1/id: 1`,
	)

	testCase(
		"See frontpage",
		t,
		m.Modes("D"),
		true,
		`meeting/1/id: 1`,
		withPerms(1, perm.MeetingCanSeeLivestream),
	)
}
