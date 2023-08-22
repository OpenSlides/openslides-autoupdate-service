package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

func TestUserModeA(t *testing.T) {
	f := collection.User{}.Modes("A")

	testCase(
		"No perms",
		t,
		f,
		false,
		`user/2/id: 2`,
		withRequestUser(1),
		withElementID(2),
	)

	testCase(
		"With anonymous",
		t,
		f,
		false,
		`user/2/id: 2`,
		withRequestUser(0),
		withElementID(2),
	)

	testCase(
		"Request user",
		t,
		f,
		true,
		`user/2/id: 2`,
		withRequestUser(1),
		withElementID(1),
	)

	testCase(
		"Can manage users",
		t,
		f,
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
		f,
		true,
		`---
		user/2/committee_ids: [5]
		user/1:
			committee_management_ids: [5]
		committee/5/user_ids: [2]
		`,
		withRequestUser(1),
		withElementID(2),
	)

	testCase(
		"Committee Manager user not in it",
		t,
		f,
		false,
		`---
		user/2/committee_ids: [5]
		user/1:
			committee_management_ids: [5]
		committee/5/user_ids: []
		`,
		withRequestUser(1),
		withElementID(2),
	)

	testCase(
		"user.can_see in meeting",
		t,
		f,
		true,
		`---
		user/2/meeting_user_ids: [20]
		meeting_user/20/meeting_id: 5
		`,
		withRequestUser(1),
		withElementID(2),
		withPerms(5, perm.UserCanSee),
	)

	testCase(
		"user.can_see not in meeting",
		t,
		f,
		false,
		`---
		user/2/meeting_user_ids: []
		`,
		withRequestUser(1),
		withElementID(2),
		withPerms(5, perm.UserCanSee),
	)

	testCase(
		"committee can manage",
		t,
		f,
		true,
		`---
		user/1:
			committee_management_ids: [7]
		committee/7/user_ids: [2]

		meeting/5/committee_id: 7
		`,
		withRequestUser(1),
		withElementID(2),
	)

	testCase(
		"committee can manage user not in meeting",
		t,
		f,
		false,
		`---
		user/2/meeting_user_ids: []
		meeting/5/committee_id: 7
		user/1:
			committee_management_ids: [7]
		committee/7/id: 7
		`,
		withRequestUser(1),
		withElementID(2),
	)

	testCase(
		"Vote delegated to",
		t,
		f,
		true,
		`---
		user/1/meeting_user_ids: [10]
		user/2/id: 2

		meeting_user:
			10:
				vote_delegated_to_id: 20
				user_id: 1
			20:
				user_id: 2
		`,
		withRequestUser(1),
		withElementID(2),
	)

	testCase(
		"Vote delegated from",
		t,
		f,
		true,
		`---
		user/1/meeting_user_ids: [10]
		user/2/id: 2

		meeting_user:
			10:
				vote_delegations_from_ids: [20]
				user_id: 1
			20:
				user_id: 2
		`,
		withRequestUser(1),
		withElementID(2),
	)

	testCase(
		"motion submitter",
		t,
		f,
		true,
		`---
		user/1/meeting_user_ids: [10]
		user/2/meeting_user_ids: [20]

		meeting_user/10:
			meeting_id: 30
		meeting_user/20:
			motion_submitter_ids: [4]
			meeting_id: 30
		
		motion_submitter/4:
			motion_id: 7
		
		motion/7:
			meeting_id: 30
			state_id: 5
		
		motion_state/5/id: 5
		`,
		withRequestUser(1),
		withElementID(2),
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"motion supporter",
		t,
		f,
		true,
		`---
		user/1/meeting_user_ids: [10]
		user/2/meeting_user_ids: [20]

		meeting_user/10:
			meeting_id: 30
		meeting_user/20:
			supported_motion_ids: [7]
			meeting_id: 30
		
		motion/7:
			meeting_id: 30
			state_id: 5
		
		motion_state/5/id: 5
		`,
		withRequestUser(1),
		withElementID(2),
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"assignment candidate",
		t,
		f,
		true,
		`---
		user/1/meeting_user_ids: [10]
		user/2/meeting_user_ids: [20]

		meeting_user/10:
			meeting_id: 30
		meeting_user/20:
			assignment_candidate_ids: [4]
			meeting_id: 30
		
		assignment_candidate/4/assignment_id: 5
		assignment/5/meeting_id: 30
		`,
		withRequestUser(1),
		withElementID(2),
		withPerms(30, perm.AssignmentCanSee),
	)

	testCase(
		"speaker",
		t,
		f,
		true,
		`---
		user/1/meeting_user_ids: [10]
		user/2/meeting_user_ids: [20]

		meeting_user/10:
			meeting_id: 30
		meeting_user/20:
			speaker_ids: [4]
			meeting_id: 30
		
		speaker/4:
			list_of_speakers_id: 5
			meeting_id: 30

		list_of_speakers/5:
			meeting_id: 30
			content_object_id: topic/10

		topic/10/meeting_id: 30
		`,
		withRequestUser(1),
		withElementID(2),
		withPerms(30, perm.ListOfSpeakersCanSee, perm.AgendaItemCanSee),
	)

	testCase(
		"vote delegated ids",
		t,
		f,
		true,
		`---
		user/1/meeting_user_ids: [10]
		user/2/meeting_user_ids: [20]

		meeting_user/10:
			meeting_id: 30
		meeting_user/20:
			vote_delegations_from_ids: [4]
			meeting_id: 30
		
		vote/4/option_id: 5
		option/5/poll_id: 6
		poll/6:
			state: published
			meeting_id: 30

		meeting/30/enable_anonymous: true
		`,
		withRequestUser(1),
		withElementID(2),
	)

	testCase(
		"chat messages",
		t,
		f,
		true,
		`---
		user/1/meeting_user_ids: [10]
		user/2/meeting_user_ids: [20]

		meeting_user/10:
			meeting_id: 30
		meeting_user/20:
			chat_message_ids: [4]
			meeting_id: 30

		meeting_user/10/group_ids: [5]
		
		meeting/30/id: 30
		
		chat_message/4:
			meeting_user_id: 20
			chat_group_id: 3
		
		chat_group/3:
			read_group_ids: [5]
			meeting_id: 30

		group/5/id: 5
		`,
		withRequestUser(1),
		withElementID(2),
	)
}

