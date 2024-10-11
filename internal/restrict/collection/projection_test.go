package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

func TestProjectionModeA(t *testing.T) {
	f := collection.Projection{}.Modes("A")

	testCase(
		"manager",
		t,
		f,
		true,
		"projection/1/meeting_id: 30",
		withPerms(30, perm.ProjectorCanManage),
	)

	testCase(
		"no perm",
		t,
		f,
		false,
		"projection/1/meeting_id: 30",
	)

	testCase(
		"linked on reference projector with no perms",
		t,
		f,
		false,
		`
		projection/1:
			meeting_id: 30
			current_projector_id: 7

		projector/7:
			used_as_reference_projector_meeting_id: 30
			meeting_id: 30
		`,
	)

	testCase(
		"linked on reference projector with meeting.can_see_autopilote",
		t,
		f,
		true,
		`
		projection/1:
			meeting_id: 30
			current_projector_id: 7

		projector/7:
			used_as_reference_projector_meeting_id: 30
			meeting_id: 30
		`,
		withPerms(30, perm.MeetingCanSeeAutopilot),
	)

	testCase(
		"linked on normal projector with projector.can_see",
		t,
		f,
		true,
		`
		projection/1:
			meeting_id: 30
			current_projector_id: 7

		projector/7/meeting_id: 30	
		`,
		withPerms(30, perm.ProjectorCanSee),
	)
}
