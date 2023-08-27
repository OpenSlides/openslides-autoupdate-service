package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
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
