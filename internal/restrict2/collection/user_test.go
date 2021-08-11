package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
)

func TestUserModeA(t *testing.T) {
	var u collection.User

	testCase(
		"No perms",
		t,
		u.Modes("A"),
		false,
		`user/2/id: 2`,
		withRequestUser(1),
		withElementID(2),
	)

	testCase(
		"Request user",
		t,
		u.Modes("A"),
		true,
		`user/2/id: 2`,
		withRequestUser(1),
		withElementID(1),
	)

	testCase(
		"Can manage users",
		t,
		u.Modes("A"),
		true,
		`---
		user/2/id: 2
		user/1/organization_management_level: can_manage_users
		`,
		withRequestUser(1),
		withElementID(2),
	)

	testCase(
		"Committee Manager",
		t,
		u.Modes("A"),
		true,
		`---
		user/2/committee_ids: [5]
		user/1:
			committee_$_management_level: ["5"]
			committee_$5_management_level: can_manage
		committee/5/user_ids: [2]
		`,
		withRequestUser(1),
		withElementID(2),
	)

	testCase(
		"Committee Manager user not in it",
		t,
		u.Modes("A"),
		false,
		`---
		user/2/committee_ids: [5]
		user/1:
			committee_$_management_level: ["5"]
			committee_$5_management_level: can_manage
		committee/5/user_ids: []
		`,
		withRequestUser(1),
		withElementID(2),
	)

	testCase(
		"user.can_see in meeting",
		t,
		u.Modes("A"),
		true,
		`---
		user/2/group_$_ids: ["5"]
		`,
		withRequestUser(1),
		withElementID(2),
		withPerms(5, perm.UserCanSee),
	)

	testCase(
		"user.can_see not in meeting",
		t,
		u.Modes("A"),
		false,
		`---
		user/2/group_$_ids: []
		`,
		withRequestUser(1),
		withElementID(2),
		withPerms(5, perm.UserCanSee),
	)

	testCase(
		"committee can manage",
		t,
		u.Modes("A"),
		true,
		`---
		user/2/group_$_ids: ["5"]
		meeting/5/committee_id: 7
		user/1:
			committee_$_management_level: ["7"]
			committee_$7_management_level: can_manage
		committee/7/id: 7
		`,
		withRequestUser(1),
		withElementID(2),
	)

	testCase(
		"committee can manage user not in meeting",
		t,
		u.Modes("A"),
		false,
		`---
		user/2/group_$_ids: []
		meeting/5/committee_id: 7
		user/1:
			committee_$_management_level: ["7"]
			committee_$7_management_level: can_manage
		committee/7/id: 7
		`,
		withRequestUser(1),
		withElementID(2),
	)

	testCase(
		"Vote delegated to",
		t,
		u.Modes("A"),
		true,
		`---
		user/1:
			vote_delegated_$_to_id: ["3"]
			vote_delegated_$3_to_id: 2
		user/2/id: 2
		`,
		withRequestUser(1),
		withElementID(2),
	)

	testCase(
		"Vote delegated from",
		t,
		u.Modes("A"),
		true,
		`---
		user/1:
			vote_delegations_$_from_ids: ["3"]
			vote_delegations_$3_from_ids: [2]
		user/2/id: 2
		`,
		withRequestUser(1),
		withElementID(2),
	)
}
