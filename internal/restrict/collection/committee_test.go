package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
)

func TestCommitteeModeA(t *testing.T) {
	var c collection.Committee

	testCase(
		"No perms",
		t,
		c.Modes("A"),
		false,
		`committee/1/id: 1`,
	)

	testCase(
		"In committee/user_ids",
		t,
		c.Modes("A"),
		true,
		`---
		committee/1/user_ids: [1]
		`,
	)

	testCase(
		"OML can_manage_users",
		t,
		c.Modes("A"),
		true,
		`---
		committee/1/id: 1
		user/1/organization_management_level: can_manage_users
		`,
	)

	testCase(
		"CML can manage in parent",
		t,
		c.Modes("A"),
		true,
		`---
		committee/5/id: 5
		committee/6/all_child_ids: [5]
		user/1:
			committee_management_ids: [6]
		`,
		withElementID(5),
	)
}

func TestCommitteeModeB(t *testing.T) {
	var c collection.Committee

	testCase(
		"OML can_manage_users",
		t,
		c.Modes("B"),
		false,
		`---
		committee/5/id: 5
		user/1/organization_management_level: can_manage_users
		`,
	)

	testCase(
		"OML can_manage_organization",
		t,
		c.Modes("B"),
		true,
		`---
		committee/5/id: 5
		user/1/organization_management_level: can_manage_organization
		`,
	)

	testCase(
		"CML can manage",
		t,
		c.Modes("B"),
		true,
		`---
		committee/5/id: 5
		user/1:
			committee_management_ids: [5]
		`,
		withElementID(5),
	)

	testCase(
		"CML can manage in parent",
		t,
		c.Modes("B"),
		true,
		`---
		committee/5/id: 5
		committee/6/all_child_ids: [5]
		user/1:
			committee_management_ids: [6]
		`,
		withElementID(5),
	)

	testCase(
		"CML can manage in parent with two",
		t,
		c.Modes("B"),
		true,
		`---
		committee/1/all_child_ids: [3,4]
		committee/3/id: 3
		committee/4/id: 4
		user/4:
			committee_management_ids: [1]
		`,
		withElementID(3),
		withRequestUser(4),
	)

	testCase(
		"In committee",
		t,
		c.Modes("B"),
		false,
		`---
		committee/5/user_ids: [1]
		`,
		withElementID(5),
	)
}
