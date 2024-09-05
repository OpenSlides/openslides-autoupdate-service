package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
)

func TestOrganizationModeA(t *testing.T) {
	f := collection.Organization{}.Modes("A")

	testCase(
		"no perms",
		t,
		f,
		true,
		``,
	)
}

func TestOrganizationModeB(t *testing.T) {
	f := collection.Organization{}.Modes("B")

	testCase(
		"anonymous",
		t,
		f,
		false,
		``,
		withRequestUser(0),
	)

	testCase(
		"logged in",
		t,
		f,
		true,
		``,
		withRequestUser(1),
	)
}

func TestOrganizationModeC(t *testing.T) {
	f := collection.Organization{}.Modes("C")

	testCase(
		"anonymous",
		t,
		f,
		false,
		``,
		withRequestUser(0),
	)

	testCase(
		"logged in",
		t,
		f,
		false,
		``,
		withRequestUser(1),
	)

	testCase(
		"OML can manage users",
		t,
		f,
		true,
		`---
		user/1/organization_management_level: can_manage_users
		`,
		withRequestUser(1),
	)
}

func TestOrganizationModeE(t *testing.T) {
	f := collection.Organization{}.Modes("E")

	testCase(
		"anonymous",
		t,
		f,
		false,
		``,
		withRequestUser(0),
	)

	testCase(
		"logged in",
		t,
		f,
		false,
		``,
		withRequestUser(1),
	)

	testCase(
		"normal user in meeting",
		t,
		f,
		false,
		`---
		meeting/7:
			admin_group_id: 8

		user/1/meeting_user_ids: [10]
		group/6/id: 6
		meeting_user/10:
			group_ids: [6]
			meeting_id: 7
		`,
		withRequestUser(1),
	)

	testCase(
		"meeting admin",
		t,
		f,
		true,
		`---
		meeting/7:
			admin_group_id: 8

		user/1/meeting_user_ids: [10]
		group/8/admin_group_for_meeting_id: 7
		meeting_user/10:
			group_ids: [8]
			meeting_id: 7
		`,
		withRequestUser(1),
	)
}
