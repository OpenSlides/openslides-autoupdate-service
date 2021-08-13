package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
)

func TestProjectionModeA(t *testing.T) {
	f := collection.Projection{}.Modes("A")

	testCase(
		"can see",
		t,
		f,
		true,
		"projection/1/meeting_id: 1",
		withPerms(1, perm.ProjectorCanSee),
	)

	testCase(
		"no perm",
		t,
		f,
		false,
		"projection/1/meeting_id: 1",
	)
}