func TestUserModeD(t *testing.T) {
	var u collection.User

	testCase(
		"No perms",
		t,
		u.Modes("D"),
		false,
		`user/2/id: 2`,
		withRequestUser(1),
		withElementID(2),
	)

	testCase(
		"OML can manage users",
		t,
		u.Modes("D"),
		true,
		`---
		user/1/organization_management_level: can_manage_users
		`,
		withRequestUser(1),
		withElementID(2),
	)

	testCase(
		"user.can_manage in meeting",
		t,
		u.Modes("D"),
		true,
		`---
		user/2/meeting_user_ids: [20]
		meeting_user/20:
			meeting_id: 5
		`,
		withRequestUser(1),
		withElementID(2),
		withPerms(5, perm.UserCanManage),
	)

	testCase(
		"user.can_manage not in meeting",
		t,
		u.Modes("D"),
		false,
		`---
		user/2/meeting_user_ids: []
		`,
		withRequestUser(1),
		withElementID(2),
		withPerms(5, perm.UserCanManage),
	)
}

func TestUserModeE(t *testing.T) {
	var u collection.User

	testCase(
		"No perms",
		t,
		u.Modes("E"),
		false,
		`user/2/id: 2`,
		withRequestUser(1),
		withElementID(2),
	)

	testCase(
		"With anonymous",
		t,
		u.Modes("E"),
		false,
		`user/2/id: 2`,
		withRequestUser(0),
		withElementID(2),
	)

	testCase(
		"OML can manage users",
		t,
		u.Modes("E"),
		true,
		`---
		user/1/organization_management_level: can_manage_users
		`,
		withRequestUser(1),
		withElementID(2),
	)

	testCase(
		"Committee Manager",
		t,
		u.Modes("E"),
		true,
		`---
		user/2/committee_ids: [5]
		user/1:
			committee_management_ids: [5]
		committee/5/user_ids: [2]
		`,
		withRequestUser(1),
		withElementID(2),
	)

	testCase(
		"Committee Manager user not in it",
		t,
		u.Modes("E"),
		false,
		`---
		user/2/committee_ids: [5]
		user/1:
			committee_management_ids: [5]
		committee/5/user_ids: []
		`,
		withRequestUser(1),
		withElementID(2),
	)

	testCase(
		"X == Y",
		t,
		u.Modes("E"),
		true,
		``,
		withRequestUser(1),
		withElementID(1),
	)

	testCase(
		"user.can_manage in meeting",
		t,
		u.Modes("E"),
		true,
		`---
		user/2/meeting_user_ids: [20]
		meeting_user/20/meeting_id: 5
		`,
		withRequestUser(1),
		withElementID(2),
		withPerms(5, perm.UserCanManage),
	)

	testCase(
		"user.can_manage not in meeting",
		t,
		u.Modes("E"),
		false,
		`---
		user/2/meeting_user_ids: []
		`,
		withRequestUser(1),
		withElementID(2),
		withPerms(5, perm.UserCanManage),
	)
}

