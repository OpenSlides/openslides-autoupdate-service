package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
)

func TestMediafileModeA(t *testing.T) {
	var m collection.Mediafile

	testCase(
		"No perms",
		t,
		m.Modes("A"),
		false,
		`---
			mediafile/1/meeting_id: 7
			meeting/7/id: 7
			`,
	)

	testCase(
		"Admin",
		t,
		m.Modes("A"),
		true,
		`---
			mediafile/1/meeting_id: 7
			meeting/7/admin_group_id: 8
			user/1/group_$7_ids: [8]
			`,
	)

	testCase(
		"In Meeting",
		t,
		m.Modes("A"),
		false,
		`---
			mediafile/1/meeting_id: 7
			meeting/7/user_ids: [1]
			`,
	)

	testCase(
		"Logo without see",
		t,
		m.Modes("A"),
		false,
		`---
			mediafile/1:
				meeting_id: 7
				used_as_logo_$_in_meeting_id: ["foo"]
			meeting/7/id: 7
			`,
	)

	testCase(
		"Logo with see",
		t,
		m.Modes("A"),
		true,
		`---
			mediafile/1:
				meeting_id: 7
				used_as_logo_$_in_meeting_id: ["foo"]
			meeting/7/user_ids: [1]
			`,
	)

	testCase(
		"On current projection with perm",
		t,
		m.Modes("A"),
		true,
		`---
			mediafile/1:
				meeting_id: 7
				projection_ids: [4]
			projection/4/current_projector_id: 5
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
				meeting_id: 7
				projection_ids: [4]
			meeting/7/id: 7
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
				meeting_id: 7
				projection_ids: [4]
			meeting/7/id: 7
			projection/4/id: 4
			`,
		withPerms(7, perm.ProjectorCanSee),
	)

	testCase(
		"mediafile can_see not public",
		t,
		m.Modes("A"),
		false,
		`---
			mediafile/1:
				meeting_id: 7
			meeting/7/id: 7
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
				meeting_id: 7
				is_public: true
			meeting/7/id: 7
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
				meeting_id: 7
				inherited_access_group_ids: [3]
			meeting/7/id: 7
			user/1/group_$7_ids: [3]
			group/3/id: 3
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
				meeting_id: 7
				inherited_access_group_ids: [3]
			meeting/7/id: 7
			user/1/group_$7_ids: [4]
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
				meeting_id: 7
				inherited_access_group_ids: [3]
			meeting/7/id: 7
			user/1/group_$7_ids: [3]
			group/3/id: 3
			`,
	)

	testCase(
		"can see lists of speakers",
		t,
		m.Modes("A"),
		true,
		`---
			mediafile/1:
				list_of_speakers_id: 3
				meeting_id: 4
			list_of_speakers/3/meeting_id: 4
			meeting/4/id: 4
			`,
		withPerms(4, perm.ListOfSpeakersCanSee),
	)
}

func TestMediafileModeB(t *testing.T) {
	var m collection.Mediafile

	testCase(
		"can see lists of speakers",
		t,
		m.Modes("B"),
		false,
		`---
		mediafile/1:
			list_of_speakers_id: 3
			meeting_id: 4
		list_of_speakers/3/meeting_id: 4
		meeting/4/id: 4
		`,
		withPerms(4, perm.ListOfSpeakersCanSee),
	)
}
