package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

func TestMeetingModeA(t *testing.T) {
	m := collection.Meeting{}.Modes("A")

	testCase(
		"locked meeting, superadmin",
		t,
		m,
		true,
		`
		meeting/30/locked_from_inside: true
		`,
		withElementID(30),
	)

	testCase(
		"No permission",
		t,
		m,
		true,
		`
		`,
		withElementID(30),
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
		"Public Access enabled",
		t,
		m.Modes("B"),
		true,
		`---
		meeting/30/enable_anonymous: true
		organization/1/enable_anonymous: true
		`,
		withElementID(30),
	)

	testCase(
		"Public access enabled only in organization",
		t,
		m.Modes("B"),
		false,
		`---
		organization/1/enable_anonymous: true
		meeting/30:
			committee_id: 3
		`,
		withElementID(30),
	)

	testCase(
		"Public Access enabled only in meeting",
		t,
		m.Modes("B"),
		false,
		`---
		meeting/30:
			enable_anonymous: true
			committee_id: 3
		`,
		withElementID(30),
	)

	testCase(
		"Public access enabled, as locked in user that was locked out",
		t,
		m.Modes("B"),
		false,
		`---
		organization/1/enable_anonymous: true
		meeting/30:
			enable_anonymous: true
			group_ids: [7]

		group/7/meeting_user_ids: [10]
		meeting_user/10:
			user_id: 1
			locked_out: true
			meeting_id: 30

		user/1/meeting_user_ids: [10]
		`,
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
		meeting_user/10:
			user_id: 1
			meeting_id: 30
		user/1/meeting_user_ids: [10]
		`,
		withElementID(30),
	)

	testCase(
		"user in meeting but locked out",
		t,
		m.Modes("B"),
		false,
		`---
		meeting/30:
			group_ids: [7]
			committee_id: 2

		group/7/meeting_user_ids: [10]
		meeting_user/10:
			user_id: 1
			locked_out: true
			meeting_id: 30

		user/1/meeting_user_ids: [10]
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
		"CML can manage, but locked out",
		t,
		m.Modes("B"),
		false,
		`---
		meeting/30:
			committee_id: 4
			group_ids: [7]

		user/1:
			committee_management_ids: [4]
			meeting_user_ids: [10]

		group/7/meeting_user_ids: [10]
		meeting_user/10:
			user_id: 1
			locked_out: true
			meeting_id: 30
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
		"OML can manage organization",
		t,
		m.Modes("B"),
		true,
		`---
		user/1/organization_management_level: can_manage_organization
		meeting/30:
			committee_id: 4
		`,
		withElementID(30),
	)

	testCase(
		"CML can manage organization, but locked out",
		t,
		m.Modes("B"),
		false,
		`---
		organization/1/enable_anonymous: true
		meeting/30:
			enable_anonymous: true
			group_ids: [7]

		group/7/meeting_user_ids: [10]
		meeting_user/10:
			user_id: 1
			locked_out: true
			meeting_id: 30

		user/1:
			meeting_user_ids: [10]
			organization_management_level: can_manage_organization
		`,
		withElementID(30),
	)

	testCase(
		"Request from public access",
		t,
		m.Modes("B"),
		false,
		`meeting/30/id: 30`,
		withRequestUser(0),
		withElementID(30),
	)

	testCase(
		"locked meeting, orga admin",
		t,
		m.Modes("B"),
		false,
		`
		user/1/organization_management_level: can_manage_organization
		meeting/30/locked_from_inside: true
		`,
		withElementID(30),
	)

	testCase(
		"locked meeting, superadmin",
		t,
		m.Modes("B"),
		false,
		`
		user/1/organization_management_level: superadmin
		meeting/30/locked_from_inside: true
		`,
		withElementID(30),
	)

	testCase(
		"locked meeting, anonymous enabled",
		t,
		m.Modes("B"),
		false,
		`
		organization/1/enable_anonymous: true
		meeting/30:
			locked_from_inside: true
			enable_anonymous: true

		`,
		withElementID(30),
		withRequestUser(0),
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

func TestMeetingModeE(t *testing.T) {
	m := collection.Meeting{}.Modes("E")

	testCase(
		"locked meeting, superadmin",
		t,
		m,
		true,
		`
		user/1/organization_management_level: superadmin
		meeting/30/locked_from_inside: true
		`,
		withElementID(30),
	)

	testCase(
		"locked meeting, orga admin",
		t,
		m,
		false,
		`
		user/1/organization_management_level: can_manage_organization
		meeting/30/locked_from_inside: true
		`,
		withElementID(30),
	)
}

func TestMeetingModeF(t *testing.T) {
	m := collection.Meeting{}.Modes("F")

	testCase(
		"locked meeting, orga manager",
		t,
		m,
		true,
		`
		user/1/organization_management_level: can_manage_organization
		meeting/30/locked_from_inside: true
		`,
		withElementID(30),
	)

	testCase(
		"locked meeting, user manager",
		t,
		m,
		false,
		`
		user/1/organization_management_level: can_manage_users
		meeting/30/locked_from_inside: true
		`,
		withElementID(30),
	)
}
