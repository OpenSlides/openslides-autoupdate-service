package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

func TestMeetingUserModeA(t *testing.T) {
	f := collection.MeetingUser{}.Modes("A")

	testCase(
		"No perms",
		t,
		f,
		false,
		`---
		user/2/id: 2
		meeting_user/20:
			user_id: 2
			meeting_id: 30
		`,
		withRequestUser(1),
		withElementID(20),
	)

	testCase(
		"Request user",
		t,
		f,
		true,
		`---
		user/1/id: 1
		meeting_user/10:
			user_id: 1
			meeting_id: 30
		`,
		withRequestUser(1),
		withElementID(10),
	)

	testCase(
		"user.can_see",
		t,
		f,
		true,
		`---
		user/2/id: 2
		meeting_user/20:
			user_id: 2
			meeting_id: 30
		`,
		withRequestUser(1),
		withElementID(20),
		withPerms(30, perm.UserCanSee),
	)

	testCase(
		"Can manage users",
		t,
		f,
		false,
		`---
		user/2/id: 2
		user/1/organization_management_level: can_manage_users
		meeting_user/20:
			user_id: 2
			meeting_id: 30
		`,
		withRequestUser(1),
		withElementID(20),
	)

	testCase(
		"Committee Manager",
		t,
		f,
		false,
		`---
		user/2/committee_ids: [5]
		user/1:
			committee_management_ids: [5]
		committee/5/user_ids: [2]
		meeting_user/20:
			user_id: 2
			meeting_id: 30
		`,
		withRequestUser(1),
		withElementID(20),
	)
}

func TestMeetingUserModeB(t *testing.T) {
	f := collection.MeetingUser{}.Modes("B")

	testCase(
		"X == Y",
		t,
		f,
		true,
		`---
		meeting_user/10/user_id: 1
		`,
		withRequestUser(1),
		withElementID(10),
	)

	testCase(
		"X != Y",
		t,
		f,
		false,
		`---
		meeting_user/20/user_id: 2
		`,
		withRequestUser(1),
		withElementID(20),
	)
}

func TestMeetingUserModeC(t *testing.T) {
	f := collection.MeetingUser{}.Modes("C")

	testCase(
		"locked meeting, orga manager",
		t,
		f,
		true,
		`
		user/1/organization_management_level: can_manage_organization
		meeting/30:
			locked_from_inside: true
			enable_anonymous: false

		user/2/id: 2

		meeting_user/20:
			user_id: 2
			meeting_id: 30
		`,
		withElementID(20),
	)
}

func TestMeetingUserModeD(t *testing.T) {
	f := collection.MeetingUser{}.Modes("D")

	testCase(
		"No perms",
		t,
		f,
		false,
		`---
		user/2/id: 2
		meeting_user/20:
			user_id: 2
			meeting_id: 5
		`,
		withRequestUser(1),
		withElementID(20),
	)

	testCase(
		"OML can manage users",
		t,
		f,
		true,
		`---
		user/1/organization_management_level: can_manage_users
		meeting_user/20:
			user_id: 2
			meeting_id: 5
		`,
		withRequestUser(1),
		withElementID(20),
	)

	testCase(
		"user.can_manage in meeting",
		t,
		f,
		true,
		`---
		user/2/meeting_user_ids: [20]
		meeting_user/20:
			meeting_id: 5
			user_id: 2
		`,
		withRequestUser(1),
		withElementID(20),
		withPerms(5, perm.UserCanManage),
	)

	testCase(
		"user.can_manage not in meeting",
		t,
		f,
		false,
		`---
		user/2/meeting_user_ids: []
		meeting_user/20:
			user_id: 2
			meeting_id: 404
		`,
		withRequestUser(1),
		withElementID(20),
		withPerms(5, perm.UserCanManage),
	)
}

func TestMeetingUserModeE(t *testing.T) {
	mode := collection.MeetingUser{}.Modes("E")

	testCase(
		"Without perms",
		t,
		mode,
		false,
		`---
		user/2/id: 2
		meeting_user/20:
			user_id: 2
			meeting_id: 5
		`,
		withElementID(20),
	)

	testCase(
		"Without perms themselves",
		t,
		mode,
		true,
		`---
		user/1/id: 1
		meeting_user/20:
			user_id: 1
			meeting_id: 5
		`,
		withElementID(20),
	)

	testCase(
		"Can see",
		t,
		mode,
		true,
		`---
		user/2/id: 2
		meeting_user/20:
			user_id: 2
			meeting_id: 5
		`,
		withPerms(5, perm.UserCanSeeSensitiveData),
		withElementID(20),
	)
}
