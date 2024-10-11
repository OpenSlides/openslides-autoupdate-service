package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

func TestMeetingMediafileModeA(t *testing.T) {
	var m collection.MeetingMediafile

	testCase(
		"No perms",
		t,
		m.Modes("A"),
		false,
		`---
		meeting_mediafile/1:
			meeting_id: 7

		meeting/7:
			id: 7
			committee_id: 404
		`,
	)

	testCase(
		"Anonymous",
		t,
		m.Modes("A"),
		true,
		`---
		organization/1/enable_anonymous: true
		meeting_mediafile/1:
			meeting_id: 7
			is_public: false
			inherited_access_group_ids: [1337]
		meeting/7:
			enable_anonymous: true
			anonymous_group_id: 1337
		`,
		withPerms(7, perm.MediafileCanSee),
		withRequestUser(0),
	)

	testCase(
		"Admin",
		t,
		m.Modes("A"),
		true,
		`---
		meeting_mediafile/1:
			meeting_id: 7
		meeting/7:
			admin_group_id: 8
			group_ids: [8]
			committee_id: 12

		group/8/meeting_user_ids: [10]
		user/1/meeting_user_ids: [10]
		meeting_user/10/group_ids: [8]
		meeting_user/10/meeting_id: 7
		meeting_user/10/user_id: 1
		`,
	)

	testCase(
		"In Meeting",
		t,
		m.Modes("A"),
		false,
		`---
		meeting_mediafile/1:
			meeting_id: 7
		
		meeting/7:
			group_ids: [2]
			committee_id: 8
		group/2/meeting_user_ids: [10]
		meeting_user/10/user_id: 1
		`,
	)

	testCase(
		"Logo with see meeting",
		t,
		m.Modes("A"),
		true,
		`---
		meeting_mediafile/3:
			meeting_id: 7
			used_as_logo_projector_main_in_meeting_id: 5
		meeting/7:
			group_ids: [2]

		group/2/meeting_user_ids: [10]
		meeting_user/10:
			user_id: 1
			meeting_id: 7
		user/1/meeting_user_ids: [10]
		`,
		withElementID(3),
	)

	testCase(
		"On current projection with perm",
		t,
		m.Modes("A"),
		true,
		`---
		meeting_mediafile/1:
			meeting_id: 7
			projection_ids: [4]

		projection/4:
			current_projector_id: 5
			meeting_id: 7

		projector/5/meeting_id: 7

		meeting/7:
			committee_id: 404
			group_ids: [2]

		group/2/meeting_user_ids: [10]
		meeting_user/10/user_id: 1
		`,
		withPerms(7, perm.ProjectorCanSee),
	)

	testCase(
		"On current projection without perm",
		t,
		m.Modes("A"),
		false,
		`---
		meeting_mediafile/1:
			meeting_id: 7
			projection_ids: [4]
		
		meeting/7:
			committee_id: 404
			group_ids: [2]

		group/2/meeting_user_ids: [10]
		meeting_user/10/user_id: 1

		projection/4/current_projector_id: 5
		`,
	)

	testCase(
		"Not on current projection with perm",
		t,
		m.Modes("A"),
		false,
		`---
		meeting_mediafile/1:
			meeting_id: 7
			projection_ids: [4]

		meeting/7:
			group_ids: [2]
			committee_id: 404
		
		group/2/meeting_user_ids: [10]
		meeting_user/10/user_id: 1

		projection/4/meeting_id: 7
		`,
		withPerms(7, perm.ProjectorCanSee),
	)

	testCase(
		"On autopilot projector with meeting.can_see_autopilot",
		t,
		m.Modes("A"),
		true,
		`---
		meeting_mediafile/1:
			meeting_id: 30
			projection_ids: [4]

		projection/4:
			current_projector_id: 7
			meeting_id: 30

		projector/7:
			used_as_reference_projector_meeting_id: 30
			meeting_id: 30

		meeting/30:
			committee_id: 404
			group_ids: [2]

		group/2/meeting_user_ids: [10]
		meeting_user/10/user_id: 1
		`,
		withPerms(30, perm.MeetingCanSeeAutopilot),
	)

	testCase(
		"meeting mediafile can_manage",
		t,
		m.Modes("A"),
		false,
		`---
		meeting_mediafile/1:
			meeting_id: 7
			inherited_access_group_ids: []
			is_public: false

		meeting/7:
			group_ids: [2]
			committee_id: 300

		group/2/meeting_user_ids: [10]
		meeting_user/10/user_id: 1
		`,
		withPerms(7, perm.MediafileCanManage),
	)

	testCase(
		"meeting mediafile can_see not public",
		t,
		m.Modes("A"),
		false,
		`---
		meeting_mediafile/1:
			meeting_id: 7
			is_public: false

		meeting/7:
			committee_id: 300
			group_ids: [2]

		group/2/meeting_user_ids: [10]
		meeting_user/10/user_id: 1
		`,
		withPerms(7, perm.MediafileCanSee),
	)

	testCase(
		"meeting mediafile can_see is public",
		t,
		m.Modes("A"),
		true,
		`---
		meeting_mediafile/1:
			meeting_id: 7
			is_public: true

		meeting/7:
			committee_id: 300
			group_ids: [2]

		group/2/meeting_user_ids: [10]
		meeting_user/10/user_id: 1
		`,
		withPerms(7, perm.MediafileCanSee),
	)

	testCase(
		"meeting mediafile can_see in inherited_access_group_ids",
		t,
		m.Modes("A"),
		true,
		`---
		meeting_mediafile/1:
			meeting_id: 7
			inherited_access_group_ids: [3]
			is_public: false

		meeting/7:
			committee_id: 300
			group_ids: [3]

		group/3/id: 3
		group/3/meeting_user_ids: [10]
		
		user/1/meeting_user_ids: [10]
		meeting_user/10/group_ids: [3]
		meeting_user/10/meeting_id: 7
		meeting_user/10/user_id: 1
		`,
		withPerms(7, perm.MediafileCanSee),
	)

	testCase(
		"meeting mediafile can_see not in inherited_access_group_ids",
		t,
		m.Modes("A"),
		false,
		`---
		meeting_mediafile/1:
			meeting_id: 7
			inherited_access_group_ids: [3]
			is_public: false

		meeting/7:
			id: 7
			committee_id: 300
		group/3/id: 3
		group/4/id: 4
		`,
		withPerms(7, perm.MediafileCanSee),
	)

	testCase(
		"meeting mediafile without perm can_see in inherited_access_group_ids",
		t,
		m.Modes("A"),
		false,
		`---
		meeting_mediafile/1:
			meeting_id: 7
			inherited_access_group_ids: [3]

		meeting/7:
			id: 7
			committee_id: 300

		group/3/id: 3
		`,
	)
}
