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
		"meeting/30/id: 30",
	)
}

func TestMeetingModeB(t *testing.T) {
	var m collection.Meeting

	testCase(
		"No perms",
		t,
		m.Modes("B"),
		false,
		`---
		meeting/30:
			id: 1
			committee_id: 300
		`,
		withElementID(30),
	)

	testCase(
		"anonymous enabled",
		t,
		m.Modes("B"),
		true,
		`meeting/30/enable_anonymous: true`,
		withElementID(30),
	)

	testCase(
		"user in meeting",
		t,
		m.Modes("B"),
		true,
		`---
		meeting/30:
			group_ids: [7]
			committee_id: 2
		
		group/7/meeting_user_ids: [10]
		meeting_user/10/user_id: 1
		`,
		withElementID(30),
	)

	testCase(
		"CML can manage",
		t,
		m.Modes("B"),
		true,
		`---
		meeting/30/committee_id: 4
		user/1/committee_management_ids: [4]
		`,
		withElementID(30),
	)

	testCase(
		"CML can manage other committee",
		t,
		m.Modes("B"),
		false,
		`---
		meeting/30/committee_id: 4
		user/1/committee_management_ids: [8]
		`,
		withElementID(30),
	)

	testCase(
		"Template meeting",
		t,
		m.Modes("B"),
		false,
		`---
		meeting/30:
			template_for_organization_id: 1
			committee_id: 4
		`,
		withElementID(30),
	)

	testCase(
		"CML can manage other meeting template meeting",
		t,
		m.Modes("B"),
		true,
		`---
		meeting/30:
			committee_id: 4
			template_for_organization_id: 16
		user/1/committee_management_ids: [8]
		`,
		withElementID(30),
	)

	testCase(
		"CML can manage organization",
		t,
		m.Modes("B"),
		true,
		`---
		user/1/organization_management_level: can_manage_organization
		meeting/30/id: 30
		`,
		withElementID(30),
	)

	testCase(
		"Request with anonymous",
		t,
		m.Modes("B"),
		false,
		`meeting/30/id: 30`,
		withRequestUser(0),
		withElementID(30),
	)
}

func TestMeetingModeC(t *testing.T) {
	var m collection.Meeting

	testCase(
		"No perms",
		t,
		m.Modes("C"),
		false,
		`meeting/30/id: 30`,
		withElementID(30),
	)

	testCase(
		"See frontpage",
		t,
		m.Modes("C"),
		true,
		`meeting/30/id: 30`,
		withPerms(30, perm.MeetingCanSeeFrontpage),
		withElementID(30),
	)
}

func TestMeetingModeD(t *testing.T) {
	var m collection.Meeting

	testCase(
		"No perms",
		t,
		m.Modes("D"),
		false,
		`meeting/30/id: 30`,
		withElementID(30),
	)

	testCase(
		"See frontpage",
		t,
		m.Modes("D"),
		true,
		`meeting/30/id: 30`,
		withPerms(30, perm.MeetingCanSeeLivestream),
		withElementID(30),
	)
}
