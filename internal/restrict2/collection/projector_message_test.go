package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
)

func TestProjectorMessageModeA(t *testing.T) {
	t.Parallel()
	f := collection.ProjectorMessage{}.Modes("A")

	testCase(
		"can see",
		t,
		f,
		true,
		"projector_message/1/meeting_id: 30",
		withPerms(30, perm.ProjectorCanSee),
	)

	testCase(
		"no perm",
		t,
		f,
		false,
		"projector_message/1/meeting_id: 30",
	)
}