func TestUserModeF(t *testing.T) {
	var u collection.User
	mode := u.Modes("F")

	testCase(
		"No perms",
		t,
		mode,
		false,
		`user/2/id: 2`,
		withRequestUser(1),
		withElementID(2),
	)

	testCase(
		"OML can manage users",
		t,
		mode,
		true,
		`---
		user/1/organization_management_level: can_manage_users
		`,
		withRequestUser(1),
		withElementID(2),
	)

	testCase(
		"X == Y",
		t,
		mode,
		false,
		``,
		withRequestUser(1),
		withElementID(1),
	)
}

func TestUserModeG(t *testing.T) {
	var u collection.User

	testCase(
		"No perms",
		t,
		u.Modes("G"),
		false,
		`user/2/id: 2`,
		withRequestUser(1),
		withElementID(2),
	)

	testCase(
		"OML can manage users",
		t,
		u.Modes("G"),
		false,
		`---
		user/1/organization_management_level: can_manage_users
		`,
		withRequestUser(1),
		withElementID(2),
	)

	testCase(
		"X == Y",
		t,
		u.Modes("G"),
		false,
		``,
		withRequestUser(1),
		withElementID(1),
	)
}

func TestUserSuperAdminModeG(t *testing.T) {
	var u collection.User

	testCase(
		"Superadmin",
		t,
		u.SuperAdmin("G"),
		false,
		``,
		withRequestUser(1),
		withElementID(2),
	)
}

func TestUserModeH(t *testing.T) {
	var u collection.User

	testCase(
		"request superadmin",
		t,
		u.Modes("H"),
		false,
		`---
		user/2:
			organization_management_level: superadmin
		`,
		withRequestUser(1),
		withElementID(2),
		withPerms(5, perm.UserCanManage),
	)

	testCase(
		"request superadmin as orga manager",
		t,
		u.Modes("H"),
		false,
		`---
		user/1/organization_management_level: can_manage_organization
		user/2:
			organization_management_level: superadmin
		`,
		withRequestUser(1),
		withElementID(2),
		withPerms(5, perm.UserCanManage),
	)

	testCase(
		"request organization manager",
		t,
		u.Modes("H"),
		false,
		`---
		user/2:
			organization_management_level: can_manage_organization
		`,
		withRequestUser(1),
		withElementID(2),
		withPerms(5, perm.UserCanManage),
	)

	testCase(
		"request organization manager as can_manage_organization",
		t,
		u.Modes("H"),
		true,
		`---
		user/1/organization_management_level: can_manage_organization
		user/2:
			organization_management_level: can_manage_organization
		`,
		withRequestUser(1),
		withElementID(2),
		withPerms(5, perm.UserCanManage),
	)

	testCase(
		"request organization user manager",
		t,
		u.Modes("H"),
		false,
		`---
		user/2:
			organization_management_level: can_manage_users
		`,
		withRequestUser(1),
		withElementID(2),
		withPerms(5, perm.UserCanManage),
	)

	testCase(
		"request organization user manager as orga manager",
		t,
		u.Modes("H"),
		true,
		`---
		user/1/organization_management_level: can_manage_organization
		user/2:
			organization_management_level: can_manage_users
		`,
		withRequestUser(1),
		withElementID(2),
		withPerms(5, perm.UserCanManage),
	)

	testCase(
		"request organization user manager as user manager",
		t,
		u.Modes("H"),
		true,
		`---
		user/1/organization_management_level: can_manage_users
		user/2:
			organization_management_level: can_manage_users
		`,
		withRequestUser(1),
		withElementID(2),
		withPerms(5, perm.UserCanManage),
	)
}
