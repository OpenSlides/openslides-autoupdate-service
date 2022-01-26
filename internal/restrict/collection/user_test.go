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
		f,
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
		f,
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
		f,
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
		f,
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
		f,
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
		f,
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
		f,
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

	testCase(
		"motion submitter",
		t,
		f,
		true,
		`---
		user/2:
			submitted_motion_$_ids: ["1"]
			submitted_motion_$1_ids: [4]
		
		motion/4:
			meeting_id: 1
			state_id: 5
		
		motion_state/5/id: 5
		`,
		withRequestUser(1),
		withElementID(2),
		withPerms(1, perm.MotionCanSee),
	)

	testCase(
		"motion supporter",
		t,
		f,
		true,
		`---
		user/2:
			supported_motion_$_ids: ["1"]
			supported_motion_$1_ids: [4]
		
		motion/4:
			meeting_id: 1
			state_id: 5
		
		motion_state/5/id: 5
		`,
		withRequestUser(1),
		withElementID(2),
		withPerms(1, perm.MotionCanSee),
	)

	testCase(
		"linked in option",
		t,
		f,
		true,
		`---
		user/2:
			option_$_ids: ["1"]
			option_$1_ids: [4]
		
		option/4/poll_id: 5
		poll/5:
			meeting_id: 1
			content_object_id: topic/5
		topic/5/meeting_id: 1
		meeting/1/enable_anonymous: true
		`,
		withRequestUser(1),
		withElementID(2),
		withPerms(1, perm.AgendaItemCanSee),
	)

	testCase(
		"assignment candidate",
		t,
		f,
		true,
		`---
		user/2:
			assignment_candidate_$_ids: ["1"]
			assignment_candidate_$1_ids: [4]
		
		assignment_candidate/4/assignment_id: 5
		assignment/5/meeting_id: 1
		`,
		withRequestUser(1),
		withElementID(2),
		withPerms(1, perm.AssignmentCanSee),
	)

	testCase(
		"speaker",
		t,
		f,
		true,
		`---
		user/2:
			speaker_$_ids: ["1"]
			speaker_$1_ids: [4]
		
		speaker/4/list_of_speakers_id: 5

		list_of_speakers/5:
			meeting_id: 1
			content_object_id: topic/10

		topic/10/meeting_id: 1
		`,
		withRequestUser(1),
		withElementID(2),
		withPerms(1, perm.ListOfSpeakersCanSee, perm.AgendaItemCanSee),
	)

	testCase(
		"poll vote",
		t,
		f,
		true,
		`---
		user/2:
			poll_voted_$_ids: ["1"]
			poll_voted_$1_ids: [4]
		
		poll/4:
			state: finished
			meeting_id: 1
			content_object_id: topic/5
		
		topic/5/meeting_id: 1
		
		meeting/1/id: 1
		`,
		withRequestUser(1),
		withElementID(2),
		withPerms(1, perm.AgendaItemCanSee),
	)

	testCase(
		"vote user ids",
		t,
		f,
		true,
		`---
		user/2:
			vote_$_ids: ["1"]
			vote_$1_ids: [4]
		
		vote/4/option_id: 5
		option/5/poll_id: 6
		poll/6:
			state: published
			meeting_id: 1

		meeting/1/enable_anonymous: true
		`,
		withRequestUser(1),
		withElementID(2),
	)

	testCase(
		"vote delegated ids",
		t,
		f,
		true,
		`---
		user/2:
			vote_delegated_vote_$_ids: ["1"]
			vote_delegated_vote_$1_ids: [4]
		
		vote/4/option_id: 5
		option/5/poll_id: 6
		poll/6:
			state: published
			meeting_id: 1

		meeting/1/enable_anonymous: true
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
		user:
			1:
				group_$1_ids: [5]

			2:
				chat_message_$_ids: ["1"]
				chat_message_$1_ids: [4]
		
		meeting/1/id: 1
		
		chat_message/4:
			user_id: 2
			chat_group_id: 3
		
		chat_group/3:
			read_group_ids: [5]
			meeting_id: 1

		group/5/id: 5
		`,
		withRequestUser(1),
		withElementID(2),
	)
}

func TestUserModeB(t *testing.T) {
	var u collection.User

	testCase(
		"X == Y",
		t,
		u.Modes("B"),
		true,
		``,
		withRequestUser(1),
		withElementID(1),
	)

	testCase(
		"X != Y",
		t,
		u.Modes("B"),
		false,
		``,
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
		user/2/group_$_ids: ["5"]
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
		user/2/group_$_ids: []
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
		u.Modes("E"),
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
		user/2/group_$_ids: ["5"]
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
		user/2/group_$_ids: []
		`,
		withRequestUser(1),
		withElementID(2),
		withPerms(5, perm.UserCanManage),
	)
}

func TestUserModeF(t *testing.T) {
	var u collection.User

	testCase(
		"No perms",
		t,
		u.Modes("F"),
		false,
		`user/2/id: 2`,
		withRequestUser(1),
		withElementID(2),
	)

	testCase(
		"OML can manage users",
		t,
		u.Modes("F"),
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
		u.Modes("F"),
		true,
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

func TestPersonalNoteSuperAdminModeG(t *testing.T) {
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
