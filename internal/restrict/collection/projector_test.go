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
		"projector/1/meeting_id: 1",
		withPerms(1, perm.ProjectorCanSee),
	)

	testCase(
		"no perm",
		t,
		f,
		false,
		"projector/1/meeting_id: 1",
	)
}
