package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-go/perm"
)

func TestProjectorModeA(t *testing.T) {
	f := collection.Projector{}.Modes("A")

	testCase(
		"can see",
		t,
		f,
		true,
		"projector/1/meeting_id: 30",
		withPerms(30, perm.ProjectorCanSee),
	)

	testCase(
		"no perm",
		t,
		f,
		false,
		"projector/1/meeting_id: 30",
	)

	testCase(
		"reference projector with projector.can_see",
		t,
		f,
		true,
		`
		projector/1:
			meeting_id: 30
			used_as_reference_projector_meeting_id: 30
		`,
		withPerms(30, perm.ProjectorCanSee),
	)

	testCase(
		"reference projector with meeting.can_see_autopilot",
		t,
		f,
		true,
		`
		projector/1:
			meeting_id: 30
			used_as_reference_projector_meeting_id: 30
		`,
		withPerms(30, perm.MeetingCanSeeAutopilot),
	)

	testCase(
		"reference projector with no perms",
		t,
		f,
		false,
		`
		projector/1:
			meeting_id: 30
			used_as_reference_projector_meeting_id: 30
		`,
	)

	testCase(
		"not reference projector with meeting.can_see_autopilot",
		t,
		f,
		false,
		`
		projector/1/meeting_id: 30
		`,
		withPerms(30, perm.MeetingCanSeeAutopilot),
	)

	testCase(
		"reference projector with meeting.can_see_autopilot but internal",
		t,
		f,
		false,
		`
		projector/1:
			meeting_id: 30
			used_as_reference_projector_meeting_id: 30
			is_internal: true
		`,
		withPerms(30, perm.MeetingCanSeeAutopilot),
	)

	testCase(
		"can see with internal",
		t,
		f,
		false,
		`
		projector/1:
			meeting_id: 30
			is_internal: true
		`,
		withPerms(30, perm.ProjectorCanSee),
	)

	testCase(
		"can manage with internal",
		t,
		f,
		true,
		`
		projector/1:
			meeting_id: 30
			is_internal: true
		`,
		withPerms(30, perm.ProjectorCanManage),
	)
}
