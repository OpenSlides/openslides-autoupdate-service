package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-go/perm"
)

func TestMediafileModeA(t *testing.T) {
	var m collection.Mediafile

	testCase(
		"No perms",
		t,
		m.Modes("A"),
		false,
		`---
		mediafile/1:
			owner_id: meeting/7
			meeting_mediafile_ids: [2]

		meeting_mediafile/2:
			meeting_id: 7

		meeting/7:
			id: 7
			committee_id: 404
		`,
	)

	testCase(
		"No perms organization file",
		t,
		m.Modes("A"),
		false,
		`---
		mediafile/1:
			owner_id: organization/1
		`,
	)

	testCase(
		"Public access organization",
		t,
		m.Modes("A"),
		false,
		`---
		mediafile/1:
			owner_id: organization/1
		`,
		withRequestUser(0),
	)

	testCase(
		"Public access published organization",
		t,
		m.Modes("A"),
		true,
		`---
		organization/1/enable_anonymous: true
		mediafile/1:
			owner_id: organization/1
			published_to_meetings_in_organization_id: 1
			meeting_mediafile_ids: [2]
		meeting_mediafile/2:
			meeting_id: 7
			mediafile_id: 1
			is_public: false
			inherited_access_group_ids: [1337]
		meeting/7:
			enable_anonymous: true
			anonymous_group_id: 1337
			committee_id: 300
		`,
		withPerms(7, perm.MediafileCanSee),
		withRequestUser(0),
	)

	testCase(
		"Public access published organization not public with token",
		t,
		m.Modes("A"),
		true,
		`---
		mediafile/1:
			owner_id: organization/1
			published_to_meetings_in_organization_id: 1
			meeting_mediafile_ids: [2]
			token: web_header
		meeting_mediafile/2:
			meeting_id: 7
			mediafile_id: 1
			is_public: false
			inherited_access_group_ids: [3]
		meeting/7:
			committee_id: 300

		group/3/id: 3
		`,
		withPerms(7, perm.MediafileCanSee),
		withRequestUser(0),
	)

	testCase(
		"Admin via meeting mediafile",
		t,
		m.Modes("A"),
		true,
		`---
		mediafile/1:
			owner_id: meeting/7
			meeting_mediafile_ids: [2]
		meeting_mediafile/2:
			meeting_id: 7
		meeting/7:
			admin_group_id: 8
			committee_id: 70
			group_ids: [8]
		committee/70/id: 70

		group/8/admin_group_for_meeting_id: 7
		group/8/meeting_user_ids: [10]
		user/1/meeting_user_ids: [10]
		meeting_user/10/group_ids: [8]
		meeting_user/10/meeting_id: 7
		meeting_user/10/user_id: 1
		`,
	)

	testCase(
		"Admin via published to orga",
		t,
		m.Modes("A"),
		true,
		`---
		mediafile/1:
			owner_id: organization/1
			published_to_meetings_in_organization_id: 1
		meeting/7:
			admin_group_id: 8
			committee_id: 70
		committee/70/id: 70

		user/1/meeting_user_ids: [10]
		meeting_user/10/group_ids: [8]
		meeting_user/10/meeting_id: 7
		group/8:
			meeting_user_ids: [10]
			admin_group_for_meeting_id: 7
		`,
	)

	testCase(
		"In Meeting",
		t,
		m.Modes("A"),
		false,
		`---
		mediafile/3:
			owner_id: meeting/7
			meeting_mediafile_ids: [2]

		meeting_mediafile/2:
			meeting_id: 7

		meeting/7:
			group_ids: [2]
			committee_id: 8
		group/2/meeting_user_ids: [10]
		meeting_user/10/user_id: 1
		`,
		withElementID(3),
	)

	testCase(
		"Logo without see",
		t,
		m.Modes("A"),
		false,
		`---
		mediafile/3:
			owner_id: meeting/7
			meeting_mediafile_ids: [2]

		meeting_mediafile/2:
			meeting_id: 7

		meeting/7:
			id: 7
			committee_id: 404
		`,
		withElementID(3),
	)

	testCase(
		"Logo with see",
		t,
		m.Modes("A"),
		true,
		`---
		mediafile/3:
			owner_id: meeting/7
			meeting_mediafile_ids: [2]

		meeting_mediafile/2:
			meeting_id: 7
			used_as_logo_projector_main_in_meeting_id: 5
		meeting/7:
			group_ids: [2]
			committee_id: 5
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
		mediafile/1:
			owner_id: meeting/7
			meeting_mediafile_ids: [2]

		meeting_mediafile/2:
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
		mediafile/1:
			owner_id: meeting/7
			meeting_mediafile_ids: [2]

		meeting_mediafile/2:
			meeting_id: 7
			projection_ids: [4]

		meeting/7:
			id: 7
			committee_id: 404

		projection/4/current_projector_id: 5
		`,
	)

	testCase(
		"On not current projection with perm",
		t,
		m.Modes("A"),
		false,
		`---
		mediafile/1:
			owner_id: meeting/7
			meeting_mediafile_ids: [2]

		meeting_mediafile/2:
			meeting_id: 7
			projection_ids: [4]

		meeting/7:
			id: 7
			committee_id: 404

		projection/4/id: 4
		`,
		withPerms(7, perm.ProjectorCanSee),
	)

	testCase(
		"mediafile can_manage",
		t,
		m.Modes("A"),
		false,
		`---
		mediafile/1:
			owner_id: meeting/7
			meeting_mediafile_ids: [2]

		meeting_mediafile/2:
			meeting_id: 7

		meeting/7:
			id: 7
			committee_id: 300
		`,
		withPerms(7, perm.MediafileCanManage),
	)

	testCase(
		"mediafile can_see not public",
		t,
		m.Modes("A"),
		false,
		`---
		mediafile/1:
			owner_id: meeting/7
			meeting_mediafile_ids: [2]

		meeting_mediafile/2:
			meeting_id: 7
			is_public: false

		meeting/7:
			id: 7
			committee_id: 300
		`,
		withPerms(7, perm.MediafileCanSee),
	)

	testCase(
		"mediafile can_see is public",
		t,
		m.Modes("A"),
		true,
		`---
		mediafile/1:
			owner_id: meeting/7
			meeting_mediafile_ids: [2]

		meeting_mediafile/2:
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
		"mediafile can_see in inherited_access_group_ids",
		t,
		m.Modes("A"),
		true,
		`---
		mediafile/1:
			owner_id: meeting/7
			meeting_mediafile_ids: [2]

		meeting_mediafile/2:
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
		"mediafile can_see not in inherited_access_group_ids",
		t,
		m.Modes("A"),
		false,
		`---
		mediafile/1:
			owner_id: meeting/7
			meeting_mediafile_ids: [2]

		meeting_mediafile/2:
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
		"mediafile without perm can_see in inherited_access_group_ids",
		t,
		m.Modes("A"),
		false,
		`---
		mediafile/1:
			owner_id: meeting/7
			meeting_mediafile_ids: [2]

		meeting_mediafile/2:
			meeting_id: 7
			inherited_access_group_ids: [3]

		meeting/7:
			id: 7
			committee_id: 300

		group/3/id: 3
		`,
	)

	testCase(
		"mediafile without perm can_see as orga admin",
		t,
		m.Modes("A"),
		true,
		`---
		user/1/organization_management_level: can_manage_organization

		mediafile/1:
			owner_id: meeting/7
			meeting_mediafile_ids: [2]

		meeting_mediafile/2:
			meeting_id: 7
			inherited_access_group_ids: [3]

		meeting/7:
			id: 7
			committee_id: 300

		group/3/id: 3
		`,
	)

	testCase(
		"mediafile without perm can_see as committee admin",
		t,
		m.Modes("A"),
		true,
		`---
		user/1/committee_management_ids: [1]

		mediafile/1:
			owner_id: organization/1
			published_to_meetings_in_organization_id: 1
		`,
	)
}
